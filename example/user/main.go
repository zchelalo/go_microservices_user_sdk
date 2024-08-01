package main

import (
	"errors"
	"fmt"
	"log"

	userSdk "github.com/zchelalo/go_microservices_user_sdk/user"
)

func main() {
	userTransport := userSdk.NewHTTPClient("http://localhost:8001", "")

	user, err := userTransport.Get("3df03810-ff9e-4435-b9bd-e90dba645b71a")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			log.Fatal("user not found: ", err.Error())
		}

		log.Fatal(err)
	}

	fmt.Println(user)
}
