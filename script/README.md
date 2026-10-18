# 🛠️ Repository Scripts

The scripts in this directory handle setup that should be consistent for
contributors, regardless of which workload they are editing.

## 🔐 Environment Files

`setup-env.sh` searches the repository for `.env.example` files and creates a
copy beside each one using the requested environment name:

```bash
make setup-env ENV=dev
# or
bash ./script/setup-env.sh dev
```

Environment names may contain letters, numbers, dashes, and underscores.
Generated files are for local use. Review them, populate required values, and
keep credentials out of Git.

The script is intentionally repository-wide: adding a new `.env.example`
automatically includes that location in the next setup run. It overwrites an
existing target file, so preserve local changes or use a separate environment
name before running it again.

The root Makefile runs this script from the repository root. Running it from a
different directory can produce incomplete results because the search is
relative to the current working directory.

## 🪝 Git Hooks

Install the repository hooks after cloning:

```bash
make setup-hooks
```

`setup-hooks.sh` installs the hooks from `script/hooks` into `.git/hooks` and
backs up existing non-symlink hooks first. Run it from the repository root.

The installed hooks currently provide:

- `pre-commit`: enforces one workload per commit and validates the commit
  boundary used by the example;
- `commit-msg`: checks the conventional commit format and supported commit
  types;
- `pre-push`: validates workload version changes and changelog expectations
  before pushing.

The hooks are local copies. Re-run `make setup-hooks` after changing the
source files or after cloning the repository on another machine.

The included hooks enforce the example repository's commit format, change
boundaries, version rules, and changelog behavior. If the template is adapted
to a different layout or release process, update the hooks rather than
silently relying on rules that no longer apply.

### 🧹 Hook Backups and Troubleshooting

If a non-symlink hook already exists, the installer copies existing hooks to a
timestamped `.git/hooks.backup.*` directory before installing the repository
versions. Check that directory if a custom local hook needs to be restored.

If a hook rejects a commit or push, read the displayed rule first. Typical
fixes are splitting changes by workload, using a conventional commit message,
updating the relevant `VERSION`, or adding the required changelog entry.

See the [workflow guide](../.github/workflows/README.md) for the corresponding
remote checks.
