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
		ref := m2.Target().Digest.String()
		p.PullConfig(repoName, ref)
		for _, layer := range m2.Layers {
			lref := layer.Digest.String()
			p.PullLayer(repoName, lref)
		}
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
	filePath := fmt.Sprintf("%s%c%s", p.localDir, os.PathSeparator, reference[7:])
	fmt.Println(filePath)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write(data)
	f.Close()
	return
}

func (p *Puller) PullLayer(repoName, reference string) (err error) {
	var data []byte
	_, data, err = p.registry.PullBlob(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := fmt.Sprintf("%s%c%s.tar.gz", p.localDir, os.PathSeparator, reference[7:])
	fmt.Println(filePath)
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write(data)
	f.Close()
	return
}
