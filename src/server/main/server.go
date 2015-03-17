package main

import (
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
)

var (
	routerIn chan *packets.RequestPacket
)

func handler(w http.ResponseWriter, r *http.Request) {
	rp := packets.RequestPacket{
		Req:  r,
		Res:  w,
		Done: make(chan bool),
	}
	routerIn <- &rp
	<-rp.Done
}

func main() {
	// Create application net
	routerIn = make(chan *packets.RequestPacket)
	a := NewApp()
	a.SetInPort("AppInput", routerIn)
	// Run
	flow.RunNet(a)
	// Serve
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9090", nil)
}
