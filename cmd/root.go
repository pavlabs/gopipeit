/*
Copyright © 2023 Artemijs Pavlovs <hi@artpav.dev>
*/
package cmd

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/afero"

	"github.com/artemijspavlovs/gopipeit/helpers"
	"github.com/spf13/cobra"
)

var ProjectName string
var GoVersion string
var GitBranch string
var RegenerateAll bool
var WithGitHubCI bool
var AFS = afero.NewOsFs()

//go:embed templates/*
var embeddedTemplates embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gopipeit",
	Short: "Generate CI configuration files with one command",
	Long: `Binary created to provide CI and local development configuration files for Go projects.
Use it to generate optimal configuration files for GitHub Actions, goreleaser, pre-commit and golangi-lint`,
	Run: func(cmd *cobra.Command, args []string) {
		metadata, err := ExtractMetadataValues()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = AFS.MkdirAll("./.github/workflows/", os.FileMode(0755))
		if err != nil {
			fmt.Println("failed to create directory ./.github/workflows: ", err)
			return
		}

		tmplts := helpers.NewTemplates()
		if WithGitHubCI {
			tmplts.WithGitHubCI()
		}

		err = GenerateConfigFromTemplates(tmplts.Pairs, metadata)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("\nThere are some workflows that manage PRs. Please enable 'GitHub Actions to create and" +
			"approve pull requests' in your GitHub repository under\n" +
			"\tSettings => Actions => General => Allow GitHub Actions to create and approve pull requests")
	},
}

func ExtractMetadataValues() (*helpers.Metadata, error) {
	meta := helpers.NewMetadata()
	if ProjectName == "" {
		fmt.Println("Project name was not set, extracting from go.mod file")
		err := meta.ExtractProjectNameFromGoModFile(AFS)
		if err != nil {
			return nil, fmt.Errorf("failed to extract project name from go.mod file: %v", err)
		}
		fmt.Println("Project name extracted from go.mod: ", meta.ProjectName)
	} else {
		fmt.Printf("Setting project name to %s\n", ProjectName)
		meta.SetProjectName(ProjectName)
	}

	if GoVersion == "" {
		fmt.Println("Go version was not set, extracting from go.mod file")
		err := meta.ExtractGoVersionFromGoModFile(AFS)
		if err != nil {
			return nil, fmt.Errorf("failed to extract go version from go.mod file: %v", err)
		}
		fmt.Println("Go version extracted from go.mod: ", meta.GoVersion)
	} else {
		fmt.Printf("Setting project name to %s\n", ProjectName)
		meta.SetProjectName(ProjectName)
	}

	if GitBranch == "" {
		fmt.Println("Git branch was not set, defaulting to `main`")
		meta.SetGitBranch("main")
	}

	if GitBranch != "" {
		fmt.Println("Git branch was provided, setting it to: ", GitBranch)
		meta.SetGitBranch(GitBranch)
	}

	return meta, nil
}

func GenerateConfigFromTemplates(s []helpers.SourceToDest, m *helpers.Metadata) error {
	var skipped int
	fs := afero.NewOsFs()
	for _, pair := range s {
		exists, _ := afero.Exists(fs, pair.ConfigDestination)
		if exists && !RegenerateAll {
			fmt.Printf("Config %s already exists, it will not be replaced\n", pair.ConfigDestination)
			skipped++
			continue
		}
		fmt.Printf("Generating config %s\n", pair.ConfigDestination)

		tmpl := template.Must(template.ParseFS(embeddedTemplates, pair.TemplateSource))
		f, err := fs.Create(pair.ConfigDestination)
		if err != nil {
			return err
		}

		err = helpers.WriteTemplateToFile(tmpl, f, m)
		if err != nil {
			return err
		}
	}

	if skipped > 0 {
		fmt.Printf("Skipped %d templates, you can use --force flag to generate and replace all templates\n", skipped)
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gopipeit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(
		&ProjectName,
		"project",
		"",
		"Project name",
	)
	rootCmd.Flags().StringVar(
		&GoVersion,
		"go-version",
		"",
		"Go Version to use in CI pipelines",
	)
	rootCmd.Flags().StringVar(
		&GitBranch,
		"default-branch",
		"",
		"Default branch in the repository ( default: main )",
	)
	rootCmd.Flags().BoolVarP(
		&RegenerateAll,
		"force",
		"f",
		false,
		"Force overwrite existing files",
	)
	rootCmd.Flags().BoolVar(
		&WithGitHubCI,
		"with-github-ci",
		false,
		"Generate GitHub action files",
	)
}
