package main

import (
	"github.com/hashicorp/atlas-go/v1"
	"testing"
)

func Test_ConvertArtifactVersionToUploadArtifactOpts(t *testing.T) {
	artifact := &atlas.ArtifactVersion{
		User: "homersimpson",
		Name: "nomad",
		Type: "amazon.ami",
		ID:   "100",
	}

	result := ConvertArtifactVersionToUploadArtifactOpts(artifact)

	if result.User != artifact.User {
		t.Fatalf("bad: %s != %s", result.User, artifact.User)
	}
	if result.Name != artifact.Name {
		t.Fatalf("bad: %s != %s", result.Name, artifact.Name)
	}
	if result.Type != artifact.Type {
		t.Fatalf("bad: %s != %s", result.Type, artifact.Type)
	}
	if result.ID != artifact.ID {
		t.Fatalf("bad: %s != %s", result.ID, artifact.ID)
	}
}
