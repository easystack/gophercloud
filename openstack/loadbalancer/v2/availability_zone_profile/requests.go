package availability_zone_profile

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAvailabilityZoneProfileListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	ID                   string `q:"id"`
	Name                 string `q:"name"`
	ProviderName         string `q:"provider_name"`
	AvailabilityZoneData string `q:"availability_zone_data"`
	Limit                int    `q:"limit"`
	Marker               string `q:"marker"`
	SortKey              string `q:"sort_key"`
	SortDir              string `q:"sort_dir"`
}

// ToAvailabilityZoneProfileListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAvailabilityZoneProfileListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// availability zone profile. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those availability zone profiles that are owned by
// the project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToAvailabilityZoneProfileListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return AvailabilityZoneProfilePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAvailabilityZoneProfileCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable name for the AvailabilityZoneProfile. Does not have to be unique.
	Name string `json:"name"`

	// The name of the provider.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the availability zone metadata.
	AvailabilityZoneData string `json:"availability_zone_data"`
}

// ToAvailabilityZoneProfileCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToAvailabilityZoneProfileCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "availability_zone_profile")
}

// Create is an operation which provisions a new availability zone profiles based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAvailabilityZoneProfileCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular AvailabilityZoneProfile based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, id), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAvailabilityZoneProfileUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable name for the AvailabilityZoneProfile. Does not have to be unique.
	Name string `json:"name"`

	// The name of the provider.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the availability zone metadata.
	AvailabilityZoneData string `json:"availability_zone_data"`
}

// ToAvailabilityZoneProfileUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToAvailabilityZoneProfileUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "availability_zone_profile")
}

// Update is an operation which modifies the attributes of the specified
// AvailabilityZoneProfile.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToAvailabilityZoneProfileUpdateMap()
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

// Delete will permanently delete a particular AvailabilityZoneProfile based on its
// unique ID.
func Delete(c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	url := resourceURL(c, id)
	resp, err := c.Delete(url, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
