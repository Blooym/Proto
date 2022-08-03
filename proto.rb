# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Proto compatability tool manager
"
  homepage "https://github.com/BitsOfAByte/proto"
  version "0.10.0"
  license "GPL-3.0-only"
  depends_on :linux

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.10.0/proto_linux_arm.zip"
      sha256 "12489841023b6b1079d9f4c4a6ef7ae099aefad8758ba2bc0202cea159d1e97c"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.10.0/proto_linux_arm64.zip"
      sha256 "bbc78335811af8ac1e731cbe71eea819470cb685182611e13e7e2a7e0a9d742c"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.10.0/proto_linux_amd64.zip"
      sha256 "ee828c1e38f444db3a904682ed0db46939982ef46f9f25960910eec700634f3b"

      def install
        bin.install "proto"
      end
    end
  end

  depends_on "gnu-tar"
end