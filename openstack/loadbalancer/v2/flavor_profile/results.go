package flavor_profile

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// FlavorProfile is the primary octavia az configuration object that
// specifies the compute az and cpu arch, etc.
type FlavorProfile struct {

	// The ID of the flavor profile.
	ID string `json:"id"`

	// The name of the flavor profile.
	Name string `json:"name"`

	// The provider this flavor profile is for.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the flavor metadata.
	FlavorData string `json:"flavor_data"`
}

// FlavorProfilePage is the page returned by a pager when traversing over a
// collection of Flavor profiles.
type FlavorProfilePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of flavor profiles has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r FlavorProfilePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"flavor_profile_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a FlavorProfilePage struct is empty.
func (r FlavorProfilePage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractFlavorProfiles(r)
	return len(is) == 0, err
}

// ExtractFlavorProfiles accepts a Page struct, specifically a FlavorProfilePage
// struct, and extracts the elements into a slice of ExtractFlavorProfiles structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractFlavorProfiles(r pagination.Page) ([]FlavorProfile, error) {
	var s struct {
		FlavorProfiles []FlavorProfile `json:"FlavorProfiles"`
	}
	err := (r.(FlavorProfilePage)).ExtractInto(&s)
	return s.FlavorProfiles, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a flavor_profile.
func (r commonResult) Extract() (*FlavorProfile, error) {
	var s struct {
		FlavorProfile *FlavorProfile `json:"flavor_profile"`
	}
	err := r.ExtractInto(&s)
	return s.FlavorProfile, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a FlavorProfile.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a FlavorProfile.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a FlavorProfile.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
