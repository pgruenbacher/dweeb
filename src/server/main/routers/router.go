package routers

import (
	"github.com/bmizerany/pat"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
	// "strings"
)

// Simple table-based request router
type Router struct {
	flow.Component
	Input    <-chan int
	Generics chan<- *packets.RequestPacket

	// State
}

func (r *Router) OnInput(i int) {
	m := pat.New()
	m.Get("/generics/:name", http.HandlerFunc(route(r.Generics)))
	m.Post("/generics/:name", http.HandlerFunc(route(r.Generics)))

	// http.HandleFunc("/generics", route(r.Generics))
	http.Handle("/", m)
	err := http.ListenAndServe("localhost:9090", nil)
	log.Error("%v", err)
}

func (r *Router) OnInputClose() {

}

func route(c chan<- *packets.RequestPacket) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("%v %s", r.Method, r.URL.Path)
		packet := makePacket(w, r)
		c <- packet
		<-packet.Done
	}
}

func makePacket(w http.ResponseWriter, r *http.Request) *packets.RequestPacket {
	rp := packets.RequestPacket{
		Req:  r,
		Res:  w,
		Done: make(chan bool),
	}
	return &rp
}
