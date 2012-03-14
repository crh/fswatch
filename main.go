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

  invoker := newInvoker(options.cmd, options.args, os.Stdout, os.Stderr)

  fsChange := make(chan bool)
  interrupt := make(chan os.Signal)

  go func() {
    for {
      select {
      case <-fsChange:
        invoker()
      case <-interrupt:
        unwatchDirs()
      }
    }
  }()

  signal.Notify(interrupt, os.Interrupt)
  fileSystemNotify(fsChange)

  if options.runInitially {
    invoker()
  }

  ok := watchDirs(options.dirs)
  if !ok {
    fmt.Fprintln(os.Stderr, "error: fsevent has failed us for the last time.")
  }
}
