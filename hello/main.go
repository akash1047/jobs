package main

import (
	"fmt"
	"os"
)

func main() {
	if user, ok := os.LookupEnv("CUSTOM_USER"); ok {
		fmt.Printf("Hi, %s\n", user)
	} else {
		fmt.Println("Hi USER_NOT_SPECIFIED")
	}

}
