package gobf

/*
#include <duckdb.h>
*/
import "C"
import "fmt"

func Hello() {
	fmt.Println("hello world")
	fmt.Println(C.GoString(C.duckdb_library_version()))
}
