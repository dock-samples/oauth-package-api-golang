package main

import (
	"fmt"
	"github.com/dock-samples/oauth-package-api-golang/authorization"
)

var (
	username    = "username"
	password    = "password"
	environment = authorization.Homologation
)

func main() {
	caradhras := authorization.New(username, password, environment)
	token, err := caradhras.GetAccessToken()
	if err != nil {
		panic(err)
	}

	// Return access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())

	token, err = caradhras.GetAccessToken()
	if err != nil {
		panic(err)
	}

	// Return the same access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())

	// Expire the access token
	caradhras.ExpireAccessToken()
	fmt.Println(caradhras.IsExpired())

	token, err = caradhras.GetAccessToken()
	if err != nil {
		panic(err)
	}

	// Return new access token
	fmt.Println(token)
	fmt.Println(caradhras.IsExpired())
}
