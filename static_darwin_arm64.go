package gobf

//#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
//#cgo LDFLAGS: -lduckdb
//#cgo LDFLAGS: -lc++ -L${SRCDIR}/darwin-arm64/
//#include <duckdb.h>
import "C"
