# 🧱 Project Template

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/rodericusifo/api-template/actions)
[![Docker](https://img.shields.io/badge/docker-ready-blue)](https://www.docker.com/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A practical starting point for projects that need a consistent repository
layout, repeatable local setup, and automated delivery.

The repository provides the surrounding structure rather than prescribing what
the project must build. The checked-in workloads and Docker Compose files form
one working example; they can be replaced without changing the repository's
core conventions.

## 🧰 What This Provides

- A predictable place for workloads, supporting resources, and maintenance
  scripts
- Make-based entry points for local setup and lifecycle commands
- Environment-file bootstrapping from checked-in examples
- Optional Docker Compose orchestration for local dependencies
- Git hooks and GitHub Actions for validation, versioning, and releases

## 📚 Guides

Read the guide closest to the change you are making:

- [Workloads](workload/README.md)
- [Infrastructure](infrastructure/README.md)
- [Repository scripts](script/README.md)
- [GitHub Actions](.github/workflows/README.md)

## 🗂️ Repository Layout

```text
.
├── workload/              Buildable or runnable project units
├── infrastructure/        Optional local dependencies
├── script/                Setup scripts and Git hooks
├── .github/workflows/     Validation and delivery automation
├── Makefile               Repository-level commands
└── stack.mk               Default workload and dependency selections
```

The included workload and infrastructure directories are an executable
example. The repository-level conventions are the reusable part of the
template.

## 🚀 Getting Started

The commands below operate the repository's included example. Adapt
`stack.mk`, the workload directories, and the infrastructure definitions when
using the template for another project.

### ✅ Requirements

- [Docker](https://docs.docker.com/get-docker/) and Docker Compose
- `make`
- Tools required by the workloads you keep or add

### ⚙️ Setup

1. Review the selections in [stack.mk](stack.mk).
2. Generate local configuration files:

   ```bash
   make setup-env ENV=dev
   ```

3. Fill in the generated values. Do not commit credentials or generated
   environment files.
4. Install the repository hooks:

   ```bash
   make setup-hooks
   ```

5. Start and stop the included example:

   ```bash
   make start ENV=dev
   make stop ENV=dev
   ```

Some included components require additional tools or credentials. Their
individual guides document those requirements.

### 🧭 Common Commands

| Command | Purpose |
| --- | --- |
| `make setup-env ENV=dev` | Create local environment files from examples |
| `make setup-hooks` | Install repository Git hooks |
| `make start ENV=dev` | Create the network and start dependencies and workloads |
| `make stop ENV=dev` | Stop Compose projects while preserving volumes |
| `make -n start ENV=dev` | Preview the commands without executing them |

`ENV` is required by lifecycle commands. Use a separate value for each local
environment when you need isolated networks, Compose configuration, and
volumes.

## 🧩 Adapting the Template

Keep the repository-level Make targets if they fit your workflow, then replace
the example workload and dependency definitions with the ones your project
needs. Update the hook and workflow assumptions when your component layout,
versioning policy, build system, or release process differs from the example.

At minimum, review these files when creating a new project from the template:

1. `stack.mk` for the units and dependencies to start.
2. `workload/Makefile` for build, generation, and container assumptions.
3. `infrastructure/*/Makefile` for supported dependency values and volumes.
4. `script/hooks/*` for commit and versioning rules.
5. `.github/workflows/*.yml` for branches, paths, registries, and secrets.

Remove example-only tooling when it is not relevant. Keeping unused services,
credentials, or publishing workflows enabled creates unnecessary maintenance
and security risk.

## 🤝 Contributing

Use a feature branch, keep changes focused, run the relevant checks, and
follow the repository's commit and release conventions. See [NOTICE](NOTICE)
and the [MIT License](LICENSE).
