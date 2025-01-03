package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#ffffff",
		"white": "#000000",
		"green": "#4fg567",
	}

	var otherColors map[string]string

	fmt.Println(colors)
	fmt.Println(otherColors) // Initialized with zero-values

	nonZeroValues := make(map[string]string) // It's empty, not zero-values

	fmt.Println(nonZeroValues)

	delete(colors, "red") // Deleteing an entry
	fmt.Println(colors)

	deleteColor(colors) // Works because is reference type, not value type and we access directly to the map and not to a copy

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func deleteColor(c map[string]string) {
	delete(c, "black")
}
