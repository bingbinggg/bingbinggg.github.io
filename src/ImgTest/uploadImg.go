package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

const UPLOAD_PATH string = "C:/Users/benben/Desktop/"

type Img struct {
	Id     bson.ObjectId `bson:"_id"`
	ImgUrl string        `bson:"imgUrl"`
}

func main() {
	http.HandleFunc("/entrance", Entrance)
	http.HandleFunc("/uploadImg", UploadImg)
	http.ListenAndServe(":8000", nil)
}

func Entrance(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("uploadImg.html")
	t.Execute(w, nil)
}

func UploadImg(w http.ResponseWriter, r *http.Request) {
	var img Img
	img.Id = bson.NewObjectId()

	r.ParseMultipartForm(1024)
	imgFile, imgHead, imgErr := r.FormFile("img")
	if imgErr != nil {
		fmt.Println(imgErr)
		return
	}
	defer imgFile.Close()

	imgFormat := strings.Split(imgHead.Filename, ".")
	img.ImgUrl = img.Id.Hex() + "." + imgFormat[len(imgFormat)-1]

	image, err := os.Create(UPLOAD_PATH + img.ImgUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer image.Close()

	_, err = io.Copy(image, imgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err = session.DB("images").C("image").Insert(img)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success to upload img")
}
