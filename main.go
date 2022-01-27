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

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())
	release, resp, err := client.Repositories.GetReleaseByTag(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, "2018-07-31")
	if err != nil {
		fmt.Printf("Repositories.GetReleaseByTag returned error: %v\n%v\n", err, resp.Body)
	}

	if len(release.Assets) < 1 {
		fmt.Printf("NO Assets")
		return
	}

	var toDownloadAsset *github.ReleaseAsset = nil
	for _, asset := range release.Assets {
		if asset.GetName() == "x86_64-full.tar.bz2" {
			toDownloadAsset = asset
		}
	}

	readCloser, _, err := client.Repositories.DownloadReleaseAsset(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, toDownloadAsset.GetID(), http.DefaultClient)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer readCloser.Close()

	outFile, err := os.Create(toDownloadAsset.GetName())
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
