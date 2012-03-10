#import <Foundation/Foundation.h>

extern NSString *dirToWatch;
extern NSString *commandToRun;
extern NSArray *argumentsToUse;

void split_out_cmd_args(int argc, char** argv);
