/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_internal_api

import (
	"encoding/json"
)

// checks if the Alleles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Alleles{}

// Alleles struct for Alleles
type Alleles struct {
	// allele
	Allele *string `json:"allele,omitempty"`
}

// NewAlleles instantiates a new Alleles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlleles() *Alleles {
	this := Alleles{}
	return &this
}

// NewAllelesWithDefaults instantiates a new Alleles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllelesWithDefaults() *Alleles {
	this := Alleles{}
	return &this
}

// GetAllele returns the Allele field value if set, zero value otherwise.
func (o *Alleles) GetAllele() string {
	if o == nil || isNil(o.Allele) {
		var ret string
		return ret
	}
	return *o.Allele
}

// GetAlleleOk returns a tuple with the Allele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alleles) GetAlleleOk() (*string, bool) {
	if o == nil || isNil(o.Allele) {
		return nil, false
	}
	return o.Allele, true
}

// HasAllele returns a boolean if a field has been set.
func (o *Alleles) HasAllele() bool {
	if o != nil && !isNil(o.Allele) {
		return true
	}

	return false
}

// SetAllele gets a reference to the given string and assigns it to the Allele field.
func (o *Alleles) SetAllele(v string) {
	o.Allele = &v
}

func (o Alleles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alleles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Allele) {
		toSerialize["allele"] = o.Allele
	}
	return toSerialize, nil
}

type NullableAlleles struct {
	value *Alleles
	isSet bool
}

func (v NullableAlleles) Get() *Alleles {
	return v.value
}

func (v *NullableAlleles) Set(val *Alleles) {
	v.value = val
	v.isSet = true
}

func (v NullableAlleles) IsSet() bool {
	return v.isSet
}

func (v *NullableAlleles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlleles(val *Alleles) *NullableAlleles {
	return &NullableAlleles{value: val, isSet: true}
}

func (v NullableAlleles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlleles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


