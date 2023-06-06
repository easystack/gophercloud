package availability_zone_profile

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// AvailabilityZoneProfile is the primary octavia az configuration object that
// specifies the compute az and cpu arch, etc.
type AvailabilityZoneProfile struct {

	// The unique ID for the AvailabilityZoneProfile.
	ID string `json:"id"`

	// Human-readable name for the AvailabilityZoneProfile. Does not have to be unique.
	Name string `json:"name"`

	// The name of the provider.
	ProviderName string `json:"provider_name"`

	// The JSON string containing the availability zone metadata.
	AvailabilityZoneData string `json:"availability_zone_data"`
}

// AvailabilityZoneProfilePage is the page returned by a pager when traversing over a
// collection of availability zone profiles.
type AvailabilityZoneProfilePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of availability zone profiles has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r AvailabilityZoneProfilePage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"availabilityzone_profile_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a AvailabilityZoneProfilePage struct is empty.
func (r AvailabilityZoneProfilePage) IsEmpty() (bool, error) {
	if r.StatusCode == 204 {
		return true, nil
	}

	is, err := ExtractAvailabilityZoneProfiles(r)
	return len(is) == 0, err
}

// ExtractAvailabilityZoneProfiles accepts a Page struct, specifically a AvailabilityZoneProfilePage
// struct, and extracts the elements into a slice of AvailabilityZoneProfiles structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractAvailabilityZoneProfiles(r pagination.Page) ([]AvailabilityZoneProfile, error) {
	var s struct {
		AvailabilityZoneProfiles []AvailabilityZoneProfile `json:"AvailabilityZoneProfiles"`
	}
	err := (r.(AvailabilityZoneProfilePage)).ExtractInto(&s)
	return s.AvailabilityZoneProfiles, err
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a availability_zone_profile.
func (r commonResult) Extract() (*AvailabilityZoneProfile, error) {
	var s struct {
		AvailabilityZoneProfile *AvailabilityZoneProfile `json:"availability_zone_profile"`
	}
	err := r.ExtractInto(&s)
	return s.AvailabilityZoneProfile, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a AvailabilityZoneProfile.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a AvailabilityZoneProfile.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a AvailabilityZoneProfile.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
