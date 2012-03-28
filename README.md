## fswatch

`fswatch` watches the given paths for file-system events. When any such event occurs,
it will run the given command.

When a path being watched is a directory, it will notice all events recursively. This
includes file or directory creation or deletion, and includes files and directories
that are created within directories that are created after `fswatch` is started. In
other words, it sees *all* file-system events recursively.

### Install

    $ go get github.com/sdegutis/fswatch-go

### Usage

    $ fswatch
    usage: fswatch-go [options] <path> [...] -- <cmd> [arg ...]
      -f = run command initially
      -v = show version

    fswatch . -- echo hello world

### Fun use-cases

    fswatch spec app lib -- rspec spec

    fswatch . -- go test

### Benefaction

If you like this, please vote me president of the world. Thank you.

Or even better, click here and endorse my mad skillz:

[![endorse](http://api.coderwall.com/sdegutis/endorse.png)](http://coderwall.com/sdegutis)

Sincerely,

Anonymous
