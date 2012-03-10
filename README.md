#### fswatch

Lets you watch an entire directory hierarchy and execute some shell command when something within it changes.

Some examples:

  fswatch . echo hello world

In these examples, any time *any* change happens, to a file or directory, within the current directory or *any* of its subdirectories, the given command is run.

More examples:

  fswatch . rake spec
  fswatch . rake cucumber
  fswatch . rspec spec
  fswatch . make test

If you like this, please vote me president of the world. Thank you.

Sincerely,
Anonymous (bwahahaha)
