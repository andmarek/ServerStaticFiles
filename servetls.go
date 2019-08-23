package main

import (
    "crypto/tls"
    "log"
    "net/http"
//    "html/template"
    "os"
    //autocert key generator
    "golang.org/x/crypto/acme/autocert"
)
// Need to define data

/*
func defaultHandler(w http.ResponseWriter, r *http.Request, tmpl *Template) {
  tmpl.Execute(w, data)
}
*/

func main() {
    var testStruct blog
    //Get the data from the blogfile and return it as a string 
    os.Setenv("BLOGFILE", "/blogs/BLOGFILE")
    // Might not be able to do dataLiason() without s:= first 
    marshallData(Post,  dataLiason("~/Code/ServerStaticFiles/"), &testStruct)

//    tmpl := template.Must(template.ParseFiles("Frontend/index.html"))

    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("helpamericathink.com"),
        Cache:      autocert.DirCache("certs"),
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//      w.Write([]byte("Hello world"))
      w.Write([]byte(testStruct.Post))
     // defaultHandler(w, r, tmpl)
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
