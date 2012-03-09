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

void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]);
void callback(ConstFSEventStreamRef streamRef, void *clientCallBackInfo, size_t numEvents, void *eventPaths, const FSEventStreamEventFlags eventFlags[], const FSEventStreamEventId eventIds[]) {
    printf("WOO!\n");
}

int main (int argc, char * argv[]) {
    argc = 3;
    argv = (char*[]){ "yes", "/Users/sdegutis/projects/go/src/github.com/sdegutis/mapstruct", "/Users/sdegutis/projects/go/src/github.com/sdegutis/blog" };
    
    printf("%d\n", argc * 1);
    
    CFMutableArrayRef pathsToWatch = CFArrayCreateMutable(NULL, 0, &kCFTypeArrayCallBacks);
    
    for (int i = 0; i < argc; i++) {
        CFArrayAppendValue(pathsToWatch, CFStringCreateWithCString(NULL, argv[i], kCFStringEncodingUTF8));
    }
    
    [NSAutoreleasePool new];
    
    CFShow(pathsToWatch);
    
    FSEventStreamRef stream = FSEventStreamCreate(NULL, callback, NULL, pathsToWatch, kFSEventStreamEventIdSinceNow, 0.5, kFSEventStreamCreateFlagFileEvents);
    FSEventStreamScheduleWithRunLoop(stream, CFRunLoopGetCurrent(), kCFRunLoopCommonModes);
    if (!FSEventStreamStart(stream)) {
        fprintf(stderr, "error: failed to run for some reason\n");
        exit(1);
    }
    
    CFRunLoopRun();
    FSEventStreamStop(stream);
    FSEventStreamInvalidate(stream);
    FSEventStreamRelease(stream);
    
    return 0;
}
