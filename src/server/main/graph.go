package main

import (
	"github.com/pgruenbacher/dweeb/src/server/main/controllers"
	"github.com/pgruenbacher/dweeb/src/server/main/routers"
	"github.com/pgruenbacher/dweeb/src/server/main/stores"
	"github.com/pgruenbacher/dweeb/src/server/main/writers"
	"github.com/pgruenbacher/goflow"
)

// A graph for our app
type App struct {
	flow.Graph

	InitTestFlag int
	FinTestFlag  chan bool
}

// A constructor that creates network structure
func NewApp() *App {
	// Create a new graph
	net := new(App)
	net.InitGraphState()
	// Add graph nodes
	net.Add(new(routers.Router), "router")
	net.Add(new(controllers.GetGeneric), "getGenericController")
	net.Add(new(controllers.PostGeneric), "postGenericController")
	net.Add(stores.NewStorage(), "storage")
	net.Add(new(writers.Responder), "responder")

	// Connect the processes
	net.Connect("router", "PostGeneric", "postGenericController", "In")
	net.Connect("router", "GetGeneric", "getGenericController", "In")
	net.Connect("getGenericController", "Out", "storage", "Get")
	net.Connect("postGenericController", "Out", "storage", "Post")
	net.Connect("storage", "Out", "responder", "In")
	// Network ports
	net.MapInPort("AppInput", "router", "In")
	return net
}

// Test for a network initializer
func (n *App) Init() {
	n.InitTestFlag = 123
}

// Test for a network finalizer
func (n *App) Finish() {
	n.InitTestFlag = 456
	n.FinTestFlag <- true
}
