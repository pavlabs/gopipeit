package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/afero"
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

func readGoMod(fs afero.Fs) ([]string, error) {
	exists, _ := afero.Exists(fs, "go.mod")
	if !exists {
		return nil, fmt.Errorf("go.mod does not exist")
	}
	f, err := fs.Open("go.mod")
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

func (cfg *Metadata) ExtractProjectNameFromGoModFile(fs afero.Fs) error {
	txt, err := readGoMod(fs)
	if err != nil {
		return fmt.Errorf("failed to read go.mod: %v", err)
	}

	for _, line := range txt {
		if strings.HasPrefix(line, "module ") {
			trimmed := strings.Replace(line, "module ", "", 1)
			t := trimmed[strings.LastIndex(trimmed, "/")+1:]
			cfg.SetProjectName(t)
			return nil
		}
	}

	if cfg.ProjectName == "" {
		return errors.New("failed to extract project name from go.mod")
	}
	return nil
}

func (cfg *Metadata) ExtractGoVersionFromGoModFile(fs afero.Fs) error {
	txt, err := readGoMod(fs)
	if err != nil {
		return fmt.Errorf("failed to read go.mod file: %v", err)
	}

	for _, line := range txt {
		if strings.HasPrefix(line, "go ") {
			t := strings.Replace(line, "go ", "", 1)
			cfg.SetGoVersion(t)
			return nil
		}
	}
	if cfg.GoVersion == "" {
		return errors.New("failed to extract go version from go.mod")
	}
	return nil
}
