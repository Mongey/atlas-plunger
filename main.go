package main

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/hashicorp/atlas-go/v1"

	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func EnsureAtlasTokenSet() {
	token := os.Getenv("ATLAS_TOKEN")

	if token == "" {
		fmt.Println("ATLAS_TOKEN env var must be set ", token)
		os.Exit(1)
	}
}

func main() {
	var artifactSlug, environment, commitSHA string

	app := cli.NewApp()
	app.Name = "atlas-plunger"
	app.Usage = "atlas artifact manager"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{{"Conor Mongey", "conor@mongey.net"}}
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:    "find",
			Aliases: []string{"f"},
			Usage:   "find stuff",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "commit",
					Usage:       "The commit you are looking for",
					EnvVar:      "ATLAS_ARETEFACT_COMMIT",
					Destination: &commitSHA,
				},
				cli.StringFlag{
					Name:        "artifact",
					Usage:       "the image that you want to promote e.g. foobar/image",
					EnvVar:      "ATLAS_ARTIFACT_SLUG",
					Destination: &artifactSlug,
				},
				cli.StringFlag{
					Name:        "environment",
					Usage:       "the environment that you want to promote to",
					EnvVar:      "ALTAS_ENVIRONMENT",
					Destination: &environment,
				},
			},
			Action: func(c *cli.Context) {
				EnsureAtlasTokenSet()

				var err error
				var art *atlas.ArtifactVersion
				var arts []*atlas.ArtifactVersion

				if artifactSlug == "" {
					err = errors.New("No artifact specified")
				}

				if (commitSHA == "") && (environment == "") {
					handleError(errors.New("-environment or -commit must be specified"))
				}

				if commitSHA != "" {
					arts, err = ArtifactsForCommit(artifactSlug, commitSHA)
					if err != nil {
						handleError(err)
					}
				}

				if environment != "" {
					art, err = LatestArtifactForEnvironment(artifactSlug, environment)

					if err != nil {
						handleError(err)
					}
					arts = append(arts, art)
				}

				for i := 0; i < len(arts); i++ {
					PrintArtifactVersion(arts[i])
					fmt.Println("-------------------")
				}
			},
		},
		{
			Name:    "promote",
			Aliases: []string{"p"},
			Usage:   "promote things",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "commit",
					Usage:       "The commit you are looking for",
					EnvVar:      "ATLAS_ARETEFACT_COMMIT",
					Destination: &commitSHA,
				},
				cli.StringFlag{
					Name:        "artifact",
					Usage:       "the image that you want to promote e.g. foobar/image",
					EnvVar:      "ATLAS_ARTIFACT_SLUG",
					Destination: &artifactSlug,
				},
				cli.StringFlag{
					Name:        "environment",
					Usage:       "the environment that you want to promote to",
					EnvVar:      "ALTAS_ENVIRONMENT",
					Destination: &environment,
				},
			},
			Action: func(c *cli.Context) {
				var err error
				var newArtifact *atlas.ArtifactVersion
				var arts []*atlas.ArtifactVersion

				EnsureAtlasTokenSet()

				if artifactSlug == "" {
					handleError(errors.New("No artifact specified"))
				}

				if (commitSHA == "") || (environment == "") {
					handleError(errors.New("-environment and -commit must be specified"))
				}

				// promote sha for slug to env
				arts, err = ArtifactsForCommit(artifactSlug, commitSHA)
				if err != nil {
					handleError(err)
				}

				if len(arts) != 1 {
					handleError(errors.New("What do I do with all these aretfacts"))
				}

				newArtifact, err = Promote(arts[0], environment)

				if err != nil {
					handleError(err)
				}
				fmt.Println("success!")
				PrintArtifactVersion(newArtifact)
			},
		},
	}

	app.Run(os.Args)
}
