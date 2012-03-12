#import "fswatch.h"
#import <assert.h>
#import <Foundation/Foundation.h>

int main(int argc, char** argv) {
  [NSAutoreleasePool new];

  // just barely too few args, without -f
  {
    split_out_cmd_args(2, (char*[]){
        "this_executable",
        "some/dir",
        });

    assert(notEnoughArgs == YES);
  }

  // just barely enough args, without -f
  {
    split_out_cmd_args(3, (char*[]){
        "this_executable",
        "some/dir",
        "pwd",
        });

    assert(notEnoughArgs == NO);
  }

  // just barely too few args, with -f
  {
    split_out_cmd_args(3, (char*[]){
        "this_executable",
        "-f",
        "some/dir",
        });

    assert(notEnoughArgs == YES);
  }

  // without -f
  {
    split_out_cmd_args(6, (char*[]){
        "this_executable",
        "some/dir",
        "echo",
        "first arg",
        "second",
        "arg",
        });

    assert(notEnoughArgs == NO);

    assert(forceFirstRun == NO);
    assert([dirToWatch isEqualTo: @"some/dir"]);

    assert([fullPathToCommandToRun isEqualTo: @"/bin/echo"]);
    assert([commandToRun isEqualTo: @"echo"]);
    NSArray *expectedArgs = [NSArray arrayWithObjects:
      @"first arg",
      @"second",
      @"arg",
      nil];
    assert([argumentsToUse isEqualTo: expectedArgs]);
  }

  // with -f
  {
    split_out_cmd_args(7, (char*[]){
        "this_executable",
        "-f",
        "some/dir",
        "echo",
        "first arg",
        "second",
        "arg",
        });

    assert(notEnoughArgs == NO);

    assert(forceFirstRun == YES);
    assert([dirToWatch isEqualTo: @"some/dir"]);

    assert([fullPathToCommandToRun isEqualTo: @"/bin/echo"]);
    assert([commandToRun isEqualTo: @"echo"]);
    NSArray *expectedArgs = [NSArray arrayWithObjects:
      @"first arg",
      @"second",
      @"arg",
      nil];
    assert([argumentsToUse isEqualTo: expectedArgs]);
  }

  assert([full_path_for(@"echo") isEqualTo: @"/bin/echo"]);
  assert([full_path_for(@"grep") isEqualTo: @"/usr/bin/grep"]);
  assert([full_path_for(@"./testrunner") isEqualTo: [[[NSFileManager defaultManager] currentDirectoryPath] stringByAppendingPathComponent: @"testrunner"]]);
  assert(full_path_for(@"blaablablabalaba") == nil);

  return 0;
}
