#include "fswatch.h"

WatchOptions split_out_cmd_args(int argc, char** argv) {
  WatchOptions options;

  argv++, --argc;

  options.forceFirstRun = (argc >= 1 && strcmp(argv[0], "-f") == 0);
  if (options.forceFirstRun)
    argv++, --argc;

  options.notEnoughArgs = (argc < 2);
  if (options.notEnoughArgs)
    return options;

  options.dirToWatch = [[NSString stringWithCString: argv[0] encoding: NSUTF8StringEncoding] retain];
  options.commandToRun = [[NSString stringWithCString: argv[1] encoding: NSUTF8StringEncoding] retain];

  options.fullPathToCommandToRun = full_path_for(options.commandToRun);

  NSMutableArray *args = [NSMutableArray array];
  for (int i = 2; i < argc; i++) {
    NSString *arg = [[NSString stringWithCString: argv[i] encoding: NSUTF8StringEncoding] retain];
    [args addObject: arg];
  }
  options.argumentsToUse = [args retain];

  return options;
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
