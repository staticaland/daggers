// A module for running the code2prompt CLI tool
//
// This module provides a function to execute the code2prompt CLI tool
// within a Dagger pipeline. code2prompt is a utility that processes
// code files and generates prompts for use with Large Language Models (LLMs).
//
// The Run function executes code2prompt with various options, allowing
// for flexible processing of files and directories. It can be called from
// the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/code-2-prompt/internal/dagger"
	"fmt"
)

type Code2Prompt struct{}

// Run executes the code2prompt CLI command
func (m *Code2Prompt) Run(
	ctx context.Context,
	// The directory containing the files to process
	source *dagger.Directory,
	// Use a custom Handlebars template file
	// +optional
	template string,
	// Include files using glob patterns (can be used multiple times)
	// +optional
	include []string,
	// Exclude files using glob patterns (can be used multiple times)
	// +optional
	exclude []string,
	// Exclude files/folders from the source tree based on exclude patterns
	// +optional
	excludeFromTree bool,
	// Display the token count of the generated prompt
	// +optional
	tokens bool,
	// Specify a tokenizer for token count
	// +optional
	encoding string,
	// The output file path (if not specified, output will be returned as a string)
	// +optional
	output string,
	// The path to files to process. If not specified, the source directory will be used
	// +optional
	path []string,
) (string, error) {
	container := dag.Container().
		From("rust").
		WithExec([]string{"cargo", "install", "code2prompt"})

	args := []string{"code2prompt"}

	// Add paths as positional arguments
	if len(path) > 0 {
		args = append(args, path...)
	} else {
		args = append(args, "/src")
	}

	if template != "" {
		args = append(args, "-t", template)
	}
	for _, pattern := range include {
		args = append(args, "--include", pattern)
	}
	for _, pattern := range exclude {
		args = append(args, "--exclude", pattern)
	}
	if excludeFromTree {
		args = append(args, "--exclude-from-tree")
	}
	if tokens {
		args = append(args, "--tokens")
	}
	if encoding != "" {
		args = append(args, "--encoding", encoding)
	}
	if output != "" {
		args = append(args, "--output", output)
	}

	result := container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args)

	if output == "" {
		return result.Stdout(ctx)
	}

	outputFile := result.File(output)
	contents, err := outputFile.Contents(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read output file: %w", err)
	}

	return contents, nil
}
