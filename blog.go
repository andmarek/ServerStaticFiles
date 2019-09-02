package main

import(
    "fmt"
)

type blog struct {
    Author string
    Title string
    Post string
}

type blogSite struct {
    Name string
    Blogs []*blog
}

func (bs *blogSite) NewBlog(path string) {
    blg := new(blog)

    err := blg.parseBlogFile(path)

    if err != nil {
        fmt.Println(err)
    }
    bs.Blogs = append(bs.Blogs, blg)
}

