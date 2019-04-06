package main

import (
	"flag"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var k8spatternsExampleGitUrl = "https://github.com/k8spatterns/examples.git"

var target string
var examplesDir string

type example struct {
	Path string `yaml: "path"`
}

type examplesList struct {
	Pattern  string    `yaml:"pattern"`
	Base     string    `yaml:"base"`
	Examples []example `yaml:"examples"`
}

func init() {
	flag.StringVar(&target, "target", "", "target directory where create resources")
}

func main() {

	flag.Parse()

	if target == "" {
		log.Fatal("no target directory provided with -target")
	}

	var err error
	examplesDir, err = ioutil.TempDir(".", "examples")
	checkError(err)
	defer os.RemoveAll(examplesDir)

	_, err = git.PlainClone(examplesDir, false, &git.CloneOptions{
		URL:   k8spatternsExampleGitUrl,
		Depth: 1,
	})
	if err != nil {
		log.Fatalf("can't clone"+k8spatternsExampleGitUrl+": %v", err)
	}

	err = filepath.Walk(examplesDir, processIndexYml)
	checkError(err)
}

func processIndexYml(path string, f os.FileInfo, err error) error {
	if f.IsDir() || f.Name() != "index.yml" {
		return nil
	}

	examples := examplesList{}
	data, err := ioutil.ReadFile(path)
	checkError(err)

	err = yaml.Unmarshal(data, &examples)
	checkError(err)

	patternDir := target + "/" + strings.Replace(examples.Pattern, " ", "", -1)
	_, err = os.Stat(patternDir)
	if os.IsNotExist(err) {
		checkError(os.Mkdir(patternDir, 0755))
	}
	for _, example := range examples.Examples {
		dest := patternDir + "/" + example.Path
		println(dest)
		copy(examplesDir+"/"+examples.Base + "/" + example.Path, dest)
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkError(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	checkError(err)
}
