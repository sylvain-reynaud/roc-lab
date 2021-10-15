package main

import "fmt"

type cashier struct {
}

func (c *cashier) execute(/* ... */) {
    if /* payment is done */ {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
		/* ... */
}