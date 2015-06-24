package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	Username string `json:username`
	Password string `json:password`
}

type errorMessage struct {
	Message string `json:message`
	Status  string `json:status`
}

type photo struct {
	ID  bson.ObjectId `bson:"_id,omitempty"`
	Url string        `json:url`
}

func sel(q ...string) (r bson.M) {
	r = make(bson.M, len(q))
	for _, s := range q {
		r[s] = 1
	}
	return
}

var database *mgo.Session

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET": //this will use for login
		{

			loggedinUser := user{"gyani", "gyani123"}

			fmt.Println(loggedinUser.Password)
			js, err := json.Marshal(loggedinUser)

			if err != nil {
				panic(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}

	case "POST":
		{

			// userName:=r.PostFormValue('userName')
			// passWord:=r.PostFormValue('password')
			// newUser := user{userName,passWord}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			// Create a new record.

		}

	case "DELETE":
		{ // Remove the record.
			fmt.Println("inside delete method")
		}

	default:
		{
			fmt.Println("no method")
		}
	}

}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": //this will use for login
		{

			loggedinUser := user{"gyani", "gyani123"}

			fmt.Println(loggedinUser.Password)
			js, err := json.Marshal(loggedinUser)

			if err != nil {
				panic(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}

	case "POST":
		{

			// userName:=r.PostFormValue('userName')
			// passWord:=r.PostFormValue('password')
			// newUser := user{userName,passWord}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			// Create a new record.

		}

	case "DELETE":
		{ // Remove the record.
			fmt.Println("inside delete method")
		}

	default:
		{
			fmt.Println("no method")
		}
	}
	switch r.Method {
	case "GET": //this will use for login
		{
			fmt.Println("inside get method")
		}

	case "POST":
		{
			fmt.Println("inside post method")
		}

	case "DELETE":
		{ // Remove the record.
			fmt.Println("inside delete method")
		}

	default:
		{
			fmt.Println("no method defined")
		}
	}

}

func main() {

	//create connection with mongo db
	database, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer database.Close()

	//mongo quries
	c := database.DB("explorePhotos").C("photos")

	// Query One
	result := photo{}
	// err = c.Find(bson.M{"url": "https://c1.staticflickr.com/1/438/19053518062_7838ff75c1_k.jpg"}).Select(bson.M{"_id": 0}).One(&result)
	err = c.Find(bson.M{"url": "https://c1.staticflickr.com/1/438/19053518062_7838ff75c1_k.jpg"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Url)

	// Insert Datas
	uc := database.DB("explorePhotos").C("users")
	err = uc.Insert(&user{Username: "gyani", Password: "123"})

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":1334", nil)

}
