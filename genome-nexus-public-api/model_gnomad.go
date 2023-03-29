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

// checks if the Gnomad type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gnomad{}

// Gnomad struct for Gnomad
type Gnomad struct {
	AlleleCount *AlleleCount `json:"alleleCount,omitempty"`
	AlleleFrequency *AlleleFrequency `json:"alleleFrequency,omitempty"`
	AlleleNumber *AlleleNumber `json:"alleleNumber,omitempty"`
	Homozygotes *Homozygotes `json:"homozygotes,omitempty"`
}

// NewGnomad instantiates a new Gnomad object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnomad() *Gnomad {
	this := Gnomad{}
	return &this
}

// NewGnomadWithDefaults instantiates a new Gnomad object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnomadWithDefaults() *Gnomad {
	this := Gnomad{}
	return &this
}

// GetAlleleCount returns the AlleleCount field value if set, zero value otherwise.
func (o *Gnomad) GetAlleleCount() AlleleCount {
	if o == nil || isNil(o.AlleleCount) {
		var ret AlleleCount
		return ret
	}
	return *o.AlleleCount
}

// GetAlleleCountOk returns a tuple with the AlleleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gnomad) GetAlleleCountOk() (*AlleleCount, bool) {
	if o == nil || isNil(o.AlleleCount) {
		return nil, false
	}
	return o.AlleleCount, true
}

// HasAlleleCount returns a boolean if a field has been set.
func (o *Gnomad) HasAlleleCount() bool {
	if o != nil && !isNil(o.AlleleCount) {
		return true
	}

	return false
}

// SetAlleleCount gets a reference to the given AlleleCount and assigns it to the AlleleCount field.
func (o *Gnomad) SetAlleleCount(v AlleleCount) {
	o.AlleleCount = &v
}

// GetAlleleFrequency returns the AlleleFrequency field value if set, zero value otherwise.
func (o *Gnomad) GetAlleleFrequency() AlleleFrequency {
	if o == nil || isNil(o.AlleleFrequency) {
		var ret AlleleFrequency
		return ret
	}
	return *o.AlleleFrequency
}

// GetAlleleFrequencyOk returns a tuple with the AlleleFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gnomad) GetAlleleFrequencyOk() (*AlleleFrequency, bool) {
	if o == nil || isNil(o.AlleleFrequency) {
		return nil, false
	}
	return o.AlleleFrequency, true
}

// HasAlleleFrequency returns a boolean if a field has been set.
func (o *Gnomad) HasAlleleFrequency() bool {
	if o != nil && !isNil(o.AlleleFrequency) {
		return true
	}

	return false
}

// SetAlleleFrequency gets a reference to the given AlleleFrequency and assigns it to the AlleleFrequency field.
func (o *Gnomad) SetAlleleFrequency(v AlleleFrequency) {
	o.AlleleFrequency = &v
}

// GetAlleleNumber returns the AlleleNumber field value if set, zero value otherwise.
func (o *Gnomad) GetAlleleNumber() AlleleNumber {
	if o == nil || isNil(o.AlleleNumber) {
		var ret AlleleNumber
		return ret
	}
	return *o.AlleleNumber
}

// GetAlleleNumberOk returns a tuple with the AlleleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gnomad) GetAlleleNumberOk() (*AlleleNumber, bool) {
	if o == nil || isNil(o.AlleleNumber) {
		return nil, false
	}
	return o.AlleleNumber, true
}

// HasAlleleNumber returns a boolean if a field has been set.
func (o *Gnomad) HasAlleleNumber() bool {
	if o != nil && !isNil(o.AlleleNumber) {
		return true
	}

	return false
}

// SetAlleleNumber gets a reference to the given AlleleNumber and assigns it to the AlleleNumber field.
func (o *Gnomad) SetAlleleNumber(v AlleleNumber) {
	o.AlleleNumber = &v
}

// GetHomozygotes returns the Homozygotes field value if set, zero value otherwise.
func (o *Gnomad) GetHomozygotes() Homozygotes {
	if o == nil || isNil(o.Homozygotes) {
		var ret Homozygotes
		return ret
	}
	return *o.Homozygotes
}

// GetHomozygotesOk returns a tuple with the Homozygotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gnomad) GetHomozygotesOk() (*Homozygotes, bool) {
	if o == nil || isNil(o.Homozygotes) {
		return nil, false
	}
	return o.Homozygotes, true
}

// HasHomozygotes returns a boolean if a field has been set.
func (o *Gnomad) HasHomozygotes() bool {
	if o != nil && !isNil(o.Homozygotes) {
		return true
	}

	return false
}

// SetHomozygotes gets a reference to the given Homozygotes and assigns it to the Homozygotes field.
func (o *Gnomad) SetHomozygotes(v Homozygotes) {
	o.Homozygotes = &v
}

func (o Gnomad) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gnomad) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlleleCount) {
		toSerialize["alleleCount"] = o.AlleleCount
	}
	if !isNil(o.AlleleFrequency) {
		toSerialize["alleleFrequency"] = o.AlleleFrequency
	}
	if !isNil(o.AlleleNumber) {
		toSerialize["alleleNumber"] = o.AlleleNumber
	}
	if !isNil(o.Homozygotes) {
		toSerialize["homozygotes"] = o.Homozygotes
	}
	return toSerialize, nil
}

type NullableGnomad struct {
	value *Gnomad
	isSet bool
}

func (v NullableGnomad) Get() *Gnomad {
	return v.value
}

func (v *NullableGnomad) Set(val *Gnomad) {
	v.value = val
	v.isSet = true
}

func (v NullableGnomad) IsSet() bool {
	return v.isSet
}

func (v *NullableGnomad) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnomad(val *Gnomad) *NullableGnomad {
	return &NullableGnomad{value: val, isSet: true}
}

func (v NullableGnomad) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnomad) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


