# git-pair

[![CI](https://github.com/inverse/git-pair/actions/workflows/ci.yml/badge.svg)](https://github.com/inverse/git-pair/actions/workflows/ci.yml)
![GitHub](https://img.shields.io/github/license/inverse/git-pair)

A tool to make it easier for git based pairing for co-authoring commits.

## Installation

### MacOS
```bash
brew install inverse/homebrew-tap/git-pair
```

### Arch: 
```bash
yay -S git-pair-bin
```

###  APT based distro:

Add the following to a new file called `/etc/apt/sources.list.d/inverse-fury.list`

 ```
 deb [trusted=yes] https://apt.fury.io/inverse/ /
```

And then install:

```bash
apt-get update && apt-get install git-pair
```

### For RPM based distro: 

Add the following to a new file called `/etc/yum.repos.d/inverse-fury.repo`

```
[fury]
name=Inverse Private Repo
baseurl=https://yum.fury.io/inverse/
enabled=1
gpgcheck=0
```

And then install:

```bash
dnf install git-pair
```

Binaries are also distributed in the [releases](https://github.com/inverse/git-pair/releases) page. Simply download the archive for your architecture and unpack and add the binary to your path.

## Usage

By default the tool will look up commit authors from the git history of the repo you are in. However, you can also maintain a list of authors within your home
directory in a file called `~/.contributors.txt`. This file uses the `Name <email>` format.

To get started run `git-pair [s]tart`, selecting the contributors for the pairing session. 

You can find out the current state of contributors running `git-pair [i]nfo`.

One you are done with that pairing session just run `git-pair [e]nd`. 

_Note: Sessions are scoped to git repositories by leveraging commit templates._

## Developing 

Built using asdf for managing required development dependencies.

### Building

```bash
task build
```

And you'll find the binary for your architecture in the `dist/` directory.

### Format

```bash
task format
```

### Lint

```bash
task lint
```
