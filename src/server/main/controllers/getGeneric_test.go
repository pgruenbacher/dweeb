package controllers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"testing"
)

func TestGetGeneric(t *testing.T) {
	request := packets.NewRequestPacket()

	d := new(GetGeneric)
	in := make(chan *packets.RequestPacket, 10)
	out := make(chan *packets.GetRequestPacket, 10)
	d.In = in
	d.Out = out
	flow.RunProc(d)
	for i := 0; i < 10; i++ {
		in <- request
		<-out
	}
	// Shutdown the component
	close(in)
	// output:
	// passed
}
