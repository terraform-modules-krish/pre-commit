# Changelog


# v0.1.9
_Released Jun 16, 2020_
### Modules affected

* `terraform-validate`

### Description

This fixes a regression bug where the `terraform-validate` hook that broke with the `tfenv` support release (`v0.1.5`).

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/32
<br>

# v0.1.8
_Released Jun 5, 2020_
### Modules affected

* `terraform-fmt`
* `terraform-validate`
* `terragrunt-hclfmt`

### Description

This fixes a regression bug where the `terraform fmt` and `terragrunt hclfmt` calls were not running on the right files after the improvement to support `tfenv` and `tgenv`.

### Special Thanks

Special thanks to @oasys  for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/29
* https://github.com/gruntwork-io/pre-commit/pull/30
<br>

# v0.1.7
_Released Jun 2, 2020_
### Modules affected

* `shellcheck`

### Description

You can now configure `shellcheck` to enable certain checks with the `--enable` CLI arg in the `pre-commit` check.

### Special Thanks

Special thanks to @06kellyjac  for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/25
<br>

# v0.1.6
_Released Jun 2, 2020_
### Modules affected

* `terraform-fmt`
* `terraform-validate`
* `terragrunt-hclfmt`

### Description

The terraform and terragrunt related precommit hooks now run the commands in the directory of each module to respect setups that depend on `tfenv` and `tgenv`.

### Special Thanks

Special thanks to @AlainODea  for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/26
<br>

# v0.1.5
_Released Apr 2, 2020_
### Modules affected

* `check-terratest-skip-env` [**NEW**]

### Description

Introduces a new pre-commit hook for checking for uncommented `os.Setenv` calls related to terratest.

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/24
<br>

# v0.1.4
_Released Mar 29, 2020_
### Modules affected

* `terragrunt-hclfmt` [**NEW**]

### Description

Introduces a new pre-commit hook for running `terragrunt hclfmt`.

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/23
<br>

# v0.1.3
_Released Feb 7, 2020_
### Modules affected

* `terraform-fmt`

### Description

Fixes a bug where the `exclude` settings were being ignored with the `terraform-fmt` hook.

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/20
<br>

# v0.1.22
_Released May 12, 2023_
## Modules affected
* `tflint`

## Description
* Updates tflint to use newer `tflint --chdir DIR --filter FILE` syntax for tflint v0.45.0+

_Note that starting with this release, you must use tflint `v0.45.0` or newer._ 

**Full Changelog**: https://github.com/gruntwork-io/pre-commit/compare/v0.1.21...v0.1.22

## Special Thanks
* Special thanks to @WolverineFan  for their contribution!

## Related Links
#95 
<br>

# v0.1.21
_Released Apr 20, 2023_
## Modules affected
* `tflint`

## Description
* Add support to use args in tflint in https://github.com/gruntwork-io/pre-commit/pull/92

**Full Changelog**: https://github.com/gruntwork-io/pre-commit/compare/v0.1.20...v0.1.21

## Special Thanks
* Special thanks to @msgongora for their contribution!
<br>

# v0.1.20
_Released Mar 31, 2023_
## Modules affected
* `golangci-lint`

## Description
* Add a `golangci-lint` hook by @robmorgan in https://github.com/gruntwork-io/pre-commit/pull/93

**Full Changelog**: https://github.com/gruntwork-io/pre-commit/compare/v0.1.19...v0.1.20
<br>

# v0.1.2
_Released Feb 7, 2020_
### Modules affected

* `markdown-link-check` [**NEW**]

### Description

