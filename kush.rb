class Kush < Formula
  desc "Cross-platform command-line SSH remote address connection tool"
  homepage "https://github.com/anigkus/kush"
  url "https://github.com/anigkus/kush/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "cce811fa7ff33874ba63bea4395f7847eb67a3ce7ebc423f0ff446a5aaa64570"
  license "Apache-2.0"
  head "https://github.com/anigkus/kush.git", branch: "main"
  depends_on "go" => :build
  def install
    system "go", "build", *std_go_args
    generate_completions_from_executable(bin/"kush", "completion")
  end
  test do
    output = shell_output("#{bin}/kush --version")
    assert_match "Version:    #{version}", output
  end
end
