export async function fetchProducts(
  products,
  query,
  stores,
  maxPage,
  fastLoad,
  strictSearch,
  page
) {
  try {
    console.log(products);
    const response = await fetch("http://localhost:5000/products/get", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        stores: [...stores],
        query,
        maxPage,
        fastLoad,
        strictSearch,
        page,
      }),
    });
    if (!response.ok) {
      console.error("HTTP error", response.status, response.statusText);
      return;
    }
    const reader = response.body.getReader();
    const decoder = new TextDecoder("utf-8");
    let chunk;

    while (!(chunk = await reader.read()).done) {
      let data = decoder.decode(chunk.value, { stream: true });
      data = data.trimEnd("\n").split("\n");
      try {
        for (let p of data) {
          if (p.trim() !== "") {
            const product = JSON.parse(p.trimEnd("\n"));
            if (product instanceof Array) {
              products.value.push(...product);
            }
            console.log("Product:", product);
          }
        }
      } catch (error) {
        console.error(data);
        console.error("Error parsing JSON:", error);
      }
    }
  } catch (error) {
    console.error(error);
  }
  console.log(products.length, "products");
}
