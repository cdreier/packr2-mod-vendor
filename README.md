# building with packr2 and vendor folder

this repository contains a minimal go project with go modules enabled. i ran `go mod vendor` to create the vendor folder

everything is running with `GO111MODULE=on`

## working builds

* `go build`
* `go build -mod vendor`
* `packr2 build`

## not working

* `packr2 build -mod vendor`

### error message

```
can't load package: package github.com/cdreier/packr2-mod-vendor/vendor: no Go files in /home/username/gows/src/github.com/cdreier/packr2-mod-vendor/vendor
Error: exit status 1
```

## also working ???

* `packr2 build -mod vendor -v`

in the verbose log i can see an explicit call to `go build -mod vendor -v`, does packr2 non-verbose handle something diffrent?