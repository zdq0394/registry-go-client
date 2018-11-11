package main

import (
	"fmt"

	r "github.com/zdq0394/registry-go-client/registry"
)

func main() {
	reg := r.NewRegistry("https://registry.docker-cn.com", "")
	data, err := reg.PullManifest("library/busybox", "latest")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
