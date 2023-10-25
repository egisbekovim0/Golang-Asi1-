package main

import (
	"fmt"
	"sync"
)

type crocodile struct {
	count int
}

var saltwaterInstance *crocodile
var nileInstance *crocodile
var once sync.Once

func GetSaltwaterCrocodile() *crocodile {
	once.Do(func() {
		saltwaterInstance = &crocodile{}
	})
	return saltwaterInstance
}

func GetNileCrocodile() *crocodile {
	once.Do(func() {
		nileInstance = &crocodile{}
	})
	return nileInstance
}

func (c *crocodile) AddOne() int {
	c.count++
	return c.count
}

func main() {
	saltwaterCrocodile1 := GetSaltwaterCrocodile()
	saltwaterCrocodile2 := GetSaltwaterCrocodile()

	nileCrocodile := GetNileCrocodile()

	saltwaterCrocodile1.AddOne()
	saltwaterCrocodile2.AddOne()
	nileCrocodile.AddOne()

	fmt.Printf("Saltwater Crocodile 1 count: %d\n", saltwaterCrocodile1.count)
	fmt.Printf("Saltwater Crocodile 2 count: %d\n", saltwaterCrocodile2.count)
	fmt.Printf("Nile Crocodile count: %d\n", nileCrocodile.count)
}
