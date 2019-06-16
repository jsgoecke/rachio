package main

import (
	rachio ".."
	"fmt"
	"os"
)

func main() {
	client, err := rachio.NewClient(os.Getenv("RACHIO_API_KEY"))
	if err != nil {
		panic(err)
	}

	person, err := client.GetPerson()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(person)
}
