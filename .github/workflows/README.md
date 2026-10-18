# ⚙️ GitHub Actions

The workflows in this directory turn the repository's local conventions into
remote checks and delivery steps. They are configured for the included
example, so review their paths, branch names, secrets, and build commands
before reusing them in another project.

## 📋 Responsibilities

- `pull-request-validator.yml` checks pull requests targeting `prod-v1`,
  including workload structure, version changes, changelogs, change
  boundaries, Go linting, and Go unit tests.
- `auto-lint.yml` is a reusable linting workflow with one job per configured
  language. It is called by the pull-request validator.
- `auto-unit-tests.yml` is a reusable unit-testing workflow with one job per
  language. It is called by the pull-request validator.
- `auto-create-or-update-changelogs.yml` maintains workload changelogs after
  `VERSION` files change on `prod-v1`.
- `auto-create-tags.yml` runs after the changelog workflow succeeds and creates
  version tags for changed workloads.
- `push-and-release-coordinator.yml` runs after tags are created. It calls
  `auto-push.yml` once for each configured registry and then calls
  `auto-create-releases.yml` after all registry jobs succeed.
- `auto-push.yml` is a reusable publishing workflow that builds each tagged
  workload image and publishes it to either Docker Hub or GHCR.
- `manual-push.yml` is a manually triggered publishing workflow that accepts a
  registry and one workload tag.
- `auto-create-releases.yml` is a reusable workflow that creates the GitHub
  release and records the resulting versions and image references.

## 🧪 Pull Request Validation

`pull-request-validator.yml` is configured for pull requests targeting
`prod-v1`. It checks the example's workload conventions, including:

- changed workloads are represented in `WORKLOAD` in `stack.mk`;
- workload changes are accompanied by the expected `VERSION` updates;
- version values follow semantic versioning and allowed increment rules;
- required `CHANGELOG.md`, `VERSION`, and container files are present; and
- a change does not unintentionally span multiple workloads;
- changed workloads pass the configured language linter; and
- changed workloads pass the configured unit-test command.

The lint and unit-test checks are implemented as generic reusable workflows
rather than duplicated inside the validator. Their language and marker file
are configured by the
`LINTING_WORKFLOWS` and `UNIT_TESTING_WORKFLOWS` environment variables in
`pull-request-validator.yml`. For example, a `go.mod` marker identifies a
Go workload. The validator passes the configuration and changed workload list
to the generic workflows, which create separate jobs for each configured
language implementation.

To add another language, add an object to both configuration arrays with its
marker file and implementation name, then add that implementation's command
to the corresponding reusable workflow.

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

The push-and-release coordinator runs after tags are created, calls the
generic registry publishing workflow through a Docker Hub/GHCR matrix, and
creates a GitHub release only after every registry entry succeeds. Adding a
registry requires adding it to the coordinator matrix and to the registry
validation, login, and image-publishing logic in `auto-push.yml`.

### 🏷️ Version and Tag Conventions

The included automation reads each selected workload's `VERSION` file and
creates tags in this form:

```text
<workload-name>/v<major>.<minor>.<patch>
```

Changelog and release workflows use those tags to identify changed workloads
and summarize version transitions.

### 📦 Container Publishing

The reusable publishing workflow builds each tagged workload from:

```text
workload/<workload-name>/Dockerfile
```

It expects registry credentials and repository permissions to be configured in
GitHub. Docker Hub uses `DOCKER_USERNAME` and `DOCKER_PASSWORD`; GHCR uses
the workflow's `GITHUB_TOKEN`. Remove these workflows if the project does not
publish containers, or replace their build steps for another packaging system.

To republish one existing workload tag manually, open **Actions → Manual Push
→ Run workflow**, select `docker-hub` or `ghcr`, and enter a tag such as:

```text
api-gateway/v1.0.3
```

Manual publishing builds the selected workload from the workflow's current
revision and publishes both the version tag and `latest`. It does not create a
Git tag, changelog entry, or GitHub release.

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
