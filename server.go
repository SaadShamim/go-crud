package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: ":8081",
	}

	http.HandleFunc("/post/", handleRequest)
	fmt.Println("listening to server")
	log.Fatal(server.ListenAndServe())
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	post, err := retrieve(id)

	if err != nil {
		return
	}

	output, err := json.Marshal(&post)

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	post := Post{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&post)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = post.create()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	post := Post{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&post)

	if err != nil {
		panic(err)
	}

	err = post.update()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))

	if err != nil {
		return
	}

	if err = deletePost(id); err != nil {
		return
	}

	w.WriteHeader(200)
	return
}
