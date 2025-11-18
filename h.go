package main

import "fmt"

func main() {


    users := make(map[string]string)

    for {

        var choice int

        fmt.Print("\n============ USER DIRECTORY ============")
        fmt.Print("\n1. Add new users.")
        fmt.Print("\n2. Update a user Email.")
        fmt.Print("\n3. Delete a user.")
        fmt.Print("\n4. Print all users.")
        fmt.Print("\n5. Exit")
        fmt.Print("\nChoose option: ")
        fmt.Scan(&choice)

        if choice == 5 {

            fmt.Print("bye ")

            break

        }

        switch choice {
			
        case 1:

            var name, email string

            fmt.Print("Enter user name: ")
            fmt.Scan(&name)

            fmt.Print("Enter user email: ")
            fmt.Scan(&email)

            users[name] = email

            fmt.Printf("User %s added successfully!\n", name)

        case 2:
            var name, newemail string

            fmt.Print("Enter Users name: ")
            fmt.Scan(&name)

            fmt.Print("Enter new email: ")
            fmt.Scan(&newemail)

            users[name] = newemail

            fmt.Printf("User %s updated successfully!\n", name) 

        case 3:
            var name string

            fmt.Print("Enter users name: ")
            fmt.Scan(&name)

            delete(users, name)

            fmt.Printf("User %s deleted successfully!\n", name) 

        case 4:

            fmt.Println("\n====== All USERS =====")

            for name, email := range users {

                fmt.Printf("User name: %s | Email: %s\n", name, email)

            }

        default: 

            fmt.Print("Invalid choice please select between 1-5 ")

        }
    }
}
