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

## Conventions

#### Packages
Are given lower case, single-word names; there should be no need for underscores or mixedCaps.
Use the package structure to help you choose good names.

#### Getters
It's neither idiomatic nor necessary to put Get into the getter's name.
If you have a field called *owner* (lower case, unexported), the getter method should be called **Owner** (upper case, exported), not *GetOwner*.
 
#### Interface names
By convention, one-method interfaces are named by the method name plus an **-er** suffix or similar modification to construct an agent noun: *Reader*, *Writer*, *Formatter*, *CloseNotifier* etc.

#### Mixed caps
Finally, the convention in Go is to use **MixedCaps** or **mixedCaps** rather than underscores to write multiword names.

## Author

Mitsiu Alejandro Carreño Sarabia
