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

// checks if the TeamData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamData{}

// TeamData struct for TeamData
type TeamData struct {
	// Team ID
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes TeamDataAttributes `json:"attributes"`
	Relationships TeamDataRelationships `json:"relationships"`
}

type _TeamData TeamData

// NewTeamData instantiates a new TeamData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamData(id string, type_ string, attributes TeamDataAttributes, relationships TeamDataRelationships) *TeamData {
	this := TeamData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewTeamDataWithDefaults instantiates a new TeamData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDataWithDefaults() *TeamData {
	this := TeamData{}
	return &this
}

// GetId returns the Id field value
func (o *TeamData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *TeamData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TeamData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *TeamData) GetAttributes() TeamDataAttributes {
	if o == nil {
		var ret TeamDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TeamData) GetAttributesOk() (*TeamDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TeamData) SetAttributes(v TeamDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *TeamData) GetRelationships() TeamDataRelationships {
	if o == nil {
		var ret TeamDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *TeamData) GetRelationshipsOk() (*TeamDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *TeamData) SetRelationships(v TeamDataRelationships) {
	o.Relationships = v
}

func (o TeamData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

func (o *TeamData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
		"relationships",
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

	varTeamData := _TeamData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamData)

	if err != nil {
		return err
	}

	*o = TeamData(varTeamData)

	return err
}

type NullableTeamData struct {
	value *TeamData
	isSet bool
}

func (v NullableTeamData) Get() *TeamData {
	return v.value
}

func (v *NullableTeamData) Set(val *TeamData) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamData) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamData(val *TeamData) *NullableTeamData {
	return &NullableTeamData{value: val, isSet: true}
}

func (v NullableTeamData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


