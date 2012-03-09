//
//  main.c
//  fswatch
//
//  Created by Steven Degutis on 3/9/12.
//  Copyright 2012 8th Light. All rights reserved.
//

#include <stdio.h>
#include <CoreServices/CoreServices.h>

#import <Foundation/Foundation.h>

NSString* watchfulCommandToRun;

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
    NSLog(@"WOO! [%@]\n", watchfulCommandToRun);
}

int main (int argc, char * argv[]) {
    argc = 4;
    argv = (char*[]){ "me", "pwd", "/Users/sdegutis/projects/go/src/github.com/sdegutis/mapstruct", "/Users/sdegutis/projects/go/src/github.com/sdegutis/blog" };

    
    [NSAutoreleasePool new];
    
    if (argc < 2) {
        printf("usage: %s command path ...", argv[0]);
        exit(1);
    }
    
    watchfulCommandToRun = [[NSString stringWithCString:argv[1] encoding:NSUTF8StringEncoding] retain];
    
    CFMutableArrayRef pathsToWatch = CFArrayCreateMutable(NULL, 0, &kCFTypeArrayCallBacks);
    
    for (int i = 2; i < argc; i++) {
        CFArrayAppendValue(pathsToWatch, CFStringCreateWithCString(NULL, argv[i], kCFStringEncodingUTF8));
    }
    
    CFShow(pathsToWatch);
    
    FSEventStreamRef stream = FSEventStreamCreate(NULL, callback, NULL, pathsToWatch, kFSEventStreamEventIdSinceNow, 0.5, kFSEventStreamCreateFlagFileEvents);
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
