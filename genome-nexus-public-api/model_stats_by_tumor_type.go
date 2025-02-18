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

// checks if the StatsByTumorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatsByTumorType{}

// StatsByTumorType struct for StatsByTumorType
type StatsByTumorType struct {
	// Median Age at Dx
	AgeAtDx *int32 `json:"ageAtDx,omitempty"`
	// Frequency Of Biallelic
	FBiallelic *float64 `json:"fBiallelic,omitempty"`
	// Frequency Of Cancer Type Count
	FCancerTypeCount *float64 `json:"fCancerTypeCount,omitempty"`
	HrdScore *HrdScore `json:"hrdScore,omitempty"`
	// Msi Score
	MsiScore *float64 `json:"msiScore,omitempty"`
	// Number Of Cancer Type Count
	NCancerTypeCount *int32 `json:"nCancerTypeCount,omitempty"`
	// Number Of Germline Homozygous Per Tumor Type
	NumberOfGermlineHomozygous *int32 `json:"numberOfGermlineHomozygous,omitempty"`
	// Number of Variants with Signature
	NumberWithSig *int32 `json:"numberWithSig,omitempty"`
	// Median TMB
	Tmb *float64 `json:"tmb,omitempty"`
	// Tumor Type
	TumorType *string `json:"tumorType,omitempty"`
}

// NewStatsByTumorType instantiates a new StatsByTumorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsByTumorType() *StatsByTumorType {
	this := StatsByTumorType{}
	return &this
}

// NewStatsByTumorTypeWithDefaults instantiates a new StatsByTumorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsByTumorTypeWithDefaults() *StatsByTumorType {
	this := StatsByTumorType{}
	return &this
}

// GetAgeAtDx returns the AgeAtDx field value if set, zero value otherwise.
func (o *StatsByTumorType) GetAgeAtDx() int32 {
	if o == nil || isNil(o.AgeAtDx) {
		var ret int32
		return ret
	}
	return *o.AgeAtDx
}

// GetAgeAtDxOk returns a tuple with the AgeAtDx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetAgeAtDxOk() (*int32, bool) {
	if o == nil || isNil(o.AgeAtDx) {
		return nil, false
	}
	return o.AgeAtDx, true
}

// HasAgeAtDx returns a boolean if a field has been set.
func (o *StatsByTumorType) HasAgeAtDx() bool {
	if o != nil && !isNil(o.AgeAtDx) {
		return true
	}

	return false
}

// SetAgeAtDx gets a reference to the given int32 and assigns it to the AgeAtDx field.
func (o *StatsByTumorType) SetAgeAtDx(v int32) {
	o.AgeAtDx = &v
}

// GetFBiallelic returns the FBiallelic field value if set, zero value otherwise.
func (o *StatsByTumorType) GetFBiallelic() float64 {
	if o == nil || isNil(o.FBiallelic) {
		var ret float64
		return ret
	}
	return *o.FBiallelic
}

// GetFBiallelicOk returns a tuple with the FBiallelic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetFBiallelicOk() (*float64, bool) {
	if o == nil || isNil(o.FBiallelic) {
		return nil, false
	}
	return o.FBiallelic, true
}

// HasFBiallelic returns a boolean if a field has been set.
func (o *StatsByTumorType) HasFBiallelic() bool {
	if o != nil && !isNil(o.FBiallelic) {
		return true
	}

	return false
}

// SetFBiallelic gets a reference to the given float64 and assigns it to the FBiallelic field.
func (o *StatsByTumorType) SetFBiallelic(v float64) {
	o.FBiallelic = &v
}

// GetFCancerTypeCount returns the FCancerTypeCount field value if set, zero value otherwise.
func (o *StatsByTumorType) GetFCancerTypeCount() float64 {
	if o == nil || isNil(o.FCancerTypeCount) {
		var ret float64
		return ret
	}
	return *o.FCancerTypeCount
}

// GetFCancerTypeCountOk returns a tuple with the FCancerTypeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetFCancerTypeCountOk() (*float64, bool) {
	if o == nil || isNil(o.FCancerTypeCount) {
		return nil, false
	}
	return o.FCancerTypeCount, true
}

// HasFCancerTypeCount returns a boolean if a field has been set.
func (o *StatsByTumorType) HasFCancerTypeCount() bool {
	if o != nil && !isNil(o.FCancerTypeCount) {
		return true
	}

	return false
}

