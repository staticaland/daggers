// A module for running the repopack CLI tool
//
// This module provides a function to execute the repopack CLI tool
// within a Dagger pipeline. repopack is a utility that packs a repository
// or specific files/directories into a single file for easy sharing or analysis.
//
// The Run function executes repopack with various options, allowing
// for flexible processing of repositories, files, and directories. It can be called from
// the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/repopack/internal/dagger"
	"fmt"
)

type Repopack struct{}

// Run executes the repopack CLI command
func (m *Repopack) Run(
	ctx context.Context,
	// The directory containing the repository to process
	source *dagger.Directory,
	// Specify patterns to include (can be used multiple times)
	// +optional
	include []string,
	// Specify patterns to ignore (can be used multiple times)
	// +optional
	ignore []string,
	// The remote repository URL to pack
	// +optional
	remote string,
	// Initialize a new configuration file (repopack.config.json)
	// +optional
	init bool,
	// The output file path (if not specified, output will be returned as a string)
	// +optional
	output string,
	// The path to directory to process. If not specified, the source directory will be used
	// +optional
	path string,
	// The style of the output (e.g., "xml")
	// +optional
	style string,
) (string, error) {
	container := dag.Container().
		From("node:slim").
		WithExec([]string{"npm", "install", "-g", "repopack"})

	args := []string{"repopack"}

	if path != "" {
		args = append(args, path)
	}

	for _, pattern := range include {
		args = append(args, "--include", pattern)
	}
	for _, pattern := range ignore {
		args = append(args, "--ignore", pattern)
	}
	if remote != "" {
		args = append(args, "--remote", remote)
	}
	if init {
		args = append(args, "--init")
	}
	if output != "" {
		args = append(args, "-o", output)
	}
	if style != "" {
		args = append(args, "--style", style)
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
