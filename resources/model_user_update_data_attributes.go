/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the UserUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateDataAttributes{}

// UserUpdateDataAttributes struct for UserUpdateDataAttributes
type UserUpdateDataAttributes struct {
	// Username
	Username *string `json:"username,omitempty"`
	// User title
	Title *string `json:"title,omitempty"`
	// User status
	Status *string `json:"status,omitempty"`
	// User avatar
	Avatar *string `json:"avatar,omitempty"`
	// User bio
	Bio *string `json:"bio,omitempty"`
}

// NewUserUpdateDataAttributes instantiates a new UserUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateDataAttributes() *UserUpdateDataAttributes {
	this := UserUpdateDataAttributes{}
	return &this
}

// NewUserUpdateDataAttributesWithDefaults instantiates a new UserUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateDataAttributesWithDefaults() *UserUpdateDataAttributes {
	this := UserUpdateDataAttributes{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserUpdateDataAttributes) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserUpdateDataAttributes) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserUpdateDataAttributes) SetUsername(v string) {
	o.Username = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserUpdateDataAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserUpdateDataAttributes) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserUpdateDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserUpdateDataAttributes) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserUpdateDataAttributes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserUpdateDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UserUpdateDataAttributes) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateDataAttributes) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UserUpdateDataAttributes) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UserUpdateDataAttributes) SetAvatar(v string) {
	o.Avatar = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *UserUpdateDataAttributes) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateDataAttributes) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *UserUpdateDataAttributes) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *UserUpdateDataAttributes) SetBio(v string) {
	o.Bio = &v
}

func (o UserUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	return toSerialize, nil
}

type NullableUserUpdateDataAttributes struct {
	value *UserUpdateDataAttributes
	isSet bool
}

func (v NullableUserUpdateDataAttributes) Get() *UserUpdateDataAttributes {
	return v.value
}

func (v *NullableUserUpdateDataAttributes) Set(val *UserUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateDataAttributes(val *UserUpdateDataAttributes) *NullableUserUpdateDataAttributes {
	return &NullableUserUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableUserUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

