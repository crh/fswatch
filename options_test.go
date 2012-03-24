package main

import "testing"
import "github.com/sdegutis/go.assert"
import "bytes"
import "strings"

func TestArgsVersion(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"-v"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 0)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "version")
  assert.StringContains(t, output.String(), Version)
  assert.True(t, strings.HasSuffix(output.String(), "\n"))
}

func TestArgsLongVersion(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"--version"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 0)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "version")
  assert.StringContains(t, output.String(), Version)
  assert.True(t, strings.HasSuffix(output.String(), "\n"))
}

func TestArgsRequestingHelp(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"-h"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestArgsBadFlag(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"-j"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestArgsMissingDash(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "echo", "hello", "world"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestArgsMissingEverything(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestArgsMissingDirectories(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"--", "pwd"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestArgsMissingCommand(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "--"}, &output)
  assert.False(t, opts.valid)

  assert.Equals(t, strings.Count(output.String(), "usage"), 1)
  assert.True(t, output.Len() > 0)
  assert.StringContains(t, output.String(), "this_program")
  assert.StringContains(t, output.String(), "-f")
  assert.StringContains(t, output.String(), " -- ")
}

func TestBasicArgs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{".", "--", "echo", "hello", "world"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.runInitially)
  assert.DeepEquals(t, opts.dirs, []string{"."})
  assert.Equals(t, opts.cmd, "echo")
  assert.DeepEquals(t, opts.args, []string{"hello", "world"})
}

func TestMultipleDirs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"spec", "features", "--", "rake", "spec", "cucumber"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.runInitially)
  assert.DeepEquals(t, opts.dirs, []string{"spec", "features"})
  assert.Equals(t, opts.cmd, "rake")
  assert.DeepEquals(t, opts.args, []string{"spec", "cucumber"})
}

func TestSimpleArgs(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"src", "--", "pwd"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.False(t, opts.runInitially)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}

func TestSimpleArgsWithInitialRun(t *testing.T) {
  var output bytes.Buffer
  opts := parseOptions("this_program", []string{"-f", "src", "--", "pwd"}, &output)
  assert.True(t, output.Len() == 0)
  assert.True(t, opts.valid)

  assert.True(t, opts.runInitially)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}
