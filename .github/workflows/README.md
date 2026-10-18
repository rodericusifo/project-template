# ⚙️ GitHub Actions

The workflows in this directory turn the repository's local conventions into
remote checks and delivery steps. They are configured for the included
example, so review their paths, branch names, secrets, and build commands
before reusing them in another project.

## 📋 Responsibilities

- `pull-request-validator.yml` checks pull requests targeting `prod-v1`,
  including required metadata, version changes, changelogs, and change
  boundaries.
- `auto-create-or-update-changelogs.yml` maintains changelogs after changes
  reach `prod-v1`.
- `auto-create-tags.yml` creates version tags for changed units.
- `auto-docker-hub-push.yml` publishes built images to Docker Hub.
- `auto-ghcr-push.yml` publishes built images to GitHub Container Registry.
- `release-coordinator.yml` waits for both image publishing jobs.
- `auto-create-releases.yml` creates the GitHub release and records the
  resulting versions and image references.

## 🧪 Pull Request Validation

`pull-request-validator.yml` is configured for pull requests targeting
`prod-v1`. It checks the example's workload conventions, including:

- changed workloads are represented in `WORKLOAD` in `stack.mk`;
- workload changes are accompanied by the expected `VERSION` updates;
- version values follow semantic versioning and allowed increment rules;
- required `CHANGELOG.md`, `VERSION`, and container files are present; and
- a change does not unintentionally span multiple workloads.

The local equivalents are installed by `make setup-hooks`, but local checks
should not be treated as a replacement for the required GitHub checks.

## 🔄 Release Flow

For the included example, the delivery chain is:

```text
pull request
  -> validation
merge to prod-v1
  -> changelog update
  -> version tag
  -> image publishing
  -> release coordination
  -> GitHub release
```

The two registry publishing workflows run in parallel after tags are created.
The coordinator prevents release creation until both have completed.

### 🏷️ Version and Tag Conventions

The included automation reads each selected workload's `VERSION` file and
creates tags in this form:

```text
<workload-name>/v<major>.<minor>.<patch>
```

Changelog and release workflows use those tags to identify changed workloads
and summarize version transitions.

### 📦 Container Publishing

The Docker Hub and GHCR workflows build each tagged workload from:

```text
workload/<workload-name>/Dockerfile
```

They expect registry credentials and repository permissions to be configured
in GitHub. Remove these workflows if the project does not publish containers,
or replace their build steps for another packaging system.

## ♻️ Reusing the Workflows

Before enabling these workflows for another project:

1. Change the target branches and path filters.
2. Replace unit names, version-file locations, and changelog assumptions.
3. Align build and packaging steps with the project's toolchain.
4. Configure registry credentials and other Actions secrets.
5. Remove publishing or release workflows that the project does not need.

Keep secrets in GitHub Actions settings. The local equivalents of the
validation rules are installed by [the repository scripts](../../script/README.md).

### 🔒 Required Repository Settings

Before enabling the delivery chain, verify:

- the target branch name matches the branch filters;
- Actions has permission to create tags, commits, and releases;
- Docker Hub and GHCR credentials are available to the publishing jobs;
- the repository allows the required workflow token permissions; and
- workload `VERSION` and `CHANGELOG.md` files follow the expected layout.

Test changes with pull request validation before enabling automated publishing.
