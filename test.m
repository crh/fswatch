#import "fswatch.h"
#import <assert.h>
#import <Foundation/Foundation.h>

int main(int argc, char** argv) {
  [NSAutoreleasePool new];

  split_out_cmd_args(6, (char*[]){
    "this_executable",
      "some/dir",
      "echo",
      "first arg",
      "second",
      "arg",
  });

  assert([dirToWatch isEqualTo: @"some/dir"]);

  assert([fullPathToCommandToRun isEqualTo: @"/bin/echo"]);
  assert([commandToRun isEqualTo: @"echo"]);
  NSArray *expectedArgs = [NSArray arrayWithObjects:
    @"first arg",
    @"second",
    @"arg",
    nil];
  assert([argumentsToUse isEqualTo: expectedArgs]);

  assert([full_path_for(@"echo") isEqualTo: @"/bin/echo"]);
  assert([full_path_for(@"grep") isEqualTo: @"/usr/bin/grep"]);
  assert([full_path_for(@"./testrunner") isEqualTo: [[[NSFileManager defaultManager] currentDirectoryPath] stringByAppendingPathComponent: @"testrunner"]]);
  assert(full_path_for(@"blaablablabalaba") == nil);

  return 0;
}
