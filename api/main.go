package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {

	//create connection with mongo db
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("explorePhotos").C("photos")
	count, err := c.Count()
	fmt.Println(count)

}
