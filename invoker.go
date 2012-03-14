package main

import "io"
import "os/exec"
import "fmt"
import "strings"

func newInvoker(cmd string, args []string, outPipe, errPipe io.Writer) func() {
  return func() {
    cmdStrings := append([]string{cmd}, args...)
    fmt.Fprintf(outPipe, "%c[34;4m%s%c[0m\n", 27, strings.Join(cmdStrings, " "), 27)
    cmd := exec.Command(cmd, args...)
    cmd.Stdout = outPipe
    cmd.Stderr = errPipe
    cmd.Run()
    fmt.Fprintln(outPipe)
  }
}
