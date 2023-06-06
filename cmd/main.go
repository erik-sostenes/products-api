package main

import "github.com/erik-sostenes/products-api/cmd/bootstrap"

func main() {
	if err := bootstrap.NewInjector(); err != nil {
		panic(err)
	}
}
