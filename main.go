package main

import "os"

func main() {
  if options := parseOptions(os.Args[0], os.Args[1:], os.Stderr); options.valid {
    realDirWatcher(options).watchDirs()
  }
}
