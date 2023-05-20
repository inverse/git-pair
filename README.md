# git-pair

A simple CLI app to make it easier for git based pairing for co-authoring commits.

Sessions are scoped to git repositories by leveraging commit templates.

## Usage

By default the tool will look up commit authors from the git history of the repo you are in. However, you can also maintain a list of authors within your home
directory in a file called `~/.contributors.txt`. This file uses the `Name <email>` format.

To get started run `git-pair [s]tart`, selecting the contributors for the pairing session. 

You can find out the current state of contributors running `git-pair [i]nfo`.

One you are done with that pairing session just run `git-pair [e]nd`. 

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
