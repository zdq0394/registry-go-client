package main

import (
	"fmt"

	"github.com/zdq0394/registry-go-client/pull"
	r "github.com/zdq0394/registry-go-client/registry"
)

func pullManifest() {
	reg := r.DefaultRegistry()
	_, _, data, err := reg.PullManifest("library/nginx", "latest")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	m2, _, _, err := reg.PullManifestAsObjects("library/nginx", "latest")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m2.Config.Digest)
		for _, l := range m2.Layers {
			fmt.Println(l.Digest)
		}
	}
}

func main() {
	puller := pull.NewPuller(nil, fmt.Sprintf("D:\\var\\lib\\docker\\images"))
	puller.Pull("library/busybox", "latest")
}
