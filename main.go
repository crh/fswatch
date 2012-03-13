package main

import "os"

func main() {
  parseOptions(os.Args[0], os.Args[1:], os.Stderr)
}
