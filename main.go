package main

import (
	"dimp/cmd"
	"dimp/pkg/app"
	"log"
)

func main() {
	appContext := &app.AppContext{}
	appContext.AddCmd(cmd.EstatesCmd)
	app, err := appContext.Build()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
