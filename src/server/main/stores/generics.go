package stores

import (
	"github.com/pgruenbacher/dweeb/src/server/main/database"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"github.com/pgruenbacher/dweeb/src/server/main/packets"
	"github.com/pgruenbacher/goflow"
	"net/http"
)

// Provides data storage and retrieval services
type Storage struct {
	flow.Component
	Post <-chan *packets.PostRequestPacket
	Get  <-chan *packets.GetRequestPacket
	Out  chan<- *packets.RequestPacket
}

var Mongo database.Mongo

// Constructs a new store server
func NewStorage() *Storage {
	Mongo.Init()
	s := new(Storage)
	s.Component.Mode = flow.ComponentModeSync
	return s
}

func (r *Storage) Init() {
	// Mongo
}

// Gets messages from the store and sends the to out
func (s *Storage) OnGet(p *packets.GetRequestPacket) {
	var err error
	var data database.Doc
	data, err = Mongo.Read(p.Id)

	if err != nil {
		p.Error(http.StatusInternalServerError, "could not read data")
	}
	p.Data = data
	p.Code = http.StatusFound
	s.Out <- p.RequestPacket
}

// Adds a new message to the store and sends status to out
func (s *Storage) OnPost(p *packets.PostRequestPacket) {
	doc := database.Doc{
		Name:    p.Name,
		Content: p.Content,
	}
	err := Mongo.Create(&doc)
	if err != nil {
		p.Error(http.StatusInternalServerError, "Could not save data")
	}
	p.Data = doc
	p.Code = http.StatusCreated
	s.Out <- p.RequestPacket
}

func (r *Storage) OnGetClose() {
}

func (r *Storage) OnPostClose() {
}

func (r *Storage) Finish() {
	log.Info("generic store closed")
}

// func (r *Storage) Finish() {
// 	// Your finalization code here
// 	fmt.Println("finishable")
// }

// func (r *Storage) Shutdown() {
//     // Close ports yourself when necessary
// }