// SetFCancerTypeCount gets a reference to the given float64 and assigns it to the FCancerTypeCount field.
func (o *StatsByTumorType) SetFCancerTypeCount(v float64) {
	o.FCancerTypeCount = &v
}

// GetHrdScore returns the HrdScore field value if set, zero value otherwise.
func (o *StatsByTumorType) GetHrdScore() HrdScore {
	if o == nil || isNil(o.HrdScore) {
		var ret HrdScore
		return ret
	}
	return *o.HrdScore
}

// GetHrdScoreOk returns a tuple with the HrdScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetHrdScoreOk() (*HrdScore, bool) {
	if o == nil || isNil(o.HrdScore) {
		return nil, false
	}
	return o.HrdScore, true
}

// HasHrdScore returns a boolean if a field has been set.
func (o *StatsByTumorType) HasHrdScore() bool {
	if o != nil && !isNil(o.HrdScore) {
		return true
	}

	return false
}

// SetHrdScore gets a reference to the given HrdScore and assigns it to the HrdScore field.
func (o *StatsByTumorType) SetHrdScore(v HrdScore) {
	o.HrdScore = &v
}

// GetMsiScore returns the MsiScore field value if set, zero value otherwise.
func (o *StatsByTumorType) GetMsiScore() float64 {
	if o == nil || isNil(o.MsiScore) {
		var ret float64
		return ret
	}
	return *o.MsiScore
}

// GetMsiScoreOk returns a tuple with the MsiScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetMsiScoreOk() (*float64, bool) {
	if o == nil || isNil(o.MsiScore) {
		return nil, false
	}
	return o.MsiScore, true
}

// HasMsiScore returns a boolean if a field has been set.
func (o *StatsByTumorType) HasMsiScore() bool {
	if o != nil && !isNil(o.MsiScore) {
		return true
	}

	return false
}

// SetMsiScore gets a reference to the given float64 and assigns it to the MsiScore field.
func (o *StatsByTumorType) SetMsiScore(v float64) {
	o.MsiScore = &v
}

// GetNCancerTypeCount returns the NCancerTypeCount field value if set, zero value otherwise.
func (o *StatsByTumorType) GetNCancerTypeCount() int32 {
	if o == nil || isNil(o.NCancerTypeCount) {
		var ret int32
		return ret
	}
	return *o.NCancerTypeCount
}

// GetNCancerTypeCountOk returns a tuple with the NCancerTypeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetNCancerTypeCountOk() (*int32, bool) {
	if o == nil || isNil(o.NCancerTypeCount) {
		return nil, false
	}
	return o.NCancerTypeCount, true
}

// HasNCancerTypeCount returns a boolean if a field has been set.
func (o *StatsByTumorType) HasNCancerTypeCount() bool {
	if o != nil && !isNil(o.NCancerTypeCount) {
		return true
	}

	return false
}

// SetNCancerTypeCount gets a reference to the given int32 and assigns it to the NCancerTypeCount field.
func (o *StatsByTumorType) SetNCancerTypeCount(v int32) {
	o.NCancerTypeCount = &v
}

// GetNumberOfGermlineHomozygous returns the NumberOfGermlineHomozygous field value if set, zero value otherwise.
func (o *StatsByTumorType) GetNumberOfGermlineHomozygous() int32 {
	if o == nil || isNil(o.NumberOfGermlineHomozygous) {
		var ret int32
		return ret
	}
	return *o.NumberOfGermlineHomozygous
}

// GetNumberOfGermlineHomozygousOk returns a tuple with the NumberOfGermlineHomozygous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetNumberOfGermlineHomozygousOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfGermlineHomozygous) {
		return nil, false
	}
	return o.NumberOfGermlineHomozygous, true
}

// HasNumberOfGermlineHomozygous returns a boolean if a field has been set.
func (o *StatsByTumorType) HasNumberOfGermlineHomozygous() bool {
	if o != nil && !isNil(o.NumberOfGermlineHomozygous) {
		return true
	}

	return false
}

// SetNumberOfGermlineHomozygous gets a reference to the given int32 and assigns it to the NumberOfGermlineHomozygous field.
func (o *StatsByTumorType) SetNumberOfGermlineHomozygous(v int32) {
	o.NumberOfGermlineHomozygous = &v
}

