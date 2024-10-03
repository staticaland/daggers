// A module for running Fabric commands
//
// This module provides a function to execute Fabric commands
// within a Dagger pipeline. Fabric is an open-source framework
// for augmenting humans using AI. It provides a modular framework
// for solving specific problems using a crowdsourced set of AI prompts
// that can be used anywhere.
//
// The Run function executes Fabric commands with various options, allowing
// for flexible execution of tasks on remote servers. It can be called from
// the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"dagger/fabric/internal/dagger"
	"fmt"
)

type Fabric struct{}

// Run executes the Fabric command
func (m *Fabric) Run(
	ctx context.Context,
	// The directory containing the Fabric configuration
	configDir *dagger.Directory,
	// OpenAI API key as a secret
	// +optional
	openAiApiKey *dagger.Secret,
	// Choose a pattern
	// +optional
	pattern string,
	// Values for pattern variables, e.g. name:John age:30
	// +optional
	variables []string,
	// Choose a context
	// +optional
	context string,
	// Choose a session
	// +optional
	session string,
	// Run setup
	// +optional
	setup bool,
	// Skip update patterns at setup
	// +optional
	setupSkipUpdatePatterns bool,
	// Stream
	// +optional
	stream bool,
	// Use the defaults of the model without sending chat options
	// +optional
	raw bool,
	// List all patterns
	// +optional
	listPatterns bool,
	// List all available models
	// +optional
	listModels bool,
	// List all contexts
	// +optional
	listContexts bool,
	// List all sessions
	// +optional
	listSessions bool,
	// Update patterns
	// +optional
	updatePatterns bool,
	// Copy to clipboard
	// +optional
	copy bool,
	// Choose model
	// +optional
	model string,
	// Output to file
	// +optional
	output string,
	// Number of latest patterns to list (default: 0)
	// +optional
	latest int,
	// Change default model
	// +optional
	changeDefaultModel bool,
	// YouTube video URL to grab transcript, comments from it and send to chat
	// +optional
	youtube string,
	// Grab transcript from YouTube video and send to chat (used by default)
	// +optional
	transcript bool,
	// Grab comments from YouTube video and send to chat
	// +optional
	comments bool,
	// Specify the Language Code for the chat, e.g. en, zh
	// +optional
	language string,
	// Scrape website URL to markdown using Jina AI
	// +optional
	scrapeURL string,
	// Search question using Jina AI
	// +optional
	scrapeQuestion string,
	// Seed to be used for LMM generation
	// +optional
	seed int,
	// Wipe context
	// +optional
	wipeContext bool,
	// Wipe session
	// +optional
	wipeSession bool,
	// Show what would be sent to the model without actually sending it
	// +optional
	dryRun bool,
) (string, error) {
	container := dag.Container().
		From("golang:alpine").
		WithExec([]string{"go", "install", "github.com/danielmiessler/fabric@latest"})

	args := []string{"fabric"}

	if pattern != "" {
		args = append(args, "-p", pattern)
	}
	for _, v := range variables {
		args = append(args, "-v", v)
	}
	if context != "" {
		args = append(args, "-C", context)
	}
	if session != "" {
		args = append(args, "--session", session)
	}
	if setup {
		args = append(args, "-S")
	}
	if setupSkipUpdatePatterns {
		args = append(args, "--setup-skip-update-patterns")
	}
	if stream {
		args = append(args, "-s")
	}
	if raw {
		args = append(args, "-r")
	}
	if listPatterns {
		args = append(args, "-l")
	}
	if listModels {
		args = append(args, "-L")
	}
	if listContexts {
		args = append(args, "-x")
	}
	if listSessions {
		args = append(args, "-X")
	}
	if updatePatterns {
		args = append(args, "-U")
	}
	if copy {
		args = append(args, "-c")
	}
	if model != "" {
		args = append(args, "-m", model)
	}
	if output != "" {
		args = append(args, "-o", output)
	}
	if latest != 0 {
		args = append(args, "-n", fmt.Sprintf("%d", latest))
	}
	if changeDefaultModel {
		args = append(args, "-d")
	}
	if youtube != "" {
		args = append(args, "-y", youtube)
	}
	if transcript {
		args = append(args, "--transcript")
	}
	if comments {
		args = append(args, "--comments")
	}
	if language != "" {
		args = append(args, "-g", language)
	}
	if scrapeURL != "" {
		args = append(args, "-u", scrapeURL)
	}
	if scrapeQuestion != "" {
		args = append(args, "-q", scrapeQuestion)
	}
	if seed != 0 {
		args = append(args, "-e", fmt.Sprintf("%d", seed))
	}
	if wipeContext {
		args = append(args, "-w")
	}
	if wipeSession {
		args = append(args, "-W")
	}
	if dryRun {
		args = append(args, "--dry-run")
	}

	result := container.
		WithDirectory("/root/.config/fabric", configDir).
		WithWorkdir("/src")

	result = result.WithExec(args)

	output, err := result.Stdout(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to execute Fabric command: %w", err)
	}

	return output, nil
}
