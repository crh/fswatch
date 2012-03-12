#include <CoreServices/CoreServices.h>
#include <stdio.h>
#include "_cgo_export.h"

static void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
static void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  watchDirsCallback();
}

int fswatch_monitor_paths(char** paths, int paths_n) {
  CFMutableArrayRef pathsToWatch = CFArrayCreateMutable(NULL, 0, &kCFTypeArrayCallBacks);

  int i;
  for (i = 0; i < paths_n; i++)
    CFArrayAppendValue(pathsToWatch, CFStringCreateWithCString(NULL, paths[i], kCFStringEncodingUTF8));

  FSEventStreamRef stream = FSEventStreamCreate(NULL,
                                                callback,
                                                NULL,
                                                pathsToWatch,
                                                kFSEventStreamEventIdSinceNow,
                                                0.1,
                                                kFSEventStreamCreateFlagFileEvents
                                                | kFSEventStreamCreateFlagNoDefer);
  FSEventStreamScheduleWithRunLoop(stream, CFRunLoopGetCurrent(), kCFRunLoopCommonModes);

  if (!FSEventStreamStart(stream)) {
    return 0;
  }

  CFRunLoopRun();

  // we NEVER get here. ever. period.
  FSEventStreamStop(stream);
  FSEventStreamInvalidate(stream);
  FSEventStreamRelease(stream);

  return 1;
}
