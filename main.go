package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"
)

func main() {
	var userName = os.Getenv("USER_NAME")
	var password = os.Getenv("PASSWORD")

	var TARGET_OWNER = os.Getenv("TARGET_OWNER")
	var TARGET_REPOSITORY = os.Getenv("TARGET_REPOSITORY")
	var TARGET_LIBRARY = os.Getenv("TARGET_LIBRARY")
	var TARGET_TAG = os.Getenv("TARGET_TAG")
	var TARGET_COMPILER = os.Getenv("TARGET_COMPILER")

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	release, resp, err := client.Repositories.GetReleaseByTag(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, TARGET_TAG)
	if err != nil {
		fmt.Printf("Repositories.GetReleaseByTag returned error: %v\n%v\n", err, resp.Body)
	}

	if len(release.Assets) < 1 {
		fmt.Printf("No Assets")
		return
	}

	// for example: sfml-v2.5.1-mingw.zip
	targetAssetName := fmt.Sprintf("%s-%s-%s.zip", TARGET_LIBRARY, TARGET_TAG, TARGET_COMPILER)
	var assetToDownload *github.ReleaseAsset = nil
	for _, asset := range release.Assets {
		if asset.GetName() == targetAssetName {
			assetToDownload = asset
		}
	}

	readCloser, _, err := client.Repositories.DownloadReleaseAsset(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, assetToDownload.GetID(), http.DefaultClient)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer readCloser.Close()

	outFile, err := os.Create(assetToDownload.GetName())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, readCloser)

	if err != nil {
		fmt.Printf("error=[%v]\n", err)
	} else {
		fmt.Printf("done\n")
	}
}
