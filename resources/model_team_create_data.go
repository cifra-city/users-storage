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

// checks if the TeamCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamCreateData{}

// TeamCreateData struct for TeamCreateData
type TeamCreateData struct {
	Type string `json:"type"`
	Attributes TeamCreateDataAttributes `json:"attributes"`
}

type _TeamCreateData TeamCreateData

// NewTeamCreateData instantiates a new TeamCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamCreateData(type_ string, attributes TeamCreateDataAttributes) *TeamCreateData {
	this := TeamCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTeamCreateDataWithDefaults instantiates a new TeamCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamCreateDataWithDefaults() *TeamCreateData {
	this := TeamCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *TeamCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TeamCreateData) GetAttributes() TeamCreateDataAttributes {
	if o == nil {
		var ret TeamCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamCreateData) GetAttributesOk() (*TeamCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TeamCreateData) SetAttributes(v TeamCreateDataAttributes) {
	o.Attributes = v
}

func (o TeamCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *TeamCreateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"attributes",
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

	varTeamCreateData := _TeamCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamCreateData)

	if err != nil {
		return err
	}

	*o = TeamCreateData(varTeamCreateData)

	return err
}

type NullableTeamCreateData struct {
	value *TeamCreateData
	isSet bool
}

func (v NullableTeamCreateData) Get() *TeamCreateData {
	return v.value
}

func (v *NullableTeamCreateData) Set(val *TeamCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamCreateData(val *TeamCreateData) *NullableTeamCreateData {
	return &NullableTeamCreateData{value: val, isSet: true}
}

func (v NullableTeamCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


