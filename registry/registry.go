package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type RegistryClient struct {
	TokenServer string
	RegServer string
}

func NewRegistryClient(tokenServer, regServer string) *RegistryClient{
	return &RegistryClient{
		TokenServer:tokenServer,
		RegServer:regServer,
	}
}

func (rc  *RegistryClient) Catalog(token AccessToken) {
	url := BuildCatalogUrl(rc.RegServer)
	fmt.Println(url)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := getBearerTokenClient(token.Token)
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
