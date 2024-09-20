package main

import "github.com/deathstarset/backend-docu-quest/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
