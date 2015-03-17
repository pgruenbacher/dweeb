package packets

import (
	"testing"
)

func TestPackets(t *testing.T) {
	requestPacket := NewRequestPacket()
	requestPacket.Req.Method = "POST"

	postRequestPacket := NewPostRequestPacket()
	getRequestPacket := NewGetRequestPacket()
	if postRequestPacket.Req.Method != "POST" {
		t.Error("fail post request packet")
	}
	if len(getRequestPacket.Id) == 0 {
		t.Error("fail get request packet")
	}
}
