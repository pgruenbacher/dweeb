package writers

import (
	"encoding/json"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
	"strconv"
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

	log.Info("%v", p.Res.Header())
	for k, v := range p.Res.Header() {
		log.Info("%v, %v", k, v)
	}
	clen, _ := strconv.Atoi(p.Res.Header().Get("Content-Length"))
	clen += len(js)
	p.Res.Header().Set("Content-Length", strconv.Itoa(clen))
	p.Res.Header().Set("Content-Type", "application/json")

	_, err = p.Res.Write(js)
	if err != nil {
		log.Error("%v", err)
	}
	p.Done <- true
}

func (r *Responder) OnInClose() {
	log.Info("Responder closed")
}
