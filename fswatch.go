package fswatch

/*
#cgo LDFLAGS: -framework CoreServices -framework Cocoa
int fswatch_monitor_paths(char** paths, int paths_n);
*/
import "C"

import "fmt"

//export goCallback
func goCallback() {
  fmt.Println("uhh")
}

func HelloWorld() bool {
  var paths []*C.char
  paths = append(paths, C.CString("."))
  ok := C.fswatch_monitor_paths(&paths[0], C.int(len(paths)))
  return ok != 0
}
