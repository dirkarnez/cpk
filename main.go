package main

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/dirkarnez/cpk/cmake"
	"github.com/google/go-github/v41/github"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	c_cppLibrariesFolder string
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	c_cppLibrariesFolder = filepath.Join(exPath, "c_cpp_libraries")

	if err := os.MkdirAll(c_cppLibrariesFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var userName = os.Getenv("USER_NAME")

	fmt.Println("Enter password (if there is private repo): ")
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	password := string(passwordBytes)
	
	//var TARGET_OWNER = "dirkarnez"          //os.Getenv("TARGET_OWNER")
	var TARGET_REPOSITORY = "sfml-prebuilt" //fetched by libraries.yml, TARGET_OWNER also set here 
	var TARGET_COMPILER = "mingw" // compiler is detected

	wd, _ := os.Getwd()
	fmt.Println(wd)

	libraries, err := cmake.FetchDependencies(filepath.Join(wd, "CMakeLists.txt"))
	if err != nil {
		fmt.Printf("cmake.FetchDependencies returned error: %v\n", err)
	}

	tp := github.BasicAuthTransport{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}

	client := github.NewClient(tp.Client())

	for _, library := range libraries {
		libraryName := library[1]
		libraryVersion := library[2]

		targetTag := fmt.Sprintf("v%s", libraryVersion)

		release, resp, err := client.Repositories.GetReleaseByTag(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, targetTag)
		if err != nil {
			fmt.Printf("Repositories.GetReleaseByTag returned error: %v\n%v\n", err, resp.Body)
		}

		if len(release.Assets) < 1 {
			fmt.Printf("No Assets")
			return
		}

		// for example: sfml-v2.5.1-mingw.zip
		targetAssetName := fmt.Sprintf("%s-%s-%s.zip", strings.ToLower(libraryName) /*TARGET_LIBRARY*/, targetTag /*TARGET_TAG*/, TARGET_COMPILER)
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

		fileName := filepath.Join(c_cppLibrariesFolder, assetToDownload.GetName())

		(func() {
			outFile, err := os.Create(fileName)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, readCloser)

			if err != nil {
				fmt.Printf("error=[%v]\n", err)
				return
			}
		})()

		location := filepath.Join(c_cppLibrariesFolder, assetToDownload.GetName()[0:len(assetToDownload.GetName())-4]) // omitting ".zip"
		// if err := os.MkdirAll(location, os.ModePerm); err != nil {
		// 	log.Fatal(err)
		// }

		if err := UnZip(location, fileName); err != nil {
			fmt.Printf("error=[%v]\n", err)
			return
		}

		e := os.Remove(fileName)
		if e != nil {
			log.Fatal(e)
		}

		// write FindXXX.cmake
		readCloser, _, err = client.Repositories.DownloadContents(context.Background(), TARGET_OWNER, TARGET_REPOSITORY, "FindSFML.cmake.tpl", nil)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(readCloser)

		outFile, err := os.Create(filepath.Join(wd, "cmake", "modules", "FindSFML.cmake"))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer outFile.Close()

		err = cmake.WriteCMakeFindModuleFile(buf.String(), location, outFile)
		if err != nil {
			fmt.Printf("error=[%v]\n", err)
		}
		fmt.Println("done")
	}
}

func UnZip(dst, fileName string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(fileName)
	if err != nil {
		return
	}
	defer zr.Close()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if len(dst) < 1 {
		if err := os.MkdirAll(dst, os.ModePerm); err != nil {
			return err
		}
	}

	// cpk should unzip file to a new directory with same name, because the name itself is conventional
	var filesInZip = zr.File
	var lenFilesInZip = len(filesInZip)
	if lenFilesInZip > 1 {
		for _, file := range filesInZip {
			path := filepath.Join(dst, file.Name)

			// 如果是目录，就创建目录
			if file.FileInfo().IsDir() {
				if err := os.MkdirAll(path, file.Mode()); err != nil {
					return err
				}
				// 因为是目录，跳过当前循环，因为后面都是文件的处理
				continue
			}

			// 获取到 Reader
			fr, err := file.Open()
			if err != nil {
				return err
			}

			// 创建要写出的文件对应的 Write
			fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}

			/*n*/
			_, err = io.Copy(fw, fr)
			if err != nil {
				return err
			}

			// 将解压的结果输出
			//fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

			// 因为是在循环中，无法使用 defer ，直接放在最后
			// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
			// 可以把它单独放在一个函数中，这里是个实验，就这样了
			fw.Close()
			fr.Close()
		}
	}

	return nil
}
