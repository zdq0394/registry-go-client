package registry

import "fmt"

func (rc *Registry) pingURL() string {
	return fmt.Sprintf("%s/v2/", rc.RegistryServer)
}

func (rc *Registry) catalogUrl() string {
	return fmt.Sprintf("%s/v2/_catalog", rc.RegistryServer)
}

func (rc *Registry) tagListURL(repoName string) string {
	return fmt.Sprintf("%s/v2/%s/tags/list", rc.RegistryServer, repoName)
}

func (rc *Registry) manifestPullURL(repoName, reference string) string {
	return fmt.Sprintf("%s/v2/%s/manifests/%s", rc.RegistryServer, repoName, reference)
}

func (rc *Registry) blobPullURL(repoName, reference string) string {
	return fmt.Sprintf("%s/v2/%s/blobs/%s", rc.RegistryServer, repoName, reference)
}
