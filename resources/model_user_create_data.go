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

// checks if the UserCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCreateData{}

// UserCreateData struct for UserCreateData
type UserCreateData struct {
	Type string `json:"type"`
	Attributes UserCreateDataAttributes `json:"attributes"`
}

type _UserCreateData UserCreateData

// NewUserCreateData instantiates a new UserCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateData(type_ string, attributes UserCreateDataAttributes) *UserCreateData {
	this := UserCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUserCreateDataWithDefaults instantiates a new UserCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateDataWithDefaults() *UserCreateData {
	this := UserCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *UserCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UserCreateData) GetAttributes() UserCreateDataAttributes {
	if o == nil {
		var ret UserCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UserCreateData) GetAttributesOk() (*UserCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UserCreateData) SetAttributes(v UserCreateDataAttributes) {
	o.Attributes = v
}

func (o UserCreateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UserCreateData) UnmarshalJSON(data []byte) (err error) {
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

	varUserCreateData := _UserCreateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCreateData)

	if err != nil {
		return err
	}

	*o = UserCreateData(varUserCreateData)

	return err
}

type NullableUserCreateData struct {
	value *UserCreateData
	isSet bool
}

func (v NullableUserCreateData) Get() *UserCreateData {
	return v.value
}

func (v *NullableUserCreateData) Set(val *UserCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateData(val *UserCreateData) *NullableUserCreateData {
	return &NullableUserCreateData{value: val, isSet: true}
}

func (v NullableUserCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


