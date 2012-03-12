package fswatch

/*
#cgo LDFLAGS: -framework CoreServices -framework Cocoa
int fswatch_monitor_paths(char** paths, int paths_n);
*/
import "C"

import "fmt"

//export watchDirsCallback
func watchDirsCallback() {
  fmt.Println("uhh")
}

func watchDirs(dirs []string) bool {
  var paths []*C.char
  for _, dir := range dirs {
    paths = append(paths, C.CString(dir))
  }
  ok := C.fswatch_monitor_paths(&paths[0], C.int(len(paths)))
  return ok != 0
}
