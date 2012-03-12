#import "fswatch.h"
#import <assert.h>
#import <Foundation/Foundation.h>

void testTooFewArgsWithoutForce() {
  WatchOptions options = split_out_cmd_args(2, (char*[]){
      "this_executable",
      "some/dir",
      });

  assert(options.notEnoughArgs == YES);
}

void testEnoughArgsWithoutForce() {
  WatchOptions options = split_out_cmd_args(3, (char*[]){
      "this_executable",
      "some/dir",
      "pwd",
      });

  assert(options.notEnoughArgs == NO);
}

void testTooFewArgsWithForce() {
  WatchOptions options = split_out_cmd_args(3, (char*[]){
      "this_executable",
      "-f",
      "some/dir",
      });

  assert(options.notEnoughArgs == YES);
}

void testAllVarsGoodWithForce() {
  WatchOptions options = split_out_cmd_args(7, (char*[]){
      "this_executable",
      "-f",
      "some/dir",
      "echo",
      "first arg",
      "second",
      "arg",
      });

  assert(options.notEnoughArgs == NO);

  assert(options.forceFirstRun == YES);
  assert([options.dirToWatch isEqualTo: @"some/dir"]);

  assert([options.fullPathToCommandToRun isEqualTo: @"/bin/echo"]);
  assert([options.commandToRun isEqualTo: @"echo"]);
  NSArray *expectedArgs = [NSArray arrayWithObjects:
    @"first arg",
    @"second",
    @"arg",
    nil];
  assert([options.argumentsToUse isEqualTo: expectedArgs]);
}

void testAllVarsGoodWithoutForce() {
  WatchOptions options = split_out_cmd_args(6, (char*[]){
      "this_executable",
      "some/dir",
      "echo",
      "first arg",
      "second",
      "arg",
      });

  assert(options.notEnoughArgs == NO);

  assert(options.forceFirstRun == NO);
  assert([options.dirToWatch isEqualTo: @"some/dir"]);

  assert([options.fullPathToCommandToRun isEqualTo: @"/bin/echo"]);
  assert([options.commandToRun isEqualTo: @"echo"]);
  NSArray *expectedArgs = [NSArray arrayWithObjects:
    @"first arg",
    @"second",
    @"arg",
    nil];
  assert([options.argumentsToUse isEqualTo: expectedArgs]);
}

void testFullPathFor() {
  assert([full_path_for(@"echo") isEqualTo: @"/bin/echo"]);
  assert([full_path_for(@"grep") isEqualTo: @"/usr/bin/grep"]);
  assert([full_path_for(@"./testrunner") isEqualTo: [[[NSFileManager defaultManager] currentDirectoryPath] stringByAppendingPathComponent: @"testrunner"]]);
  assert(full_path_for(@"blaablablabalaba") == nil);
}

int main(int argc, char** argv) {
  [NSAutoreleasePool new];

  testTooFewArgsWithoutForce();
  testTooFewArgsWithForce();

  testEnoughArgsWithoutForce();

  testAllVarsGoodWithForce();
  testAllVarsGoodWithoutForce();

  testFullPathFor();

  return 0;
}
