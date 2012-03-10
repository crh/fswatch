#include <stdio.h>
#import <CoreServices/CoreServices.h>
#import <Foundation/Foundation.h>

#include "fswatch.h"

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  printf("\e[34;4m%s %s\e[0m\n", [commandToRun UTF8String], [[argumentsToUse componentsJoinedByString: @" "] UTF8String]);
  NSTask *task = [NSTask launchedTaskWithLaunchPath: fullPathToCommandToRun
                                          arguments: argumentsToUse];
  [task waitUntilExit];
  printf("\n");
}

int main (int argc, char** argv) {
  [NSAutoreleasePool new];

  if (argc < 3) {
    printf("usage: %s dir cmd arg1 arg2 argn...\n", argv[0]);
    exit(1);
  }

  split_out_cmd_args(argc, argv);

  if (fullPathToCommandToRun == nil) {
    fprintf(stderr, "error: could not find executable '%s'\n", [commandToRun UTF8String]);
    exit(1);
  }

  CFArrayRef pathsToWatch = (CFArrayRef)[NSArray arrayWithObject: dirToWatch];
  FSEventStreamRef stream = FSEventStreamCreate(NULL,
                                                callback,
                                                NULL,
                                                pathsToWatch,
                                                kFSEventStreamEventIdSinceNow,
                                                0.1,
                                                kFSEventStreamCreateFlagFileEvents | kFSEventStreamCreateFlagNoDefer);
  FSEventStreamScheduleWithRunLoop(stream, CFRunLoopGetCurrent(), kCFRunLoopCommonModes);
  if (!FSEventStreamStart(stream)) {
    fprintf(stderr, "error: failed to run for some reason\n");
    exit(1);
  }

  CFRunLoopRun();

  // we NEVER get here. ever. period.
  FSEventStreamStop(stream);
  FSEventStreamInvalidate(stream);
  FSEventStreamRelease(stream);

  return 0;
}
