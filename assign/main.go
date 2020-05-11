package main

import (
	supermarket "assign/first"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type item struct {
	Item	string
	Value	float64
}

Var Items []item

func main() {
	
	supermarket.Get("apple")
	supermarket.Post("orange", 29.7)
	supermarket.Update("apple", 21.1)
	supermarket.Delete("orange")

	itemdata:= Items{item{Item:Apple,Value:20.4 },
                    item{Item:Orange,Value:21.4}}

	http.HandleFunc("/testget", testget) // each request calls handler
	http.HandleFunc("/testpost", testpost)
	http.HandleFunc("/testput", testput)
	http.HandleFunc("/testdelete", testdelete)
	log.Fatal(http.ListenAndServe(":9455", nil))
}

func process(w http.ResponseWriter, cr chan *http.Request) {
	r := <-cr
	var s item
	json.NewDecoder(r.Body).Decode(&s)
	json.NewEncoder(w).Encode(s)

}

//l
func testget(w http.ResponseWriter, r *http.Request) {
	cr := make(chan *http.Request, 1)
	cr <- r
	var pleasewait sync.WaitGroup
	pleasewait.Add(1)

	go func() {
		defer pleasewait.Done()
		process(w, cr)
	}()
	pleasewait.Wait()
	w.WriteHeader(200)
}

func testpost(w http.ResponseWriter, r *http.Request) {
	var newitem item
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter item and value to create")
	}
	json.Unmarshal(reqBody, &newitem)
	Items = append(Items, newitem)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newitem)
}
func testput(w http.ResponseWriter, r *http.Request) {
	var newitem item
	err := json.NewDecoder(r.Body).Decode(&newitem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Println("Product Info - Updated")
		fmt.Println("item: ", newitem.Item)
		fmt.Println("value: ", newitem.Value)
		w.WriteHeader(http.StatusOK)
		result, _ := json.Marshal(newitem)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}

//l
 func testdelete(w http.ResponseWriter, r *http.Request) {
 	 vars := mux.Vars(r)
 	 item := vars["item"]
 		for index, itemdata := range Items {
 		if itemdata.Item == item {
 			Items = append(Items[:index], Items[index+1:]...)
 		}
 	}
 	w.Write(Items)
 }

