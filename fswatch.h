#import <Foundation/Foundation.h>

NSString* full_path_for(NSString* file);

typedef struct _WatchOptions {
  BOOL notEnoughArgs;
  BOOL forceFirstRun;

  NSString *dirToWatch;

  NSString *commandToRun;
  NSString *fullPathToCommandToRun;

  NSArray *argumentsToUse;
} WatchOptions;

WatchOptions split_out_cmd_args(int argc, char** argv);
