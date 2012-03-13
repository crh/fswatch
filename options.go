package fswatch

import "flag"

type options struct {
  forceFirst bool
  cmd string
  dirs []string
  args []string
}

func split(args []string) (dirs []string, cmdArgs []string) {
  for i, v := range args {
    if v == "-" {
      dirs = args[:i]
      cmdArgs = args[i+1:]
      return
    }
  }
  return
}

func parseOptions(args []string) options {
  var forceFirst bool

  fs := flag.NewFlagSet("uhh", flag.ExitOnError)
  fs.BoolVar(&forceFirst, "f", false, "run the command initially")
  fs.Parse(args)

  args = fs.Args()

  dirs, cmdArgs := split(args)
  cmd, cmdArgs := cmdArgs[0], cmdArgs[1:]

  return options{
    forceFirst: forceFirst,
    dirs: dirs,
    args: cmdArgs,
    cmd: cmd,
  }
}
