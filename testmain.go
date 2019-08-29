package main


import (
    "fmt"
)
func myWalkFn(path string, info os.FileInfo, err error) error {
  if err != nil {
      fmt.Printf("Could not walk filepath %s\n Maybe it doesn't exist?", path)
      return err
  }
  if !(info.isDir()) {

  }
}

func main () {
    paths := getBlogFiles(".")
    tbs := blogSite{}
    //Correct syntax?
    for i in range paths {
        parseBlogFiles(paths[i])
    }
    //test blog site
    //create empty blogsite
    tbs.Name = "Divided We Stand"
    //adds new blog to the tbs
    tbs.NewBlog()

    fmt.Println((*tbs.Blogs[0]).Title)
    fmt.Println((*tbs.Blogs[0]).Author)
    fmt.Println((*tbs.Blogs[0]).Post)

}
