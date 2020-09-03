package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Que es esta importacion?
)

type Cart struct {
	NameOwner string `json:"nameOwner"` // Porque los atributos deben estar en mayusucla (publicos)? No respeta encapsulamiento?
	ID        int    `json:"id"`
	Email     string `json:"email"`
}

type Item struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Cost int    `json:"cost"`
	Desc string `json:"desc"`
}

type CartItem struct {
	IDItem     int `json:"idItem"`
	IDCart     int `json:"idCart"`
	NumerItems int `json:"numberItems"`
	CostItems  int `json:"costItems"`
}

type Carts []Cart

type Items []Item

type CartItems []CartItem

func allCarts(w http.ResponseWriter, r *http.Request) {
	carts := Carts{
		Cart{NameOwner: "Andres", ID: 1, Email: "a.ariza@abc.com"},
		Cart{NameOwner: "Felipe", ID: 2, Email: "f.ariza@abc.com"},
	}

	fmt.Println("Endpoint Hit: allCarts")
	json.NewEncoder(w).Encode(carts)
}

func allItems(w http.ResponseWriter, r *http.Request) {
	items := Items{
		Item{Name: "Producto 1", ID: 1, Cost: 11000, Desc: "Este es el producto 1"},
		Item{Name: "Producto 2", ID: 2, Cost: 21000, Desc: "Este es el producto 2"},
	}

	fmt.Println("Endpoint Hit: allItems")
	json.NewEncoder(w).Encode(items)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/cart", allCarts).Methods("GET")
	router.HandleFunc("/item", allItems).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
