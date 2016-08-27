load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(
    default_visibility = ["//:__subpackages__"],
)

go_binary(
    name = "primebench",
    srcs = ["primebench.go"],
    deps = [
        "//:go_default_library",
        "//examples/isprime:go_default_library",
    ],
)
