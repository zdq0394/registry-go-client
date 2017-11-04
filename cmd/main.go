package main

import (
	"fmt"

	"github.com/zdq0394/registry-go-client/registry"
	_ "github.com/docker/distribution/manifest/schema2"
)

const(
	TokenServer = "https://authgate-dev.cloudappl.com/v2/token"
	RegServer = "https://reg-dev.cloudappl.com"
)

func getRegistryClient() *registry.RegistryClient{
	return registry.NewRegistryClient(TokenServer, RegServer)
}

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

func getCatalog(userName, pass string) {
	scopes := []string{
		"registry:catalog:*",
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	c.Catalog(t)
}

func getRepoTagsByName(userName, pass string, repoName string) {
	scopes := []string{
		fmt.Sprintf("repository:%s:*", repoName),
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	c.ListTags(repoName, t)
}

func getManifestOfImage(repoName, tag string, userName, pass string) {
	scopes := []string{
		fmt.Sprintf("repository:%s:*", repoName),
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	b, err:=c.PullManifest(repoName, tag, t)
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func getImageConfigBlob(repoName, tag string, userName, pass string) {
	scopes := []string{
		fmt.Sprintf("repository:%s:*", repoName),
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	_, b, _ := c.PullBlob(repoName, tag, t)
	fmt.Println(string(b))
}

func getImageSize(repoName, tag string, userName, pass string) {
	scopes := []string{
		fmt.Sprintf("repository:%s:*", repoName),
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	m, d, _:=c.PullManifestAsObjects(repoName, tag, t)
	var size int64
	size += d.Size
	for _, r:= range m.References() {
		size += r.Size
	}
	fmt.Printf("Size in Bytes: %d Bytes\n", size)
	fmt.Printf("Size in MB: %d MBytes\n", size/(1024*1024))
}

func main() {
	userName := "test"
	pass := "you-never-know"
	//getToken(userName, "keadmin", []string{"repository:library/redis:*"})
	//getCatalog(userName, pass)
	//getRepoTagsByName(userName, pass, )
	//getManifestOfImage("library/redis", "latest", userName, pass)
	//getImageConfigBlob("library/redis", "sha256:481995377a044d40ca3358e4203fe95eca1d58b98a1d4c2d9cec51c0c4569613", userName, pass)
	getImageSize("library/redis", "latest", userName, pass)
}
