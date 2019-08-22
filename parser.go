package main
import (
    "bufio"
//    "bytes"
    "os"
    "fmt"
    "log"
    "strings"
    "path/filepath"

    )
// Struct for blog and page
type blog struct {
  Author string
  Title string
  Date string
  /* Not sure if we want this as a string*/
  content string
}

type Page struct {
  BlogList []blog
}

func dataLiason(fpath string) {
  fmt.Println("Entered the appropriate function")
  fmt.Println(fpath)

  err := filepath.Walk("./blogs", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Printf("path, err: %q:, %v\n", path, err)
      return err
    }

    // Find the file 
    if info.Name() == "BLOGFILE" && info.IsDir() == false  {
     fmt.Printf("Found file !")
    }

    file, err := os.Open("blogs/BLOGFILE")

    if err != nil {
      log.Fatal(err)
    }

    defer file.Close()

    //For parsing whole file 
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    //For building lines 
    var s strings.Builder
    for scanner.Scan() {
        s.Write(scanner.Bytes() + " ")
        fmt.Println(s.String())
    }
    s.Reset()

    return err
  })

  if err != nil {
    log.Fatal("Could not find file")
  }
}

func main() {
  os.Setenv("BLOGFILE", "/blogs/BLOGFILE")
  fmt.Println("env" + os.Getenv("BLOGFILE"))

  fmt.Println("Begin main function")
  dataLiason("~/Code/ServerStaticFiles/")

}

