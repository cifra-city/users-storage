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

// checks if the TeamUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamUpdateData{}

// TeamUpdateData struct for TeamUpdateData
type TeamUpdateData struct {
	// Team ID
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes TeamCreateDataAttributes `json:"attributes"`
}

type _TeamUpdateData TeamUpdateData

// NewTeamUpdateData instantiates a new TeamUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamUpdateData(id string, type_ string, attributes TeamCreateDataAttributes) *TeamUpdateData {
	this := TeamUpdateData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewTeamUpdateDataWithDefaults instantiates a new TeamUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamUpdateDataWithDefaults() *TeamUpdateData {
	this := TeamUpdateData{}
	return &this
}

// GetId returns the Id field value
func (o *TeamUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TeamUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamUpdateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TeamUpdateData) GetAttributes() TeamCreateDataAttributes {
	if o == nil {
		var ret TeamCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamUpdateData) GetAttributesOk() (*TeamCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TeamUpdateData) SetAttributes(v TeamCreateDataAttributes) {
	o.Attributes = v
}

func (o TeamUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *TeamUpdateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varTeamUpdateData := _TeamUpdateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamUpdateData)

	if err != nil {
		return err
	}

	*o = TeamUpdateData(varTeamUpdateData)

	return err
}

type NullableTeamUpdateData struct {
	value *TeamUpdateData
	isSet bool
}

func (v NullableTeamUpdateData) Get() *TeamUpdateData {
	return v.value
}

func (v *NullableTeamUpdateData) Set(val *TeamUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamUpdateData(val *TeamUpdateData) *NullableTeamUpdateData {
	return &NullableTeamUpdateData{value: val, isSet: true}
}

func (v NullableTeamUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


