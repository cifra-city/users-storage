/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TeamCreateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamCreateDataAttributes{}

// TeamCreateDataAttributes struct for TeamCreateDataAttributes
type TeamCreateDataAttributes struct {
	// Team name
	Name string `json:"name"`
	// Team description
	Description string `json:"description"`
}

type _TeamCreateDataAttributes TeamCreateDataAttributes

// NewTeamCreateDataAttributes instantiates a new TeamCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreateDataAttributes(name string, description string) *TeamCreateDataAttributes {
	this := TeamCreateDataAttributes{}
	this.Name = name
	this.Description = description
	return &this
}

// NewTeamCreateDataAttributesWithDefaults instantiates a new TeamCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateDataAttributesWithDefaults() *TeamCreateDataAttributes {
	this := TeamCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *TeamCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *TeamCreateDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TeamCreateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TeamCreateDataAttributes) SetDescription(v string) {
	o.Description = v
}

func (o TeamCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamCreateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *TeamCreateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTeamCreateDataAttributes := _TeamCreateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamCreateDataAttributes)

	if err != nil {
		return err
	}

	*o = TeamCreateDataAttributes(varTeamCreateDataAttributes)

	return err
}

type NullableTeamCreateDataAttributes struct {
	value *TeamCreateDataAttributes
	isSet bool
}

func (v NullableTeamCreateDataAttributes) Get() *TeamCreateDataAttributes {
	return v.value
}

func (v *NullableTeamCreateDataAttributes) Set(val *TeamCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreateDataAttributes(val *TeamCreateDataAttributes) *NullableTeamCreateDataAttributes {
	return &NullableTeamCreateDataAttributes{value: val, isSet: true}
}

func (v NullableTeamCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


