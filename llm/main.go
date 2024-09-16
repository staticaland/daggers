// A module for running the llm CLI tool
//
// This module provides a function to execute the llm CLI tool
// within a Dagger pipeline. llm is a command-line utility for
// interacting with Large Language Models (LLMs).
//
// The Run function executes llm with various options, allowing
// for flexible prompting and interaction with LLMs. It can be called from
// the Dagger CLI or from one of the SDKs.

package main

import (
	"context"
	"fmt"
	"strings"

	"dagger/llm/internal/dagger"
)

type Llm struct{}

// Run executes the llm CLI command
func (m *Llm) Run(
	ctx context.Context,
	// The prompt to send to the LLM
	prompt string,
	// System prompt to use
	// +optional
	system string,
	// Model to use
	// +optional
	model string,
	// Key/value options for the model (format: "key=value")
	// +optional
	options []string,
	// Template to use
	// +optional
	template string,
	// Parameters for template (format: "key=value")
	// +optional
	params []string,
	// Do not stream output
	// +optional
	noStream bool,
	// Don't log to database
	// +optional
	noLog bool,
	// Log prompt and response to the database
	// +optional
	log bool,
	// Continue the most recent conversation
	// +optional
	continueConversation bool,
	// Continue the conversation with the given ID
	// +optional
	conversationID string,
	// API key to use
	// +optional
	apiKey string,
	// Save prompt with this template name
	// +optional
	save string,
	// OpenAI API key as a secret
	// +optional
	openAiApiKey *dagger.Secret,
) (string, error) {
	container := dag.Container().
		From("python:slim").
		WithExec([]string{"pip", "install", "llm"})

	if openAiApiKey != nil {
		container = container.WithSecretVariable("OPENAI_API_KEY", openAiApiKey)
	}

	args := []string{"llm", "prompt"}

	if system != "" {
		args = append(args, "--system", system)
	}
	if model != "" {
		args = append(args, "--model", model)
	}
	for _, option := range options {
		args = append(args, "--option", option)
	}
	if template != "" {
		args = append(args, "--template", template)
	}
	for _, param := range params {
		args = append(args, "--param", param)
	}
	if noStream {
		args = append(args, "--no-stream")
	}
	if noLog {
		args = append(args, "--no-log")
	}
	if log {
		args = append(args, "--log")
	}
	if continueConversation {
		args = append(args, "--continue")
	}
	if conversationID != "" {
		args = append(args, "--conversation", conversationID)
	}
	if save != "" {
		args = append(args, "--save", save)
	}

	// Add the prompt as the last argument
	args = append(args, prompt)

	result := container.WithExec(args)

	output, err := result.Stdout(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to execute llm command: %w", err)
	}

	return strings.TrimSpace(output), nil
}
