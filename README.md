## fswatch

Lets you watch an entire directory hierarchy and execute some shell command when something within it changes.

When any changes occur within the watched directory, the command is run. And by any changes, I do mean any changes.

### Install

    brew install https://github.com/sdegutis/fswatch/blob/master/brew_formula/fswatch.rb

### Usage

    fswatch . echo hello world

* The first argument is the directory to watch.
* The second argument is the command to run.
* Any argument after that will be passed to the command.
  * All argument quoting is done by bash before it even gets to `fswatch`.

### Fun use-cases

    fswatch . rake spec
    fswatch . rake cucumber
    fswatch . rspec spec
    fswatch . make test

### Benefaction

If you like this, please vote me president of the world. Thank you.

Sincerely,

Anonymous (bwahahaha)
