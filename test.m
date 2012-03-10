#import "fswatch.h"
#import <assert.h>
#import <Foundation/Foundation.h>

int main(int argc, char** argv) {
  [NSAutoreleasePool new];

  split_out_cmd_args(6, (char*[]){
    "this_executable",
      "some/dir",
      "a_command",
      "first arg",
      "second",
      "arg",
  });

  assert([dirToWatch isEqualTo: @"some/dir"]);

  assert([commandToRun isEqualTo: @"a_command"]);
  NSArray *expectedArgs = [NSArray arrayWithObjects:
    @"first arg",
    @"second",
    @"arg",
    nil];
  assert([argumentsToUse isEqualTo: expectedArgs]);

  return 0;
}
