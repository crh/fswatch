package main

import "os"
import "os/signal"
import "fmt"

import "github.com/sdegutis/mac-fsevents"

const Version = "3.0"

/*
dependencies:

argv[0]
argv[1:]
stdout
stderr
invoke()
watchDirs()
fileSystemNotify()
signal.Notify()
*/

func main() {
  // get the options
  options := parseOptions(os.Args[0], os.Args[1:], os.Stderr)
  if !options.valid {
    return
  }

  // setup the command
  cmd := command{
    name: options.cmd,
    args: options.args,
    outPipe: os.Stdout,
    errPipe: os.Stderr,
  }
  exec := decorate(cmd, invoke)

  // invoke it at first if required
  if options.runInitially {
    exec()
  }

  // register for path events
  fsChange := fsevents.WatchPaths(options.dirs)
  if fsChange == nil {
    fmt.Fprintln(os.Stderr, "error: fsevent has failed us for the last time.")
    return
  }

  // register for sigint events
  interrupt := make(chan os.Signal)
  signal.Notify(interrupt, os.Interrupt)

  // respond to events
  MainLoop:
  for {
    select {
    case <-fsChange:
      exec()
    case <-interrupt:
      break MainLoop
    }
  }
}
