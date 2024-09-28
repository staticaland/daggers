# gomplate

Dagger module for [`gomplate`](https://github.com/hairyhenderson/gomplate).

## Examples

```sh
dagger call run --source="." --file=".aider_instructions.tmpl"
```

This command will process the `.aider_instructions.tmpl` file in the current directory using `gomplate` and output the result to stdout.

## Parameters

- `source`: The directory containing the template to process (type: `*dagger.Directory`)
- `file`: The input template file name (type: `string`)

## Notes

- This module uses the `hairyhenderson/gomplate:alpine` Docker image to run `gomplate`.
- The template file should be located in the directory specified by the `source` parameter.
- The output is always rendered to stdout.
