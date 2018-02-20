# Go Tour
 
This is where i'll add all the basic knowledge from the go tour

## Commands
```
$ go get <github.com/user/project>      //download and install packages and dependencies
```

```
$ go build      // test that the package compiles (compile packages and dependencies)
```

```
$ go install    // build and install that program 
```

```
$ go test       // run the test
```

```
$ go run <file> // compile and run Go program
```

```
$ go doc        // show documentation for package or symbol
```

```
$ godoc <package>   // Show the documentation (godoc regexp | grep parse)
```
## Testing

Create a file *_test.go, containing functions named "TestXXX" with signature (t *testing.T)
**If the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.**

## VS Code Plugins
https://github.com/Microsoft/vscode-go
- [ ] go-outline
- [ ] go-symbols
- [ ] goreturns
- [x] gorename - rename in all proyect "cambiar todas las ocurrencias" (go get -v golang.org/x/tools/cmd/gorename)
- [x] guru -  renombrar funciones, buscar definiciones - (go get -v golang.org/x/tools/cmd/guru)
- [x] godef - renombrar funciones, buscar definiciones (go get -v github.com/rogpeppe/godef)
- [x] golint - go linter (go get -u github.com/golang/lint/golint)
- [x] gopkgs - list your installed Go packages for import (go get -v github.com/uudashr/gopkgs/cmd/gopkgs)
- [x] gocode - autocompletion (go get -v github.com/nsf/gocode)


## Conventions

> The **var** statement declares a list of variables; as in function argument lists, the type is last.

> Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.


#### Packages
Are given lower case, single-word names; there should be no need for underscores or mixedCaps.
Use the package structure to help you choose good names.
**Programs start running in package _main_**
A name is exported if it begins with a capital letter _Pi_ from the math package

#### Getters
It's neither idiomatic nor necessary to put Get into the getter's name.
If you have a field called *owner* (lower case, unexported), the getter method should be called **Owner** (upper case, exported), not *GetOwner*.
 
#### Interface names
By convention, one-method interfaces are named by the method name plus an **-er** suffix or similar modification to construct an agent noun: *Reader*, *Writer*, *Formatter*, *CloseNotifier* etc.

#### Mixed caps
Finally, the convention in Go is to use **MixedCaps** or **mixedCaps** rather than underscores to write multiword names.

## Author

Mitsiu Alejandro Carreño Sarabia