Adds a new pre-commit hook to run [markdown-link-check](https://github.com/tcort/markdown-link-check) on the changed files.


### Related links

* https://github.com/gruntwork-io/pre-commit/pull/19
<br>

# v0.1.19
_Released Mar 10, 2023_
## Modules affected
* `sentinel-fmt`

## Description
* Add sentinel fmt hook by @thepoppingone in https://github.com/gruntwork-io/pre-commit/pull/89

## Special thanks
* @thepoppingone made their first contribution in https://github.com/gruntwork-io/pre-commit/pull/89

**Full Changelog**: https://github.com/gruntwork-io/pre-commit/compare/v0.1.18...v0.1.19
<br>

# v0.1.18
_Released Jan 31, 2023_
### Modules affected

* `mdlink-check.sh`

### Description

Ignore hash links when running `markdown-link-check` , which is not capable of correctly parsing HTML anchor tags.

### Special thanks

Special thanks to @endrec  for the contribution!

### Related links
* https://github.com/gruntwork-io/pre-commit/pull/86
<br>

# v0.1.17
_Released Oct 27, 2021_
### Modules affected

* `tflint`


### Description

Initialise `tflint` to install plugins.

### Special thanks

Special thanks to @ pauloconnor  for the contribution!

### Related links
* https://github.com/gruntwork-io/pre-commit/pull/62
<br>

# v0.1.16
_Released Oct 5, 2021_
### Modules affected

* `packer-validate`


### Description

Fix a bug where `packer-validate` was not targeting correct files.

### Special thanks

Special thanks to @tpdownes  for the contribution!

### Related links
* https://github.com/gruntwork-io/pre-commit/pull/61
<br>

# v0.1.15
_Released Sep 27, 2021_
### Modules affected

* `packer-validate` **[NEW MODULE]**


### Description

A new hook was added to validate packer files.

### Special thanks

Special thanks to @queglay  for the contribution!

### Related links
* https://github.com/gruntwork-io/pre-commit/pull/58

<br>

# v0.1.14
_Released Sep 23, 2021_
### Modules affected

* `terraform-fmt`


### Description

* Save the original value of `$PATH` and change it back after.

### Special thanks

Special thanks to @joshschmitter for the contribution!

### Related links
* https://github.com/gruntwork-io/pre-commit/pull/55
* https://github.com/gruntwork-io/pre-commit/pull/56
<br>

# v0.1.13
_Released Sep 21, 2021_
### Modules affected

* `go-fmt`
* `goimports`
* `helmlint`
* `mdlink-check`
* `terragrunt-hclfmt`
* `yapf`

### Description

* Use `/usr/bin/env bash` in all the scripts, to improve portability.

### Special thanks

Special thanks to @alias-dev for the contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/54
<br>

# v0.1.12
_Released Dec 10, 2020_
### Modules affected

* `terraform-fmt`
* `terraform-validate`

### Description

* The `terraform-validate` hook will now (a) set the `TF_IN_AUTOMATION` variable to reduce Terraform output that isn't relevant in automation, (b) print out each directory it's running in so if you hit an error, you know where to look, (c) save errors until the end, so `validate` runs in all modules, rather than exiting on the first error.
* The `terraform-fmt` hook will now (a) run with `-diff -check` so the differences and affected files are printed to logs, rather than written to disk and (b) save errors until the end, so `fmt` runs in all modules, rather than exiting on the first error.

### Special thanks

Special thanks to @davidalger for the contributions!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/45
* https://github.com/gruntwork-io/pre-commit/pull/46
<br>

# v0.1.11
_Released Nov 25, 2020_
### Modules affected

* `terraform-fmt`
* `terraform-validate`
* `tflint`

### Description

Use `/usr/bin/env bash` instead of `/bin/bash` for better portability.

### Special thanks

Special thanks to @parkalla86 for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/44
<br>

# v0.1.10
_Released Aug 17, 2020_
### Modules affected

* `helmlint`

### Description

`helmlint` hook will now run serially.

### Special thanks

Special thanks to @shmileee for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/35
<br>

# v0.1.1
_Released Jan 6, 2020_
### Modules affected

* `tflint` [**NEW**]

### Description

Adds a new pre-commit hook to run `tflint` on the changed files.


### Special Thanks

Special thanks to @Tensho for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/16
<br>

# v0.1.0
_Released Dec 26, 2019_
### Modules affected

* `terraform-validate` [**NEW**]

### Description

Adds a new pre-commit hook to run `terraform validate`.

Note that starting this release, the minimal version supported is `pre-commit` version 1.13. You must upgrade your runtime environment to at least v1.13 to use the hooks.

### Special Thanks

Special thanks to @Tensho for their contribution!

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/15
<br>

# v0.0.9
_Released May 8, 2019_
### Modules affected

* `gofmt`

### Description

* Fixes a bug in the `gofmt` hook where the package input was not prefixed with `./`, so golang was treating the input like a package as opposed to a directory.

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/13
<br>

# v0.0.8
_Released Feb 22, 2019_
### Modules affected

* `shellcheck`

### Description

* Updated the `shellcheck` pre-commit hook to only match *sh shebangs (e.g., sh, bash, zsh, etc). It used to (inadvertently) match a much broader range of files, such as bats tests.

### Related links

* https://github.com/gruntwork-io/pre-commit/pull/11
<br>

# v0.0.7
_Released Feb 8, 2019_
## Hooks affected

- `helmlint`

## Description

- Add support for a chart that intentionally leaves required variables out. Specifically, the pre-commit hook will now look for an additional `values.yaml` file called `linter_values.yaml`, which will be merged with the charts `values.yaml`. See [the README](https://github.com/gruntwork-io/pre-commit/tree/v0.0.7#linter_valuesyaml) for more information.

## Reference

https://github.com/gruntwork-io/pre-commit/pull/9
<br>

# v0.0.6
_Released Feb 8, 2019_
## Hooks affected

- `helmlint`

## Description

- Updates the `helmlint` hook so that debug logs are logged to `stderr`.
- Updates to the `helmlint` script for better maintainability.

## Reference

https://github.com/gruntwork-io/pre-commit/pull/8
<br>

# v0.0.5
_Released Feb 6, 2019_
## Hooks affected

- `helmlint` [**NEW**]

## Description

- Introduces a new pre-commit hook for running `helm lint` on your helm charts.

## Reference

https://github.com/gruntwork-io/pre-commit/pull/7
<br>

# v0.0.4
_Released Jan 23, 2019_
## Hooks affected

- `yapf`

## Description

- Fix bug in `yapf` hook where it did not exclude files in `.tox` correctly.

## Reference

https://github.com/gruntwork-io/pre-commit/pull/6
<br>

# v0.0.3
_Released Jan 23, 2019_
## Hooks affected

- `yapf` [**NEW**]

## Description

- Add a new precommit hook for `yapf` that can be used to format python files.

## Reference

https://github.com/gruntwork-io/pre-commit/pull/5
<br>

# v0.0.2
_Released May 16, 2018_
https://github.com/gruntwork-io/pre-commit/pull/2: Add `goimports` hook.
<br>

# v0.0.1
_Released May 16, 2018_
https://github.com/gruntwork-io/pre-commit/pull/1: First release!
<br>

