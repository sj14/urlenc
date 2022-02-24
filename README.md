# urlenc

A simple tool to convert an URL-encoded or URL-decoded input string.

## Installation

### Precompiled Binaries

Binaries are available for all major platforms. See the [releases](https://github.com/sj14/urlenc/releases) page.

### Homebrew

Using the [Homebrew](https://brew.sh/) package manager for macOS:

```console
brew install sj14/tap/urlenc
```

### Manually

It's also possible to install the current development snapshot with `go install`:

```console
go install github.com/sj14/urlenc
```

## Usage

```text
Usage of urlenc:
  -decode
        perform decoding instead of encoding of the given input
```

## Examples

### Encode 

```
$ urlenc info@example.com
info%40example.com
```

```
$ echo info@example.com | urlenc
info%40example.com
```

### Decode

```
$ urlenc -decode info%40example.com
info@example.com
```

```
$ echo info%40example.com | urlenc
info@example.com
```

## FAQ

Q: Are you ashamed of yourself for writing such a small tool in Go which takes so much disk space?  
A: Yes.
