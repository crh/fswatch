package main

import "flag"

type options struct {
  forceFirst bool
  cmd string
  dirs []string
  args []string

  valid bool
}

func split(args []string) (dirs, cmdArgs []string, success bool) {
  for i, v := range args {
    if v == "-" {
      return args[:i], args[i+1:], true
    }
  }
  return nil, nil, false
}

func parseOptions(args []string) options {
  var opts options

  fs := flag.NewFlagSet("uhh", flag.ExitOnError)
  fs.BoolVar(&opts.forceFirst, "f", false, "run the command initially")
  fs.Parse(args)
  args = fs.Args()

  opts.dirs, opts.args, opts.valid = split(args)
  if opts.valid {
    opts.cmd, opts.args = opts.args[0], opts.args[1:]
  }

  return opts
}
