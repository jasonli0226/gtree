# gTree

Latest Version: `1.1.9`

Updated Log:

```txt
1.0.0       init version
1.0.1       added search command
1.0.2       update search command - added wildcard search
1.0.3       added colorization
1.0.4       update search command - added `--no-recursive` \
            and fixed reading dir as file issue
1.0.5       gtree search - fixed recursive function wrongly invoked
1.0.6       gtree search - fixed not working without `--pattern` issue
1.0.7       gtree search - added search report

1.1.7       upgrade to `github.com/urfave/cli/v2`
1.1.8       update string flags to string slice flags
1.1.9       fixed bugs - negative number input for search numOfLine

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

## Todo

- [x] (2021) fix issue: list - ignore-folder option

- [x] (2021) list - ignore option

- [x] (2021) remove - Remove directories/files with specified pattern/matches

- [ ] list - ignore-folder (with pattern/matches)

- [ ] list - add report for file count

- [x] (2022-03) search - search for specified files

- [x] (2022-03) search - search for specified file patterns/matches

&nbsp;
