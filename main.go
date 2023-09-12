package main

import (
	"acm-payments/acm"
	"flag"
	"fmt"
)



func main() {
	flag.Parse()
	result := flag.Arg(0)
	amcC := acm.NewAcmCompany()
	data := amcC.LoadFile(result)
	doPayments := amcC.ProcessData(data)

	for key,val := range doPayments {
		fmt.Printf("Name %s total to earn %f$\n", key,val)
	}
	
}