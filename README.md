# GoPipeIt

An opinionated ci/cd stack generator for go projects.

Generates config files for [golangci-lint](https://github.com/golangci/golangci-lint),
[goreleaser](https://github.com/goreleaser/goreleaser), [pre-commit](https://github.com/pre-commit/pre-commit)
and CI pipeline configuration files for GitHub actions.

### Examples

```shell
# generate local configuration files, skip file if it's already present
gopipeit

# generate (overwriting existing) configuration files
# with GitHub action workflows
gopipeit --force --with-github-ci
```

### Features

- Generate config files for local development ([`golangci-lint`](https://github.com/golangci/golangci-lint), [`pre-commit`](https://github.com/pre-commit/pre-commit))
- Generate config files for GitHub actions

---

built with [`Cobra`](https://github.com/spf13/cobra) and [`pterm`](https://github.com/pterm/pterm) 🖤
