package main

import "fmt"

func main() {

	var choice int

	fmt.Print("============ USER DIRECTORY ============")

	fmt.Print("\n1. Add new users. ")

	fmt.Print("\n2. Update a user Email. ")

	fmt.Print("\n3. Delete a user. ")

	fmt.Print("\n4. Print all users. ")

	fmt.Print("\n5. Choose option: ")
	fmt.Scan(&choice)

	if choice == 5 {
		fmt.Print("BYE ")

	}

}
