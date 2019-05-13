package main

import (
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
    "fmt"
)

// Get Category Prices for Multiple Stores
func getCategoryPricesForStores(w http.ResponseWriter, r *http.Request) {
    stores := r.URL.Query() // Store list of stores to consider into a var
    params := mux.Vars(r)   // Used to get category_id from endpoint URL
    fmt.Println(stores)

    //fmt.Println(params["category_id"])
    // Query Firebase for category_id given
    resp, err := http.Get("https://honey4groceries.firebaseio.com/categories/" + params["category_id"] + "/store to price.json")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("500 - Something bad happened!"))
        //fmt.Println("Hello error")
    }
    defer resp.Body.Close()
    /*
    body, err := ioutil.ReadAll(resp.Body)
    map := &map[string]interface{}{}
    json.Unmarshal(body, map)
    desiredValue := map["desiredKey"]
    */
    
    // Parse the JSON Response
    body, _ := ioutil.ReadAll(resp.Body)
    var result map[string]interface{}
    json.Unmarshal(body, &result)
    fmt.Println(result)

    //fmt.Println(body)
    //result2 := result.(map[string]interface{})
    //storesToPrices := result[].(map[string]interface{})

    var storePrices map[string]string
    storePrices = make(map[string]string)

    // Struct to Unmarshal Store Name
    type StoreName struct {
        Name string
     }

     // Struct to Unmarshal Price
     type Price struct {
        Price string
     }

    // Loop through store_ids from Foursquare
    for _, store_id := range stores["stores"] {
        // Query Firebase for Store Name
        fmt.Println("Stoer id is", store_id)

        resp, err =http.Get("https://honey4groceries.firebaseio.com/stores/venueID1.json")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("500 - Something bad happened!"))
        }
        defer resp.Body.Close()

        body, _ = ioutil.ReadAll(resp.Body)
        fmt.Println(string(body))
        // Get Store name from JSON Response
        var storeName StoreName
        //json.NewDecoder(body).Decode(storeName)

        fmt.Println("storename is", storeName)
        json.Unmarshal(body, &storeName)

        // Query Firebase for Price
        //fmt.Println("hello", result[store_id].(string))
        
        resp, err = http.Get("https://honey4groceries.firebaseio.com/prices/priceID.json")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("500 - Something bad happened!"))
        }
        defer resp.Body.Close()

        // Get Price from JSON Response
        var price Price
        //json.NewDecoder(resp.Body).Decode(price)
        body, _ = ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &price)
        fmt.Println("price from qeury", body)

        // Add Store Name and Price pair to storePrices map
        storePrices[storeName.Name] = price.Price
    }

    json.NewEncoder(w).Encode(storePrices)
}

// Get Category Prices for a Single Store
func getCategoryPricesForStore(w http.ResponseWriter, r *http.Request) {

}

// Get Item Prices for Multiple Stores
func getItemPricesForStores(w http.ResponseWriter, r *http.Request) {

}

// Main Function
func main() {
	r := mux.NewRouter()
    //fmt.Println("Hello")
    // Subrouter for the categories/category_id/prices Endpoint
    categoryPrices := r.PathPrefix("/categories/{category_id}/prices").Subrouter()

    // Route Handlers for the categories/category_id/prices Endpoint
    categoryPrices.HandleFunc("", getCategoryPricesForStores).Methods("GET").Queries("stores", "{store_ids}")
    categoryPrices.HandleFunc("", getCategoryPricesForStore).Methods("GET").Queries("store", "{store_id}")


    // Subrouter for the items/item_id/prices Endpoint
    itemPrices := r.PathPrefix("/items/{item_id}/prices").Subrouter()

    // Route Handlers for the items/item_id/prices Endpoint
    itemPrices.HandleFunc("", getItemPricesForStores).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
