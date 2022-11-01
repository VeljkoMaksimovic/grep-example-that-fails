package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TESTS_TO_RUN"))
}
