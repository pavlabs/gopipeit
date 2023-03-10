package helpers

import (
	"os"
	"text/template"
)

func WriteTemplateToFile(t *template.Template, f *os.File, cfg *Metadata) error {
	err := t.Execute(f, cfg)
	if err != nil {
		return err
	}
	return nil
}

// CreateDirPath taken from https://stackoverflow.com/a/59961623
func CreateDirPath(p string) error {
	err := os.MkdirAll(p, 0775)
	if err != nil {
		return err
	}
	return nil
}
