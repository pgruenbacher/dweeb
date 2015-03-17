package writers

import (
	"encoding/json"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
)

// Simple JSON response generator

type Responder struct {
	flow.Component
	In <-chan *packets.RequestPacket
}

// Processes a request packet and sends the response JSON
func (r *Responder) OnIn(p *packets.RequestPacket) {
	js, err := json.Marshal(p.Data)
	if err != nil {
		log.Error("%v", err)
		p.Error(http.StatusInternalServerError, "Could not marshal JSON")
		return
	}
	_, err = p.Res.Write(js)
	if err != nil {
		log.Error("%v", err)
	}
	p.Done <- true
}

func (r *Responder) OnInClose() {
}
