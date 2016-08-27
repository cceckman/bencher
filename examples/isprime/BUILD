load("@io_bazel_rules_go//go:def.bzl", "go_test", "go_library")

package(
    default_visibility = ["//:__subpackages__"],
)

go_library(
    name = "go_default_library",
    srcs = ["isprime.go"],
)

go_test(
    name = "isprime_test",
    srcs = ["isprime_test.go"],
    library = ":go_default_library",
)
