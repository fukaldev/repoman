/*
Copyright © 2024 Furkan Kalabalik furkankalabalikk@gmail.com
*/
package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/repoman/cmd"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	cmd.Execute()
}
