package database

import (
	"errors"
	"github.com/pgruenbacher/dweeb/src/server/main/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"os"
	"time"
)

type Mongo struct {
	// state
	sess *mgo.Session
	uri  string
}

type Doc struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"msg"`
	Count   int           `bson:"count"`
	Content *url.Values   `bson:"content"`
	Created time.Time     `bson:"created"`
	Updated time.Time     `bson:"updated"`
}

const genericsTable = "generics"
const mongoURI = "mongodb://localhost"

// Need to initialize the session
func (m *Mongo) Init() (err error) {
	m.uri = os.Getenv("MONGOHQ_URL")
	if m.uri == "" {
		m.uri = mongoURI
	}

	m.sess, err = mgo.Dial(m.uri)
	if err != nil {
		log.Error("%v", err)
		os.Exit(1)
	}

	m.sess.SetSafe(&mgo.Safe{})
	return err
}

func setObjectId(id string) (bson.ObjectId, error) {
	if bson.IsObjectIdHex(id) {
		return bson.ObjectIdHex(id), nil
	} else {
		return "", errors.New("invalid object id")
	}
}

func (m *Mongo) Read(id string) (first Doc, err error) {
	objectId, _ := setObjectId(id)
	err = m.sess.DB("dweeb").C("generics").Find(bson.M{"_id": objectId}).One(&first)
	return first, err
}

func (m *Mongo) Create(doc *Doc) (err error) {
	doc.Created = time.Now()
	doc.Id = bson.NewObjectId()
	err = m.sess.DB("dweeb").C("generics").Insert(doc)
	return
}

func (m *Mongo) Update(id string, doc *Doc) (err error) {
	doc.Updated = time.Now()
	objectId, _ := setObjectId(id)
	err = m.sess.DB("dweeb").C("generics").Update(bson.M{"_id": objectId}, doc)
	return
}

func (m *Mongo) Destroy(id string) (err error) {
	objectId, _ := setObjectId(id)
	err = m.sess.DB("dweeb").C("generics").Remove(bson.M{"_id": objectId})
	return
}

func main() {

}
