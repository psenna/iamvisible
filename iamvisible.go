package main

import (
	"fmt"

	"github.com/psenna/iamvisible/app"
)

func main() {
	fmt.Println("starting server")
	server := app.GetServer()

	err := server.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}

}
