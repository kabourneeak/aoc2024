# Advent of Code 2024

Playing around with the [Advent of Code 2024](https://adventofcode.com/2024) in [Go](https://go.dev/).

## Installing Go in WSL

```sh
# download the binary installer
curl -O -L https://go.dev/dl/go1.23.3.linux-amd64.tar.gz

# install, as per https://go.dev/doc/install
sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz

# apply path modification on each login
cat << 'EOF' >> ~/.profile

# Add Go binary directory to PATH for Go development
export PATH="$PATH:/usr/local/go/bin"
EOF

# install Go debugger
go install -v github.com/go-delve/delve/cmd/dlv@latest
```

## Build

```sh
go build ./cmd/aoccli/
```

## Run

The basic run command is as follows, which will display usage information for the tool.

```sh
./aoccli
```

To understand where to put your test input,
use the `aoccli env` command
and note the `Input dir` path.

```sh
./aoccli env
Binary name: /home/user/code/aoc2024/aoccli
Binary path: /home/user/code/aoc2024
Current dir: /home/user/code/aoc2024
Input dir: /home/user/code/aoc2024/inputs
```

## Test

```sh
go test ./...
```
