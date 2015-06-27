package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

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
	Status      string  `json:status`
	NextPage    int     `json:nextPage`
	TotalPhotos int     `json:totalPhotos`
	TotalPages  int     `json:totalPages`
	Photos      []photo `json:photos`
}

type photo struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Url      string        `json:url`
	ThumbUrl string        `json:thumbUrl`
	Title    string        `json:title`
}

type vote struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	PhotoId string        `json:photoId`
	UserId  string        `json:userId`
	Value   int           `json:value`
}

type votesResponse struct {
	Status string `json:status`
	Vote   vote   `json:vote`
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

			var currentPage int
			var err error
			perPagePhotos := int(51)
			//get current page
			requestPage := r.FormValue("current_page")

			if len(requestPage) > 0 {
				currentPage, err = strconv.Atoi(requestPage)

				if err != nil {
					panic(err)
				}
			} else {
				currentPage = 1
			}

			c := database.DB("explorePhotos").C("photos")

			//get the total number of photos
			totalPhotos, err := c.Count()

			if err != nil {
				panic(err)
			}

			// fmt.Println(currentPage) //r

			// fmt.Println(reflect.TypeOf(totalPhotos).Kind()) //r

			totalNumberOfPhotos := float64(totalPhotos)

			perPagePhotosC := float64(perPagePhotos)
			totalNumberOfPages := math.Ceil(totalNumberOfPhotos / perPagePhotosC)

			//if given page is more then total number of pages
			if currentPage > int(totalNumberOfPages) {

				returnError := errorMessage{"unsuccess", "page doesn't exit"}
				js, err := json.Marshal(returnError)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			} else {
				// fmt.Println(totalNumberOfPages) //r

				//calculate offset
				offset := (currentPage - 1) * perPagePhotos

				result := []photo{}
				err = c.Find(bson.M{}).Skip(offset).Limit(51).All(&result)
				if err != nil {
					panic(err)
				}

				var nextPage int
				if currentPage+1 <= int(totalNumberOfPages) {

					nextPage = currentPage + 1
				} else {
					nextPage = -1
				}

				response := photosResponse{"success", nextPage, int(totalNumberOfPhotos), int(totalNumberOfPages), result}

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

func votesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		{

			photoId := r.PostFormValue("photo_id")
			userId := r.PostFormValue("user_id")
			value := r.PostFormValue("value")

			objectId := bson.NewObjectId()

			n, errr := strconv.Atoi(value)

			if errr != nil {
				panic(errr)
			}

			newVotes := vote{objectId, photoId, userId, n}

			// fmt.Println(newVotes)

			//check if user is alreayd exist
			uv := database.DB("explorePhotos").C("votes")
			result := vote{}
			err := uv.Find(bson.M{"photoId": photoId, "userId": userId}).One(&result)
			if err != nil {
				fmt.Println(err)
			}

			if result == (vote{}) { //register user here

				err = uv.Insert(&vote{Id: objectId, PhotoId: photoId, UserId: userId, Value: n})
				if err != nil {
					panic(err)
				}
				//update the count in photos doc

				returnMessage := votesResponse{"success", newVotes}
				js, err := json.Marshal(returnMessage)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)

			} else { //send message user is already exist
				returnError := errorMessage{"unsuccess", "user already voted"}
				js, err := json.Marshal(returnError)
				if err != nil {
					panic(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			}

		}

	case "DELETE":
		{ // Remove the record.

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
	http.HandleFunc("/votes", votesHandler)
	http.ListenAndServe(":1334", nil)

}
