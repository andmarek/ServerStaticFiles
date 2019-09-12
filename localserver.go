package main

import (
	"fmt"
	"html/template"

	//    "log"
	"net/http"
)

func main() {
	//Template things

	tpl := template.Must(template.ParseFiles("tmpl/nev.html", "tmpl/site.html", "tmpl/blogtmpl.html"))
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

	testHand := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("the test handlder ran")
		tpl.ExecuteTemplate(w, "site", tbs)
		//		io.WriteString(w, "dogs")
	}

	http.HandleFunc("/", testHand)
	http.ListenAndServe(":9090", nil)
}
