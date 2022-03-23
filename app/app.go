package app

import (
	"log"
	"time"

	"github.com/LitusLuca/Taco/window"
)

var a *Application

type Application struct {
	window.IWindow
	running bool
	lastframeTime float32
}

func App() *Application {
	if a != nil {
		return a
	}
	a = new(Application)
	props := window.WindowProps{Title: "TACO ENGINE !!!", Width: 500, Height: 500}
	err := window.Init(props)
	if err != nil {
		panic(err)
	}
	a.IWindow = window.Get()
	//TODO: init renderer
	//TODO: init ImGui?

	return a
}

func (a *Application) Run() {
	a.running = true
	a.lastframeTime = 0.0
	for a.running {
		//calculate frameTime and dT
		currentTime := a.GetTime()
		dt := currentTime - a.lastframeTime
		a.lastframeTime = currentTime

		//update layer
		a.IWindow.OnUpdate()
		log.Printf("Update: %v", dt)
		time.Sleep(1 * time.Second)
	}
}

//TODO OnEvent and OnWindowCloseEvent
