#include "fswatch.h"

BOOL notEnoughArgs;

NSString *dirToWatch;
NSString *commandToRun;
NSString *fullPathToCommandToRun;
NSArray *argumentsToUse;
BOOL forceFirstRun;

void split_out_cmd_args(int argc, char** argv) {
  argv++, --argc;

  forceFirstRun = (argc >= 1 && strcmp(argv[0], "-f") == 0);
  if (forceFirstRun)
    argv++, --argc;

  notEnoughArgs = (argc < 2);
  if (notEnoughArgs)
    return;

  dirToWatch = [[NSString stringWithCString: argv[0] encoding: NSUTF8StringEncoding] retain];
  commandToRun = [[NSString stringWithCString: argv[1] encoding: NSUTF8StringEncoding] retain];

  fullPathToCommandToRun = full_path_for(commandToRun);

  NSMutableArray *args = [NSMutableArray array];
  for (int i = 2; i < argc; i++) {
    NSString *arg = [[NSString stringWithCString: argv[i] encoding: NSUTF8StringEncoding] retain];
    [args addObject: arg];
  }
  argumentsToUse = [args retain];
}

NSString* full_path_for(NSString* file) {
  if ([file rangeOfString: @"/"].length > 0) {
    NSString *path = [[[NSFileManager defaultManager] currentDirectoryPath] stringByAppendingPathComponent: file];
    if ([[NSFileManager defaultManager] isExecutableFileAtPath: path])
      return [[path stringByStandardizingPath] retain];
    else
      return nil;
  }

  NSString *lookupPaths = [NSString stringWithCString:getenv("PATH") encoding:NSUTF8StringEncoding];
  NSArray *paths = [lookupPaths componentsSeparatedByString:@":"];
  for (NSString* path in paths) {
    path = [path stringByAppendingPathComponent: file];
    if ([[NSFileManager defaultManager] isExecutableFileAtPath: path])
      return [path retain];
  }
  return nil;
}
