package main

import (
	"fmt"
	"log"

	"github.com/liush5/gologger"
)

const (
	LOG_NAME    string = "test"
	LOG_DIR     string = "/Users/liushaohai/test/log"
	LOG_LEVEL   string = "debug"
	LOG_CONSOLE bool   = true
)

func init() {
	gologger.InitLogger(LOG_DIR, LOG_NAME, LOG_LEVEL, LOG_CONSOLE)
}

func main() {
	test("你好", 1)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//	startGin()
	startGin1()
}

func test(args ...interface{}) {
	var in int64
	var str string

	for _, arg := range args {
		switch arg.(type) {
		case string:
			fmt.Println(arg, "is a string value.")
			str = arg.(string)
		case int64:
			fmt.Println(arg, "is an int64 value.")
			in = arg.(int64)

		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}

	fmt.Printf("i=%d str=%s \n", in, str)
	gologger.Debug("djkdjkfj===")

}
