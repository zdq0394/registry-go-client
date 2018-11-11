package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/docker/distribution"
	_ "github.com/docker/distribution/manifest/schema1"
	_ "github.com/docker/distribution/manifest/schema2"
	"github.com/docker/docker/image"
)

const (
	MediaTypeManifest = "application/vnd.docker.distribution.manifest.v2+json"
)

func (rc *Registry) ListTags(repoName string) ([]string, error) {
	url := rc.tagListURL(repoName)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp, err := rc.Client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(b))
	return nil, nil
}

func (rc *Registry) PullManifest(repoName, reference string) ([]byte, error) {
	url := rc.manifestPullURL(repoName, reference)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.Header.Add(http.CanonicalHeaderKey("Accept"), MediaTypeManifest)
	resp, err := rc.Client.Do(r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}

func (rc *Registry) PullManifestAsObjects(repoName, reference string) (m distribution.Manifest, d distribution.Descriptor, err error) {
	var b []byte
	b, err = rc.PullManifest(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	m, d, err = distribution.UnmarshalManifest(MediaTypeManifest, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	return m, d, err
}

func (rc *Registry) PullBlob(repoName, reference string) (size int64, data []byte, err error) {
	url := rc.blobPullURL(repoName, reference)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}

	resp, err := rc.Client.Do(r)
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}
	if resp.StatusCode == http.StatusOK {
		contengLength := resp.Header.Get(http.CanonicalHeaderKey("Content-Length"))
		size, err = strconv.ParseInt(contengLength, 10, 64)
		if err != nil {
			return
		}
		data, err = ioutil.ReadAll(resp.Body)
	}
	defer resp.Body.Close()
	return
}

func (rc *Registry) PullBlobAsObject(repoName, reference string) (i image.Image, err error) {
	_, data, err := rc.PullBlob(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &i)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
