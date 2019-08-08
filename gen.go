package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	//  "json"
)

/* This defines a blog post */
type BlogPost struct {
	Title   string
	Author  string
	Date    string
	Content string
}

type PostPage struct {
	PageTitle string
	Posts     []BlogPost
}

/*
func catchJSON(w http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)

}
*/

func main() {
	// Create the mux to handle http requests
//	mux := http.NewServeMux()
	//	rh := http.RedirectHandler("")
	/* Begin user given data */
	reader := bufio.NewReader(os.Stdin)
	/* Ask the user what to put in */

	fmt.Print("Title: ")
	userTitle, _ := reader.ReadString('\n')

	fmt.Print("Author: ")
	userAuthor, _ := reader.ReadString('\n')

	fmt.Print("Date: ")
	userDate, _ := reader.ReadString('\n')

	fmt.Print("Content: ")
	userContent, _ := reader.ReadString('\n')

	/* Generates a new template */
	tmpl := template.Must(template.ParseFiles("bloglayout.html"))

	/* Data to populate template */
	data := PostPage{
		//Page Title
		PageTitle: "Divided We Stand",
		// List of Posts on the page
		Posts: []BlogPost{
			{Title: userTitle, Author: userAuthor, Date: userDate, Content: userContent},
			{Title: "Enemy of the State", Author: "Nick Walsh", Date: "7/29/19", Content: "Dogs. That's it."},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.ListenAndServe("localhost:8080", nil)
}
