package main

import "fmt"
import "strings"
import "github.com/sdegutis/shattr"

func decorate(cmd command, exec func(command)) func() {
  return func() {
    cmdStrings := append([]string{cmd.name}, cmd.args...)

    output := shattr.NewWriter(cmd.outPipe, shattr.Underline, shattr.Blue)
    fmt.Fprintln(output, strings.Join(cmdStrings, " "))

    exec(cmd)

    fmt.Fprintln(cmd.outPipe)
  }
}
