package main

import "testing"
import "github.com/sdegutis/assert"
import "bytes"
import "strings"

func TestDecorateExecutesCommand(t *testing.T) {
  var output, errput bytes.Buffer

  ex := command{
    name: "pwd",
    args: []string{},
    outPipe: &output,
    errPipe: &errput,
  }

  var ranCmd *command = nil

  fakeRun := func(cmd command) {
    ranCmd = &cmd
  }

  decorate(ex, fakeRun)()

  assert.DeepEquals(t, ex, *ranCmd)
}

func TestDecoratePrintsPrettily(t *testing.T) {
  var output, errput bytes.Buffer

  ex := command{
    name: "echo",
    args: []string{"hello", "world"},
    outPipe: &output,
    errPipe: &errput,
  }

  decorate(ex, func(cmd command){})()

  assert.True(t, strings.HasSuffix(output.String(), "\n\n"))
  assert.StringContains(t, output.String(), "echo hello world")
}
