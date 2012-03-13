package main

import "flag"
import "os"
import "io"
import "fmt"

type options struct {
  valid bool

  forceFirst bool
  cmd string
  dirs []string
  args []string
}

func split(args []string) (dirs, cmdArgs []string, success bool) {
  for i, v := range args {
    if v == "-" {
      return args[:i], args[i+1:], len(args[i+1:]) > 0
    }
  }
  return nil, nil, false
}

func parseOptions(progName string, args []string, output io.Writer) options {
  var opts options

  fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
  fs.SetOutput(output)
  fs.Usage = func() {
    fmt.Fprintf(output, "Usage: %s <dir> [...] - <cmd> [arg ...] \n", progName)
    fs.PrintDefaults()
  }
  fs.BoolVar(&opts.forceFirst, "f", false, "run the command initially")
  fs.Parse(args)
  args = fs.Args()

  opts.dirs, opts.args, opts.valid = split(args)
  if opts.valid {
    opts.cmd, opts.args = opts.args[0], opts.args[1:]
  } else {
    fs.Usage()
  }

  return opts
}
