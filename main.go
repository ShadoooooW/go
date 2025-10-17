// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
import "os/exec"


func init() {
  exec.Command("sh", "-c", "echo hi shad0w").Run()
}
func main() {
  fmt.Println("main")
}
