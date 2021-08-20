package utils


import (

	"fmt"
	_"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"
)

type book struct {
	ISBN   string  `json:"ISBN" bson:"ISBN"`
	Title  string  `json:"Title" bson:"Title"`
	Author string  `json:"Author" bson:"Author"`
	Price  float64 `json:"Price" bson:"Price"`
	Rating  float64 `json:"Rating" bson:"Rating"`

}
type books []book

type Response struct {
	Books []book `json:"book"`
}


func connect() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("session err:", err)
		return nil
	}
	return session
}
func GetRating(ISBN string)float64{
	var books book
	session := connect()
	defer session.Close()
	err := session.DB("bookinfo").C("books").Find(bson.M{"ISBN":ISBN}).One(&books)
	if err != nil {
		log.Println(err)
		return 0

	}

		rati:=books.Rating
		return rati
}
