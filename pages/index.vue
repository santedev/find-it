<script setup>
const input = ref("");
const products = ref([
    {
        "name": "Nicpro-Carpenter-Mechanical-Construction-Woodworking",
        "price": "11.99",
        "old_price": "13.99",
        "picture": "https://m.media-amazon.com/images/I/71oK1YxSLYL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "2,962",
        "product_link": "https://www.amazon.com//Nicpro-Carpenter-Mechanical-Construction-Woodworking/dp/B09WYNGWCM",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Saker-Adjustable-Precisely-Duplicator-Woodworking",
        "price": "29.97",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71YyjBCyyEL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "25,389",
        "product_link": "https://www.amazon.com//Saker-Adjustable-Precisely-Duplicator-Woodworking/dp/B083FBRJPM",
        "rating": "4.5",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Cartman-148Piece-General-Household-Plastic",
        "price": "34.99",
        "old_price": "39.99",
        "picture": "https://m.media-amazon.com/images/I/91JQZ+jOi4L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "22,477",
        "product_link": "https://www.amazon.com//Cartman-148Piece-General-Household-Plastic/dp/B01NCJTSW7",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "WORKPRO-9-piece-Anti-slip-Ratcheting-Combination",
        "price": "45.99",
        "old_price": "5.11",
        "picture": "https://m.media-amazon.com/images/I/81bQcmsTBIL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "226",
        "product_link": "https://www.amazon.com//WORKPRO-9-piece-Anti-slip-Ratcheting-Combination/dp/B0BJDNP5VD",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "WORKPRO-Socket-Adapters-Universal-Joints-Impact-Coupler-Socket-Accessories",
        "price": "39.99",
        "old_price": "49.99",
        "picture": "https://m.media-amazon.com/images/I/71yEjXJNy7L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "678",
        "product_link": "https://www.amazon.com//WORKPRO-Socket-Adapters-Universal-Joints-Impact-Coupler-Socket-Accessories/dp/B0B6N7KKMH",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Goldblatt-Multi-Tool-Commercial-Baseboard-Remodeling",
        "price": "26.99",
        "old_price": "31.99",
        "picture": "https://m.media-amazon.com/images/I/61djX+hCTwL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "7,740",
        "product_link": "https://www.amazon.com//Goldblatt-Multi-Tool-Commercial-Baseboard-Remodeling/dp/B083SG7ZT8",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEWALT-DCK340C2-20V-3-Tool-Combo",
        "price": "224.99",
        "old_price": "270.99",
        "picture": "https://m.media-amazon.com/images/I/71ZwQkucwOL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "718",
        "product_link": "https://www.amazon.com//DEWALT-DCK340C2-20V-3-Tool-Combo/dp/B01D94IED8",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "CRAFTSMAN-VERSASTACK-Mechanics-Drive-CMMT45306",
        "price": "129.00",
        "old_price": "199.00",
        "picture": "https://m.media-amazon.com/images/I/81lNhitd2KL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "6,420",
        "product_link": "https://www.amazon.com//CRAFTSMAN-VERSASTACK-Mechanics-Drive-CMMT45306/dp/B0CMJF6DYD",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "HORUSDY-16-Piece-Chisel-Including-Chisels",
        "price": "19.94",
        "old_price": "20.99",
        "picture": "https://m.media-amazon.com/images/I/81FIhSOXFvL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,129",
        "product_link": "https://www.amazon.com//HORUSDY-16-Piece-Chisel-Including-Chisels/dp/B0BRPVRC77",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEWALT-DCS334B-Brushless-Tool-Only",
        "price": "133.29",
        "old_price": "239.00",
        "picture": "https://m.media-amazon.com/images/I/71fbAzlbkZL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "7,588",
        "product_link": "https://www.amazon.com//DEWALT-DCS334B-Brushless-Tool-Only/dp/B07JPFHQKG",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Klein-Tools-32500HD-Multi-Bit-Screwdriver",
        "price": "21.97",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/4133BYQDfUL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,425",
        "product_link": "https://www.amazon.com//Klein-Tools-32500HD-Multi-Bit-Screwdriver/dp/B0C9V6NYZ4",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEWALT-14-Piece-Degree-Plastic-DWA1184",
        "price": "9.98",
        "old_price": "16.99",
        "picture": "https://m.media-amazon.com/images/I/812mViBQXeL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "14,144",
        "product_link": "https://www.amazon.com//DEWALT-14-Piece-Degree-Plastic-DWA1184/dp/B07DH55YRX",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DeWalt-DWMT73803-Mechanics-Tool-Piece",
        "price": "145.34",
        "old_price": "154.99",
        "picture": "https://m.media-amazon.com/images/I/71EXsP2lG0L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "16,704",
        "product_link": "https://www.amazon.com//DeWalt-DWMT73803-Mechanics-Tool-Piece/dp/B00PXN00BS",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEWALT-DCS331B-20-Volt-Li-Ion-Tool",
        "price": "141.01",
        "old_price": "205.00",
        "picture": "https://m.media-amazon.com/images/I/61N5Jf+VGZL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "7,246",
        "product_link": "https://www.amazon.com//DEWALT-DCS331B-20-Volt-Li-Ion-Tool/dp/B007NVSSFS",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Assorted-Sizes-12-Resistant-ANOSON",
        "price": "6.69",
        "old_price": "0.02",
        "picture": "https://m.media-amazon.com/images/I/81RPjE2FmgL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "3,031",
        "product_link": "https://www.amazon.com//Assorted-Sizes-12-Resistant-ANOSON/dp/B0C2Z4L3S6",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Nut-Driver-Impact-Bit-Set",
        "price": "24.99",
        "old_price": "29.99",
        "picture": "https://m.media-amazon.com/images/I/81D4r+TcGFL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "136",
        "product_link": "https://www.amazon.com//Nut-Driver-Impact-Bit-Set/dp/B0CD6Q5GLL",
        "rating": "4.4",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Extractor-Set-Christmas-Stocking-Stuffers-Stripped",
        "price": "9.99",
        "old_price": "2.00",
        "picture": "https://m.media-amazon.com/images/I/61gp+upu7vL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "2,856",
        "product_link": "https://www.amazon.com//Extractor-Set-Christmas-Stocking-Stuffers-Stripped/dp/B0BTNT4225",
        "rating": "4.2",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Anfrere-Standard-Rethreading-Threads-80-Piece",
        "price": "49.55",
        "old_price": "0.62",
        "picture": "https://m.media-amazon.com/images/I/81-nNTFH14L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "760",
        "product_link": "https://www.amazon.com//Anfrere-Standard-Rethreading-Threads-80-Piece/dp/B098KFYB47",
        "rating": "4.3",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "AXELECT-Pcs-Clip-Pliers-Set",
        "price": "8.99",
        "old_price": "9.99",
        "picture": "https://m.media-amazon.com/images/I/71+EJN-IIDL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,363",
        "product_link": "https://www.amazon.com//AXELECT-Pcs-Clip-Pliers-Set/dp/B08ZMZV2WR",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Coquimbo-Rechargeable-360%C2%B0Rotate-Flashlight-Inspection",
        "price": "14.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71SBVIY+BxL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "20,413",
        "product_link": "https://www.amazon.com//Coquimbo-Rechargeable-360%C2%B0Rotate-Flashlight-Inspection/dp/B07G9X19G1",
        "rating": "4.4",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Bucket-Boss-10030-Bucketeer-BTO",
        "price": "16.48",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/51USgQXzqbL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "11,027",
        "product_link": "https://www.amazon.com//Bucket-Boss-10030-Bucketeer-BTO/dp/B00GK4TOWK",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Screwdriver-Light-weight-Quick-change-Extension-Carabiner",
        "price": "9.99",
        "old_price": "13.99",
        "picture": "https://m.media-amazon.com/images/I/81PENMH-wKL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "7,229",
        "product_link": "https://www.amazon.com//Screwdriver-Light-weight-Quick-change-Extension-Carabiner/dp/B0B2ZPTMW6",
        "rating": "4.5",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Impact-Socket-Adapter-Set-Tools",
        "price": "6.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71Hjjpq9XRL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "326",
        "product_link": "https://www.amazon.com//Impact-Socket-Adapter-Set-Tools/dp/B0BZM98WTQ",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "toolant-Wrenches-Patented-Magnetizing-Demagnetizing",
        "price": "39.99",
        "old_price": "1.82",
        "picture": "https://m.media-amazon.com/images/I/615ZynRJt8L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "161",
        "product_link": "https://www.amazon.com//toolant-Wrenches-Patented-Magnetizing-Demagnetizing/dp/B0CP8TPSDZ",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Automatic-Premium-Adjustable-Plastics-karmiero",
        "price": "4.99",
        "old_price": "6.99",
        "picture": "https://m.media-amazon.com/images/I/61NbP9i4h-L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "657",
        "product_link": "https://www.amazon.com//Automatic-Premium-Adjustable-Plastics-karmiero/dp/B0BR8B8XKJ",
        "rating": "4.1",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Depinning-Electrical-Connector-Extractor-Automotive",
        "price": "9.98",
        "old_price": "18.48",
        "picture": "https://m.media-amazon.com/images/I/8195xgzASmS._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "3,570",
        "product_link": "https://www.amazon.com//Depinning-Electrical-Connector-Extractor-Automotive/dp/B0922GF9W8",
        "rating": "4.3",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Measuring-Stainless-Micrometer-Fraction-Household",
        "price": "19.98",
        "old_price": "26.99",
        "picture": "https://m.media-amazon.com/images/I/613ogMkDwQS._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "6,190",
        "product_link": "https://www.amazon.com//Measuring-Stainless-Micrometer-Fraction-Household/dp/B08Y8M54M9",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "KingTool-General-Toolbox-Storage-Household",
        "price": "109.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/711pTh86caS._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,640",
        "product_link": "https://www.amazon.com//KingTool-General-Toolbox-Storage-Household/dp/B08VWCZ9H8",
        "rating": "4.5",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Klein-Tools-32933-Holding-Feature",
        "price": "19.98",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/41jyyMCiOtL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "534",
        "product_link": "https://www.amazon.com//Klein-Tools-32933-Holding-Feature/dp/B0CQ3SHF3W",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEWALT-Right-Angle-Attachment-Impact",
        "price": "19.75",
        "old_price": "24.99",
        "picture": "https://m.media-amazon.com/images/I/711+trHkZ1L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "19,085",
        "product_link": "https://www.amazon.com//DEWALT-Right-Angle-Attachment-Impact/dp/B07NQS465R",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "DEKOPRO-Combination-Package-Plastic-Toolbox",
        "price": "79.99",
        "old_price": "99.99",
        "picture": "https://m.media-amazon.com/images/I/71PexuvCDNL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "6,728",
        "product_link": "https://www.amazon.com//DEKOPRO-Combination-Package-Plastic-Toolbox/dp/B0723914V1",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "KAFAHOM-Organizer-Organizers-Storage-Screwdriver",
        "price": "59.99",
        "old_price": "74.90",
        "picture": "https://m.media-amazon.com/images/I/81wRxczINnL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "744",
        "product_link": "https://www.amazon.com//KAFAHOM-Organizer-Organizers-Storage-Screwdriver/dp/B0BBGG9G7D",
        "rating": "4.5",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Daxiongmao-Borescope-Endoscope-Waterproof-Inspection",
        "price": "41.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71jKyLDsS9L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,605",
        "product_link": "https://www.amazon.com//Daxiongmao-Borescope-Endoscope-Waterproof-Inspection/dp/B0C3R381FC",
        "rating": "4.3",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Klein-Tools-80016-Circuit-Accessories",
        "price": "54.98",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/61nUGm7CuJL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "7,351",
        "product_link": "https://www.amazon.com//Klein-Tools-80016-Circuit-Accessories/dp/B08VWCSG7G",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "20736A-Inspection-Telescoping-Extendable-Flashlight",
        "price": "15.97",
        "old_price": "18.99",
        "picture": "https://m.media-amazon.com/images/I/71bnINPwTlL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "5,673",
        "product_link": "https://www.amazon.com//20736A-Inspection-Telescoping-Extendable-Flashlight/dp/B09FVQD7WN",
        "rating": "4.5",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Dewalt-DCD771C2-Cordless-Lithium-Ion-Compact",
        "price": "99.00",
        "old_price": "179.00",
        "picture": "https://m.media-amazon.com/images/I/81okj20ws8L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "43,939",
        "product_link": "https://www.amazon.com//Dewalt-DCD771C2-Cordless-Lithium-Ion-Compact/dp/B00ET5VMTU",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "SUNHZMCKP-Magnetic-Screwdriver-Precision-Replaceable",
        "price": "39.99",
        "old_price": "45.99",
        "picture": "https://m.media-amazon.com/images/I/715TDVJ8ADL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,549",
        "product_link": "https://www.amazon.com//SUNHZMCKP-Magnetic-Screwdriver-Precision-Replaceable/dp/B09MYSWLGT",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Collapsible-Magnetic-Foldable-Silicone-Organizes",
        "price": "15.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71FOWbd0-uL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "190",
        "product_link": "https://www.amazon.com//Collapsible-Magnetic-Foldable-Silicone-Organizes/dp/B0B7JBHK2N",
        "rating": "4.4",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Endoscope-Camera-Light-Inspection-Semi-Rigid",
        "price": "42.99",
        "old_price": "49.99",
        "picture": "https://m.media-amazon.com/images/I/71LyENHO39L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "748",
        "product_link": "https://www.amazon.com//Endoscope-Camera-Light-Inspection-Semi-Rigid/dp/B0CFF8Z1NF",
        "rating": "4.7",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "SUNALLY-Organizer-Cordless-Floating-Mounted",
        "price": "29.99",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71o8X7YUBEL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "374",
        "product_link": "https://www.amazon.com//SUNALLY-Organizer-Cordless-Floating-Mounted/dp/B0C4F2KRBV",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "LEATHERMAN-Wingman-Multitool-Stainless-Steel",
        "price": "69.95",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71rL3PDIFWL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "9,181",
        "product_link": "https://www.amazon.com//LEATHERMAN-Wingman-Multitool-Stainless-Steel/dp/B005DI0XM4",
        "rating": "4.8",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "AWLKIM-Meat-Thermometer-Digital-Waterproof",
        "price": "8.49",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/71GGC2pcr6L._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "6,432",
        "product_link": "https://www.amazon.com//AWLKIM-Meat-Thermometer-Digital-Waterproof/dp/B0CL5XT968",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "2457-20-Milwaukee-Cordless-Lithium-Ion-344Y584H351707",
        "price": "94.89",
        "old_price": "",
        "picture": "https://m.media-amazon.com/images/I/51XZepStbqL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "3,108",
        "product_link": "https://www.amazon.com//2457-20-Milwaukee-Cordless-Lithium-Ion-344Y584H351707/dp/B01MEDKJB8",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    },
    {
        "name": "Repairing-Removing-Wallpaper-Scraping-Applying",
        "price": "8.98",
        "old_price": "12.99",
        "picture": "https://m.media-amazon.com/images/I/610FY792AbL._AC_UL320_.jpg",
        "count_sold": "",
        "count_rated": "1,592",
        "product_link": "https://www.amazon.com//Repairing-Removing-Wallpaper-Scraping-Applying/dp/B0B6HLRF7G",
        "rating": "4.6",
        "store": "amazon",
        "free_shipping": ""
    }
]);
const debounce = (func, wait) => {
  let timeout;
  return function (...args) {
    clearTimeout(timeout);
    timeout = setTimeout(() => func.apply("", args), wait);
  };
};

