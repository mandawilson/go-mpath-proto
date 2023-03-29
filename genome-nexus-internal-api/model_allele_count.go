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

// checks if the AlleleCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlleleCount{}

// AlleleCount struct for AlleleCount
type AlleleCount struct {
	Ac int32 `json:"ac"`
	AcAfr int32 `json:"ac_afr"`
	AcAmr int32 `json:"ac_amr"`
	AcAsj int32 `json:"ac_asj"`
	AcEas int32 `json:"ac_eas"`
	AcFin int32 `json:"ac_fin"`
	AcNfe int32 `json:"ac_nfe"`
	AcOth int32 `json:"ac_oth"`
	AcSas int32 `json:"ac_sas"`
}

// NewAlleleCount instantiates a new AlleleCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlleleCount(ac int32, acAfr int32, acAmr int32, acAsj int32, acEas int32, acFin int32, acNfe int32, acOth int32, acSas int32) *AlleleCount {
	this := AlleleCount{}
	this.Ac = ac
	this.AcAfr = acAfr
	this.AcAmr = acAmr
	this.AcAsj = acAsj
	this.AcEas = acEas
	this.AcFin = acFin
	this.AcNfe = acNfe
	this.AcOth = acOth
	this.AcSas = acSas
	return &this
}

// NewAlleleCountWithDefaults instantiates a new AlleleCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlleleCountWithDefaults() *AlleleCount {
	this := AlleleCount{}
	return &this
}

// GetAc returns the Ac field value
func (o *AlleleCount) GetAc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ac
}

// GetAcOk returns a tuple with the Ac field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ac, true
}

// SetAc sets field value
func (o *AlleleCount) SetAc(v int32) {
	o.Ac = v
}

// GetAcAfr returns the AcAfr field value
func (o *AlleleCount) GetAcAfr() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcAfr
}

// GetAcAfrOk returns a tuple with the AcAfr field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcAfrOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcAfr, true
}

// SetAcAfr sets field value
func (o *AlleleCount) SetAcAfr(v int32) {
	o.AcAfr = v
}

// GetAcAmr returns the AcAmr field value
func (o *AlleleCount) GetAcAmr() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcAmr
}

// GetAcAmrOk returns a tuple with the AcAmr field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcAmrOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcAmr, true
}

// SetAcAmr sets field value
func (o *AlleleCount) SetAcAmr(v int32) {
	o.AcAmr = v
}

// GetAcAsj returns the AcAsj field value
func (o *AlleleCount) GetAcAsj() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcAsj
}

// GetAcAsjOk returns a tuple with the AcAsj field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcAsjOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcAsj, true
}

// SetAcAsj sets field value
func (o *AlleleCount) SetAcAsj(v int32) {
	o.AcAsj = v
}

// GetAcEas returns the AcEas field value
func (o *AlleleCount) GetAcEas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcEas
}

// GetAcEasOk returns a tuple with the AcEas field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcEasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcEas, true
}

// SetAcEas sets field value
func (o *AlleleCount) SetAcEas(v int32) {
	o.AcEas = v
}

// GetAcFin returns the AcFin field value
func (o *AlleleCount) GetAcFin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcFin
}

// GetAcFinOk returns a tuple with the AcFin field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcFinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcFin, true
}

// SetAcFin sets field value
func (o *AlleleCount) SetAcFin(v int32) {
	o.AcFin = v
}

// GetAcNfe returns the AcNfe field value
func (o *AlleleCount) GetAcNfe() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcNfe
}

// GetAcNfeOk returns a tuple with the AcNfe field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcNfeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcNfe, true
}

// SetAcNfe sets field value
func (o *AlleleCount) SetAcNfe(v int32) {
	o.AcNfe = v
}

// GetAcOth returns the AcOth field value
func (o *AlleleCount) GetAcOth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcOth
}

// GetAcOthOk returns a tuple with the AcOth field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcOthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcOth, true
}

// SetAcOth sets field value
func (o *AlleleCount) SetAcOth(v int32) {
	o.AcOth = v
}

// GetAcSas returns the AcSas field value
func (o *AlleleCount) GetAcSas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcSas
}

// GetAcSasOk returns a tuple with the AcSas field value
// and a boolean to check if the value has been set.
func (o *AlleleCount) GetAcSasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcSas, true
}

// SetAcSas sets field value
func (o *AlleleCount) SetAcSas(v int32) {
	o.AcSas = v
}

func (o AlleleCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlleleCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ac"] = o.Ac
	toSerialize["ac_afr"] = o.AcAfr
	toSerialize["ac_amr"] = o.AcAmr
	toSerialize["ac_asj"] = o.AcAsj
	toSerialize["ac_eas"] = o.AcEas
	toSerialize["ac_fin"] = o.AcFin
	toSerialize["ac_nfe"] = o.AcNfe
	toSerialize["ac_oth"] = o.AcOth
	toSerialize["ac_sas"] = o.AcSas
	return toSerialize, nil
}

type NullableAlleleCount struct {
	value *AlleleCount
	isSet bool
}

func (v NullableAlleleCount) Get() *AlleleCount {
	return v.value
}

func (v *NullableAlleleCount) Set(val *AlleleCount) {
	v.value = val
	v.isSet = true
}

func (v NullableAlleleCount) IsSet() bool {
	return v.isSet
}

func (v *NullableAlleleCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlleleCount(val *AlleleCount) *NullableAlleleCount {
	return &NullableAlleleCount{value: val, isSet: true}
}

func (v NullableAlleleCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlleleCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


