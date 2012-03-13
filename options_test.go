package main

import "testing"
import "github.com/sdegutis/assert"

func TestArgsMissingDash(t *testing.T) {
  opts := parseOptions([]string{".", "echo", "hello", "world"})
  assert.False(t, opts.valid)
}

func TestBasicArgs(t *testing.T) {
  opts := parseOptions([]string{".", "-", "echo", "hello", "world"})
  assert.True(t, opts.valid)
  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"."})
  assert.Equals(t, opts.cmd, "echo")
  assert.DeepEquals(t, opts.args, []string{"hello", "world"})
}

func TestMultipleDirs(t *testing.T) {
  opts := parseOptions([]string{"spec", "features", "-", "rake", "spec", "cucumber"})
  assert.True(t, opts.valid)
  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"spec", "features"})
  assert.Equals(t, opts.cmd, "rake")
  assert.DeepEquals(t, opts.args, []string{"spec", "cucumber"})
}

func TestSimpleArgs(t *testing.T) {
  opts := parseOptions([]string{"src", "-", "pwd"})
  assert.True(t, opts.valid)
  assert.False(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}

func TestSimpleArgsWithForce(t *testing.T) {
  opts := parseOptions([]string{"-f", "src", "-", "pwd"})
  assert.True(t, opts.valid)
  assert.True(t, opts.forceFirst)
  assert.DeepEquals(t, opts.dirs, []string{"src"})
  assert.Equals(t, opts.cmd, "pwd")
  assert.Equals(t, len(opts.args), 0)
}
