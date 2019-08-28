package main

import (
//    "time"
    "errors"
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Semantics to use as delimiters more or less.
const blogBeginSemantic = "$BEGIN$"
const blogEndSemantic = "$FINISH$"
const authorSemantic = "$AUTHOR$"
const titleSemantic = "$TITLE$"
const contentSemantic = "$CONTENT$"

// parseBlogFile parses files in $BLOGDIR and puts them into instances of a
// 'blog' struct.  The blog file follows a particular, and rather simple format
// to make parsing easier.  This may be revamped for future use and maybe I'll
// swap this out for a more efficient markdown parser or something, but this is
// just something I hacked up in my spare time.
func (bp *blog) parseBlogFile() error {
    file, err := os.Open("blogs/piece1")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()

    snr := bufio.NewScanner(file)
    snr.Split(bufio.ScanLines)

    var s strings.Builder

    snr.Scan()

    //Check if there's a BEGIN semantic
    if(snr.Text() != "$BEGIN$") {
        fmt.Printf("Not begin semantic\n")
        err := errors.New("Begin semantic error!")
        return err
    } else  {
        fmt.Printf("Found being semantic\n")
    }

    for snr.Scan() {
        if snr.Text() == blogBeginSemantic{
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
