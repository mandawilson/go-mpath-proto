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

// checks if the Cosmic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cosmic{}

// Cosmic struct for Cosmic
type Cosmic struct {
	// alt
	Alt *string `json:"alt,omitempty"`
	// chrom
	Chrom *string `json:"chrom,omitempty"`
	// cosmic_id
	CosmicId *string `json:"cosmicId,omitempty"`
	Hg19 *Hg19 `json:"hg19,omitempty"`
	// _license
	License *string `json:"license,omitempty"`
	// mut_freq
	MutFreq *float64 `json:"mutFreq,omitempty"`
	// mut_nt
	MutNt *string `json:"mutNt,omitempty"`
	// ref
	Ref *string `json:"ref,omitempty"`
	// tumor_site
	TumorSite *string `json:"tumorSite,omitempty"`
}

// NewCosmic instantiates a new Cosmic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCosmic() *Cosmic {
	this := Cosmic{}
	return &this
}

// NewCosmicWithDefaults instantiates a new Cosmic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCosmicWithDefaults() *Cosmic {
	this := Cosmic{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *Cosmic) GetAlt() string {
	if o == nil || isNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetAltOk() (*string, bool) {
	if o == nil || isNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *Cosmic) HasAlt() bool {
	if o != nil && !isNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *Cosmic) SetAlt(v string) {
	o.Alt = &v
}

// GetChrom returns the Chrom field value if set, zero value otherwise.
func (o *Cosmic) GetChrom() string {
	if o == nil || isNil(o.Chrom) {
		var ret string
		return ret
	}
	return *o.Chrom
}

// GetChromOk returns a tuple with the Chrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetChromOk() (*string, bool) {
	if o == nil || isNil(o.Chrom) {
		return nil, false
	}
	return o.Chrom, true
}

// HasChrom returns a boolean if a field has been set.
func (o *Cosmic) HasChrom() bool {
	if o != nil && !isNil(o.Chrom) {
		return true
	}

	return false
}

// SetChrom gets a reference to the given string and assigns it to the Chrom field.
func (o *Cosmic) SetChrom(v string) {
	o.Chrom = &v
}

// GetCosmicId returns the CosmicId field value if set, zero value otherwise.
func (o *Cosmic) GetCosmicId() string {
	if o == nil || isNil(o.CosmicId) {
		var ret string
		return ret
	}
	return *o.CosmicId
}

// GetCosmicIdOk returns a tuple with the CosmicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetCosmicIdOk() (*string, bool) {
	if o == nil || isNil(o.CosmicId) {
		return nil, false
	}
	return o.CosmicId, true
}

// HasCosmicId returns a boolean if a field has been set.
func (o *Cosmic) HasCosmicId() bool {
	if o != nil && !isNil(o.CosmicId) {
		return true
	}

	return false
}

// SetCosmicId gets a reference to the given string and assigns it to the CosmicId field.
func (o *Cosmic) SetCosmicId(v string) {
	o.CosmicId = &v
}

// GetHg19 returns the Hg19 field value if set, zero value otherwise.
func (o *Cosmic) GetHg19() Hg19 {
	if o == nil || isNil(o.Hg19) {
		var ret Hg19
		return ret
	}
	return *o.Hg19
}

// GetHg19Ok returns a tuple with the Hg19 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetHg19Ok() (*Hg19, bool) {
	if o == nil || isNil(o.Hg19) {
		return nil, false
	}
	return o.Hg19, true
}

// HasHg19 returns a boolean if a field has been set.
func (o *Cosmic) HasHg19() bool {
	if o != nil && !isNil(o.Hg19) {
		return true
	}

	return false
}

