package main

/*
#cgo LDFLAGS: -framework CoreServices
#include <stdlib.h>
int fswatch_monitor_paths(char** paths, int paths_n);
void fswatch_unwatch_dirs();
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
  ok := C.fswatch_monitor_paths(&paths[0], C.int(len(paths)))
  return ok != 0
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
