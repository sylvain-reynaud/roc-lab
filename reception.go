package main

import "fmt"

type reception struct {
}

func (r *reception) execute(/* ... */) {
    if /* registration is done */ {
        fmt.Println("Patient registration already done")
        /* ... */
    }
    fmt.Println("Reception registering patient")
		/* ... */
}