package main

import (
	"fmt"

	"github.com/cisco-cx/informer/example"
	informer "github.com/cisco-cx/informer/v1"
)

func main() {
	info := informer.NewInfoService(
		example.Program,
		example.License,
		example.URL,
		example.BuildUser,
		example.BuildDate,
		example.GoVersion,
		example.Version,
		example.Revision,
		example.Branch,
	)
	fmt.Println(info)
}
