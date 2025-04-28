# k8s-metrics-repro

Minimal reproduction of Bazel build errors when using `k8s.io/metrics` with Gazelle.

## Error

When running `bazel build //...`, you get errors like:

```
ERROR: no such package '@@gazelle++go_deps+io_k8s_metrics//k8s.io/apimachinery/pkg/runtime': BUILD file not found in directory 'k8s.io/apimachinery/pkg/runtime'
```

## Fix

Uncomment these lines in `MODULE.bazel`:

```bazel
go_deps.gazelle_default_attributes(
    directives = [
        "gazelle:go_generate_proto false",
    ],
)
```
