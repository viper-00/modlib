package main

import (
	"fmt"

	"github.com/zhong-my/modlib"
	"github.com/zhong-my/modlib/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", modlib.Config())
	fmt.Println(clientlib.Hello())
}
