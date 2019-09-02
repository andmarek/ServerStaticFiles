package main

import (
	"fmt"
)

func main() {
	paths, _ := getBlogFiles(".")

	tbs := blogSite{}

  fmt.Println(len(paths))

	for i := range paths {
      tbs.NewBlog(paths[i])
	}

	//Print our paths to test it
	fmt.Println(paths)
	//test blog site
	//create empty blogsite
	tbs.Name = "Divided We Stand"

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
