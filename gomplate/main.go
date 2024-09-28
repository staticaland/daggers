// A module for running the gomplate CLI tool
//
// This module provides a function to execute the gomplate CLI tool
// within a Dagger pipeline. gomplate is a template renderer which supports
// various data sources and has many built-in functions.
//
// The Run function executes gomplate with various options, allowing
// for flexible template rendering. It can be called from the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/gomplate/internal/dagger"
	"fmt"
)

type Gomplate struct{}

// Run executes the gomplate CLI command with a single input file and renders to stdout
func (m *Gomplate) Run(
	ctx context.Context,
	// The directory containing the template to process
	source *dagger.Directory,
	// The input template file
	file string,
) (string, error) {
	container := dag.Container().
		From("hairyhenderson/gomplate:alpine")

	args := []string{"gomplate", "--file", file}

	result := container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args)

	output, err := result.Stdout(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to execute gomplate: %w", err)
	}

	return output, nil
}
