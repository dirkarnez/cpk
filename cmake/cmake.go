package cmake

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type cmakeFindModuleFile struct {
	Location string
}

func WriteCMakeFindModuleFile(templateContent, locationNoEndingSlash string, w io.Writer) error {
	cmakeFindModuleFile := cmakeFindModuleFile{Location: strings.ReplaceAll(locationNoEndingSlash, `\`, `/`)}

	tmpl, err := template.New("test").Parse(templateContent)
	if err != nil {
		panic(err)
	}
	return tmpl.Execute(w, cmakeFindModuleFile)
}

func FetchDependencies(CMakeListsTXT_filePath string) ([][]string, error) {
	file, err := os.Open(CMakeListsTXT_filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	content := buf.String()

	// commented find_package should be ignored
	regexp, _ := regexp.Compile(`(?im:^find_package\(([^\s]+)\s+([^\s]+).*$)`)

	return regexp.FindAllStringSubmatch(content, -1), nil
}
