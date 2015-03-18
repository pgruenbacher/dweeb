package routers

import (
	// "github.com/bmizerany/pat"
	// "github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
	// "strings"
)

// Simple table-based request router
type Router struct {
	flow.Component
	Init     chan bool
	Generics chan *packets.RequestPacket

	// State
}

func route(c chan *packets.RequestPacket) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		packet := makePacket(w, r)
		c <- packet
		<-packet.Done
	}
}

func (r *Router) OnInit() {

	http.HandleFunc("/generics", route(r.Generics))
	http.ListenAndServe("localhost:9090", nil)
}

func makePacket(w http.ResponseWriter, r *http.Request) *packets.RequestPacket {
	rp := packets.RequestPacket{
		Req:  r,
		Res:  w,
		Done: make(chan bool),
	}
	return &rp
}
