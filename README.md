# crane-plugin-xks-to-ocp

Skeleton of a [Crane](https://github.com/konveyor/crane) transform plugin for
migrations from XKS to OpenShift.

The plugin currently passes every resource through unchanged. It returns no
patches, whiteouts, or new resources and declares no optional flags.

## Build

Requires Go 1.24 or later.

```sh
make build
```

This creates `./crane-plugin-xks-to-ocp`. The generated binary is ignored by Git.
The `crane-plugin-` prefix allows Crane to discover the binary in its plugin
directory.

## Run

The plugin uses JSON on stdin and stdout. Crane invokes it once per resource.

Request plugin metadata:

```sh
echo '{}' | ./crane-plugin-xks-to-ocp
```

Submit a resource:

```sh
echo '{"apiVersion":"v1","kind":"ConfigMap","metadata":{"name":"example"}}' | ./crane-plugin-xks-to-ocp
```

The transform response is `{}`, which tells Crane to keep the resource unchanged.

To use it with Crane, place the binary in a plugin directory and pass that
directory to `crane transform --plugin-dir /path/to/plugins`.

## Development

- `main.go` connects the plugin to the Crane CLI protocol handler.
- `xks-to-ocp/plugin.go` defines metadata and the transform entry point.
- `go.mod` and `go.sum` pin dependencies.

Run the Go checks with `make test`.

See the [Crane transform plugin guide](https://github.com/migtools/crane-plugins/blob/main/docs/creating-crane-transform-plugin.md)
for the protocol and examples of future transforms.
