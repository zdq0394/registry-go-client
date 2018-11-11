package main

import (
	r "github.com/zdq0394/registry-go-client/registry"
)

func main() {
	reg := r.NewRegistry("https://registry.docker-cn.com", "")
	reg.Catalog()
}
