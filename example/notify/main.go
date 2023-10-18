package main

import (
	"context"
	"fmt"
	"os"

	"github.com/starry-axul/notifit-go-sdk/notify"
)

func main() {
	trans := notify.NewHttpClient("http://ec2-54-226-191-87.compute-1.amazonaws.com:8100", "")

	if err := trans.Push(context.Background(), "Mi message", "test 123", ""); err != nil {
		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("ok")

}
