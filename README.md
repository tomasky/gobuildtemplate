go-build-template
===

### ABOUT
This is a go build template with glide and gb tools

### USAGE

```
$ glide --debug up --use-gopath
$ mkdir vendor/src && cd vendor/src && ln -s ../github.com github.com 
$ cd ../..
$ gb build 
$ gb test -v 
```
