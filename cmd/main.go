package main

import (
	"fmt"

	_ "github.com/docker/distribution/manifest/schema2"
	"github.com/zdq0394/registry-go-client/registry"
)

const (
	TokenServer = "https://authgate.qiniu.com/v2/token"
	RegServer   = "https://reg.qiniu.com"
)

func getRegistryClient() *registry.RegistryClient {
	return registry.NewRegistryClient(TokenServer, RegServer)
}

func getToken(userName, pass string, scopes []string) (registry.AccessToken, error) {
	accessToken, err := registry.GetTokenByBasicAuth(TokenServer, userName, pass, scopes)
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
	b, err := c.PullManifest(repoName, tag, t)
	if err != nil {
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
	i, _ := c.PullBlobAsObject(repoName, tag, t)
	fmt.Println(i.Architecture)
	fmt.Println(i.Config.Cmd)
	fmt.Println(i.Config.Entrypoint)
}

func getImageSize(repoName, tag string, userName, pass string) {
	scopes := []string{
		fmt.Sprintf("repository:%s:*", repoName),
	}
	t, _ := getToken(userName, pass, scopes)
	c := getRegistryClient()
	m, d, _ := c.PullManifestAsObjects(repoName, tag, t)
	var size int64
	size += d.Size
	for _, r := range m.References() {
		size += r.Size
	}
	fmt.Printf("Size in Bytes: %d Bytes\n", size)
	fmt.Printf("Size in MB: %f MBytes\n", float64(size)/(1024*1024))

}

func main() {
	userName := "admin"
	pass := "AdMiN@KiRk"
	//getToken(userName, "keadmin", []string{"repository:library/redis:*"})
	//getCatalog(userName, pass)
	//getRepoTagsByName(userName, pass, )
	//getManifestOfImage("library/redis", "latest", userName, pass)
	//getImageConfigBlob("library/redis", "sha256:1fb7b6c8c0d0713640c99dc75f7f39849cb9fc5619c1ba4ff6da286e6af759ee", userName, pass)
	getImageSize("u-1380469282-default/mysql", "5.6", userName, pass)

}
