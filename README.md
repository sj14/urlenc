# urlenc

A simple tool to convert an URL-encoded or URL-decoded input string.

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
