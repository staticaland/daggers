# filestoprompt

## Examples

Run on a local directory:

```sh
dagger call run --source=".." --path="LICENSE" --path="README.md" --cxml
```

Run on a GitHub repository:

```sh
dagger call run --source="https://github.com/staticaland/dotfiles.git" --path="README.md" --cxml
```
