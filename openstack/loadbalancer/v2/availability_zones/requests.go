package availability_zones

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAvailabilityZoneListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API.
type ListOpts struct {
	Name                      string `q:"name"`
	Description               string `q:"description"`
	AvailabilityZoneProfileID string `q:"availability_zone_profile_id"`
	Enabled                   bool   `q:"enabled"`
	Limit                     int    `q:"limit"`
	Marker                    string `q:"marker"`
	SortKey                   string `q:"sort_key"`
	SortDir                   string `q:"sort_dir"`
}

// ToAvailabilityZoneListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAvailabilityZoneListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// availability zone. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
//
// Default policy settings return only those availability zones that are owned by
// the project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := rootURL(c)
	if opts != nil {
		query, err := opts.ToAvailabilityZoneListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return AvailabilityZonePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAvailabilityZoneCreateMap() (map[string]interface{}, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable description for the AvailabilityZone.
	Description string `json:"description"`

	// Human-readable name for the AvailabilityZone. Does not have to be unique.
	Name string `json:"name"`

	// The UUID of a availability zone profile if set.
	AvailabilityZoneProfileID string `json:"availability_zone_profile_id"`

	// If the resource is available for use.
	Enabled bool `json:"enabled"`
}

// ToAvailabilityZoneCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToAvailabilityZoneCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "availability_zone")
}

// Create is an operation which provisions a new availabilityzone based on the
// configuration defined in the CreateOpts struct. Once the request is
// validated and progress has started on the provisioning process, a
// CreateResult will be returned.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToAvailabilityZoneCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	resp, err := c.Post(rootURL(c), b, &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// Get retrieves a particular availabilityzone based on its unique ID.
func Get(c *gophercloud.ServiceClient, name string) (r GetResult) {
	resp, err := c.Get(resourceURL(c, name), &r.Body, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAvailabilityZoneUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable description for the AvailabilityZone.
	Description string `json:"description"`

	// Human-readable name for the AvailabilityZone. Does not have to be unique.
	Name string `json:"name"`

	// The UUID of a availability zone profile if set.
	AvailabilityZoneProfileID string `json:"availability_zone_profile_id"`

	// If the resource is available for use.
	Enabled bool `json:"enabled"`
}

// ToAvailabilityZoneUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToAvailabilityZoneUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "availability_zone")
}

// Update is an operation which modifies the attributes of the specified
// AvailabilityZone.
func Update(c *gophercloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToAvailabilityZoneUpdateMap()
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

// Delete will permanently delete a particular AvailabilityZone based on its
// unique name.
func Delete(c *gophercloud.ServiceClient, name string) (r DeleteResult) {
	url := resourceURL(c, name)
	resp, err := c.Delete(url, nil)
	_, r.Header, r.Err = gophercloud.ParseResponse(resp, err)
	return
}
