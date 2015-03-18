package routers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
	"strings"
)

// Simple table-based request router
type Splitter struct {
	flow.Component
	In   <-chan *packets.RequestPacket
	Get  chan<- *packets.RequestPacket
	Post chan<- *packets.RequestPacket

	// State
	BasePath string
}

// Request handler
func (r *Splitter) OnIn(p *packets.RequestPacket) {

	path := p.Req.URL.Path

	if strings.Index(path, r.BasePath) == 0 {
		path = path[len(r.BasePath):]
	}

	var fwd chan<- *packets.RequestPacket
	ok := true

	if p.Req.Method == "POST" {
		fwd = r.Post
	} else if p.Req.Method == "GET" {
		fwd = r.Get
	}
	if !ok {
		p.Error(http.StatusBadRequest, "Unknown path: "+path)
	} else {
		fwd <- p
	}
}
