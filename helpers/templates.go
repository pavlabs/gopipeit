package helpers

var gitHubConfigs = []SourceToDest{
	{
		TemplateSource:    "templates/github/release.yaml.tmpl",
		ConfigDestination: ".github/workflows/release.yaml",
	},
	{
		TemplateSource:    "templates/github/golangci-lint.yaml.tmpl",
		ConfigDestination: ".github/workflows/golangci-lint.yaml",
	},
	{
		TemplateSource:    "templates/github/commitlint.yaml.tmpl",
		ConfigDestination: ".github/workflows/commitlint.yaml",
	},
}

var localConfigs = []SourceToDest{
	{
		TemplateSource:    "templates/golangci.yaml.tmpl",
		ConfigDestination: "./.golangci.yaml",
	},
	{
		TemplateSource:    "templates/goreleaser.yaml.tmpl",
		ConfigDestination: "./.goreleaser.yaml",
	},
	{
		TemplateSource:    "templates/pre-commit-config.yaml.tmpl",
		ConfigDestination: "./.pre-commit-config.yaml",
	},
	//{
	//	TemplateSource:    "templates/dockerfile.tmpl",
	//	ConfigDestination: "./Dockerfile",
	// },
}

type SourceToDest struct {
	TemplateSource    string
	ConfigDestination string
}

type Templates struct {
	Pairs []SourceToDest
}

func NewTemplates() *Templates {
	return &Templates{
		Pairs: localConfigs,
	}
}

func (t *Templates) AddPair(pair SourceToDest) {
	t.Pairs = append(t.Pairs, pair)
}

func (t *Templates) AddSlice(pairs []SourceToDest) {
	t.Pairs = append(t.Pairs, pairs...)
}

func (t *Templates) WithGitHubCI() {
	t.Pairs = append(t.Pairs, gitHubConfigs...)
}
