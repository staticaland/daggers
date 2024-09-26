# markdownlintcli2

Dagger module for [`markdownlint-cli2`](https://github.com/DavidAnson/markdownlint-cli2).

## Examples

Run on a local directory:

```sh
dagger call run --source=".." --globs="."
```

Run on a GitHub repository:

```sh
dagger call run --source="https://github.com/renovatebot/renovate" --globs="."
```
