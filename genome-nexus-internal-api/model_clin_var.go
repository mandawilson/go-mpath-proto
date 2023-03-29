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

// checks if the ClinVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClinVar{}

// ClinVar struct for ClinVar
type ClinVar struct {
	// allele_id
	AlleleId *int32 `json:"alleleId,omitempty"`
	// alt
	Alt *string `json:"alt,omitempty"`
	// chrom
	Chrom *string `json:"chrom,omitempty"`
	// cytogenic
	Cytogenic *string `json:"cytogenic,omitempty"`
	Gene *Gene `json:"gene,omitempty"`
	Hg19 *Hg19 `json:"hg19,omitempty"`
	Hg38 *Hg38 `json:"hg38,omitempty"`
	Hgvs *Hgvs `json:"hgvs,omitempty"`
	// license
	License *string `json:"license,omitempty"`
	Rcv []Rcv `json:"rcv,omitempty"`
	// variant_id
	VariantId *int32 `json:"variantId,omitempty"`
}

// NewClinVar instantiates a new ClinVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClinVar() *ClinVar {
	this := ClinVar{}
	return &this
}

// NewClinVarWithDefaults instantiates a new ClinVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClinVarWithDefaults() *ClinVar {
	this := ClinVar{}
	return &this
}

// GetAlleleId returns the AlleleId field value if set, zero value otherwise.
func (o *ClinVar) GetAlleleId() int32 {
	if o == nil || isNil(o.AlleleId) {
		var ret int32
		return ret
	}
	return *o.AlleleId
}

// GetAlleleIdOk returns a tuple with the AlleleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetAlleleIdOk() (*int32, bool) {
	if o == nil || isNil(o.AlleleId) {
		return nil, false
	}
	return o.AlleleId, true
}

// HasAlleleId returns a boolean if a field has been set.
func (o *ClinVar) HasAlleleId() bool {
	if o != nil && !isNil(o.AlleleId) {
		return true
	}

	return false
}

// SetAlleleId gets a reference to the given int32 and assigns it to the AlleleId field.
func (o *ClinVar) SetAlleleId(v int32) {
	o.AlleleId = &v
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *ClinVar) GetAlt() string {
	if o == nil || isNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetAltOk() (*string, bool) {
	if o == nil || isNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *ClinVar) HasAlt() bool {
	if o != nil && !isNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *ClinVar) SetAlt(v string) {
	o.Alt = &v
}

// GetChrom returns the Chrom field value if set, zero value otherwise.
func (o *ClinVar) GetChrom() string {
	if o == nil || isNil(o.Chrom) {
		var ret string
		return ret
	}
	return *o.Chrom
}

// GetChromOk returns a tuple with the Chrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetChromOk() (*string, bool) {
	if o == nil || isNil(o.Chrom) {
		return nil, false
	}
	return o.Chrom, true
}

// HasChrom returns a boolean if a field has been set.
func (o *ClinVar) HasChrom() bool {
	if o != nil && !isNil(o.Chrom) {
		return true
	}

	return false
}

// SetChrom gets a reference to the given string and assigns it to the Chrom field.
func (o *ClinVar) SetChrom(v string) {
	o.Chrom = &v
}

// GetCytogenic returns the Cytogenic field value if set, zero value otherwise.
func (o *ClinVar) GetCytogenic() string {
	if o == nil || isNil(o.Cytogenic) {
		var ret string
		return ret
	}
	return *o.Cytogenic
}

// GetCytogenicOk returns a tuple with the Cytogenic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetCytogenicOk() (*string, bool) {
	if o == nil || isNil(o.Cytogenic) {
		return nil, false
	}
	return o.Cytogenic, true
}

// HasCytogenic returns a boolean if a field has been set.
func (o *ClinVar) HasCytogenic() bool {
	if o != nil && !isNil(o.Cytogenic) {
		return true
	}

	return false
}

// SetCytogenic gets a reference to the given string and assigns it to the Cytogenic field.
func (o *ClinVar) SetCytogenic(v string) {
	o.Cytogenic = &v
}

// GetGene returns the Gene field value if set, zero value otherwise.
func (o *ClinVar) GetGene() Gene {
	if o == nil || isNil(o.Gene) {
		var ret Gene
		return ret
	}
	return *o.Gene
}

// GetGeneOk returns a tuple with the Gene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetGeneOk() (*Gene, bool) {
	if o == nil || isNil(o.Gene) {
		return nil, false
	}
	return o.Gene, true
}

// HasGene returns a boolean if a field has been set.
func (o *ClinVar) HasGene() bool {
	if o != nil && !isNil(o.Gene) {
		return true
	}

	return false
}

// SetGene gets a reference to the given Gene and assigns it to the Gene field.
func (o *ClinVar) SetGene(v Gene) {
	o.Gene = &v
}

// GetHg19 returns the Hg19 field value if set, zero value otherwise.
func (o *ClinVar) GetHg19() Hg19 {
	if o == nil || isNil(o.Hg19) {
		var ret Hg19
		return ret
	}
	return *o.Hg19
}

// GetHg19Ok returns a tuple with the Hg19 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetHg19Ok() (*Hg19, bool) {
	if o == nil || isNil(o.Hg19) {
		return nil, false
	}
	return o.Hg19, true
}

