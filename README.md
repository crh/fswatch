## fswatch

When `fswatch some_dir echo hello world` is run, print "hello world" any time anything changes within `some_dir`.
And when I say anything, I mean anything. Here are some examples, to clarify:

* Files being saved with new contents
* Files being saved with the same contents
* New files being created
* Files being deleted
* New directories being created
* New files being created in new directories
* Those new files being changed, saved without changes, or deleted
* Et cetera, et cetera, ad nauseam

Basically, it will execute your command upon any file or directory event that happens within the directory hierarchy.

### Install

    brew install https://raw.github.com/sdegutis/fswatch/master/brew_formula/fswatch.rb

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

Anonymous
