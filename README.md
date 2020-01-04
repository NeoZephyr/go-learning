bazel query //...
bazel run --action_env=GOPROXY=https://goproxy.io //:gazelle

bazel run --action_env=GOPROXY=https://goproxy.io //basic:basic
bazel test --action_env=GOPROXY=https://goproxy.io //basic:go_default_test --test_output=all

bazel run //:gazelle -- update-repos github.com/segmentio/kafka-g
bazel run //:gazelle -- update-repos -from_file=go.mod
