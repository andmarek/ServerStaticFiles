package main
import (
    "bufio"
    "os"
    "fmt"
    "log"
    )

func main() {

  file, err := os.Open("BLOGFILE")
    if err != nil {
      log.Fatal(err)
    }
  defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    success := scanner.Scan()

    if success == false {
      err = scanner.Err()

        if err == nil {
          log.Println("Scan completed and reached EOF")
        } else {
          log.Fatal(err)
        }
    }

  fmt.Println("First word found: ", scanner.Text())
}
