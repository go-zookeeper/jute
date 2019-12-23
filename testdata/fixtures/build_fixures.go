package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-zookeeper/jute/generate"
	"github.com/go-zookeeper/jute/generate/gogen"
)

type fixture struct {
	name    string
	files   []string
	options *gogen.Options
}

var fixtures = []fixture{
	{"test", []string{"testdata/test.jute"}, nil},
	{"zookeeper",
		[]string{"testdata/zookeeper.jute"},
		&gogen.Options{
			ImportPathPrefix: "github.com/go-zookeeper/zk/internal",
			ModuleMap: []gogen.ModuleMap{
				{
					Re:   regexp.MustCompile("org.apache.zookeeper"),
					Repl: "", // drop thre prefix
				},
			},
		},
	},
}

func buildFixture(f fixture) error {
	outputDir := filepath.Join("testdata/fixtures", f.name)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	files, err := generate.ParseFiles(f.files...)
	if err != nil {
		return err
	}

	return gogen.Generate(outputDir, files, f.options)
}

func main() {
	if _, err := os.Stat("go.mod"); err != nil {
		log.Fatal("this should be ran at the root of `jute` repository: \ngo run testdata/fixtures/build_fixtures.go`")
	}

	total := len(fixtures)
	for i, f := range fixtures {
		log.Printf("building fixture %s (%d/%d)", f.name, i, total)
		if err := buildFixture(f); err != nil {
			log.Fatalf("failed to build fixture %s: %v", f.name, err)
		}
	}
}
