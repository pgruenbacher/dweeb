package writers

import (
	"fmt"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http/httptest"
	"time"
)

func Example() {
	writer := httptest.NewRecorder()
	data := map[string]interface{}{
		"more": "stuff",
	}
	request := packets.RequestPacket{
		Code: 404,
		Res:  writer,
		Data: data,
	}

	d := new(Responder)
	in := make(chan *packets.RequestPacket, 10)
	d.In = in
	flow.RunProc(d)
	for i := 0; i < 1; i++ {
		in <- &request

	}
	time.Sleep(1000)
	fmt.Println(request.Res.(*httptest.ResponseRecorder).Body)
	// Shutdown the component
	close(in)
	// output:
	// {"more":"stuff"}
}
