package main

import (
	"fmt"

	"github.com/zhong-my/modlib"
	"github.com/zhong-my/modlib/internal/auth"
	"github.com/zhong-my/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", modlib.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
