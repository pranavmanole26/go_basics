// package main

// import (
// 	"fmt"
// 	"path"
// )

// func main() {
// 	dir, file := path.Split("go_basics/enum_basics/enum_basics.go")

// 	fmt.Printf("dir: %s, file: %s\n", dir, file)
// 	fmt.Printf("file addr: %d\n", &file)
// 	code, dir, file := "abc", "xyz", "pqr"

// 	fmt.Printf("code: %s, file: %s\n", code, file)
// 	fmt.Printf("file addr: %d\n", &file)
// }

package main

import (
	"fmt"
	"os"
)

type user struct {
	username string
	password string
}

var users = []user{
	{username: "shobha", password: "123"},
	{username: "pranav", password: "456"},
}

func main() {

	if len(os.Args) == 3 {
		un := os.Args[1]
		ps := os.Args[2]
		// fmt.Println(users)
		loggedIn := false
		for _, u := range users {
			if un == u.username && ps == u.password {
				fmt.Printf("Welcome %s", un)
				loggedIn = true
				break
			}
		}
		if !loggedIn {
			fmt.Println("Wrong credentials")
		}
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments.")
	} else {
		fmt.Println("Too few arguments.")
	}

}
