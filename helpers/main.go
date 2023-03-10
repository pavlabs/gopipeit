package helpers

import (
	"text/template"

	"github.com/spf13/afero"
)

func WriteTemplateToFile(t *template.Template, f afero.File, cfg *Metadata) error {
	err := t.Execute(f, cfg)
	if err != nil {
		return err
	}
	return nil
}
