package main

import "os"
import "os/signal"
import "fmt"

const Version = "3.0"

func main() {
  options := parseOptions(os.Args[0], os.Args[1:], os.Stderr)
  if !options.valid {
    return
  }

  cmd := command{
    name: options.cmd,
    args: options.args,
    outPipe: os.Stdout,
    errPipe: os.Stderr,
  }

  fsChange := make(chan bool)
  interrupt := make(chan os.Signal)

  go func() {
    for {
      select {
      case <-fsChange:
        decorate(cmd, invoke)
      case <-interrupt:
        unwatchDirs()
      }
    }
  }()

  signal.Notify(interrupt, os.Interrupt)
  fileSystemNotify(fsChange)

  if options.runInitially {
    decorate(cmd, invoke)
  }

  ok := watchDirs(options.dirs)
  if !ok {
    fmt.Fprintln(os.Stderr, "error: fsevent has failed us for the last time.")
  }
}
