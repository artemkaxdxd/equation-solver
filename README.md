# equation-solver

### Description
An application written on Go programming language to solve quadratic equations ( *a * x^2 + b * x + c = 0* ).

## Build and start using source
For this application to work, you'll need to install Go environment. See [go.dev](https://go.dev/) for installation instructions.

When you have installed Go, clone this repository
```
git clone https://github.com/artemkaxdxd/equation-solver.git
```
#### Interactive Mode
To start in interactive mode simply type *go run* command with the name of the source file.

```
go run solver.go
```
#### Non-interactive Mode
In non-interactive mode you need to have a file in .txt format containing three numbers (more or less numbers result in an error) divided by spaces. 
Example of file structure:
```
2 1 -3
```
To start the application type the same command as in interactive mode, but add a file name to the end. 
For example file.txt:
```
go run solver.go file.txt
```
## Revert commit

[ee283f16cb6fcff9f884066d0ff656ad9a430713](https://github.com/artemkaxdxd/equation-solver/commit/ee283f16cb6fcff9f884066d0ff656ad9a430713)
