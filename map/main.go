package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[int]string)

	//delete(colors, 10)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bff34",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}
