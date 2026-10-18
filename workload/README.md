# 🧩 Workloads

This directory is the template's home for buildable or runnable project
units. A unit might be a command-line tool, worker, library, plugin,
or another independently developed part of a project.

The included repository example currently contains:

- `api-gateway`
- `auth-service`
- `author-service`
- `book-service`
- `category-service`

Those names and their implementation are sample content. When adopting the
template, replace them with the units that make sense for the project rather
than treating this directory as a prescribed architecture.

## 📦 Included Example

Each example unit has its own build metadata, container definition, version,
changelog, environment examples, and Go module. The shared
[workload Makefile](Makefile) can:

- generate certificates for the example units;
- generate their protobuf clients; and
- build, start, or stop their Compose projects.

These targets are implementation details of the included example. A project
using another language or build system should replace them with equivalent
targets, not add language-specific workarounds around them.

### 📄 Unit Files

The current example generally expects each unit to provide:

| File or directory | Role |
| --- | --- |
| `Dockerfile` | Container build instructions |
| `docker-compose.yaml` | Local runtime definition |
| `Makefile` | Unit-specific generation or build targets |
| `VERSION` | Semantic version consumed by release automation |
| `CHANGELOG.md` | Version history maintained by automation |
| `env/.env.example` | Configuration template |
| `go.mod` and `go.sum` | Go module metadata for the example |

Not every project needs every file. Keep only the files required by the
selected language, build system, and packaging model.

## ▶️ Lifecycle

From the repository root:

```bash
make setup-env ENV=dev
make start ENV=dev
make stop ENV=dev
```

The root Makefile starts infrastructure before this directory's units. The
units selected by `WORKLOAD` in [stack.mk](../stack.mk) are processed in
the listed order.

To work on one unit, use its local Makefile. For example, the included
protobuf-based unit exposes:

```bash
make -C workload/auth-service gen-proto
```

The shared lifecycle passes these values into each Compose invocation:

- `WORKLOAD` - the selected unit name
- `ENV` - the environment file suffix
- `SLUG` - the project prefix used for names and networks
- `VERSION` - the value read from the unit's `VERSION` file

The included example generates certificates and protobuf clients before
building containers. A different project can replace those prerequisites with
compilation, packaging, code generation, migrations, or no pre-step at all.

The example requires `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` for
protobuf generation. Install the Go plugins with:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

The workload Makefiles automatically include the Go binary directory reported
by `go env GOPATH` when locating these plugins. If `protoc` itself is missing,
install it with the package manager for your operating system.

## ➕ Adding or Replacing a Unit

1. Create a directory with the build and runtime files required by the unit.
2. Add its name to the appropriate selection in `stack.mk`.
3. Add an environment example if it needs configuration.
4. Add version and changelog files if it participates in the release
   automation.
5. Update the shared Makefile only when the unit needs shared lifecycle work.
6. Add or update workflow path filters if the unit is built or released by CI.

Before committing a new unit, verify that its local Makefile works from both
the unit directory and the repository root. Confirm that its environment
file, version file, container context, and network aliases agree with the
corresponding Compose definition.

Keep generated files and local credentials out of commits. The repository
hooks and workflows currently assume one independently versioned unit per
change; adjust those rules if that does not match the project.
