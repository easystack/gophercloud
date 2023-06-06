package flavor

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Flavor is the primary octavia configuration object that
// specifies the az configuration.
type Flavor struct {
	// The unique ID for the Flavor.
	ID string `json:"id"`

	// Human-readable description for the Flavor.
	Description string `json:"description"`

	// Human-readable name for the Flavor. Does not have to be unique.
	Name string `json:"name"`

	// The UUID of a flavor profile id.
	FlavorProfileID string `json:"flavor_profile_id"`

	// If the resource is available for use.
	Enabled bool `json:"enabled"`
}

// FlavorPage is the page returned by a pager when traversing over a
// collection of flavors.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of flavor has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r FlavorPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"flavor_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a FlavorPage struct is empty.
func (r FlavorPage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractFlavors(r)
	return len(is) == 0, err
}

// ExtractFlavors accepts a Page struct, specifically a FlavorPage
// struct, and extracts the elements into a slice of Flavor structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var s struct {
		Flavors []Flavor `json:"flavors"`
	}
	err := (r.(FlavorPage)).ExtractInto(&s)
	return s.Flavors, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a flavor.
func (r commonResult) Extract() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"flavors"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Flavor.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Flavor.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Flavor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
