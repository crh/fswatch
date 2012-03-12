package fswatch

import "testing"

func TestSomething(t *testing.T) {
  watchDirs([]string{ ".", "../blog" })
}
