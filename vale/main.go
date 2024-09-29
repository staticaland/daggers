// A module for running the Vale CLI tool
//
// This module provides a function to execute the Vale CLI tool
// within a Dagger pipeline. Vale is a syntax-aware linter for prose
// that supports various markup formats.
//
// The Run function executes Vale with various options, allowing
// for flexible linting of documentation. It can be called from the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/vale/internal/dagger"
)

type Vale struct{}

// Run executes the Vale CLI command with the provided options
func (m *Vale) Run(
	ctx context.Context,
	// The directory containing the files to lint
	source *dagger.Directory,
	// +optional
	config string,
	// +optional
	ext string,
	// +optional
	filter string,
	// +optional
	glob string,
	// +optional
	ignoreSyntax bool,
	// +optional
	noExit bool,
	// +optional
	noWrap bool,
	// +optional
	output string,
) (string, error) {
	container := dag.Container().
		From("jdkato/vale")

	args := []string{"vale"}

	if config != "" {
		args = append(args, "--config", config)
	}
	if ext != "" {
		args = append(args, "--ext", ext)
	}
	if filter != "" {
		args = append(args, "--filter", filter)
	}
	if glob != "" {
		args = append(args, "--glob", glob)
	}
	if ignoreSyntax {
		args = append(args, "--ignore-syntax")
	}
	if noExit {
		args = append(args, "--no-exit")
	}
	if noWrap {
		args = append(args, "--no-wrap")
	}
	if output != "" {
		args = append(args, "--output", output)
	}

	args = append(args, ".")

	return container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args).
		Stdout(ctx)
}
