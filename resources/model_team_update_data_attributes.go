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

// checks if the TeamUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamUpdateDataAttributes{}

// TeamUpdateDataAttributes struct for TeamUpdateDataAttributes
type TeamUpdateDataAttributes struct {
	// Team name
	Name *string `json:"name,omitempty"`
	// Team description
	Description *string `json:"description,omitempty"`
}

// NewTeamUpdateDataAttributes instantiates a new TeamUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamUpdateDataAttributes() *TeamUpdateDataAttributes {
	this := TeamUpdateDataAttributes{}
	return &this
}

// NewTeamUpdateDataAttributesWithDefaults instantiates a new TeamUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamUpdateDataAttributesWithDefaults() *TeamUpdateDataAttributes {
	this := TeamUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamUpdateDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamUpdateDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamUpdateDataAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamUpdateDataAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

func (o TeamUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableTeamUpdateDataAttributes struct {
	value *TeamUpdateDataAttributes
	isSet bool
}

func (v NullableTeamUpdateDataAttributes) Get() *TeamUpdateDataAttributes {
	return v.value
}

func (v *NullableTeamUpdateDataAttributes) Set(val *TeamUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamUpdateDataAttributes(val *TeamUpdateDataAttributes) *NullableTeamUpdateDataAttributes {
	return &NullableTeamUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableTeamUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


