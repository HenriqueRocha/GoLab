// Points in a F1 race; first version

package main

import "fmt"

func main() {
	var place int
	fmt.Printf("Which place? ")
	fmt.Scanf("%d", &place)
	if place == 1 {
		fmt.Printf("Won 10 points.\n") 
	} else if place == 2 {
		fmt.Printf("Won 8 points.\n")
	} else if place == 3 {
		fmt.Printf("Won 6 points.\n")
	} else if place == 4 {
		fmt.Printf("Won 5 points.\n")
	} else if place == 5 {
		fmt.Printf("Won 4 points.\n")
	} else if place == 6 {
		fmt.Printf("Won 3 points.\n")
	} else if place == 7 {
		fmt.Printf("Won 2 points.\n")
	} else if place == 8 {
		fmt.Printf("Won 1 point.\n")
	} else {
		fmt.Printf("Didn't win points.\n") 
	}
}

