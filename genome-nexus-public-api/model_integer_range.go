/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"encoding/json"
)

// checks if the IntegerRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegerRange{}

// IntegerRange struct for IntegerRange
type IntegerRange struct {
	End *int32 `json:"end,omitempty"`
	Start *int32 `json:"start,omitempty"`
}

// NewIntegerRange instantiates a new IntegerRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegerRange() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// NewIntegerRangeWithDefaults instantiates a new IntegerRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegerRangeWithDefaults() *IntegerRange {
	this := IntegerRange{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *IntegerRange) GetEnd() int32 {
	if o == nil || isNil(o.End) {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetEndOk() (*int32, bool) {
	if o == nil || isNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *IntegerRange) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *IntegerRange) SetEnd(v int32) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *IntegerRange) GetStart() int32 {
	if o == nil || isNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerRange) GetStartOk() (*int32, bool) {
	if o == nil || isNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *IntegerRange) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *IntegerRange) SetStart(v int32) {
	o.Start = &v
}

func (o IntegerRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegerRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableIntegerRange struct {
	value *IntegerRange
	isSet bool
}

func (v NullableIntegerRange) Get() *IntegerRange {
	return v.value
}

func (v *NullableIntegerRange) Set(val *IntegerRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegerRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegerRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegerRange(val *IntegerRange) *NullableIntegerRange {
	return &NullableIntegerRange{value: val, isSet: true}
}

func (v NullableIntegerRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegerRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


