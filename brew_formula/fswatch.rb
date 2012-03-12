require 'formula'

class Fswatch < Formula
  homepage 'https://github.com/sdegutis/fswatch'

  url 'https://github.com/sdegutis/fswatch/tarball/v2.1'
  md5 'ad9c4bd3fdc7884aef67bc78c77d8d40'

  head 'https://github.com/sdegutis/fswatch.git'

  def install
    system "make"
    bin.install "fswatch"
  end
end
