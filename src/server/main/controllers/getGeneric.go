package controllers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
)

// Default controller for message retreival
type GetGeneric struct {
	flow.Component
	In  <-chan *packets.RequestPacket
	Out chan<- *packets.GetRequestPacket
}

// Simple controller routing
func (c *GetGeneric) OnIn(p *packets.RequestPacket) {
	err := p.Req.ParseForm()
	if err != nil {
		log.Error("%v", err)
	}
	c.Out <- &packets.GetRequestPacket{
		RequestPacket: p,
		Id:            p.Req.FormValue("id"),
		Name:          p.Req.FormValue("name"),
	}
}

func (r *GetGeneric) OnInClose() {
}
