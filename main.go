package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func missingVarEnv(envArr ...string) error {
	for _, v := range envArr {
		if v == "" {
			return fmt.Errorf("missing var")
		}
	}
	return nil
}
func stringNotAllVoid(strs ...string) error {
	for _, str := range strs {
		str = strings.Trim(str, " ")
		if len(str) > 0 {
			return nil
		}
	}
	return fmt.Errorf("missing all the strings")
}

type Product struct {
	Name         string `json:"name"`
	Price        string `json:"price"`
	OldPrice     string `json:"old_price"`
	Picture      string `json:"picture"`
	CountSold    string `json:"count_sold"`
	CountRated   string `json:"count_rated"`
	ProductLink  string `json:"product_link"`
	Rating       string `json:"rating"`
	Store        string `json:"store"`
	FreeShipping string `json:"free_shipping"`
}

type AliexpressProduct struct {
	ProductId     string `json:"productId"`
	Title         string `json:"displayTitle"`
	OriginalPrice string `json:"originalPrice.formattedPrice"`
	SalePrice     string `json:"salePrice.formattedPrice"`
	ImageURL      string `json:"imgUrl"`
}

type AliexpressJson struct {
	Data struct {
		Data struct {
			Root struct {
				Fields struct {
					Mods struct {
						ItemList struct {
							Content []struct {
								ProductID string `json:"productId"`
								Image     struct {
									ImgURL string `json:"imgUrl"`
								} `json:"image"`
								Title struct {
									SeoTitle     string `json:"seoTitle"`
									DisplayTitle string `json:"displayTitle"`
								} `json:"title"`
								Evaluation struct {
									StarRating float64 `json:"starRating"`
								} `json:"evaluation"`
								Trade struct {
									TradeDesc string `json:"tradeDesc"`
								} `json:"trade"`
								Prices struct {
									CurrencySymbol string `json:"currencySymbol"`
									OriginalPrice  struct {
										PriceType      string  `json:"priceType"`
										CurrencyCode   string  `json:"currencyCode"`
										MinPrice       float64 `json:"minPrice"`
										MinPriceType   int     `json:"minPriceType"`
										FormattedPrice string  `json:"formattedPrice"`
										Cent           int     `json:"cent"`
									} `json:"originalPrice"`
									SalePrice struct {
										Discount         int     `json:"discount"`
										MinPriceDiscount int     `json:"minPriceDiscount"`
										PriceType        string  `json:"priceType"`
										CurrencyCode     string  `json:"currencyCode"`
										MinPrice         float64 `json:"minPrice"`
										MinPriceType     int     `json:"minPriceType"`
										FormattedPrice   string  `json:"formattedPrice"`
										Cent             int     `json:"cent"`
									} `json:"salePrice"`
								} `json:"prices"`
							} `json:"content"`
						} `json:"itemList"`
					} `json:"mods"`
				} `json:"fields"`
			} `json:"root"`
		} `json:"data"`
	} `json:"data"`
}

