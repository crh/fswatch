package main

import "testing"
import "github.com/sdegutis/assert"

func TestWatcherCallsInvoke(t *testing.T) {
  observer := make(chan []PathEvent)
  ch := make(chan int)

  go func() {
    for {
      <-observer
      ch <- 1
    }
  }()

  fileSystemNotify(observer)

  i := 0
  watchDirsCallback()
  i += <-ch
  watchDirsCallback()
  i += <-ch

  assert.Equals(t, i, 2)
}
