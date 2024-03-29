package main

import (
	"bufio"
  "errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Semantics to use as delimiters more or less.
const (
	blogBeginSemantic = "$BEGIN$"
	blogEndSemantic   = "$FINISH$"
	authorSemantic    = "$AUTHOR$"
	titleSemantic     = "$TITLE$"
	contentSemantic   = "$CONTENT$"

	blogPrefix = "piece"
)

//Problem: confused as to where the path is coming from
func getBlogFiles(dirPath string) ([]string, error) {
   var pathList []string
	//Maybe "." because we want it to look at cur dir THEN the dir
	err := filepath.Walk("blogs/", func(path string, f os.FileInfo, err error) error {
   // path = dirPath
		if err != nil {
			fmt.Printf("Ya'll got an error\n")
			return err
		}
		if !(f.IsDir()) && strings.Contains(f.Name(), blogPrefix) {
			pathList = append(pathList, path)
			//Remove later
			fmt.Println(pathList) //
		}
    return nil
	})

  if err != nil {

      return pathList, nil
  }

	return pathList, nil
}

// parseBlogFile parses files in $BLOGDIR and puts them into instances of a
// 'blog' struct.  The blog file follows a particular, and rather simple format
// to make parsing easier.  This may be revamped for future use and maybe I'll
// swap this out for a more efficient markdown parser or something, but this is
// just something I hacked up in my spare time.
func (bp *blog) parseBlogFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	snr := bufio.NewScanner(file)
	snr.Split(bufio.ScanLines)

	var s strings.Builder

	snr.Scan()

	//Check if there's a BEGIN semantic
	if snr.Text() != "$BEGIN$" {
		fmt.Printf("Not begin semantic\n")
		err := errors.New("Begin semantic error!")
		return err
	}

	for snr.Scan() {
		if snr.Text() == blogBeginSemantic {
			continue
		}
		if snr.Text() == authorSemantic {
			snr.Scan()
			bp.Author = snr.Text()
		}
		if snr.Text() == titleSemantic {
			snr.Scan()
			bp.Title = snr.Text()
		}
		if snr.Text() == contentSemantic {
			for snr.Scan() {
				if snr.Text() == blogEndSemantic {
					break
				}
				s.Write(snr.Bytes())
			}
			bp.Post = s.String()
		}
	}
	return nil
}
