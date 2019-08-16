package main

import (
    "crypto/tls"
    "log"
    "net/http"

    "golang.org/x/crypto/acme/autocert"
)

func main() {
    certManager := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("helpamericathink.com"),
        Cache:      autocert.DirCache("certs"),
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello world"))
    })

    server := &http.Server{
        //Addr:":https",
        Addr:":https",
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
        },
    }


    go http.ListenAndServe(":http",certManager.HTTPHandler(nil))
    log.Fatal(server.ListenAndServeTLS("",""))
}
