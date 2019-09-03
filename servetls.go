package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "html/template"

    //autocert key generator
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    //Template things
    tpl := template.Must(template.ParseFiles("tmpl/testindex.tmpl"))

    // Create paths slice with a list of string paths
    paths, _ := getBlogFiles(".")

    //New "The Blog Site"
    tbs := blogSite{}

    // NewBlog takes in a string argument
    for i := range paths {
        tbs.NewBlog(paths[i])
    }

    tbs.Name = "Divided We Stand"


// Just server things xD
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("helpamericathink.com"),
        Cache:      autocert.DirCache("certs"),
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      tpl.Execute(w, tbs)
    })

    server := &http.Server{
        Addr:":https",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }

    go http.ListenAndServe(":http",certManager.HTTPHandler(nil))
    log.Fatal(server.ListenAndServeTLS("",""))
}
