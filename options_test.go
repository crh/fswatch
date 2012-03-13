package main

import "testing"
import "github.com/sdegutis/assert"
import "bytes"

func TestArgsMissingDash(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "echo", "hello", "world"}, &output)
  assert.False(t, opts.valid)

  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "Usage")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " - ")
}

func TestArgsMissingEverything(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{}, &output)
  assert.False(t, opts.valid)

  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "Usage")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " - ")
}

func TestArgsMissingCommand(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "-"}, &output)
  assert.False(t, opts.valid)

  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "Usage")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " - ")
}

func TestBasicArgs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "-", "echo", "hello", "world"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"."})
  assert.Equals(t, opts.cmd, "echo")
  assert.DeepEquals(t, opts.args, []string{"hello", "world"})
}

func TestMultipleDirs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"spec", "features", "-", "rake", "spec", "cucumber"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"spec", "features"})
  assert.Equals(t, opts.cmd, "rake")
  assert.DeepEquals(t, opts.args, []string{"spec", "cucumber"})
}

func TestSimpleArgs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"src", "-", "pwd"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}

func TestSimpleArgsWithForce(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"-f", "src", "-", "pwd"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.True(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}
