package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, val := range d {
		fmt.Println("key ", i, " = ", val)
	}
}
