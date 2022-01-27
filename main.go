package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"
)

func main() {
	var userName = os.Getenv("USER_NAME")
	var password = os.Getenv("PASSWORD")
	var repository = os.Getenv("REPOSITORY")
	var path = os.Getenv("PATH")

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())

	readCloser, _, err := client.Repositories.DownloadContents(context.Background(), userName, repository, path, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	outFile, err := os.Create(getFileName(path))
	// handle err
	defer outFile.Close()
	_, err = io.Copy(outFile, readCloser)

	if err != nil {
		fmt.Printf("error=[%v]\n", err)
	} else {
		fmt.Printf("done")
	}
}

func getFileName(path string) string {
	return path[strings.LastIndex(path, "/")+1:]
}
