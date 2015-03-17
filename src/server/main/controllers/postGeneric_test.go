package controllers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"testing"
)

func TestFirst(t *testing.T) {
	request := packets.NewRequestPacket()
	request.Req.Method = "POST"
	d := new(PostGeneric)
	in := make(chan *packets.RequestPacket, 10)
	out := make(chan *packets.PostRequestPacket, 10)
	d.In = in
	d.Out = out
	flow.RunProc(d)
	for i := 0; i < 10; i++ {
		in <- request
		i2 := <-out
		if i2.Req.Method != "POST" {
			t.Error("postGeneric failed")
		}
	}

	close(in)
}
