package main

/*
#cgo LDFLAGS: -framework CoreServices
#include <stdlib.h>
int fswatch_monitor_paths(char** paths, int paths_n);
void fswatch_unwatch_dirs();
void CFRunLoopRun();
*/
import "C"
import "unsafe"

var fileSystemChangeObservers []chan bool

func fileSystemNotify(ch chan bool) {
  fileSystemChangeObservers = append(fileSystemChangeObservers, ch)
}

func watchDirs(dirs []string) bool {
  var paths []*C.char
  for _, dir := range dirs {
    path := C.CString(dir)
    defer C.free(unsafe.Pointer(path))
    paths = append(paths, path)
  }

  successChan := make(chan bool)

  go func(successChan chan bool) {
    ok := C.fswatch_monitor_paths(&paths[0], C.int(len(paths))) != 0

    if ok {
      successChan <- true
      C.CFRunLoopRun()
    } else {
      successChan <- false
    }
  }(successChan)

  return <-successChan
}

func unwatchDirs() {
  C.fswatch_unwatch_dirs()
}

//export watchDirsCallback
func watchDirsCallback() {
  for _, ch := range fileSystemChangeObservers {
    ch <- true
  }
}
