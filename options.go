package main

import "flag"
import "io"
import "fmt"

type options struct {
  valid bool

  runInitially bool
  cmd string
  dirs []string
  args []string
}

func dashIndex(args []string) int {
  for i, v := range args {
    if v == "--" {
      return i
    }
  }
  return -1
}

func split(args []string) (dirs, cmdArgs []string, success bool) {
  i := dashIndex(args)

  if i <= 0 || i == len(args)-1 {
    return nil, nil, false
  }

  return args[:i], args[i+1:], true
}

func parseOptions(progName string, args []string, output io.Writer) options {
  var opts options
  var showVersion bool

  fs := flag.NewFlagSet(progName, flag.ContinueOnError)
  fs.SetOutput(output)
  fs.Usage = func() {
    fmt.Fprintf(output, "usage: %s <dir> [...] -- <cmd> [arg ...] \n", progName)
    fmt.Fprintf(output, "  -f = run command initially\n")
    fmt.Fprintf(output, "  -v = show version\n")
  }
  fs.BoolVar(&opts.runInitially, "f", false, "")
  fs.BoolVar(&showVersion, "v", false, "")
  fs.BoolVar(&showVersion, "version", false, "")
  err := fs.Parse(args)

  if showVersion {
    fmt.Fprintf(output, "version %s\n", Version)
    return opts
  }

  args = fs.Args()
  opts.dirs, opts.args, opts.valid = split(args)

  if opts.valid {
    opts.cmd, opts.args = opts.args[0], opts.args[1:]
  } else if err == nil {
    fs.Usage()
  }

  return opts
}
