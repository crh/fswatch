package main

import "testing"
import "github.com/sdegutis/go.assert"
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

  fullCommand :=  "echo hello world"

  assert.StringContains(t, output.String(), fullCommand)

  outstr := output.String()
  indexAferCmd := strings.Index(outstr, fullCommand) + len(fullCommand)
  assert.True(t, strings.Count(outstr[indexAferCmd:], "\n") == 2)
}
