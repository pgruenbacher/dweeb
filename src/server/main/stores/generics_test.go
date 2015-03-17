package stores

import (
	"fmt"
	"github.com/pgruenbacher/dweeb/src/server/main/database"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
)

// var proc *Storage

func init() {
	// proc = NewStorage()
	// flow.Register("store", proc)
	// flow.Annotate("store", flow.ComponentInfo{
	// 	Description: "Doubles its input",
	// })
}

func Example() {
	// data := map[string]interface{}{
	// 	"more": "stuff",
	// }
	req := packets.RequestPacket{}
	req2 := packets.RequestPacket{}
	postRequest := packets.PostRequestPacket{
		RequestPacket: &req,
		Name:          "a name test",
	}
	getRequest := packets.GetRequestPacket{
		RequestPacket: &req2,
		Name:          "a asdf",
	}

	store := NewStorage()

	post := make(chan *packets.PostRequestPacket, 10)
	get := make(chan *packets.GetRequestPacket, 10)
	out := make(chan *packets.RequestPacket, 10)
	store.Post = post
	store.Get = get
	store.Out = out
	flow.RunProc(store)

	var tempId string
	for i := 0; i < 1; i++ {
		post <- &postRequest
		response, ok := <-out
		if !ok {
			fmt.Println("not ok")
		}
		tempId = response.Data.(database.Doc).Id.Hex()
		fmt.Println(response.Code)
		getRequest.Id = tempId
	}
	for i := 0; i < 1; i++ {
		get <- &getRequest
		response2, ok := <-out
		if !ok {
			fmt.Println("not ok")
		}
		fmt.Println(response2.Code)
		fmt.Println(response2.Data.(database.Doc).Name)
	}
	close(get)
	close(post)

	//output:
	// 201
	// 302
	// a name test
}
