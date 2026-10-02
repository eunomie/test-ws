package greetings

import (
	"dagger.io/dagger/core"

	"dagger/greetings/internal/dagger"
	"dagger/greetings/internal/dagger/clients/greetings"
)

type Greetings struct {
	// Image the build runs on.
	BaseImageAddress string

	// +private
	Source *dagger.Directory
}

func New(
	ws *dagger.Workspace,
	// +default="alpine:3.21"
	baseImageAddress string,
) *Greetings {
	return &Greetings{
		BaseImageAddress: baseImageAddress,
		Source: ws.Directory("/", dagger.WorkspaceDirectoryOpts{
			Exclude: []string{"**/.git", "**/.dagger", "**/vendor"},
		}),
	}
}

// A container with the source mounted, ready to build on.
func (m *Greetings) Container() *dagger.Container {
	return core.NewContainer().
		From(m.BaseImageAddress).
		WithDirectory("/src", m.Source).
		WithWorkdir("/src")
}

func (m *Greetings) SayHello(ws *dagger.Workspace, name string) *dagger.Container {
	return greetings.New(ws).Container().WithExec([]string{"echo", "hello", name})
}
