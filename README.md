***WARNING: THIS REPO IS AN AUTO-GENERATED COPY.*** *This repo has been copied from [Gruntwork’s](https://gruntwork.io/) GitHub repositories so that you can consume it from your company’s own internal Git repositories. This copy is automatically created and updated by the `repo-copier` CLI tool. If you need to make changes to this repo, you should make the changes in a separate fork, and NOT make changes directly in this repo, as otherwise, the `repo-copier` will overwrite your changes! Please see the `repo-copier` [documentation](https://github.com/terraform-modules-krish/repo-copier) for more information on how the code is copied, how cross-references are updated, how the changelog is handled, etc.*

***

_You may find it valuable to view the following resources in the original repo. If these links give you a 404, visit https://app.gruntwork.io to gain access or email support@gruntwork.io if you need assistance._

[Home Page](https://github.com/gruntwork-io/pre-commit/) |
[Pull Requests](https://github.com/gruntwork-io/pre-commit/pulls) |
[Issues](https://github.com/gruntwork-io/pre-commit/issues) |
[Releases and Assets](https://github.com/gruntwork-io/pre-commit/releases)

_Alternatively, you can view a copied version of the resources listed above._

[Pull Requests](https://github.com/terraform-modules-krish/pre-commit/blob/master/.github/PULL_REQUESTS.md) |
[Issues](https://github.com/terraform-modules-krish/pre-commit/blob/master/.github/ISSUES.md) |
[ChangeLog](https://github.com/terraform-modules-krish/pre-commit/blob/master/.github/CHANGELOG.md)

***

[![Maintained by Gruntwork.io](https://img.shields.io/badge/maintained%20by-gruntwork.io-%235849a6.svg)](https://gruntwork.io/?ref=repo_pre-commit)

# Pre-commit hooks

This repo defines Git pre-commit hooks intended for use with [pre-commit](http://pre-commit.com/). The currently
supported hooks are:

* **terraform-fmt**: Automatically run `terraform fmt` on all Terraform code (`*.tf` files).
* **shellcheck**: Run [`shellcheck`](https://www.shellcheck.net/) to lint files that contain a bash [shebang](https://en.wikipedia.org/wiki/Shebang_(Unix))
* **gofmt**: Automatically run `gofmt` on all Golang code (`*.go` files).
* **golint**: Automatically run `golint` on all Golang code (`*.go` files)
* **yapf**: Automatically run [`yapf`](https://github.com/google/yapf) on all python code (`*.py` files).




## General Usage

In each of your repos, add a file called `.pre-commit-config.yml` with the following contents:

```yaml
repos:
  - repo: https://github.com/terraform-modules-krish/pre-commit
    rev: <VERSION> # Get the latest from: https://github.com/gruntwork-io/pre-commit/releases
    hooks:
      - id: terraform-fmt
      - id: shellcheck
      - id: gofmt
      - id: golint
```

Next, have every developer: 

1. Install [pre-commit](http://pre-commit.com/). E.g. `brew install pre-commit`.
1. Run `pre-commit install` in the repo.

That’s it! Now every time you commit a code change (`.tf` file), the hooks in the `hooks:` config will execute.




## Running Against All Files At Once


### Example: Formatting all files

If you'd like to format all of your code at once (rather than one file at a time), you can run:

```bash
pre-commit run terraform-fmt --all-files
```



### Example: Enforcing in CI

If you'd like to enforce all your hooks, you can configure your CI build to fail if the code doesn't pass checks by
adding the following to your build scripts:

```bash
pip install pre-commit
pre-commit install
pre-commit run --all-files
```

If all the hooks pass, the last command will exit with an exit code of 0. If any of the hooks make changes (e.g.,
because files are not formatted), the last command will exit with a code of 1, causing the build to fail.





## License

This code is released under the Apache 2.0 License. Please see [LICENSE](LICENSE) and [NOTICE](NOTICE) for more details.

Copyright &copy; 2018 Gruntwork, Inc.
