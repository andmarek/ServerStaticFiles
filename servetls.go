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

    tpl := template.Must(template.ParseFiles("tmpl/nev.html","tmpl/site.tmpl"))
    //tpl := template.Must(template.New("tmpl/site.tmpl").ParseGlob("tmpl/*.tmpl"))
    paths, _ := getBlogFiles(".")

    tbs := blogSite{}

    for i := range paths {
        tbs.NewBlog(paths[i])
    }

    tbs.Name = "Divided We Stand"


// Just server things xD
//Cert
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("helpamericathink.com"),
        Cache:      autocert.DirCache("certs"),
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      //;tpl.Execute(w, tbs)
      tpl.ExecuteTemplate(w, "site", tbs)
    })

/*    http.Handle("/static/", http.StripPrefix("/static/",
 *    http.FileServer(http.Dir("static")))) */

    server := &http.Server{
        Addr:":https",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }
//go routine for listen and serve 
    go http.ListenAndServe(":http",certManager.HTTPHandler(nil))
    log.Fatal(server.ListenAndServeTLS("",""))
}
