package main

import "fmt"

type medical struct {
}

func (m *medical) execute(/* ... */) {
    if /* medicine is done */ {
        fmt.Println("Medicine already given to patient")
        /* ... */
    }
    fmt.Println("Medical giving medicine to patient")
    /* ... */
}