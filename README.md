# Y
[![Travis
branch](https://img.shields.io/travis/emil2k/y.svg?style=flat)](https://travis-ci.org/emil2k/y)
[![Coverage
Status](https://img.shields.io/coveralls/emil2k/y.svg?style=flat)](https://coveralls.io/r/emil2k/y)

**WARNING: This is a work in progress, if you want to help jump in.**

A Swiss Army knife for vending your own Go packages.

## Installation

```
go install github.com/emil2k/y -o vend
```

## Usage

### `vend init`

For the package in the current working directory copies all external packages
into the specified `[directory]`, while updating all the import paths. External
packages are packages that are not located in a subdirectory or a standard
package. The specified `[directory]` is created if necessary.

The packages are copied into a subdirectory specified by the package name. If
multiple dependencies have the same package name the command will fail and
provide all the duplicates, the user should use the `vend cp` command to place
those packages in unique directories before running `vend init` again to process
the other packages.

```
  vend init [directory]

  -f=false: forces copy, replaces destination folder
  -i=false: ignore hidden files, files starting with a dot
  -v=false: detailed output
```

Example :

```
vend init ./lib
```

### `vend cp`

Copies the package in the `[from]` import path or directory to the `[to]`
directory, updating the necessary import paths for the package in the current
working directory.

```
  vend cp [from] [to]

  -f=false: forces copy, replaces destination folder
  -i=false: ignore hidden files, files starting with a dot
  -v=false: detailed output
```

Example :

```
vend cp image/png ./lib/mypng
```

### `vend mv`

Moves the package in the `[from]` path or directory to the `[to]` directory,
updating the necessary import paths for the package in the current working
directory. The `mv` subcommand cannot be used with standard packages, use
`cp` instead.

```
  vend mv [from] [to]

  -f=false: forces move, replaces destination folder
  -i=false: ignore hidden files, files starting with a dot
  -v=false: detailed output
```

Example :

```
vend mv ./lib/pq ./lib/postgresql
```

### `vend name [path] [name]`

Changes the package name of the package specified by the `[path]` import path or
directory to the `[name]`, updating all the [qualified
identifiers](https://golang.org/ref/spec#Qualified_identifiers) for the package
in the current working directory. Qualified identifiers aren't modified if the
package name is defined during import. The `name` subcommand cannot be used with
standard packages, you must first `cp` the package out of the `GOROOT`.

Example :

```
vend name ./lib/mypq mypq
```

### `vend each [command]`

Changes to the directory of each dependency, outside of the standard library,
for the package in the current working directory and runs the `[command]`.

Example :

```
vend each go test -v .
```

### `vend list`

Lists all the dependencies of the package specified by the `[path]`, if ommitted
defaults to the current working directory. The `[path]` can be specified
relative to the current working directory or as an import path resolved through
the `GOPATH`.

```
  vend list [arguments] [path]

  -c=true: output child packages, stationed inside subdirectories
  -q=false: outputs only import paths
  -s=true: output standard library packages
  -t=true: include test files
  -v=false: outputs import details
```

### `vend info`

Print out information regarding the package specified by the `[path]`, if
ommitted defaults to the current working directory. The `[path]` can be
specified relative to the current working directory or as an import path
resolved through the `GOPATH`.

```
  vend info [arguments] [path]

  -v=false: detailed output
```