const searchProducts = debounce(
  async (
    query,
    stores = [],
    maxPage = 1,
    fastLoad = 10,
    strictSearch = true,
    page = 1
  ) => {
    if (query == null || query.length == 0) {
      throw new Error("query is null or empty, in function searchProducts");
    }
    products.value = [];
    await fetchProducts(
      products,
      query,
      stores,
      maxPage,
      fastLoad,
      strictSearch,
      page
    );
  },
  200
);
</script>
<template>
  <!--
  <sign-in-google />
  <nuxt-link to="/">home</nuxt-link>
  <nuxt-link to="/login">login</nuxt-link> -->
  <div class="mt-4 mb-3 flex flex-wrap justify-around">
    <img class="m-0 p-0 pb-1 sm:pr-40" src="/svg/find-it.svg" alt="logo">
    <form class="flex justify-center gap-1" method="post" action="Submit" @submit.prevent="searchProducts(input)">
      <input
      type="text"
      class="bg-gray-300 input p-2 text-gray-800 block"
      placeholder="search"
      v-model="input"
      />
      <button class="p-2 border border-gray-600 leading-none" type="submit">send</button>
    </form>
  </div>
  <div
    class="grid-products p-2 grid xl:grid-cols-5 lg:grid-cols-4 sm:grid-cols-3 grid-cols-1  bg-slate-100"
  >
    <a
      class="product-class sm:w-fit w-full mx-auto flex flex-col justify-between cursor-pointer bg-white"
      v-for="(product, index) in products"
      :key="index"
      :href="product.product_link"
    >
      <div class="flex relative align-center justify-center h-max flex-1">
        <img class="store-svg mr-3 mt-4" :src="`/svg/${product.store}.svg`" v-if="product.store.length > 0"></img>
        <img
          loading="lazy"
          class="m-auto"
          v-if="product.picture.length > 0"
          :src="product.picture"
        />
      </div>
      <div
        class="flex flex-col justify-evenly border-t-gray-400 p-3 border-t min-h-44 align-middle"
      >
        <h1 class="text-sm" v-if="product.name.length > 0">
          {{ product.name }}
        </h1>
        <span class="flex flex-wrap gap-2" v-if="product.price.length > 0 && product.old_price.length > 0">
          <p class="text-xl">${{ product.price }}</p>
          <p class="line-through text-gray-400">${{ product.old_price }}</p>
        </span>
        <span v-else-if="product.price.length > 0"
          ><p class="text-xl">${{ product.price }}</p>
        </span>
        <p v-if="product.count_rated.length > 0">
          {{ product.count_rated }} people rated this product
        </p>
        <p class="text-lime-500" v-if="product.free_shipping.length > 0">
          {{ product.free_shipping }}
        </p>        
        <p class="text-amber-500" v-if="product.rating.length > 0">{{ product.rating }} stars</p>
      </div>
    </a>
  </div>
</template>
<style scoped>
.input {
  width: clamp(13.75rem, 40vw, 50rem);
}
.grid-products {
  gap: 0.5rem;
}
.product-class {
  min-width: 180px;
}
.store-svg {
  position: absolute;
  top: 0;
  right: 0;
}
</style>
