## fswatch

`fswatch` watches the given paths for file-system events. When any such event occurs,
it will run the given command.

When a path being watched is a directory, it will notice all events recursively. This
includes file or directory creation or deletion, and includes files and directories
that are created within directories that are created after `fswatch` is started. In
other words, it sees *all* file-system events recursively.

It does this using the native FSEvents API on Mac OS X. That means it's efficient and
doesn't stress your system with polling. Also it probably requires at least Lion (10.7).
But it's not a gem or node module, so it doesn't require anything other than Go.

### Install

    $ go get github.com/sdegutis/fswatch

### Usage

    $ fswatch
    usage: fswatch [options] path [...] -- cmd [arg ...]
      -f = run command initially
      -v = show version

    $ fswatch . -- echo hello world

### Fun use-cases

    $ fswatch spec app lib -- rspec spec
    $ fswatch . -- go test

### Benefaction

If you like this, please vote me president of the world. Thank you.

Or even better, click here and endorse my mad skillz:

[![endorse](http://api.coderwall.com/sdegutis/endorse.png)](http://coderwall.com/sdegutis)

Sincerely,

Anonymous
