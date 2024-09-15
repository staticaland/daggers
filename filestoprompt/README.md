# filestoprompt

Dagger module for [`files-to-prompt`](https://github.com/simonw/files-to-prompt).

## Examples

Run on a local directory:

```sh
dagger call run --source=".." --path="LICENSE" --path="README.md" --cxml
```

Run on a GitHub repository:

```sh
dagger call run --source="https://github.com/staticaland/dotfiles.git" --path="README.md" --cxml
```

Run on a GitHub repository, ignoring certain files:

```sh
dagger call run --source="https://github.com/simonw/files-to-prompt.git" --ignore='!*py'
```

Pipe into `llm`:

Not implemented yet.

Pipe into `strip-tags`:

Not implemented yet.

What about `code2prompt`?

https://github.com/mufeedvh/code2prompt?tab=readme-ov-file#installation
