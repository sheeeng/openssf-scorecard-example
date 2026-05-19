# openssf-scorecard-example

[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/sheeeng/openssf-scorecard-example/badge)](https://scorecard.dev/viewer/?uri=github.com/sheeeng/openssf-scorecard-example)

See:

- [ossf/scorecard](https://github.com/ossf/scorecard)
- [ossf/scorecard-action](https://github.com/ossf/scorecard-action)

## Miscellenous

```shell
nix run nixpkgs#markdown-link-check -- README.md
nix run nixpkgs#yamlfmt -- .github/workflows/
nix run nixpkgs#yamllint -- .github/workflows/
```

```shell
find . -name "*.md" -exec nix run 'nixpkgs#markdown-link-check' -- {} \;
find . -name "*.md" | xargs --replace={} nix run 'nixpkgs#markdown-link-check' -- {}
rg --files --type markdown | xargs --replace={} nix run 'nixpkgs#markdown-link-check' -- {}
```
