package database

import (
	"fmt"
	"net/url"
)

var db Mongo

func init() {
	db.Init()
}

func Example() {
	data2 := make(map[string][]string)
	data2["key1"] = []string{"a", "b", "c"}
	data2["key2"] = []string{"alone"}
	data := url.Values(data2)
	doc := Doc{
		Content: &data,
		Name:    "name1",
	}

	err := db.Create(&doc)

	doc.Name = "name2"
	err = db.Update(doc.Id.Hex(), &doc)
	if err != nil {
		fmt.Println(err)
	}

	doc, err = db.Read(doc.Id.Hex())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(doc.Name)

	err = db.Destroy(doc.Id.Hex())
	if err != nil {
		fmt.Println(err)
	}
	// output:
	// name2
}