// GetNumberWithSig returns the NumberWithSig field value if set, zero value otherwise.
func (o *StatsByTumorType) GetNumberWithSig() int32 {
	if o == nil || isNil(o.NumberWithSig) {
		var ret int32
		return ret
	}
	return *o.NumberWithSig
}

// GetNumberWithSigOk returns a tuple with the NumberWithSig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetNumberWithSigOk() (*int32, bool) {
	if o == nil || isNil(o.NumberWithSig) {
		return nil, false
	}
	return o.NumberWithSig, true
}

// HasNumberWithSig returns a boolean if a field has been set.
func (o *StatsByTumorType) HasNumberWithSig() bool {
	if o != nil && !isNil(o.NumberWithSig) {
		return true
	}

	return false
}

// SetNumberWithSig gets a reference to the given int32 and assigns it to the NumberWithSig field.
func (o *StatsByTumorType) SetNumberWithSig(v int32) {
	o.NumberWithSig = &v
}

// GetTmb returns the Tmb field value if set, zero value otherwise.
func (o *StatsByTumorType) GetTmb() float64 {
	if o == nil || isNil(o.Tmb) {
		var ret float64
		return ret
	}
	return *o.Tmb
}

// GetTmbOk returns a tuple with the Tmb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetTmbOk() (*float64, bool) {
	if o == nil || isNil(o.Tmb) {
		return nil, false
	}
	return o.Tmb, true
}

// HasTmb returns a boolean if a field has been set.
func (o *StatsByTumorType) HasTmb() bool {
	if o != nil && !isNil(o.Tmb) {
		return true
	}

	return false
}

// SetTmb gets a reference to the given float64 and assigns it to the Tmb field.
func (o *StatsByTumorType) SetTmb(v float64) {
	o.Tmb = &v
}

// GetTumorType returns the TumorType field value if set, zero value otherwise.
func (o *StatsByTumorType) GetTumorType() string {
	if o == nil || isNil(o.TumorType) {
		var ret string
		return ret
	}
	return *o.TumorType
}

// GetTumorTypeOk returns a tuple with the TumorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsByTumorType) GetTumorTypeOk() (*string, bool) {
	if o == nil || isNil(o.TumorType) {
		return nil, false
	}
	return o.TumorType, true
}

// HasTumorType returns a boolean if a field has been set.
func (o *StatsByTumorType) HasTumorType() bool {
	if o != nil && !isNil(o.TumorType) {
		return true
	}

	return false
}

// SetTumorType gets a reference to the given string and assigns it to the TumorType field.
func (o *StatsByTumorType) SetTumorType(v string) {
	o.TumorType = &v
}

func (o StatsByTumorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatsByTumorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AgeAtDx) {
		toSerialize["ageAtDx"] = o.AgeAtDx
	}
	if !isNil(o.FBiallelic) {
		toSerialize["fBiallelic"] = o.FBiallelic
	}
	if !isNil(o.FCancerTypeCount) {
		toSerialize["fCancerTypeCount"] = o.FCancerTypeCount
	}
	if !isNil(o.HrdScore) {
		toSerialize["hrdScore"] = o.HrdScore
	}
	if !isNil(o.MsiScore) {
		toSerialize["msiScore"] = o.MsiScore
	}
	if !isNil(o.NCancerTypeCount) {
		toSerialize["nCancerTypeCount"] = o.NCancerTypeCount
	}
	if !isNil(o.NumberOfGermlineHomozygous) {
		toSerialize["numberOfGermlineHomozygous"] = o.NumberOfGermlineHomozygous
	}
	if !isNil(o.NumberWithSig) {
		toSerialize["numberWithSig"] = o.NumberWithSig
	}
	if !isNil(o.Tmb) {
		toSerialize["tmb"] = o.Tmb
	}
	if !isNil(o.TumorType) {
		toSerialize["tumorType"] = o.TumorType
	}
	return toSerialize, nil
}

type NullableStatsByTumorType struct {
	value *StatsByTumorType
	isSet bool
}

func (v NullableStatsByTumorType) Get() *StatsByTumorType {
	return v.value
}

func (v *NullableStatsByTumorType) Set(val *StatsByTumorType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsByTumorType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsByTumorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsByTumorType(val *StatsByTumorType) *NullableStatsByTumorType {
	return &NullableStatsByTumorType{value: val, isSet: true}
}

func (v NullableStatsByTumorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsByTumorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


