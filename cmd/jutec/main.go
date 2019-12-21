package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-zookeeper/jute/generate"
	"github.com/go-zookeeper/jute/generate/gogen"
)

type moduleMaps []gogen.ModuleMap

func (mm moduleMaps) String() string {
	sb := strings.Builder{}
	for i, m := range mm {
		sb.WriteString(m.String())
		if i < len(mm)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func (mm *moduleMaps) Set(s string) (err error) {
	p := strings.Split(s, ":")
	if len(p) != 2 {
		return fmt.Errorf("maps must be in the form of <regular expression>:<replacement>")
	}

	m := gogen.ModuleMap{}
	m.Re, err = regexp.Compile(p[0])
	if err != nil {
		return fmt.Errorf("invalid map: %w", err)
	}

	if p[1] == "-" {
		m.Skip = true
	} else {
		m.Repl = p[1]
	}

	*mm = append(*mm, m)
	return nil
}

func main() {

	lang := flag.String("lang", "go", "language for output files")
	outputDir := flag.String("outDir", "./jute-gen", "directory to generate files in")
	// Go options
	goJuteImport := flag.String("go.juteImport", "github.com/go-zookeeper/jute/lib/go/jute", "override for encode/decoder library import path")
	goPkgPrefix := flag.String("go.prefix", "", "prefix for generated packages (not used in directory creation)")
	goModuleMap := moduleMaps{}
	flag.Var(&goModuleMap, "go.moduleMap", "map of jute module to go package")

	flag.Parse()

	files, err := generate.ParseFiles(flag.Args()...)
	if err != nil {
		log.Fatalf("failed to parse files: %v", err)
	}

	switch *lang {
	case "go", "golang":
		err = gogen.Generate(*outputDir, files, &gogen.Options{
			ImportPathPrefix: *goPkgPrefix,
			JuteImport:       *goJuteImport,
			ModuleMap:        goModuleMap,
		})
	default:
		log.Fatalf("unknown language: %s", *lang)
	}

	if err != nil {
		log.Fatalf("failed to generate: %v\n", err)
	}
}
