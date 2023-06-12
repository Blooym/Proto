# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Proto compatability tool manager
"
  homepage "https://github.com/Blooym/proto"
  version "1.1.2"
  license "GPL-3.0-only"

  depends_on "gnu-tar"
  depends_on :linux

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/Blooym/proto/releases/download/v1.1.2/proto_1.1.2_linux_amd64.zip"
      sha256 "2f1d20b25182c27b1bdb0196aa06217e6ab73727b3fd2c6285e738cde0f53271"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/Blooym/proto/releases/download/v1.1.2/proto_1.1.2_linux_armv6.zip"
      sha256 "800fc9da370cd2197852b4ef9001a36d01d8eef0b5bae99887e4a53100c449e8"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/Blooym/proto/releases/download/v1.1.2/proto_1.1.2_linux_arm64.zip"
      sha256 "3e251708da5779d5d6c8781fd0908329b2cdf24c19aaa4cf01bcb44ce0527cd8"

      def install
        bin.install "proto"
      end
    end
  end
end
