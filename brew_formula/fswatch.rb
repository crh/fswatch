require 'formula'

class Fswatch < Formula
  homepage 'https://github.com/sdegutis/fswatch'

  url 'https://github.com/sdegutis/fswatch/tarball/v2.1'
  md5 '2c9b0bc0b6ade5cd3c302d5743c0ade1'

  head 'https://github.com/sdegutis/fswatch.git'

  def install
    system "make"
    bin.install "fswatch"
  end
end
