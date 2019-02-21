package main

import "fmt"

type Pizza struct {
	Toppings []string
	Origin   string
	Size     string
	Feeds    int
	Price    float32
}

func bake(toppings []string, size string, origin string) Pizza {
	var price float32
	var feeds int
	if size == "small" {
		price = 9.99
		feeds = 2
	} else if size == "medium" {
		price = 14.99
		feeds = 4
	} else if size == "large" {
		price = 19.99
		feeds = 6
	}
	price += float32(len(toppings) * 2)

	newPizza := Pizza{Toppings: toppings, Origin: origin, Size: size, Feeds: feeds, Price: price}
	return newPizza
}

func main() {
	size := "large"
	origin := "Pizza Hut"
	toppings := []string{"Mushrooms", "Ham", "Drugs"}
	newPizza := bake(toppings, size, origin)
	fmt.Println(newPizza)
}
