package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Registry struct {
	RegistryServer string
	Client         *http.Client
}

func NewRegistry(regServer string, token string) *Registry {
	r := &Registry{
		RegistryServer: regServer,
		Client:         getDefaultHTTPClient(),
	}
	if token != "" {
		r.Client = getBearerTokenClient(token)
	}
	return r
}

func (rc *Registry) Catalog() {
	url := rc.catalogUrl()
	fmt.Println(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := getDefaultHTTPClient()
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var cataLogResp CatalogResp
	err = json.Unmarshal(b, &cataLogResp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cataLogResp.Repositories)
}
