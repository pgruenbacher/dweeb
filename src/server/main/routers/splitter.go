package routers

import (
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"testing"
)

func TestFirst(t *testing.T) {

	postPacket := packets.NewRequestPacket()
	postPacket.Req.URL.Path = "/generic"
	postPacket.Req.Method = "POST"
	getPacket := packets.NewRequestPacket()
	getPacket.Req.URL.Path = "/generic"

	d := new(Router)
	postGeneric := make(chan *packets.RequestPacket, 10)
	getGeneric := make(chan *packets.RequestPacket, 10)
	in := make(chan *packets.RequestPacket, 10)
	d.In = in
	d.GetGeneric = getGeneric
	d.PostGeneric = postGeneric
	flow.RunProc(d)

	for i := 0; i < 10; i++ {
		in <- postPacket
		in <- getPacket
		for i := 0; i < 1; i++ {
			i2 := <-getGeneric
			i3 := <-postGeneric
			if i2.Req.Method != "GET" || i3.Req.Method != "POST" {
				t.Error("failed router")
			}
		}
	}

	close(in)
}
