# repopack

Dagger module for [`repopack`](https://github.com/mufeedvh/repopack).

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
dagger call run --remote="https://github.com/example/repo.git" --output="repo_packed.zip"
```

Initialize a new configuration file:

```sh
dagger call run --source=".." --init
```
