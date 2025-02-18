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

// checks if the PtmAnnotation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PtmAnnotation{}

// PtmAnnotation struct for PtmAnnotation
type PtmAnnotation struct {
	Annotation [][]PostTranslationalModification `json:"annotation,omitempty"`
	License *string `json:"license,omitempty"`
}

// NewPtmAnnotation instantiates a new PtmAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtmAnnotation() *PtmAnnotation {
	this := PtmAnnotation{}
	return &this
}

// NewPtmAnnotationWithDefaults instantiates a new PtmAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtmAnnotationWithDefaults() *PtmAnnotation {
	this := PtmAnnotation{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *PtmAnnotation) GetAnnotation() [][]PostTranslationalModification {
	if o == nil || isNil(o.Annotation) {
		var ret [][]PostTranslationalModification
		return ret
	}
	return o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtmAnnotation) GetAnnotationOk() ([][]PostTranslationalModification, bool) {
	if o == nil || isNil(o.Annotation) {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *PtmAnnotation) HasAnnotation() bool {
	if o != nil && !isNil(o.Annotation) {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given [][]PostTranslationalModification and assigns it to the Annotation field.
func (o *PtmAnnotation) SetAnnotation(v [][]PostTranslationalModification) {
	o.Annotation = v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *PtmAnnotation) GetLicense() string {
	if o == nil || isNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtmAnnotation) GetLicenseOk() (*string, bool) {
	if o == nil || isNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *PtmAnnotation) HasLicense() bool {
	if o != nil && !isNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *PtmAnnotation) SetLicense(v string) {
	o.License = &v
}

func (o PtmAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PtmAnnotation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Annotation) {
		toSerialize["annotation"] = o.Annotation
	}
	if !isNil(o.License) {
		toSerialize["license"] = o.License
	}
	return toSerialize, nil
}

type NullablePtmAnnotation struct {
	value *PtmAnnotation
	isSet bool
}

func (v NullablePtmAnnotation) Get() *PtmAnnotation {
	return v.value
}

func (v *NullablePtmAnnotation) Set(val *PtmAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullablePtmAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullablePtmAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtmAnnotation(val *PtmAnnotation) *NullablePtmAnnotation {
	return &NullablePtmAnnotation{value: val, isSet: true}
}

func (v NullablePtmAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtmAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


