package flavor_profile

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorProfileListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	ID           string `q:"id"`
	Name         string `q:"name"`
	ProviderName string `q:"provider_name"`
	FlavorData   string `q:"flavor_data"`
	Limit        int    `q:"limit"`
	Marker       string `q:"marker"`
	SortKey      string `q:"sort_key"`
	SortDir      string `q:"sort_dir"`
}

// ToFlavorProfileListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorProfileListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// flavor profile. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those flavor profiles that are owned by
// the project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToFlavorProfileListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return FlavorProfilePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToFlavorProfileCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable name for the FlavorProfile. Does not have to be unique.
	Name string `json:"name"`

	// The name of the provider.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the flavor metadata.
	FlavorData string `json:"flavor_data"`
}

// ToFlavorProfileCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToFlavorProfileCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "flavorprofile")
}

// Create is an operation which provisions a new flavor profiles based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToFlavorProfileCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular FlavorProfile based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToFlavorProfileUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable name for the FlavorProfile. Does not have to be unique.
	Name string `json:"name"`

	// The name of the provider.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the flavor metadata.
	FlavorData string `json:"flavor_data"`
}

// ToFlavorProfileUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToFlavorProfileUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "flavorprofile")
}

// Update is an operation which modifies the attributes of the specified
// FlavorProfile.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToFlavorProfileUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Put(resourceURL(c, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Delete will permanently delete a particular FlavorProfile based on its
// unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	url := resourceURL(c, id)
	resp, err := c.Delete(url, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
