# lcase

lcase is a simple command line tool that lowercases the given file(s).
i made this because i like to lowercase my readme and other files, but i dont wanna to do it manually. 

## usage

```bash
lcase file1.txt file2.md file3.py ...
```

output
```
lowercased all (3) file(s)
```

## installation

- run `go install github.com/dragsbruh/lcase`
- or if you dont have `go`, see [releases](https://github.com/dragsbruh/lcase/releases)

## example

this readme is a example of lcase. see the first commit.

or if you want to see one live

![lcase demo](assets/demo.gif)

## known issues

- after lcaseing the files, a trailing newline is added to the end of file (if not already there)
- its written in golang, so the source code is forced to have uppercase letters
