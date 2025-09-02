package main

import "are-you-die/database"

func main() {
	err := database.Open()
	if err != nil {
		panic(err)
	}
}
