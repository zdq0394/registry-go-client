package registry

import "fmt"

func buildPingURL(endpoint string) string {
	return fmt.Sprintf("%s/v2/", endpoint)
}

func BuildCatalogUrl(endpoint string) string {
	return fmt.Sprintf("%s/v2/_catalog", endpoint)
}

func BuildTagListURL(endpoint, repoName string) string {
	return fmt.Sprintf("%s/v2/%s/tags/list", endpoint, repoName)
}

func BuildManifestURL(endpoint, repoName, reference string) string {
	return fmt.Sprintf("%s/v2/%s/manifests/%s", endpoint, repoName, reference)
}

func BuildBlobURL(endpoint, repoName, reference string) string {
	return fmt.Sprintf("%s/v2/%s/blobs/%s", endpoint, repoName, reference)
}
