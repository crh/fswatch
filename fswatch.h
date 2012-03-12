#import <Foundation/Foundation.h>

typedef struct _WatchOptions {
  BOOL notEnoughArgs;
  BOOL forceFirstRun;

  NSString *dirToWatch;

  NSString *commandToRun;
  NSString *fullPathToCommandToRun;

  NSArray *argumentsToUse;
} WatchOptions;

WatchOptions split_out_cmd_args(int argc, char** argv);
NSString* full_path_for(NSString* file);
