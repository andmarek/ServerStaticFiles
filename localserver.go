package main

import (
	"fmt"
	"html/template"
	"io"

	//    "log"
	"net/http"
)

func main() {
	//Template things

	//Loading blog first, nav then since site puts it together
	tpl := template.Must(template.ParseFiles("tmpl/blogtmpl.html", "tmpl/real_nav.html", "tmpl/site.html"))

	//    /Users/andrewmarek/ServerStaticFiles/tmpl/blogtmpl.html
	paths, _ := getBlogFiles(".")

	tbs := blogSite{}

	for i := range paths {
		tbs.NewBlog(paths[i])
	}

	tbs.Name = "Divided We Stand"

	for i := range tbs.Blogs {
		fmt.Println("title:" + tbs.Blogs[i].Title)
		fmt.Println("Post:" + tbs.Blogs[i].Post)
	}
	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entered the home handler")
		tpl.ExecuteTemplate(w, "tpl", tbs)
	}

	// About page
	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entered the about handler")
		//	tpl.ExecuteTemplate(w, "", "data structure")

	}

	// Handles any contact page interaction
	contactHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")

	}

	http.HandleFunc("/", homeHandler)

	// Maybe use http.Handler() for this
	http.HandleFunc("/blog{n}", blogHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":9090", nil)
}
