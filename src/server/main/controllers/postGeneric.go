package controllers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
)

// Default controller for message retreival
type PostGeneric struct {
	flow.Component
	In  <-chan *packets.RequestPacket
	Out chan<- *packets.PostRequestPacket
}

// Simple controller routing
func (c *PostGeneric) OnIn(p *packets.RequestPacket) {
	err := p.Req.ParseForm()
	if err != nil {
		log.Error("%v", err)
	}

	c.Out <- &packets.PostRequestPacket{
		Name:          p.Req.FormValue("name"),
		Content:       &p.Req.Form,
		RequestPacket: p,
	}
}

func (r *PostGeneric) OnInClose() {
}
