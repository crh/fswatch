package main

/*
#cgo LDFLAGS: -framework CoreServices
#include <stdlib.h>
int fswatch_monitor_paths(char** paths, int paths_n);
*/
import "C"

import "fmt"
import "unsafe"

type dirWatcherInterface interface {
  watchDirs()
  callback()
}

var runningDirWatcher dirWatcherInterface

type realDirWatcher options

func (dw realDirWatcher) watchDirs() bool {
  var paths []*C.char
  for _, dir := range dw.dirs {
    path := C.CString(dir)
    defer C.free(unsafe.Pointer(path))
    paths = append(paths, path)
  }
  ok := C.fswatch_monitor_paths(&paths[0], C.int(len(paths)))
  return ok != 0
}

//export watchDirsCallback
func watchDirsCallback() {
  fmt.Println("uhh")
}
