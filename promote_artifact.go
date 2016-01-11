package main

import (
	"fmt"
	"github.com/hashicorp/atlas-go/v1"
	"strconv"
	"time"
)

func Promote(artifact_to_promote *atlas.ArtifactVersion, destinationEnvironment string) (*atlas.ArtifactVersion, error) {
	client := atlas.DefaultClient()

	timestamp := time.Now().Unix()

	metadata_to_overwrite := map[string]string{
		"environment": destinationEnvironment,
		"promoted_at": strconv.FormatInt(timestamp, 10),
	}

	newArtifact := ConvertArtifactVersionToUploadArtifactOpts(artifact_to_promote)

	for k, v := range metadata_to_overwrite {
		newArtifact.Metadata[k] = v
	}

	av, err := client.UploadArtifact(newArtifact)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error Uploading artifact %s", err))
		return nil, err
	}

	fmt.Println(av)

	return av, nil
}
