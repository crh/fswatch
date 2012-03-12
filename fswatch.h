#import <Foundation/Foundation.h>

extern BOOL notEnoughArgs;

extern NSString *dirToWatch;
extern NSString *commandToRun;
extern NSString *fullPathToCommandToRun;
extern NSArray *argumentsToUse;
extern BOOL forceFirstRun;

void split_out_cmd_args(int argc, char** argv);
NSString* full_path_for(NSString* file);
