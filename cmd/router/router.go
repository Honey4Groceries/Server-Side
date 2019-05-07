package main

import (
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
)

// Get Category Prices for Multiple Stores
func getCategoryPricesForStores(w http.ResponseWriter, r *http.Request) {
    stores := r.URL.Query() // Store list of stores to consider into a var
    params := mux.Vars(r)   // Used to get category_id from endpoint URL

    // Query Firebase for category_id given
    resp, _ := http.Get("https://honey4groceries.firebaseio.com/categories/" + params["category_id"] + "/storeToPrice.json")
    //if err != nil {
      //  return err
    //}
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
    for _, store_id := range stores["store"] {
        // Query Firebase for Store Name
        resp, _ =http.Get("https://honey4groceries.firebaseio.com/stores/" + store_id + "/name.json")
        //if err != nil {
        //    return err
       // }
        defer resp.Body.Close()

        // Get Store name from JSON Response
        var storeName StoreName
        json.NewDecoder(resp.Body).Decode(storeName)
        //json.Unmarshal([]byte(resp), &storeName)

        // Close Response Body
        defer resp.Body.Close()

        // Query Firebase for Price
        
        resp, _ = http.Get("https://honey4groceries.firebaseio.com/prices/" + result[store_id].(string) + "/price.json")
        //if err != nil {
          //  return err
        //}
        defer resp.Body.Close()

        // Get Price from JSON Response
        var price Price
        body, _ = ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &price)

        // Close Response Body
        defer resp.Body.Close()

        // Add Store Name and Price pair to storePrices map
        storePrices[storeName.Name] = price.Price
    }

    json.NewEncoder(w).Encode(storePrices)
    //return nil
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
