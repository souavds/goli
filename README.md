# Goli

Goli is a CLI Tool for running Go commands with colorized output.

> Note: Goli is still a WIP. It has very basic commands and limitations. Feel free to contribute with the project.

## Installing

Using Goli is very easy. First, use `go get` to install the latest version of the library. This command will install the `goli` executable along with the library and its dependencies:

```shell
go get github.com/arthurvdiniz/goli
```

Add `goli` to your binaries folder

```shell
sudo ln -s ${GOPATH}/bin/goli /usr/local/bin/goli 
```

## Commands

### Test

For running test command put your arguments and flags inside single quotes.

```shell
goli test './... -v'
```

After running the command you'll get an example of output:

![Text output command example](/assets/images/test-command.png "Text output command example")
