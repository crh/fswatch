#include <stdio.h>
#import <CoreServices/CoreServices.h>
#import <Foundation/Foundation.h>

#include "fswatch.h"

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  void (^runTask)() = clientCallBackInfo;
  runTask();
}

int main (int argc, char** argv) {
  [NSAutoreleasePool new];

  WatchOptions options = split_out_cmd_args(argc, argv);

  if (options.notEnoughArgs) {
    printf("usage: %s dir [-f] cmd arg1 arg2 argn...\n"
           "   -f = also run command initially\n", argv[0]);
    exit(1);
  }

  if (options.fullPathToCommandToRun == nil) {
    fprintf(stderr, "error: could not find executable '%s'\n", [options.commandToRun UTF8String]);
    exit(1);
  }

  NSArray *args = [[NSArray arrayWithObject: options.commandToRun] arrayByAddingObjectsFromArray: options.argumentsToUse];
  const char* fullCommandString = [[args componentsJoinedByString: @" "] UTF8String];

  void (^runTask)() = [^{
    printf("\e[34;4m%s\e[0m\n", fullCommandString);
    NSTask *task = [NSTask launchedTaskWithLaunchPath: options.fullPathToCommandToRun
                                            arguments: options.argumentsToUse];
    [task waitUntilExit];
    printf("\n");
  } copy];

  FSEventStreamContext ctx;
  ctx.version = 0;
  ctx.info = runTask;
  ctx.retain = NULL;
  ctx.release = NULL;
  ctx.copyDescription = NULL;

  CFArrayRef pathsToWatch = (CFArrayRef)[NSArray arrayWithObject: options.dirToWatch];
  FSEventStreamRef stream = FSEventStreamCreate(NULL,
                                                callback,
                                                &ctx,
                                                pathsToWatch,
                                                kFSEventStreamEventIdSinceNow,
                                                0.1,
                                                kFSEventStreamCreateFlagNoDefer);
  FSEventStreamScheduleWithRunLoop(stream, CFRunLoopGetCurrent(), kCFRunLoopCommonModes);
  if (!FSEventStreamStart(stream)) {
    fprintf(stderr, "error: failed to run for some reason\n");
    exit(1);
  }

  if (options.forceFirstRun)
    runTask();

  CFRunLoopRun();

  // we NEVER get here. ever. period.
  FSEventStreamStop(stream);
  FSEventStreamInvalidate(stream);
  FSEventStreamRelease(stream);

  return 0;
}
