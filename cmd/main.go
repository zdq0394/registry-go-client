package main

import (
	"fmt"

	"github.com/zdq0394/registry-go-client/registry"
)

func getToken(userName, pass string, scopes []string) (registry.AccessToken, error) {
	tokenEnp := "https://authgate-dev.cloudappl.com/v2/token"
	accessToken, err := registry.GetTokenByBasicAuth(tokenEnp, userName, pass, scopes)
	if err != nil {
		fmt.Println(err)
		return accessToken, err
	}
	fmt.Println(accessToken.Token)
	return accessToken, nil

}

func getCatalog(endpoint string, userName, pass string) {
	scopes := []string{
		"registry:catalog:*",
	}
	t, _ := getToken(userName, pass, scopes)
	registry.Catalog(endpoint, t)
}

func getRepoTagsByName(endpoint string, userName, pass string, repoName string) {
	scopes := []string{
		"repository:library/redis:*",
	}
	t, _ := getToken(userName, pass, scopes)
	registry.ListTags(endpoint, repoName, t)
}
func getManifestOfImage(endpoint, repoName, tag string, userName, pass string) {
	scopes := []string{
		"repository:library/redis:*",
	}
	t, _ := getToken(userName, pass, scopes)
	registry.PullManifest(endpoint, repoName, tag, t)
}
func getImageSize(endpoint, repoName, tag string, userName, pass string) {
	scopes := []string{
		"repository:library/redis:*",
	}
	t, _ := getToken(userName, pass, scopes)
	//b, _ := registry.PullManifest(endpoint, repoName, tag, t)
	_, b, _ := registry.PullBlob(endpoint, repoName, tag, t)
	fmt.Println(string(b))
}
func main() {
	userName := "test"
	pass := "you-never-know"
	//getToken(userName, "keadmin", []string{"repository:library/redis:*"})
	endpoint := "https://reg-dev.cloudappl.com"
	getCatalog(endpoint, userName, pass)
	//getRepoTagsByName(endpoint, userName, pass, )
	//getManifestOfImage(endpoint, "library/redis", "latest", userName, pass)
	//getImageSize(endpoint, "library/redis", "sha256:481995377a044d40ca3358e4203fe95eca1d58b98a1d4c2d9cec51c0c4569613", userName, pass)
}
