package main

import (
	"fmt"
	"qurush/database"
)

func main() {
	database.ConnectionDB()
	fmt.Println("Hello")
}
