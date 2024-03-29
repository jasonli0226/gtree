# gTree

Latest Version: `1.1.14`

Updated Log:

```txt
1.0.0       init version
1.0.1       added search command
1.0.2       update search command - added wildcard search
1.0.3       added colorization for mac
1.0.4       update search command - added `--no-recursive` \
            and fixed reading dir as file issue
1.0.5       search - fixed recursive function wrongly invoked
1.0.6       search - fixed not working without `--pattern` issue
1.0.7       search - added search report

1.1.7       upgrade to `github.com/urfave/cli/v2`
1.1.8       update string flags to string slice flags
1.1.9       fixed bugs - negative number input for search numOfLine
1.1.10      added colorization for windows
1.1.11      added checking for symlinks, symlinks are not allowed \
            added line number for search
1.1.12      search - added copy mode
1.1.13      added new feature - make directory / files
1.1.14      search - added default ignore directory, deleted "file" flag

```

&nbsp;

## Issue List

- [ ] Handling on Special Char

```bash
gtree sc -f -t main() -p \*.ts

```

&nbsp;

## Build

```bash
go build -o build

# Mac Build in Windows
docker run --rm -v "${PWD}:/usr/src/app" -w /usr/src/app golang:1.17 env GOOS=darwin GOARCH=amd64 go build -o build

```

&nbsp;
