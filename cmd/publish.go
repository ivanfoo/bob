package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/xanzy/go-gitlab"
	"gopkg.in/urfave/cli.v1"
)

type Project struct {
	Name        string
	Description string
	Token       string
}

func Publish(c *cli.Context) error {
	p := NewProject(
		c.String("name"),
		c.String("description"),
		c.String("token"),
	)

	p.createOnGitlab()

	return nil

}

func NewProject(name, description, token string) *Project {
	if token == "" {
		log.Fatal("Missing --token value")
	}

	if name == "" {
		pwd, _ := os.Getwd()
		_, name = filepath.Split(pwd)
	}

	return &Project{
		Name:        name,
		Description: description,
		Token:       token,
	}
}

func (p *Project) createOnGitlab() error {
	g := gitlab.NewClient(nil, p.Token)

	_, _, err := g.Projects.CreateProject(&gitlab.CreateProjectOptions{
		Name:                 gitlab.String(p.Name),
		Description:          gitlab.String(p.Description),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(true),
		VisibilityLevel:      gitlab.VisibilityLevel(gitlab.PrivateVisibility),
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
