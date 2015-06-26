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

type userResponse struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Username string        `json:username`
	Password string        `json:password`
}

type errorMessage struct {
	Status  string `json:status`
	Message string `json:message`
}

type successMessage struct {
	Status string `json:status`
	User   userResponse
}

type photosResponse struct {
	Status string  `json:status`
	Photos []photo `json:photos`
}

type photo struct {
	ID  bson.ObjectId `bson:"_id,omitempty"`
	Url string        `json:url`
}

var database *mgo.Session

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		result := userResponse{}
		uc := database.DB("explorePhotos").C("users")
		err := uc.Find(bson.M{"username": username, "password": password}).One(&result)
		if err != nil {
			fmt.Println(err)
		}

		if result != (userResponse{}) { //register user here

			existingUser := userResponse{result.Id, username, password}

			returnMessage := successMessage{"success", existingUser}
			js, err := json.Marshal(returnMessage)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		} else { //invalid user name and password
			returnError := errorMessage{"unsuccess", "user doesn't exist"}
			js, err := json.Marshal(returnError)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}

	} else {
		returnError := errorMessage{"unsuccess", "method not defind"}
		js, err := json.Marshal(returnError)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		{
			//later will change this to id
			username := r.FormValue("username")

			c := database.DB("explorePhotos").C("users")
			result := userResponse{}
			err := c.Find(bson.M{"username": username}).One(&result)
			if err != nil {
				returnError := errorMessage{"unsuccess", "no user found"}
				js, err := json.Marshal(returnError)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			} else {
				response := successMessage{"success", result}

				js, err := json.Marshal(response)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			}
		}

	case "POST":
		{

			username := r.PostFormValue("username")
			password := r.PostFormValue("password")
			objectId := bson.NewObjectId()
			newUser := userResponse{objectId, username, password}

			//check if user is alreayd exist
			uc := database.DB("explorePhotos").C("users")
			result := user{}
			err := uc.Find(bson.M{"username": username, "password": password}).One(&result)
			if err != nil {
				fmt.Println(err)
			}

			if result == (user{}) { //register user here

				err = uc.Insert(&userResponse{Id: objectId, Username: username, Password: password})
				if err != nil {
					panic(err)
				}
				returnMessage := successMessage{"success", newUser}
				js, err := json.Marshal(returnMessage)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)

			} else { //send message user is already exist
				returnError := errorMessage{"unsuccess", "user already registerd"}
				js, err := json.Marshal(returnError)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			}

		}

	default:
		{
			returnError := errorMessage{"unsuccess", "method not defind"}
			js, err := json.Marshal(returnError)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}
	}

}

func photosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		{

			c := database.DB("explorePhotos").C("photos")
			result := []photo{}
			err := c.Find(bson.M{}).All(&result)
			if err != nil {
				panic(err)
			}

			response := photosResponse{"success", result}

			js, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
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
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	database = session

	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/photos", photosHandler)
	http.ListenAndServe(":1334", nil)

}
