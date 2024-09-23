# Prompts

This directory contains LLM prompts and tools for various maintenance tasks.

I use Claude for prompts but will switch to CLI or Python code soon.

## Update prompt XML template

```sh
uv run update_xml.py '../..' 'xml/release-please-config.xml'
```

```text
Reading inline script metadata from: update_xml.py
Manifest file updated: xml/release-please-config.xml
Components included: llm, code2prompt, filestoprompt, repopack
```

## Format the Python code

```sh
uvx black update_xml.py
```

## See also

See [`dbohdan/structured-text-tool`](https://github.com/dbohdan/structured-text-tools?tab=readme-ov-file#xml).

Can use Dasel to modify the prompts.
