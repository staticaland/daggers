// A module for running the markdownlint-cli2 tool
//
// This module provides a function to execute the markdownlint-cli2 tool
// within a Dagger pipeline. markdownlint-cli2 is a command-line interface
// for the markdownlint library, which checks markdown files for style issues.

package main

import (
	"context"
	"fmt"
	"strings"

	"dagger/markdownlintcli-2/internal/dagger"
)

type Markdownlintcli2 struct{}

// Run executes the markdownlint-cli2 command
func (m *Markdownlintcli2) Run(
	ctx context.Context,
	// The directory containing the files to lint
	source *dagger.Directory,
	// Glob expressions for files to lint
	globs []string,
	// Path to a configuration file
	// +optional
	config string,
	// Fix issues if possible
	// +optional
	fix bool,
	// Ignore the "globs" property in the top-level options object
	// +optional
	noGlobs bool,
) (string, error) {
	container := dag.Container().From("davidanson/markdownlint-cli2:latest")

	args := []string{"markdownlint-cli2"}

	if len(globs) > 0 {
		args = append(args, globs...)
	} else {
		args = append(args, "/src")
	}

	if config != "" {
		args = append(args, "--config", config)
	}
	if fix {
		args = append(args, "--fix")
	}
	if noGlobs {
		args = append(args, "--no-globs")
	}

	result := container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args)

	output, err := result.Stdout(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to execute markdownlint-cli2 command: %w", err)
	}

	return strings.TrimSpace(output), nil
}
