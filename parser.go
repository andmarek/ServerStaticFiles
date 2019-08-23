package main
import (
    "bufio"
    "os"
    "fmt"
    "log"
    "strings"
    "path/filepath"

    )
type ContentType int

const (
  Author ContentType = 0
  Title  ContentType = 1
  Date   ContentType = 2
  Post   ContentType = 3
)

// Struct for blog and page
type blog struct {
  Author string
  Title string
  Date string
  /* Not sure if we want this as a string*/
  Post string
}

type Page struct {
  BlogList []blog
}

// dataLiason takes text from "BLOGFILE" and returns it as a string.
func dataLiason(fpath string) string {
  fmt.Println("Path: " + fpath)

  var blogContent string

  err := filepath.Walk("./blogs", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Printf("path, err: %q:, %v\n", path, err)
      return err
    }

    // Find the file 
    if info.Name() == "BLOGFILE" && info.IsDir() == false  {
     fmt.Printf("Found file: %s\n", info.Name())
    }

    file, err := os.Open("blogs/BLOGFILE")
    if err != nil {
      log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var s strings.Builder

    for scanner.Scan() {
        s.Write(scanner.Bytes())
    }

    blogContent = s.String()

    s.Reset()

    return err
  })

  if err != nil {
    log.Fatal("Could not find file\n")
  }

  return blogContent
}
// marshallData takes in the type of content, the data (a string) and a
// reference of a blog struct - placing the data into the blog. 
func marshallData(d ContentType, data string, blogInstance *blog) {
  if d == Author {
    blogInstance.Author = data
    fmt.Println("[Type]: Author\n")
  }

  if d == Date {
    blogInstance.Date = data
    fmt.Println("[Type]: Date\n")
  }

  if d == Title {
    blogInstance.Title = data
    fmt.Println("[Type]: Title\n")
  }

  if d == Post {
    blogInstance.Post = data
    fmt.Println("[Type]: Content\n")
  }
}
/*
func main() {
  var testStruct blog

  os.Setenv("BLOGFILE", "/blogs/BLOGFILE")
  s := dataLiason("~/Code/ServerStaticFiles/")
  marshallData(Post, s, &testStruct)
  fmt.Println(s)
  fmt.Println("testStruct blog post: %s", testStruct.Post)

}
*/
