package registry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/docker/distribution/manifest/schema1"
)

func (rc *Registry) PullBlob(repoName, reference string) (size int64, data []byte, err error) {
	url := rc.blobPullURL(repoName, reference)
	fmt.Println(url)
	var r *http.Request
	r, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var resp *http.Response
	resp, err = rc.Client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		contengLength := resp.Header.Get(http.CanonicalHeaderKey("Content-Length"))
		size, err = strconv.ParseInt(contengLength, 10, 64)
		if err != nil {
			return
		}
		data, err = ioutil.ReadAll(resp.Body)
	}
	return
}
