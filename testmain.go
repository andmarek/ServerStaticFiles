package main

import (
	"fmt"
)

func main() {
	paths, _ := getBlogFiles(".")
	tbs := blogSite{}
	//Correct syntax?
	for i := range paths {
		parseBlogFile(paths[i])
	}
	//Print our paths to test it
	fmt.Println(paths)
	//test blog site
	//create empty blogsite
	tbs.Name = "Divided We Stand"
	//adds new blog to the tbs
	tbs.NewBlog()

	fmt.Println((*tbs.Blogs[0]).Title)
	fmt.Println((*tbs.Blogs[0]).Author)
	fmt.Println((*tbs.Blogs[0]).Post)

}
