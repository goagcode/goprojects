package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opeing JSON file:", err)
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("Erro decoding JSON: ", err)
		return
	}
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filenam)
	if err != nil {
		fmt.Println("Error opening JSON file: ", errr)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data: ", err)
	}
	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(post)
	}
}
