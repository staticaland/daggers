# code2prompt

## Examples

Run on a local directory:

```sh
dagger call run --source=".." --path="LICENSE" --path="README.md"
```

Run on a GitHub repository:

```sh
dagger call run --source="https://github.com/simonw/files-to-prompt.git" --path="README.md"
```
