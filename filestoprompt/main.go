// A module for running the files-to-prompt CLI tool
//
// This module provides a function to execute the files-to-prompt CLI tool
// within a Dagger pipeline. files-to-prompt is a utility that concatenates
// a directory full of files into a single prompt for use with Large Language
// Models (LLMs).
//
// The Run function executes files-to-prompt with various options, allowing
// for flexible processing of files and directories. It can be called from
// the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/filestoprompt/internal/dagger"
	"fmt"
)

type Filestoprompt struct{}

// Run executes the files-to-prompt CLI command
func (m *Filestoprompt) Run(
    ctx context.Context,
    // The directory containing the files to process
    source *dagger.Directory,
    // Include files and folders starting with . (hidden files and directories)
    // +optional
    includeHidden bool,
    // Ignore .gitignore files and include all files
    // +optional
    ignoreGitignore bool,
    // Specify patterns to ignore (can be used multiple times)
    // +optional
    ignore []string,
    // Output in Claude XML format
    // +optional
    cxml bool,
    // The output file path (if not specified, output will be returned as a string)
    // +optional
    output string,
    // The path to files to process. If not specified, the source directory will be used
    // +optional
    path []string,
) (string, error) {
    container := dag.Container().
        From("python:slim").
        WithExec([]string{"pip", "install", "files-to-prompt"})

    args := []string{"files-to-prompt"}

    // Add paths as positional arguments
    if len(path) > 0 {
    	args = append(args, path...)
    } else {
    	args = append(args, "/src")
    }

    if includeHidden {
        args = append(args, "--include-hidden")
    }
    if ignoreGitignore {
        args = append(args, "--ignore-gitignore")
    }
    for _, pattern := range ignore {
        args = append(args, "--ignore", pattern)
    }
    if cxml {
        args = append(args, "--cxml")
    }
    if output != "" {
        args = append(args, "-o", output)
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
