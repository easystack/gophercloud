package flavor

import "github.com/gophercloud/gophercloud"

const (
	rootPath     = "lbaas"
	resourcePath = "flavors"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, name string) string {
	return c.ServiceURL(rootPath, resourcePath, name)
}
