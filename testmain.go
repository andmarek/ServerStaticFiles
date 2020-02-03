package main

import (
    "html/template"
    "net/http"
    "fmt"
    "log"
)

func main() {

    tpl := template.Must(template.ParseFiles("tmpl/testindex.html"))

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

//    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
//    http.Handle("/", http.FileServer(http.Dir("css/")))


/* good stuff */
    paths, _ := getBlogFiles(".")

    tbs := blogSite{}

    tbs.Name = "Divided We Stand"

    fmt.Println(len(paths))

    for i := range paths {
        tbs.NewBlog(paths[i])
    }
/***/

    helloHandler := func(w http.ResponseWriter, req *http.Request) {
        tpl.Execute(w, tbs)
    }

    http.HandleFunc("/", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))

    //Print our paths to test it
    fmt.Println(paths)
    //test blog site
    //create empty blogsite

    // No blog ?
    //Index out of bounds rn
    fmt.Println(len(tbs.Blogs))

    fmt.Println((*tbs.Blogs[0]).Title)
    fmt.Println((*tbs.Blogs[0]).Author)
    fmt.Println((*tbs.Blogs[0]).Post)

    fmt.Println((*tbs.Blogs[1]).Title)
    fmt.Println((*tbs.Blogs[1]).Author)
    fmt.Println((*tbs.Blogs[1]).Post)

}
