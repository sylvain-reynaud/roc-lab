package main

import "fmt"

type doctor struct {
}

func (d *doctor) execute(/* ... */) {
    if /* doctorCheckUp is done */ {
        fmt.Println("Doctor checkup already done")
        /* ... */
    }
    fmt.Println("Doctor checking patient")
    /* ... */
}