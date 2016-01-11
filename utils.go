package main

import (
	"fmt"
	"github.com/hashicorp/atlas-go/v1"
	"os"
)

func PrintArtifactVersion(artifactVersion *atlas.ArtifactVersion) {

	fmt.Println(artifactVersion.Tag)
	fmt.Println("\t type:", artifactVersion.Type)
	fmt.Println("\t id :", artifactVersion.ID)
	fmt.Println("\t version :", artifactVersion.Version)
	fmt.Println("\t slug :", artifactVersion.Slug)
	fmt.Println("\t tags? :", artifactVersion.Tag)

	for k, v := range artifactVersion.Metadata {
		fmt.Println("\t", k, ":", v)
	}
}

func ConvertArtifactVersionToUploadArtifactOpts(artifactVersion *atlas.ArtifactVersion) *atlas.UploadArtifactOpts {
	return &atlas.UploadArtifactOpts{
		User:     artifactVersion.User,
		Name:     artifactVersion.Name,
		Type:     artifactVersion.Type,
		ID:       artifactVersion.ID,
		Metadata: artifactVersion.Metadata,
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: 🚑  🚨 ")
		fmt.Println(err)
		os.Exit(0)
	}
}