// HasHg19 returns a boolean if a field has been set.
func (o *ClinVar) HasHg19() bool {
	if o != nil && !isNil(o.Hg19) {
		return true
	}

	return false
}

// SetHg19 gets a reference to the given Hg19 and assigns it to the Hg19 field.
func (o *ClinVar) SetHg19(v Hg19) {
	o.Hg19 = &v
}

// GetHg38 returns the Hg38 field value if set, zero value otherwise.
func (o *ClinVar) GetHg38() Hg38 {
	if o == nil || isNil(o.Hg38) {
		var ret Hg38
		return ret
	}
	return *o.Hg38
}

// GetHg38Ok returns a tuple with the Hg38 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetHg38Ok() (*Hg38, bool) {
	if o == nil || isNil(o.Hg38) {
		return nil, false
	}
	return o.Hg38, true
}

// HasHg38 returns a boolean if a field has been set.
func (o *ClinVar) HasHg38() bool {
	if o != nil && !isNil(o.Hg38) {
		return true
	}

	return false
}

// SetHg38 gets a reference to the given Hg38 and assigns it to the Hg38 field.
func (o *ClinVar) SetHg38(v Hg38) {
	o.Hg38 = &v
}

// GetHgvs returns the Hgvs field value if set, zero value otherwise.
func (o *ClinVar) GetHgvs() Hgvs {
	if o == nil || isNil(o.Hgvs) {
		var ret Hgvs
		return ret
	}
	return *o.Hgvs
}

// GetHgvsOk returns a tuple with the Hgvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetHgvsOk() (*Hgvs, bool) {
	if o == nil || isNil(o.Hgvs) {
		return nil, false
	}
	return o.Hgvs, true
}

// HasHgvs returns a boolean if a field has been set.
func (o *ClinVar) HasHgvs() bool {
	if o != nil && !isNil(o.Hgvs) {
		return true
	}

	return false
}

// SetHgvs gets a reference to the given Hgvs and assigns it to the Hgvs field.
func (o *ClinVar) SetHgvs(v Hgvs) {
	o.Hgvs = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *ClinVar) GetLicense() string {
	if o == nil || isNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetLicenseOk() (*string, bool) {
	if o == nil || isNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *ClinVar) HasLicense() bool {
	if o != nil && !isNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *ClinVar) SetLicense(v string) {
	o.License = &v
}

// GetRcv returns the Rcv field value if set, zero value otherwise.
func (o *ClinVar) GetRcv() []Rcv {
	if o == nil || isNil(o.Rcv) {
		var ret []Rcv
		return ret
	}
	return o.Rcv
}

// GetRcvOk returns a tuple with the Rcv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetRcvOk() ([]Rcv, bool) {
	if o == nil || isNil(o.Rcv) {
		return nil, false
	}
	return o.Rcv, true
}

// HasRcv returns a boolean if a field has been set.
func (o *ClinVar) HasRcv() bool {
	if o != nil && !isNil(o.Rcv) {
		return true
	}

	return false
}

// SetRcv gets a reference to the given []Rcv and assigns it to the Rcv field.
func (o *ClinVar) SetRcv(v []Rcv) {
	o.Rcv = v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *ClinVar) GetVariantId() int32 {
	if o == nil || isNil(o.VariantId) {
		var ret int32
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClinVar) GetVariantIdOk() (*int32, bool) {
	if o == nil || isNil(o.VariantId) {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *ClinVar) HasVariantId() bool {
	if o != nil && !isNil(o.VariantId) {
		return true
	}

	return false
}

// SetVariantId gets a reference to the given int32 and assigns it to the VariantId field.
func (o *ClinVar) SetVariantId(v int32) {
	o.VariantId = &v
}

func (o ClinVar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClinVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlleleId) {
		toSerialize["alleleId"] = o.AlleleId
	}
	if !isNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !isNil(o.Chrom) {
		toSerialize["chrom"] = o.Chrom
	}
	if !isNil(o.Cytogenic) {
		toSerialize["cytogenic"] = o.Cytogenic
	}
	if !isNil(o.Gene) {
		toSerialize["gene"] = o.Gene
	}
	if !isNil(o.Hg19) {
		toSerialize["hg19"] = o.Hg19
	}
	if !isNil(o.Hg38) {
		toSerialize["hg38"] = o.Hg38
	}
	if !isNil(o.Hgvs) {
		toSerialize["hgvs"] = o.Hgvs
	}
	if !isNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !isNil(o.Rcv) {
		toSerialize["rcv"] = o.Rcv
	}
	if !isNil(o.VariantId) {
		toSerialize["variantId"] = o.VariantId
	}
	return toSerialize, nil
}

type NullableClinVar struct {
	value *ClinVar
	isSet bool
}

func (v NullableClinVar) Get() *ClinVar {
	return v.value
}

func (v *NullableClinVar) Set(val *ClinVar) {
	v.value = val
	v.isSet = true
}

func (v NullableClinVar) IsSet() bool {
	return v.isSet
}

func (v *NullableClinVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClinVar(val *ClinVar) *NullableClinVar {
	return &NullableClinVar{value: val, isSet: true}
}

func (v NullableClinVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClinVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


