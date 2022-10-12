# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Proto compatability tool manager
"
  homepage "https://github.com/BitsOfAByte/proto"
  version "0.12.0"
  license "GPL-3.0-only"

  depends_on "gnu-tar"
  depends_on :linux

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.0/proto_0.12.0_linux_armv6.zip"
      sha256 "2e2d0f9e1014d5cc98ab1b5b43ed2affead10b799d7d7734dcebf56bb046f18e"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.0/proto_0.12.0_linux_amd64.zip"
      sha256 "821b63a7a489f9b58d79df4c7eb5e85916304857362907bb7364649820bd35b9"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.0/proto_0.12.0_linux_arm64.zip"
      sha256 "5cd6a821296dca82b4f990a05c1778f6a4396814fc2c16e6aa67c22ebbc03b09"

      def install
        bin.install "proto"
      end
    end
  end
end
