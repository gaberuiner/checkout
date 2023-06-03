package main

import (
	"log"
	"net/http"

	"route256/checkout/internal/client/loms/clientcreateorder"
	"route256/checkout/internal/client/loms/clientstocks"
	"route256/checkout/internal/domain/addtocartdomain"
	"route256/checkout/internal/domain/purchasedomain"
	"route256/checkout/internal/handlers/addtocart"
	"route256/checkout/internal/handlers/deletefromcart"
	"route256/checkout/internal/handlers/listcart"
	"route256/checkout/internal/handlers/purchase"
	"route256/libs/srvwrapper"
)

const port = ":8080"

func main() {
	// err := config.Init()
	// if err != nil {
	// 	log.Fatal("ERR: ", err)
	// }
	model_addtocart := addtocartdomain.New(clientstocks.New())

	hand_addtocart := &addtocart.Handler{
		Model: model_addtocart,
	}

	hand_delete_from_cart := &deletefromcart.Handler{}
	hand_list_cart := &listcart.Handler{}
	model_purchase := purchasedomain.New(clientcreateorder.New())
	hand_purchase := &purchase.Handler{
		Model: model_purchase,
	}

	http.Handle("/addToCart", srvwrapper.New(hand_addtocart.Handle))
	http.Handle("/deleteFromCart", srvwrapper.New(hand_delete_from_cart.Handle))
	http.Handle("/listCart", srvwrapper.New(hand_list_cart.Handle))
	http.Handle("/purchase", srvwrapper.New(hand_purchase.Handle))

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ERR: ", err)
	}
}
