package main

import (
	"log"
	"net/http"

	"google.golang.org/api/option"
)

const ServiceAccountLocation = "./bm-agent-service-account-credentials.json"

var opts = []option.ClientOption{
	option.WithCredentialsFile(ServiceAccountLocation),
	option.WithScopes("https://www.googleapis.com/auth/businessmessages"),
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	server := http.NewServeMux()

	server.HandleFunc("POST /callback", handler)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Println("error is ", err)
		return
	}

	return
}
