# repopack

Dagger module for [`repopack`](https://github.com/yamadashy/repopack).

## Examples

Run on a local directory:

```sh
dagger call run --source=".."
```

Run with include and ignore patterns:

```sh
dagger call run --source=".." --include="*.go" --ignore="vendor/*" --output="output.zip"
```

Run on a remote repository:

```sh
dagger call run --remote="https://github.com/renovatebot/renovate" --path="docs/usage"
```

Initialize a new configuration file:

```sh
dagger call run --source=".." --init
```
