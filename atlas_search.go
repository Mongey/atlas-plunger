package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/atlas-go/v1"
	"github.com/imdario/mergo"
	"strings"
)

func LatestArtifactForEnvironment(artifact string, environment string) (artifactVersion *atlas.ArtifactVersion, err error) {

	searchOptions := &atlas.ArtifactSearchOpts{
		Version: "latest",
		Metadata: map[string]string{
			"environment": environment,
		},
	}

	artifacts, err := findArtifacts(artifact, searchOptions)
	if err != nil {
		fmt.Println("findArtifacts error")
		return nil, err
	}

	if len(artifacts) > 1 {
		return nil, errors.New(fmt.Sprintf("Mutiple artifacts %s", len(artifacts)))
	}

	return artifacts[0], nil
}

func ArtifactsForCommit(artifact string, sha string) (artifactVersions []*atlas.ArtifactVersion, err error) {

	searchOptions := &atlas.ArtifactSearchOpts{
		Metadata: map[string]string{
			"git_commit":        sha,
			"git_commit_status": "clean",
		},
	}

	return findArtifacts(artifact, searchOptions)
}

func findArtifacts(fullArtifactSlug string, inputSearchData *atlas.ArtifactSearchOpts) (artifactVersion []*atlas.ArtifactVersion, err error) {

	client := atlas.DefaultClient()

	s := strings.Split(fullArtifactSlug, "/")
	if len(s) < 2 {
		return nil, err
	}
	organisation, artifact_name := s[0], s[1]

	defaultSearchOptions := atlas.ArtifactSearchOpts{
		User: organisation,
		Name: artifact_name,
		Type: "amazon.ami",
	}

	if err := mergo.Merge(inputSearchData, defaultSearchOptions); err != nil {
		fmt.Println(err)
		return nil, err
	}

	var artifacts []*atlas.ArtifactVersion
	artifacts, err = client.ArtifactSearch(inputSearchData)

	if err != nil {
		return nil, err
	}

	if len(artifacts) < 1 {
		return nil, errors.New(fmt.Sprintf("No artifacts found for %s/%s with search options ", organisation, artifact_name, inputSearchData))
	}

	return artifacts, nil
}
