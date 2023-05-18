# Go Pair

A simple CLI app to make it easier for pairing for co-authoring commits.
## Usage

by default it will look up commit authors from the git history of the repo you are in buy you can also maintain a list of authors within your home
directory in a file called `~/.contributors.txt`. Using the `Name <email>` format.

To get started run `git-pair begin`, selecting the contributors in the pairing session. 


One you are done with that pairing session just run `git-pair end`. You can find out the current state of contributors running `git-pair statius`..


## Developing 

Built using asdf simply install dependencies using that or a compatible tool.

### Building

```bash
task build
```

### Format

```bash
task format
```

### Lint

```bash
task lint
```