// https://www.aliexpress.com/w/wholesale-watch-wood.html
// https://www.amazon.com/s?k=watch+wood
// https://www.alibaba.com/trade/search?tab=all&SearchText=watch+wood
// https://www.temu.com/search_result.html?search_key=watch%20wood
// https://listado.mercadolibre.com.co/watch-wood
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if err := missingVarEnv(port); err != nil {
		log.Fatal(err)
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/products/get", func(w http.ResponseWriter, r *http.Request) {
		var reqData struct {
			Stores       []string `json:"stores"`
			Query        string   `json:"query"`
			MaxPage      int      `json:"maxPage"`
			FastLoad     int      `json:"fastLoad"`
			StrictSearch bool     `json:"strictSearch"`
			Page         int      `json:"page"`
		}
		err := json.NewDecoder(r.Body).Decode(&reqData)
		if err != nil {
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}
		stores := []string{"aliexpress", "amazon", "alibaba", "mercadolibre"}
		query, maxPage, fastLoad, strictSearch, page := "", 1, 10, false, 1
		if len(reqData.Query) > 0 {
			query = reqData.Query
		}
		if reqData.MaxPage > 0 {
			maxPage = reqData.MaxPage
		}
		if reqData.FastLoad > 0 && reqData.FastLoad <= 10 {
			fastLoad = reqData.FastLoad
		}
		if reqData.StrictSearch {
			strictSearch = reqData.StrictSearch
		}
		if len(reqData.Stores) > 0 {
			stores = make([]string, len(reqData.Stores))
			copy(stores, reqData.Stores)
		}
		if reqData.Page > 0 {
			page = reqData.Page
		}
		getProducts(w, query,
			maxPage, fastLoad, strictSearch, page, stores...)
	})

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	handler := cors.Handler(mux)
	fmt.Println("Server is running at http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func getProducts(w http.ResponseWriter, query string, maxPages int, fastLoad int, strictSearch bool, page int, stores ...string) {
	if fastLoad > 10 || fastLoad < 1 {
		fmt.Println("FastLoad param is not between the range from 0 to 10")
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	ch := make(chan []Product)
	done := make(chan struct{})
	var wg sync.WaitGroup

	for _, store := range stores {
		wg.Add(1)
		go func(store string) {
			defer wg.Done()
			scrapeStore(store, query, maxPages, fastLoad, strictSearch, page, ch)
		}(store)
	}

	go func() {
		wg.Wait()
		close(done)
	}()
loop:
	for {
		select {
		case <-done:
			close(ch)
			break loop
		case p, ok := <-ch:
			if !ok {
				break loop
			}
			if err := json.NewEncoder(w).Encode(p); err != nil {
				http.Error(w, "Failed to encode products", http.StatusInternalServerError)
				return
			}
			flusher.Flush()
		}
	}
}

func scrapeStore(store, query string, maxPages, fastLoad int, strictSearch bool, page int, mainCh chan<- []Product) {
	ch := make(chan []Product)
	c := colly.NewCollector(
		colly.UserAgent("ScraperBot/1.0 (+https://mywebsite.com/contact)"),
		colly.AllowURLRevisit(),
	)

	parallelism := 1
	delayReq := 1
	switch store {
	case "mercadolibre":
		parallelism = 5
		delayReq = 2
	}

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*" + store + ".com*",
		Parallelism: parallelism,
		Delay:       time.Duration(delayReq) * time.Second,
		RandomDelay: 500 * time.Millisecond,
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL: %s failed with response: %v and error: %v\n", r.Request.URL, r, err)
		if r.Request.Ctx.GetAny("retry_count") == nil {
			r.Request.Ctx.Put("retry_count", 1)
		}
		retryCountInterface := r.Request.Ctx.GetAny("retry_count")
		var retryCount int
		if retryCountInterface != nil {
			retryCount = retryCountInterface.(int)
		}
		if retryCount >= 3 {
			return
		}
		r.Request.Ctx.Put("retry_count", retryCount+1)
		r.Request.Retry()
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	go func() {
		defer close(ch)
		switch store {
		case "mercadolibre":
			scrapeMercadolibre(c, query, maxPages, fastLoad, strictSearch, store, page, ch)
		case "alibaba":
			scrapeAlibaba(c, query, maxPages, store, page, ch)
		case "amazon":
			scrapeAmazon(c, query, maxPages, store, page, ch)
		case "aliexpress":
			scrapeAliexpress(c, query, maxPages, strictSearch, store, page, ch)
		default:
			close(ch)
		}
	}()

loop:
	for {
		select {
		case <-ctx.Done():
			close(ch)
			break loop
		case p, ok := <-ch:
			if !ok {
				break loop
			}
			mainCh <- p
		}
	}
}

func scrapeMercadolibre(c *colly.Collector, query string, maxPages int, fastLoad int, strictSearch bool, store string, page int, ch chan<- []Product) {
	fetchCount, pageCount, limit, page := 0, 0, 0, page
	sent := false
	var products []Product
	url := "https://listado.mercadolibre.com.co/" + strings.Join(strings.Split(query, " "), "-")
	if page > 1 {
		num := 50 * (page - 1)
		url = url + "_Desde_" + strconv.Itoa(num+1) + "_NoIndex_True"
	}
	 
	var mu sync.Mutex
	var wgOn sync.WaitGroup
	c.OnHTML("li", func(e *colly.HTMLElement) {
		productLink := e.ChildAttr("a.poly-component__title", "href")
		name := e.ChildText("a.poly-component__title h2")
		countedRated := e.ChildText("span.poly-reviews__total")
		picture := e.ChildAttr("img.poly-component__picture", "src")
		price := e.ChildText("div.poly-price__current span.andes-money-amount__fraction")
		rating := e.ChildText("span.poly-reviews__rating")
		oldPrice := e.ChildText("s.andes-money-amount--previous span.andes-money-amount__fraction")
		freeShipping := e.ChildText("div.poly-component__shipping")
		if picture == "" {
			picture = e.ChildAttr("div.andes-carousel-snapped__slide img", "data-src")
		}
		if picture == "" {
			return
		}
		if productLink == "" {
			productLink = e.ChildAttr("a.ui-search-item__group__element", "href")
		}
		if len(countedRated) > 0 {
			countedRated = countedRated[1:1]
		}
		if name == "" {
			name = e.ChildText("a.ui-search-item__group__element h2")
		}
		if price == "" {
			price = e.ChildAttr("div.ui-search-price span.andes-money-amount", "aria-label")
		}
		if len(price) > 1 {
			price = strings.Split(price, " ")[0]
			price = parsePrice(price)
		}
		if freeShipping == "" {
			freeShipping = e.ChildText("p.ui-meliplus-pill span.ui-pb-highlight")
		}
		p := Product{
			Name:         name,
			Price:        price,
			Picture:      picture,
			Rating:       rating,
			ProductLink:  productLink,
			CountRated:   countedRated,
			OldPrice:     oldPrice,
			Store:        store,
			FreeShipping: freeShipping,
		}
		if p.Picture[0] == 'h' {
			mu.Lock()
			products = append(products, p)
			mu.Unlock()
			return
		}
		rdm := rand.Intn(fastLoad) + 1
		if rdm != fastLoad && p.Picture[0] == 'd' && !strictSearch {
			products = append(products, p)
			return
		}
		if (strictSearch && p.Picture[0] == 'd') || (p.Name == "") || (p.Price == "" && strictSearch) {
			return
		}
		if p.Picture[0] == 'd' {
			wgOn.Add(1)
			go func(p Product) {
				defer wgOn.Done()
				ch := make(chan string)
				go fetchProductImage(p.ProductLink, c.Clone(), ch)
				p.Picture = <-ch
				err := stringNotAllVoid(p.Name, p.Price, p.Picture, p.Rating, p.ProductLink, p.CountRated, p.OldPrice, p.FreeShipping)
				if err != nil {
					fmt.Printf("Error in data validation: %v\n", err)
					return
				}
				mu.Lock()
				products = append(products, p)
				mu.Unlock()
			}(p)
		}
		if len(products) >= 10 && !sent {
			wgOn.Wait()
			mu.Lock()
			ch <- products
			products = make([]Product, 0, 50)
			sent = true
			mu.Unlock()
			return
		}
	})

	c.OnHTML(".andes-pagination__button.andes-pagination__button--next a.andes-pagination__link", func(e *colly.HTMLElement) {
		if len(products) < 5 && limit < 2 {
			limit++
			time.Sleep(1 * time.Second)
			reFetchCh := make(chan []Product)
			reFetch(fetchCount, query, maxPages, fastLoad, strictSearch, store, page, reFetchCh)
			fetchCount++
			for productsArr := range reFetchCh {
				mu.Lock()
				for _, p := range productsArr {
					duplicate := false
					for _, sp := range products {
						if p.Name == sp.Name && sp.Store == store {
							duplicate = true
							break
						}
					}
					if !duplicate {
						products = append(products, p)
					}
				}
				mu.Unlock()
			}
		}
		pageCount++
		if pageCount < maxPages && pageCount != 0 {
			nextPage := e.Attr("href")
			e.Request.Visit(nextPage)
		}
	})
	c.Visit(url)
	wgOn.Wait()
	ch <- products
}

func scrapeAlibaba(c *colly.Collector, query string, maxPages int, store string, page int, ch chan<- []Product) {
	pageCount, limit, page := 0, 0, page
	sent := false
	var products []Product
	url := "https://www.alibaba.com/trade/search?tab=all&SearchText=" + strings.Join(strings.Split(query, " "), "+")
	if page > 1 {
		url = url + "&page=" + strconv.Itoa(page)
	}
	
	c.OnHTML("div.organic-list div.fy23-search-card", func(e *colly.HTMLElement) {
		name := e.ChildAttr("h2.search-card-e-title a", "href")
		if len(name) > 0 {
			name = strings.Join(
				strings.Split(
					strings.Split(
						strings.Split(name, "/")[4],
						"_")[0],
					"-"),
				" ")
		}
		productLink := e.ChildAttr("h2.search-card-e-title a", "href")
		if len(productLink) > 0 {
			productLink = "https:" + productLink
		}
		price := e.ChildText("div.search-card-e-price-main")
		picture := e.ChildAttr("div.search-card-e-slider__wrapper img", "src")
		if len(picture) > 0 {
			picture = "https:" + picture
		}
		rating := e.ChildText("span.search-card-e-review strong")
		countRated := e.ChildText("span.search-card-e-review span")
		oldPrice := e.ChildText("div.search-card-e-price__list span.search-card-e-price__original")
		err := stringNotAllVoid(name, price, picture, rating, productLink, countRated, oldPrice)
		if err != nil {
			fmt.Printf("Error in data validation: %v\n", err)
			return
		}
		product := Product{
			Name:        name,
			Price:       price,
			Picture:     picture,
			Rating:      rating,
			ProductLink: productLink,
			CountRated:  countRated,
			OldPrice:    oldPrice,
			Store:       store,
		}

		products = append(products, product)
		if len(products) >= 10 && !sent {
			ch <- products
			products = make([]Product, 0, 50)
			sent = true
		}
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		if len(products) == 0 && limit < 4 {
			limit++
			time.Sleep(1 * time.Second)
			e.Request.Visit(e.Request.URL.String())
			return
		}
		pageCount++
		if pageCount < maxPages && pageCount != 0 {
			page++
			nextPage := url + "&page=" + strconv.Itoa(page)
			e.Request.Visit(nextPage)
		}
	})
	c.Visit(url)
	ch <- products
}

func scrapeAliexpress(c *colly.Collector, query string, maxPages int, strictSearch bool, store string, page int, ch chan<- []Product) {
	pageCount, limit, page := 0, 0, page
	sent := false
	var products []Product
	url := "https://aliexpress.com/w/wholesale-" + strings.Join(strings.Split(query, " "), "-") + ".html"
	if page > 1 {
		url = url + "?page=" + strconv.Itoa(page)
	}

	c.OnHTML("script", func(e *colly.HTMLElement) {
		scriptContent := e.Text
		if strings.Contains(scriptContent, `{ data: {"hierarchy":{"root":`) {

			dataParsed := strings.Trim(
				strings.Split(scriptContent, "window._dida_config_._init_data_=")[1],
				" ")

			jsonData := strings.Replace(dataParsed, "{ data:", `{"data":`, 1)

			var result AliexpressJson

			err := json.Unmarshal([]byte(jsonData), &result)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}
			productList := result.Data.Data.Root.Fields.Mods.ItemList.Content

			for _, p := range productList {
				title := p.Title.SeoTitle
				productId := p.ProductID
				price := p.Prices.OriginalPrice.FormattedPrice
				if price == "" && strictSearch {
					continue
				}
				if len(price) > 0 {
					price = strings.Replace(
						strings.Split(price, " ")[1],
						",", ".", 1)[1:]
				}
				oldPrice := p.Prices.SalePrice.FormattedPrice
				if len(oldPrice) > 0 {
					oldPrice = strings.Replace(
						strings.Split(oldPrice, " ")[1],
						",", ".", 1)[1:]
				}
				imageUrl := p.Image.ImgURL
				if len(imageUrl) > 0 {
					imageUrl = "https:" + imageUrl
				}
				starRating := strconv.FormatFloat(p.Evaluation.StarRating, 'f', 1, 64)
				countSold := p.Trade.TradeDesc
				if len(countSold) > 0 {
					countSold = strings.Split(countSold, " ")[0]
				}

				productLink := "https://www.aliexpress.com/item/" + productId + ".html"

				product := Product{
					Name:        title,
					Price:       price,
					Picture:     imageUrl,
					Rating:      starRating,
					ProductLink: productLink,
					CountSold:   countSold,
					OldPrice:    oldPrice,
					Store:       store,
				}

				err = stringNotAllVoid(product.Name, product.Price, product.Picture, product.Rating, product.ProductLink, product.CountSold, product.OldPrice)
				if err != nil {
					fmt.Printf("Error in data validation: %v\n", err)
					continue
				}

				products = append(products, product)
				if len(products) >= 10 && !sent {
					ch <- products
					products = make([]Product, 0, 50)
					sent = true
				}
			}
		}
	})
	c.OnHTML("html", func(e *colly.HTMLElement) {
		if limit < 2 && len(products) <= 5 {
			limit++
			e.Request.Retry()
		}
		pageCount++
		if pageCount < maxPages && pageCount != 0 {
			page++
			nextPage := url + "?page=" + strconv.Itoa(page)
			e.Request.Visit(nextPage)
		}
	})
	c.Visit(url)
	ch <- products
}

