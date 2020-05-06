# monodiff

**IMPORTANT: You can detect modified part of monorepo by using GitHub Action's [on.<push|pull_request>.paths](https://help.github.com/en/actions/reference/workflow-syntax-for-github-actions#onpushpull_requestpaths) more easily.**

monodiff is a simple and language-agnostic tool to detect modified part of monorepo.

Once you declare sub-projects and their dependencies in `monodiff.json`, monodiff reads a list of changed files, typically result of `git diff —name-only origin/master`, from standard input, then outputs affected sub-projects by the changes. Using its output, you can easily build or deploy only affected projects in CI/CD.

## Real-world example

See the following example repoistory to use monodiff with Gradle multi-project build:

https://github.com/orangain/monodiff-example-multi-project

## Install

### Pre-compiled binary

Download from the [releases page](https://github.com/orangain/monodiff/releases) and put it onto your `$PATH`.

### Docker

Docker image is available at Docker Hub.

https://hub.docker.com/r/orangain/monodiff

## Configuration

Create a file `monodiff.json` at the root directory of Git repository. In the `monodiff.json`, you need to declare sub-projects and their dependencies.

This is an example `monodiff.json`.

```json
{
  "apps/account-app": {
    "deps": ["build.gradle.kts", "settings.gradle.kts", "libs/greeter", "libs/profile"]
  },
  "apps/inventory-app": {
    "deps": ["build.gradle.kts", "settings.gradle.kts", "libs/profile"]
  }
}
```

Each key, e.g. `apps/account-app`, represents a directory of the sub-project. When a file under the directory is changed, the sub-project is determined to be changed.

Additionally, each sub-project can have a `deps` list. The `deps` represents dependencies. Each item of the `deps` is a path to a file or directory. When the file or a file under the directory is changed, the sub-project is also determined to be changed.

All the paths in the `monodiff.json` should be a relative path to the root of Git repository.

## CLI Usage

```
$ monodiff --help
NAME:
   monodiff - Simple tool to detect modified part of monorepo

USAGE:
   monodiff [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --prefix value     Prefix of output line
   --suffix value     Suffix of output line
   --separator value  Path separator of output line
   --help, -h         show help (default: false)
```

## License

MIT License.
