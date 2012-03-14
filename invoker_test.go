package main

import "testing"
import "github.com/sdegutis/assert"
import "bytes"
import "strings"
import "os"

func TestInvokerExecutesCommand(t *testing.T) {
  var output, errput bytes.Buffer
  invoker := newInvoker("pwd", []string{}, &output, &errput)
  invoker()

  pwd, _ := os.Getwd()
  assert.StringContains(t, output.String(), pwd)
}

func TestInvokerPrintsPrettily(t *testing.T) {
  var output, errput bytes.Buffer
  invoker := newInvoker("echo", []string{"hello", "world"}, &output, &errput)
  invoker()

  assert.True(t, strings.HasSuffix(output.String(), "\n\n"))
  assert.StringContains(t, output.String(), "echo hello world")
}
