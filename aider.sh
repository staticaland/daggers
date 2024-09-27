#!/bin/bash

# Set default values if environment variables are not set
SOURCE_FILE=${SOURCE_FILE:-"llm/main.go"}
CLI_NAME=${CLI_NAME:-"markdownlint-cli2"}
TARGET_FILE=${TARGET_FILE:-"markdownlint2cli/main.go"}
USAGE_FILE=${USAGE_FILE:-"markdownlint2cli/USAGE.md"}
DOCKER_IMAGE=${DOCKER_IMAGE:-"davidanson/markdownlint-cli2:latest"}

# Create a temporary file for the gomplate output
TEMP_FILE=$(mktemp)

# Run gomplate with environment variables and save to temp file
gomplate -f - > "$TEMP_FILE" << EOF
{{- \$sourceFile := env.Getenv "SOURCE_FILE" | default "llm/main.go" -}}
{{- \$targetFile := env.Getenv "TARGET_FILE" | default "markdownlint2cli/main.go" -}}
{{- \$dockerImage := env.Getenv "DOCKER_IMAGE" | default "davidanson/markdownlint-cli2:latest" -}}
{{- \$cliName := env.Getenv "CLI_NAME" | default "markdownlint-cli2" -}}
{{- \$usageFile := env.Getenv "USAGE_FILE" | default "USAGE.md" -}}

Modify the file \`{{ \$targetFile }}\` to mirror the patterns found in \`{{ \$sourceFile }}\`. Use \`{{ \$dockerImage }}\` as the Docker image. Ensure that the structure and logic in \`{{ \$targetFile }}\` closely follow the patterns established in \`{{ \$sourceFile }}\`, adapting them as necessary for the {{ \$cliName }} context. When referencing the Docker image in your code, use the exact string "{{ \$dockerImage }}". Review the contents of \`{{ \$usageFile }}\` to ensure that the implementation in \`{{ \$targetFile }}\` accurately reflects the documented CLI usage and functionality of {{ \$cliName }}.
EOF

# Run aider with the generated message and specified files
uvx --from aider-chat aider --message-file "$TEMP_FILE" "$SOURCE_FILE" "$TARGET_FILE"

# Clean up the temporary file
rm "$TEMP_FILE"
