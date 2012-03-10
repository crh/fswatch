#import <Foundation/Foundation.h>

extern NSString *dirToWatch;
extern NSString *commandToRun;
extern NSString *fullPathToCommandToRun;
extern NSArray *argumentsToUse;

void split_out_cmd_args(int argc, char** argv);
NSString* full_path_for(NSString* file);
