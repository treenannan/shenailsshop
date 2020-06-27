package facebook

import (
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
	PAGEACCESSTOKEN = "EAAlZB3L45KPYBAEDidhC0ZB4ReEvKBMFC05IjLzwPhDiQD6Ro7MxiXlyTJif1FU0KRZB37oB3ym3KLsZB9Pus7jUQ2CRFJvFUz3ALOkZBSBZAPO8ZCUqXpccan5PchM1IXxCUoYj0SOx0rhDicaest26CVsR8CRidSZCNFNKYDjJZBwZDZD"
)

func InitFacebookAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndpoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndpoint).Methods("POST")
	if err := http.ListenAndServe("0.0.0.0:80", r); err != nil {
		log.Fatal(err)
	}
}
