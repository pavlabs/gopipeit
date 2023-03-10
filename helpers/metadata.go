package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Metadata struct {
	ProjectName string
	GoVersion   string
	GitBranch   string
}

func NewMetadata() *Metadata {
	return &Metadata{}
}

func (cfg *Metadata) SetProjectName(n string) {
	cfg.ProjectName = n
}

func (cfg *Metadata) SetGoVersion(n string) {
	cfg.GoVersion = n
}

func (cfg *Metadata) SetGitBranch(n string) {
	cfg.GitBranch = n
}

func readGoMod() ([]string, error) {
	f, err := os.Open("go.mod")
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var txt []string
	for sc.Scan() {
		txt = append(txt, sc.Text())
	}
	f.Close()
	return txt, nil
}

func (cfg *Metadata) ExtractProjectNameFromGoModFile() {
	txt, err := readGoMod()
	if err != nil {
		log.Fatal("failed to read go.mod file")
	}

	for _, line := range txt {
		if strings.HasPrefix(line, "module ") {
			trimmed := strings.Replace(line, "module ", "", 1)
			t := trimmed[strings.LastIndex(trimmed, "/")+1:]
			cfg.SetProjectName(t)
			break
		}
	}

	if cfg.ProjectName == "" {
		log.Fatal("failed to extract project name from go.mod")
	}
}

func (cfg *Metadata) ExtractGoVersionFromGoModFile() {
	txt, err := readGoMod()
	if err != nil {
		log.Fatal("failed to read go.mod file")
	}

	for _, line := range txt {
		if strings.HasPrefix(line, "go ") {
			t := strings.Replace(line, "go ", "", 1)
			cfg.SetGoVersion(t)
			break
		}
	}
	if cfg.GoVersion == "" {
		log.Fatal("failed to extract go version from go.mod")
	}
}
