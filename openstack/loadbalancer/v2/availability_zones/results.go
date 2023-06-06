package availability_zones

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// AvailabilityZone is the primary octavia configuration object that
// specifies the az configuration.
type AvailabilityZone struct {
	// Human-readable description for the AvailabilityZone.
	Description string `json:"description"`

	// Human-readable name for the AvailabilityZone. Does not have to be unique.
	Name string `json:"name"`

	// The UUID of a availability zone profile if set.
	AvailabilityZoneProfileID string `json:"availability_zone_profile_id"`

	// If the resource is available for use.
	Enabled bool `json:"enabled"`
}

// AvailabilityZonePage is the page returned by a pager when traversing over a
// collection of availability zones.
type AvailabilityZonePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of availability zones has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r AvailabilityZonePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"availability_zones_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a AvailabilityZonePage struct is empty.
func (r AvailabilityZonePage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractAvailabilityZones(r)
	return len(is) == 0, err
}

// ExtractAvailabilityZones accepts a Page struct, specifically a AvailabilityZonePage
// struct, and extracts the elements into a slice of AvailabilityZone structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractAvailabilityZones(r pagination.Page) ([]AvailabilityZone, error) {
	var s struct {
		AvailabilityZones []AvailabilityZone `json:"availability_zones"`
	}
	err := (r.(AvailabilityZonePage)).ExtractInto(&s)
	return s.AvailabilityZones, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a availabilityzone.
func (r commonResult) Extract() (*AvailabilityZone, error) {
	var s struct {
		AvailabilityZone *AvailabilityZone `json:"availability_zone"`
	}
	err := r.ExtractInto(&s)
	return s.AvailabilityZone, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a AvailabilityZone.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a AvailabilityZone.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a AvailabilityZone.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
