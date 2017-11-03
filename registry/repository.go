package registry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	MediaTypeManifest = "application/vnd.docker.distribution.manifest.v2+json"
)

func ListTags(endpoint, repoName string, token AccessToken) ([]string, error) {
	url := BuildTagListURL(endpoint, repoName)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := getBearerTokenClient(token.Token)
	resp, err := client.Do(r)
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

func PullManifest(endpoint, repoName, reference string, token AccessToken) ([]byte, error) {
	url := BuildManifestURL(endpoint, repoName, reference)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.Header.Add(http.CanonicalHeaderKey("Accept"), MediaTypeManifest)
	client := getBearerTokenClient(token.Token)
	resp, err := client.Do(r)
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

func PullBlob(endpoint, repoName, digest string, token AccessToken) (size int64, data []byte, err error) {
	url := BuildBlobURL(endpoint, repoName, digest)
	fmt.Println(url)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}
	client := getBearerTokenClient(token.Token)
	resp, err := client.Do(r)
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
