# golangcliscaffold

Learn how to create your first Golang CLI tool with Cobra and Viper. This step-by-step guide will show you how to bootstrap your project, add commands and flags, and create a usable CLI tool. By the end of this article, you'll have a working CLI tool that you can use to automate your tasks.

## Tutorial Article

[Step By Step: Using Cobra and Viper to Create Your First Golang CLI Tool](https://medium.com/@matt.wiater/step-by-step-using-cobra-and-viper-to-create-your-first-golang-cli-tool-8050d7675093)

## Command

### getsize

`go build -o bin/getsize && ./bin/getsize --help`

```
This command will display the size of a directory with several different options.

Usage:
  getsize [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dirs        Show the largest directories in the given path.
  files       Show the largest files in the given path.
  help        Help about any command

Flags:
  -d, --debug           Display debugging output in the console. (default: false)
  -h, --help            help for getsize
      --highlight int   Highlight files/directories over this threshold, in MB (default 500)
  -p, --path string     Define the path to scan. (default "/home/matt")
  -v, --verbose         Display more verbose output in console output. (default: false)

Use "getsize [command] --help" for more information about a command.
``````

## Subcommands

### getsize files

`go build -o bin/getsize && ./bin/getsize files --help`

```
Quickly scan a directory and find large files. . Use the flags below to target the output.

Usage:
  getsize files [flags]

Flags:
  -f, --filecount int     Limit the number of files returned (default 10)
  -h, --help              help for files
      --minfilesize int   Minimum size for files in search in MB. (default 50)

Global Flags:
  -d, --debug           Display debugging output in the console. (default: false)
      --highlight int   Highlight files/directories over this threshold, in MB (default 500)
  -p, --path string     Define the path to scan. (default "/home/matt")
  -v, --verbose         Display more verbose output in console output. (default: false)
```

### getsize dirs

`go build -o bin/getsize && ./bin/getsize dirs --help`

```
 
Quickly scan a directory and find large directories. Use the flags below to target the output.

Usage:
  getsize dirs [flags]

Flags:
      --depth int        Depth of directory tree to display (default 2)
  -h, --help             help for dirs
      --mindirsize int   Only display directories larger than this threshold in MB. (default 100)

Global Flags:
  -d, --debug           Display debugging output in the console. (default: false)
      --highlight int   Highlight files/directories over this threshold, in MB (default 500)
  -p, --path string     Define the path to scan. (default "/home/matt")
  -v, --verbose         Display more verbose output in console output. (default: false)

```