package main

import (
	"fmt"
	"net/http"
    "encoding/json"
	"github.com/gorilla/mux"
)

// Get Category Prices for Multiple Stores
func getCategoryPricesForStores(w http.ResponseWriter, r *http.Request) {
    stores := r.URL.Query() // Store list of stores to consider into a var
    params := mux.Vars(r)   // Used to get category_id from endpoint URL

    // Query Firebase for category_id given
    resp, err := http.Get("https://honey4groceries.firebaseio.com/categories/"
                            + params["category_id"] + "/storeToPrice.json")

    // Parse the JSON Response
    var storesToPrices map[string]interface{}
    json.Unmarshal([]byte(resp), &storesToPrices)

    // Close Response Body
    defer resp.Body.Close()

    // Query for Foursquare store_ids
    foursquare := r.URL.Query()

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
    for _, v := range foursquare {
        // Query Firebase for Store Name
        resp, err :=
        http.Get("https://honey4groceries.firebaseio.com/stores/" + v +
        "/name.json")

        // Get Store name from JSON Response
        var storeName StoreName
        json.Unmarshal([]byte(resp), &storeName)

        // Close Response Body
        defer resp.Body.Close()

        // Query Firebase for Price
        resp, err := http.Get("https://honey4groceries.firebaseio.com/prices/" +
    storesToPrices[v] + "/price.json")

        // Get Price from JSON Response
        var price Price
        json.Unmarshal([]byte(resp), &price)

        // Close Response Body
        defer resp.Body.Close()

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

    // Subrouter for the categories/category_id/prices Endpoint
    categoryPrices := r.PathPrefix("/categories/{category_id}/prices")
                       .Subrouter()

    // Route Handlers for the categories/category_id/prices Endpoint
    categoryPrices.Handler(getCategoryPricesForStores).Methods("GET")
    categoryPrices.Handler(getCategoryPricesForStore).Methods("GET")
        .Queries("store", "{store_id}")


    // Subrouter for the items/item_id/prices Endpoint
    itemPrices := r.PathPrefix("/items/{item_id}/prices")

    // Route Handlers for the items/item_id/prices Endpoint
    itemPrices.Handler(getItemPricesForStores).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
