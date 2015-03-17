package main

import (
	"fmt"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"testing"
)

// Create new Network for Test
func TestNetwork(t *testing.T) {
	packetInput := packets.NewRequestPacket()
	packetInput.Req.Method = "POST"
	gPI := packets.NewRequestPacket()

	n := NewApp()
	// Ports
	in := make(chan *packets.RequestPacket)

	n.SetInPort("AppInput", in)
	// Exported state
	n.FinTestFlag = make(chan bool)
	flow.RunNet(n)

	in <- packetInput
	<-packetInput.Done
	fmt.Println("posted")
	in <- gPI
	<-gPI.Done
	fmt.Println("getted")

	close(in)
	// Wait for finalization signal
	<-n.FinTestFlag

	fmt.Println("waiting for done")

	if n.InitTestFlag != 456 {
		t.Errorf("Finish: %d != %d", n.InitTestFlag, 456)
	}

	<-n.Wait()
}

// Initialization
