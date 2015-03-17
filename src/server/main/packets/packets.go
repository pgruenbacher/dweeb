package packets

import (
	"encoding/json"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type RequestPacket struct {
	Req  *http.Request
	Res  http.ResponseWriter
	Code int
	Data interface{}
	Done chan bool
}

// Immediately pops the request with error response
func (p *RequestPacket) Error(code int, msg string) {
	log.Error("%v:%v", code, msg)
	p.Res.WriteHeader(code)
	js, _ := json.Marshal(Error{Code: code, Msg: msg})
	p.Res.Write(js)
	p.Done <- true
}

type GetRequestPacket struct {
	*RequestPacket
	Id   string
	Name string
}

type PostRequestPacket struct {
	*RequestPacket
	Content *url.Values
	Name    string
}

type Error struct {
	Code int
	Msg  string
}

func newFormValues() url.Values {
	values := url.Values{}
	values.Set("key1", "value1")
	values.Add("array1", "value1")
	values.Add("array1", "value2")
	values.Add("name", "sample name")
	values.Add("id", "55077361dc3a6f4c09000001")
	return values
}

func NewRequestPacket() *RequestPacket {
	urlPath := url.URL{
		Path: "/",
	}
	request := http.Request{
		Method: "GET",
		URL:    &urlPath,
		Form:   newFormValues(),
	}
	writer := httptest.NewRecorder()
	return &RequestPacket{
		Code: 200,
		Res:  writer,
		Req:  &request,
		Done: make(chan bool),
	}

}

func NewPostRequestPacket() *PostRequestPacket {
	requestPacket := NewRequestPacket()
	requestPacket.Req.Method = "POST"
	return &PostRequestPacket{
		RequestPacket: requestPacket,
		Content:       &requestPacket.Req.Form,
		Name:          "sample packet",
	}
}

func NewGetRequestPacket() *GetRequestPacket {
	requestPacket := NewRequestPacket()
	requestPacket.Req.Form.Add("id", "asdlfkj314")
	return &GetRequestPacket{
		RequestPacket: requestPacket,
		Id:            requestPacket.Req.Form.Get("id"),
		Name:          "sample packet",
	}
}
