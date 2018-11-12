package registry

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest/schema2"
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

func (rc *Registry) PullManifest(repoName, reference string) (digest, mediaType string, payload []byte, err error) {
	url := rc.manifestPullURL(repoName, reference)
	fmt.Println(url)
	var r *http.Request
	r, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Header.Add(http.CanonicalHeaderKey("Accept"), MediaTypeManifest)
	var resp *http.Response
	resp, err = rc.Client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode == http.StatusOK {
		digest = resp.Header.Get(http.CanonicalHeaderKey("Docker-Content-Digest"))
		mediaType = resp.Header.Get(http.CanonicalHeaderKey("Content-Type"))
		payload = b
		return
	}
	err = errors.New(fmt.Sprintf("Http return %d:%s", resp.StatusCode, string(b)))
	return
}

func (rc *Registry) PullManifestAsObjects(repoName, reference string) (m2 *schema2.DeserializedManifest, m distribution.Manifest, d distribution.Descriptor, err error) {
	_, _, b, err := rc.PullManifest(repoName, reference)
	if err != nil {
		fmt.Println(err)
		return
	}
	m, d, err = distribution.UnmarshalManifest(MediaTypeManifest, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	m2, ok := m.(*schema2.DeserializedManifest)
	if !ok {
		err = errors.New("Not type: schema2.DeserializedManifest")
		return
	}
	return m2, m, d, err
}
