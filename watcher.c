#include <CoreServices/CoreServices.h>
#include <stdio.h>
#include "_cgo_export.h"

static void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
static void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  watchDirsCallback();
}

static FSEventStreamRef stream;

void fswatch_unwatch_dirs() {
  FSEventStreamStop(stream);
  FSEventStreamInvalidate(stream);
  FSEventStreamRelease(stream);
}

int fswatch_monitor_paths(char** paths, int paths_n) {
  CFMutableArrayRef pathsToWatch = CFArrayCreateMutable(NULL, 0, &kCFTypeArrayCallBacks);

  int i;
  for (i = 0; i < paths_n; i++)
    CFArrayAppendValue(pathsToWatch, CFStringCreateWithCString(NULL, paths[i], kCFStringEncodingUTF8));

  stream = FSEventStreamCreate(NULL,
                               callback,
                               NULL,
                               pathsToWatch,
                               kFSEventStreamEventIdSinceNow,
                               0.1,
                               kFSEventStreamCreateFlagNoDefer);
  FSEventStreamScheduleWithRunLoop(stream, CFRunLoopGetCurrent(), kCFRunLoopCommonModes);

  return FSEventStreamStart(stream);
}
