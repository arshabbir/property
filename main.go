package main

import (
	"log"

	"github.com/arshabbir/propertymod/app"
)

func main() {
	pApp := app.NewPropertyApp()

	if err := pApp.StartApp(); err != nil {
		log.Println("error while starting the propery app..", err.Error())
	}
}
