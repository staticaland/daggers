// A generated module for Releaseplease functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/releaseplease/internal/dagger"
	"log"
	"time"
)

type Releaseplease struct{}

// Is it a Release PR being merged? Run release.
// Is it a normal commit? Run PR.
// See https://github.com/googleapis/release-please/blob/main/docs/cli.md
// Important options:
// --token
// --repo-url

// release-please release-pr
//   --token=$GITHUB_TOKEN \
//   --repo-url=<owner>/<repo> [extra options]

// release-please github-release \
//   --token=$GITHUB_TOKEN
//   --repo-url=<owner>/<repo> [extra options]

// See https://github.com/dagger/dagger/pull/8468

// Install the release-please CLI tool
func (m *Releaseplease) Install() *dagger.Container {

	container := dag.Container().
		From("node:slim").
		WithExec([]string{"npm", "install", "-g", "release-please"})

	return container
}

// Run the release-please CLI tool with the release-pr command
func (m *Releaseplease) Pr(
	ctx context.Context,
	// GitHub token with repo write permissions
	token *dagger.Secret,
	// GitHub repository in the format of `<owner>/<repo>`
	repo string,
) *dagger.Container {

	container := m.Install()

	t, err := token.Plaintext(ctx)

	if err != nil {
		log.Fatalf("failed to get token plaintext: %v", err)
	}

	container = container.WithSecretVariable("GITHUB_TOKEN", token).
		WithEnvVariable("CACHEBUSTER", time.Now().String()).
		WithExec([]string{"release-please", "release-pr", "--repo-url", repo, "--token", t})

	return container
}

// Run the release-please CLI tool with the github-release command
func (m *Releaseplease) Release(
	ctx context.Context,
	// GitHub token with repo write permissions
	// +optional
	token *dagger.Secret,
	// GitHub repository in the format of `<owner>/<repo>`
	repo string,
) *dagger.Container {

	container := m.Install()

	t, err := token.Plaintext(ctx)

	if err != nil {
		log.Fatalf("failed to get token plaintext: %v", err)
	}

	container = container.WithSecretVariable("GITHUB_TOKEN", token).
		WithEnvVariable("CACHEBUSTER", time.Now().String()).
		WithExec([]string{"release-please", "github-release", "--repo-url", repo, "--token", t})

	return container
}
