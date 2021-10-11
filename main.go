package main

import (
	"Chise1/gomicro/Services"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main() {
	user := Services.UserService{}
	endpoint := Services.GenUserEndpoint(user)
	serverHandler := httptransport.NewServer(endpoint, Services.DecodeUserRequest, Services.EncodeUserResponse)
	http.ListenAndServe(":8080", serverHandler)
}