// SetHg19 gets a reference to the given Hg19 and assigns it to the Hg19 field.
func (o *Cosmic) SetHg19(v Hg19) {
	o.Hg19 = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *Cosmic) GetLicense() string {
	if o == nil || isNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetLicenseOk() (*string, bool) {
	if o == nil || isNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *Cosmic) HasLicense() bool {
	if o != nil && !isNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *Cosmic) SetLicense(v string) {
	o.License = &v
}

// GetMutFreq returns the MutFreq field value if set, zero value otherwise.
func (o *Cosmic) GetMutFreq() float64 {
	if o == nil || isNil(o.MutFreq) {
		var ret float64
		return ret
	}
	return *o.MutFreq
}

// GetMutFreqOk returns a tuple with the MutFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetMutFreqOk() (*float64, bool) {
	if o == nil || isNil(o.MutFreq) {
		return nil, false
	}
	return o.MutFreq, true
}

// HasMutFreq returns a boolean if a field has been set.
func (o *Cosmic) HasMutFreq() bool {
	if o != nil && !isNil(o.MutFreq) {
		return true
	}

	return false
}

// SetMutFreq gets a reference to the given float64 and assigns it to the MutFreq field.
func (o *Cosmic) SetMutFreq(v float64) {
	o.MutFreq = &v
}

// GetMutNt returns the MutNt field value if set, zero value otherwise.
func (o *Cosmic) GetMutNt() string {
	if o == nil || isNil(o.MutNt) {
		var ret string
		return ret
	}
	return *o.MutNt
}

// GetMutNtOk returns a tuple with the MutNt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetMutNtOk() (*string, bool) {
	if o == nil || isNil(o.MutNt) {
		return nil, false
	}
	return o.MutNt, true
}

// HasMutNt returns a boolean if a field has been set.
func (o *Cosmic) HasMutNt() bool {
	if o != nil && !isNil(o.MutNt) {
		return true
	}

	return false
}

// SetMutNt gets a reference to the given string and assigns it to the MutNt field.
func (o *Cosmic) SetMutNt(v string) {
	o.MutNt = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Cosmic) GetRef() string {
	if o == nil || isNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetRefOk() (*string, bool) {
	if o == nil || isNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Cosmic) HasRef() bool {
	if o != nil && !isNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Cosmic) SetRef(v string) {
	o.Ref = &v
}

// GetTumorSite returns the TumorSite field value if set, zero value otherwise.
func (o *Cosmic) GetTumorSite() string {
	if o == nil || isNil(o.TumorSite) {
		var ret string
		return ret
	}
	return *o.TumorSite
}

// GetTumorSiteOk returns a tuple with the TumorSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cosmic) GetTumorSiteOk() (*string, bool) {
	if o == nil || isNil(o.TumorSite) {
		return nil, false
	}
	return o.TumorSite, true
}

// HasTumorSite returns a boolean if a field has been set.
func (o *Cosmic) HasTumorSite() bool {
	if o != nil && !isNil(o.TumorSite) {
		return true
	}

	return false
}

// SetTumorSite gets a reference to the given string and assigns it to the TumorSite field.
func (o *Cosmic) SetTumorSite(v string) {
	o.TumorSite = &v
}

func (o Cosmic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cosmic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !isNil(o.Chrom) {
		toSerialize["chrom"] = o.Chrom
	}
	if !isNil(o.CosmicId) {
		toSerialize["cosmicId"] = o.CosmicId
	}
	if !isNil(o.Hg19) {
		toSerialize["hg19"] = o.Hg19
	}
	if !isNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !isNil(o.MutFreq) {
		toSerialize["mutFreq"] = o.MutFreq
	}
	if !isNil(o.MutNt) {
		toSerialize["mutNt"] = o.MutNt
	}
	if !isNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !isNil(o.TumorSite) {
		toSerialize["tumorSite"] = o.TumorSite
	}
	return toSerialize, nil
}

type NullableCosmic struct {
	value *Cosmic
	isSet bool
}

func (v NullableCosmic) Get() *Cosmic {
	return v.value
}

func (v *NullableCosmic) Set(val *Cosmic) {
	v.value = val
	v.isSet = true
}

func (v NullableCosmic) IsSet() bool {
	return v.isSet
}

func (v *NullableCosmic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCosmic(val *Cosmic) *NullableCosmic {
	return &NullableCosmic{value: val, isSet: true}
}

func (v NullableCosmic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCosmic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


