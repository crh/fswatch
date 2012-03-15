package main

import "io"
import "os/exec"

type command struct {
  name string
  args []string
  outPipe, errPipe io.Writer
}

func invoke(cmd command) {
  command := exec.Command(cmd.name, cmd.args...)
  command.Stdout = cmd.outPipe
  command.Stderr = cmd.errPipe
  command.Run()
}
