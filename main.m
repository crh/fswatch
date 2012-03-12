#include <stdio.h>
#import <CoreServices/CoreServices.h>
#import <Foundation/Foundation.h>

#include "fswatch.h"

void runTask() {
  static const char* fullCommandString;
  if (fullCommandString == NULL) {
    NSArray *args = [[NSArray arrayWithObject: commandToRun] arrayByAddingObjectsFromArray: argumentsToUse];
    fullCommandString = [[args componentsJoinedByString: @" "] UTF8String];
  }

  printf("\e[34;4m%s\e[0m\n", fullCommandString);
  NSTask *task = [NSTask launchedTaskWithLaunchPath: fullPathToCommandToRun
                                          arguments: argumentsToUse];
  [task waitUntilExit];
  printf("\n");
}

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  runTask();
}

int main (int argc, char** argv) {
  [NSAutoreleasePool new];

  split_out_cmd_args(argc, argv);

  if (notEnoughArgs) {
    printf("usage: %s dir [-f] cmd arg1 arg2 argn...\n"
           "  -f,\trun command immediately as well", argv[0]);
    exit(1);
  }

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

  if (forceFirstRun)
    runTask();

  CFRunLoopRun();

  // we NEVER get here. ever. period.
  FSEventStreamStop(stream);
  FSEventStreamInvalidate(stream);
  FSEventStreamRelease(stream);

  return 0;
}
