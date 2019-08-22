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

func dataLiason(fpath string) string {
  fmt.Println("Entered the appropriate function")
  fmt.Println(fpath)
  var blogContent string
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
        s.Write(scanner.Bytes())
    }
    blogContent = s.String()
//    fmt.Println(s.String())
    s.Reset()

    return err
  })

  if err != nil {
    log.Fatal("Could not find file")
  }
  return blogContent
}

func marshallData(d ContentType, data string, blogInstance *blog) {
//Place the data into different struct fields
  if d == Author {
    blogInstance.Author = data
    fmt.Println("Author")
  }

  if d == Date {
    blogInstance.Date = data
    fmt.Println("Date")
  }

  if d == Title {
    blogInstance.Title = data
    fmt.Println("Title")
  }

  if d == Post {
    blogInstance.Post = data
    fmt.Println("Post")
  }
}

func main() {
  var testPage Page
  var testStruct blog

  os.Setenv("BLOGFILE", "/blogs/BLOGFILE")
  s := dataLiason("~/Code/ServerStaticFiles/")
  marshallData(Post, s, &testStruct)
  fmt.Println(s)
  fmt.Println("testStruct blog post: %s", testStruct.Post)

}

