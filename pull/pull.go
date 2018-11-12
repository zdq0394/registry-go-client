package pull

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/docker/image"
	r "github.com/zdq0394/registry-go-client/registry"
)

type Puller struct {
	registry *r.Registry
	localDir string
}

func NewPuller(reg *r.Registry, localDir string) *Puller {
	if reg == nil {
		reg = r.DefaultRegistry()
	}
	return &Puller{
		registry: reg,
		localDir: localDir,
	}
}

func (p *Puller) Pull(repoName, reference string) error {
	m2, _, _, err := p.registry.PullManifestAsObjects(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		ref := m2.Config.Digest.String()
		fmt.Println(ref)
		d, _, _ := p.PullConfig(repoName, ref)
		filePath := fmt.Sprintf("%s%c%s", p.localDir, os.PathSeparator, ref[7:])
		fmt.Println(filePath)
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return err
		}
		f.Write(d)
		f.Close()
	}
	return nil
}

func (p *Puller) PullConfig(repoName, reference string) (data []byte, i image.Image, err error) {
	_, data, err = p.registry.PullBlob(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	return
}

// func (p *Puller) PullLayer(repoName, reference string) (err error) {
// 	var data []byte
// 	_, data, err = p.registry.PullBlob(repoName, reference)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
