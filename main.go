package main

import (
	 "github.com/leaanthony/mewn"
	 "github.com/wailsapp/wails"
	
)
 
func main() {

 	js := mewn.String("./frontend/dist/my-app/main.js")
	css := mewn.String("./frontend/dist/my-app/styles.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "angular-wails-demo",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(getProcessList)
	app.Bind(getOSinfo)
	app.Run()
 

}
