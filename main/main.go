package main

import (
	"fmt"

	"github.com/transfashion/goedcmeg"
)

func main() {
	cfg := goedcmeg.Config{}

	edc := goedcmeg.NewEdc(cfg)

	res, err := edc.Sale(goedcmeg.SaleTransaction{})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