func scrapeAmazon(c *colly.Collector, query string, maxPages int, store string, page int, ch chan<- []Product) {
	pageCount, page := 0, page
	sent := false
	var products []Product
	url := "https://www.amazon.com/s?k=" + strings.Join(strings.Split(query, " "), "+")
	if page > 1 {
		url = url + "&page=" + strconv.Itoa(page)
	}
	c.OnHTML("div.a-section", func(e *colly.HTMLElement) {
		productLink := e.ChildAttr("div.s-product-image-container a", "href")
		name := e.ChildText("h2.a-size-mini.a-spacing-non a span")
		if name == "" && len(productLink) > 0 {
			name = strings.Split(productLink[1:], "/")[0]
		}
		if name == "" {
			return
		}
		price := e.ChildText("span.a-price-whole") + e.ChildText("span.a-price-fraction")
		if price == "" || len(price) >= 18 {
			return
		}
		if len(productLink) > 0 {
			productLink = "https://www.amazon.com/" + productLink
		}
		picture := e.ChildAttr("div.a-section img", "src")
		rating := e.ChildText("a i span.a-icon-alt")
		if len(rating) > 0 {
			rating = strings.Split(rating, " ")[0]
		}
		countRated := e.ChildText("div.s-csa-instrumentation-wrapper span a span")
		oldPrice := e.ChildText("span.a-price.a-text-price")
		if len(oldPrice) > 1 {
			oldPrice = oldPrice[1:]
			oldPrice = strings.Split(oldPrice, "$")[0]
		}
		err := stringNotAllVoid(name, price, picture, productLink)
		if err != nil {
			fmt.Printf("Error in data validation: %v\n", err)
			return
		}
		product := Product{
			Name:        name,
			Price:       price,
			Picture:     picture,
			Rating:      rating,
			ProductLink: productLink,
			CountRated:  countRated,
			OldPrice:    oldPrice,
			Store:       store,
		}
		repeated := false
		for _, p := range products {
			p := p
			if name == p.Name && p.Store == store {
				repeated = true
				break
			}
		}
		if !repeated {
			products = append(products, product)
		}
		if len(products) >= 10 && !sent {
			ch <- products
			products = make([]Product, 0, 50)
			sent = true
		}
	})

	c.OnHTML("a.s-pagination-next", func(e *colly.HTMLElement) {
		pageCount++
		if pageCount < maxPages && pageCount != 0 {
			nextPage := e.Attr("href")
			e.Request.Visit(nextPage)
		}
	})
	c.Visit(url)
	ch <- products
}

func fetchProductImage(url string, c *colly.Collector, ch chan string) {
	c.OnHTML("img.ui-pdp-image.ui-pdp-gallery__figure__image", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
		ch <- imgSrc
	})
	c.Visit(url)
}

func reFetch(count int, query string,
	maxPages int, fastLoad int, strictSearch bool, store string, page int, ch chan<- []Product) {
	if count >= 1 {
		ch <- []Product{}
		return
	}

	reFetchCh := make(chan []Product)
	go func() {
		defer close(reFetchCh)
		scrapeStore(store, query, maxPages, fastLoad, strictSearch, page, reFetchCh)
	}()

	for p := range reFetchCh {
		ch <- p
	}
}

func parsePrice(price string) string {
	if strings.ContainsAny(price, ".") {
		return price
	}
	var builder strings.Builder
	num := len(price) % 3
	i := num

	if num != 0 {
		builder.WriteString(price[0:num] + ".")
	}

	for i < len(price) {
		builder.WriteString(price[i : i+3])
		i += 3
		if i < len(price) {
			builder.WriteString(".")
		}
	}
	return builder.String()
}
