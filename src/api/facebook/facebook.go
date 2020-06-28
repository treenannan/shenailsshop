package facebook

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// FACEBOOKAPIURL - facebook graph api url
	FACEBOOKAPIURL = "https://graph.facebook.com/%s/me/messages?access_token=%s"
	// FACEBOOKAPIVER - version facebook graph api
	FACEBOOKAPIVER = "7.0"

	APPACCESSTOKEN  = "2672761339717878|NS02Lt-7mAP1bGN55mbqhUx3rfw"
	PAGEACCESSTOKEN = "EAAlZB3L45KPYBABhbQjpebCZCogoXqYbmKUTC0dvhUQJjVpjse2zjkMkdCxFSTiAzN7DMCWZAUPwspawJplA6TOj2Op0NQkB5v4Kp4oSc97ZAW38ELsLSneiU2sAbjJ0PubqqGevZAMHQMzC6ZABMvi0ZCm2bmjwuyBUFjwTT64XAZDZD"
)

func InitFacebookAPI() {
	fmt.Println("InitFacebookAPI")
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndpoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndpoint).Methods("POST")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}
