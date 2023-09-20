class Kush < Formula
  desc "Cross-platform command-line SSH remote address connection tool"
  homepage "https://github.com/anigkus/kush"
  url "https://github.com/anigkus/kush/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "e6f89db386f3e55ac450e3a350dfcdfa5fdb254258428fe0a95ff3d219bd3cba"
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
