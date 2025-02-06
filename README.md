# go-branch-fun

`main` branch builds like so:
```
CGO_LDFLAGS="-lc++ -lduckdb -L/path/to/lib/darwin_arm64" go build -tags=static_lib 
CGO_LDFLAGS="-lc++ -lduckdb -L/path/to/lib/darwin_arm64" go test -tags=static_lib
# github.com/taniabogatsch/go-branch-fun.test
ld: warning: ignoring duplicate libraries: '-lc++', '-lduckdb'
hello world
v1.1.3
PASS
ok      github.com/taniabogatsch/go-branch-fun  0.457s

```