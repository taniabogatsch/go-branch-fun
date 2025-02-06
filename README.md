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

On `darwin-arm64`, default building is possible like so:
```
go test
hello world
v1.1.3
PASS
ok      github.com/taniabogatsch/go-branch-fun  0.439s
```

This works like a charm, with a second go repo:
```
package main

import (
	gobf "github.com/taniabogatsch/go-branch-fun"
)

func main() {
	gobf.Hello()
}
```

And the `go.mod` like so:
```
module latenight

go 1.23.4

require github.com/taniabogatsch/go-branch-fun v0.1.1-0.20250206103040-335d572f80c7
```

Likely there is some more magic to check out different branches,
but for now I used 
```
go get github.com/taniabogatsch/go-branch-fun@darwin-arm64
```