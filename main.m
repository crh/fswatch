#include <stdio.h>
#import <CoreServices/CoreServices.h>
#import <Foundation/Foundation.h>

#include "fswatch.h"

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
  NSTask *task = [NSTask launchedTaskWithLaunchPath: commandToRun
                                          arguments: argumentsToUse];
  //NSLog(@"WOO! [%@] %@\n", commandToRun, argumentsToUse);
}

int main (int argc, char** argv) {
    [NSAutoreleasePool new];

    if (argc < 2) {
        printf("usage: %s command path ...", argv[0]);
        exit(1);
    }

    split_out_cmd_args(argc, argv);

    CFArrayRef pathsToWatch = (CFArrayRef)[NSArray arrayWithObject: dirToWatch];
    FSEventStreamRef stream = FSEventStreamCreate(NULL, callback, NULL, pathsToWatch, kFSEventStreamEventIdSinceNow, 0, kFSEventStreamCreateFlagFileEvents);
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
