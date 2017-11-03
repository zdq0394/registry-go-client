package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetTokenByBasicAuth(tokenEndpoint, userName, password string, scopes []string) (AccessToken, error) {
	var accessToken AccessToken
	u, err := url.Parse(tokenEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	q := u.Query()
	//issuer := "authgate-token-issuer"
	service := "token-service"
	q.Add("service", service)
	if scopes != nil {
		for _, scope := range scopes {
			q.Add("scope", scope)
		}
	}

	u.RawQuery = q.Encode()
	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return accessToken, err
	}
	r.SetBasicAuth(userName, password)
	client := getDefaultHTTPClient()
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return accessToken, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return accessToken, err
	}
	err = json.Unmarshal(b, &accessToken)
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}
