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

// checks if the MyVariantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MyVariantInfo{}

// MyVariantInfo struct for MyVariantInfo
type MyVariantInfo struct {
	ClinVar *ClinVar `json:"clinVar,omitempty"`
	Cosmic *Cosmic `json:"cosmic,omitempty"`
	Dbsnp *Dbsnp `json:"dbsnp,omitempty"`
	GnomadExome *Gnomad `json:"gnomadExome,omitempty"`
	GnomadGenome *Gnomad `json:"gnomadGenome,omitempty"`
	// hgvs
	Hgvs *string `json:"hgvs,omitempty"`
	Mutdb *Mutdb `json:"mutdb,omitempty"`
	Query *string `json:"query,omitempty"`
	Snpeff *Snpeff `json:"snpeff,omitempty"`
	// variant
	Variant *string `json:"variant,omitempty"`
	Vcf *Vcf `json:"vcf,omitempty"`
	// version
	Version *int32 `json:"version,omitempty"`
}

// NewMyVariantInfo instantiates a new MyVariantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyVariantInfo() *MyVariantInfo {
	this := MyVariantInfo{}
	return &this
}

// NewMyVariantInfoWithDefaults instantiates a new MyVariantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyVariantInfoWithDefaults() *MyVariantInfo {
	this := MyVariantInfo{}
	return &this
}

// GetClinVar returns the ClinVar field value if set, zero value otherwise.
func (o *MyVariantInfo) GetClinVar() ClinVar {
	if o == nil || isNil(o.ClinVar) {
		var ret ClinVar
		return ret
	}
	return *o.ClinVar
}

// GetClinVarOk returns a tuple with the ClinVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetClinVarOk() (*ClinVar, bool) {
	if o == nil || isNil(o.ClinVar) {
		return nil, false
	}
	return o.ClinVar, true
}

// HasClinVar returns a boolean if a field has been set.
func (o *MyVariantInfo) HasClinVar() bool {
	if o != nil && !isNil(o.ClinVar) {
		return true
	}

	return false
}

// SetClinVar gets a reference to the given ClinVar and assigns it to the ClinVar field.
func (o *MyVariantInfo) SetClinVar(v ClinVar) {
	o.ClinVar = &v
}

// GetCosmic returns the Cosmic field value if set, zero value otherwise.
func (o *MyVariantInfo) GetCosmic() Cosmic {
	if o == nil || isNil(o.Cosmic) {
		var ret Cosmic
		return ret
	}
	return *o.Cosmic
}

// GetCosmicOk returns a tuple with the Cosmic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetCosmicOk() (*Cosmic, bool) {
	if o == nil || isNil(o.Cosmic) {
		return nil, false
	}
	return o.Cosmic, true
}

// HasCosmic returns a boolean if a field has been set.
func (o *MyVariantInfo) HasCosmic() bool {
	if o != nil && !isNil(o.Cosmic) {
		return true
	}

	return false
}

// SetCosmic gets a reference to the given Cosmic and assigns it to the Cosmic field.
func (o *MyVariantInfo) SetCosmic(v Cosmic) {
	o.Cosmic = &v
}

// GetDbsnp returns the Dbsnp field value if set, zero value otherwise.
func (o *MyVariantInfo) GetDbsnp() Dbsnp {
	if o == nil || isNil(o.Dbsnp) {
		var ret Dbsnp
		return ret
	}
	return *o.Dbsnp
}

// GetDbsnpOk returns a tuple with the Dbsnp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetDbsnpOk() (*Dbsnp, bool) {
	if o == nil || isNil(o.Dbsnp) {
		return nil, false
	}
	return o.Dbsnp, true
}

// HasDbsnp returns a boolean if a field has been set.
func (o *MyVariantInfo) HasDbsnp() bool {
	if o != nil && !isNil(o.Dbsnp) {
		return true
	}

	return false
}

// SetDbsnp gets a reference to the given Dbsnp and assigns it to the Dbsnp field.
func (o *MyVariantInfo) SetDbsnp(v Dbsnp) {
	o.Dbsnp = &v
}

// GetGnomadExome returns the GnomadExome field value if set, zero value otherwise.
func (o *MyVariantInfo) GetGnomadExome() Gnomad {
	if o == nil || isNil(o.GnomadExome) {
		var ret Gnomad
		return ret
	}
	return *o.GnomadExome
}

// GetGnomadExomeOk returns a tuple with the GnomadExome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetGnomadExomeOk() (*Gnomad, bool) {
	if o == nil || isNil(o.GnomadExome) {
		return nil, false
	}
	return o.GnomadExome, true
}

// HasGnomadExome returns a boolean if a field has been set.
func (o *MyVariantInfo) HasGnomadExome() bool {
	if o != nil && !isNil(o.GnomadExome) {
		return true
	}

	return false
}

// SetGnomadExome gets a reference to the given Gnomad and assigns it to the GnomadExome field.
func (o *MyVariantInfo) SetGnomadExome(v Gnomad) {
	o.GnomadExome = &v
}

// GetGnomadGenome returns the GnomadGenome field value if set, zero value otherwise.
func (o *MyVariantInfo) GetGnomadGenome() Gnomad {
	if o == nil || isNil(o.GnomadGenome) {
		var ret Gnomad
		return ret
	}
	return *o.GnomadGenome
}

// GetGnomadGenomeOk returns a tuple with the GnomadGenome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetGnomadGenomeOk() (*Gnomad, bool) {
	if o == nil || isNil(o.GnomadGenome) {
		return nil, false
	}
	return o.GnomadGenome, true
}

// HasGnomadGenome returns a boolean if a field has been set.
func (o *MyVariantInfo) HasGnomadGenome() bool {
	if o != nil && !isNil(o.GnomadGenome) {
		return true
	}

	return false
}

// SetGnomadGenome gets a reference to the given Gnomad and assigns it to the GnomadGenome field.
func (o *MyVariantInfo) SetGnomadGenome(v Gnomad) {
	o.GnomadGenome = &v
}

// GetHgvs returns the Hgvs field value if set, zero value otherwise.
func (o *MyVariantInfo) GetHgvs() string {
	if o == nil || isNil(o.Hgvs) {
		var ret string
		return ret
	}
	return *o.Hgvs
}

// GetHgvsOk returns a tuple with the Hgvs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetHgvsOk() (*string, bool) {
	if o == nil || isNil(o.Hgvs) {
		return nil, false
	}
	return o.Hgvs, true
}

// HasHgvs returns a boolean if a field has been set.
func (o *MyVariantInfo) HasHgvs() bool {
	if o != nil && !isNil(o.Hgvs) {
		return true
	}

	return false
}

// SetHgvs gets a reference to the given string and assigns it to the Hgvs field.
func (o *MyVariantInfo) SetHgvs(v string) {
	o.Hgvs = &v
}

// GetMutdb returns the Mutdb field value if set, zero value otherwise.
func (o *MyVariantInfo) GetMutdb() Mutdb {
	if o == nil || isNil(o.Mutdb) {
		var ret Mutdb
		return ret
	}
	return *o.Mutdb
}

// GetMutdbOk returns a tuple with the Mutdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetMutdbOk() (*Mutdb, bool) {
	if o == nil || isNil(o.Mutdb) {
		return nil, false
	}
	return o.Mutdb, true
}

// HasMutdb returns a boolean if a field has been set.
func (o *MyVariantInfo) HasMutdb() bool {
	if o != nil && !isNil(o.Mutdb) {
		return true
	}

	return false
}

// SetMutdb gets a reference to the given Mutdb and assigns it to the Mutdb field.
func (o *MyVariantInfo) SetMutdb(v Mutdb) {
	o.Mutdb = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MyVariantInfo) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MyVariantInfo) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MyVariantInfo) SetQuery(v string) {
	o.Query = &v
}

// GetSnpeff returns the Snpeff field value if set, zero value otherwise.
func (o *MyVariantInfo) GetSnpeff() Snpeff {
	if o == nil || isNil(o.Snpeff) {
		var ret Snpeff
		return ret
	}
	return *o.Snpeff
}

// GetSnpeffOk returns a tuple with the Snpeff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetSnpeffOk() (*Snpeff, bool) {
	if o == nil || isNil(o.Snpeff) {
		return nil, false
	}
	return o.Snpeff, true
}

// HasSnpeff returns a boolean if a field has been set.
func (o *MyVariantInfo) HasSnpeff() bool {
	if o != nil && !isNil(o.Snpeff) {
		return true
	}

	return false
}

// SetSnpeff gets a reference to the given Snpeff and assigns it to the Snpeff field.
func (o *MyVariantInfo) SetSnpeff(v Snpeff) {
	o.Snpeff = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *MyVariantInfo) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *MyVariantInfo) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *MyVariantInfo) SetVariant(v string) {
	o.Variant = &v
}

// GetVcf returns the Vcf field value if set, zero value otherwise.
func (o *MyVariantInfo) GetVcf() Vcf {
	if o == nil || isNil(o.Vcf) {
		var ret Vcf
		return ret
	}
	return *o.Vcf
}

// GetVcfOk returns a tuple with the Vcf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetVcfOk() (*Vcf, bool) {
	if o == nil || isNil(o.Vcf) {
		return nil, false
	}
	return o.Vcf, true
}

// HasVcf returns a boolean if a field has been set.
func (o *MyVariantInfo) HasVcf() bool {
	if o != nil && !isNil(o.Vcf) {
		return true
	}

	return false
}

// SetVcf gets a reference to the given Vcf and assigns it to the Vcf field.
func (o *MyVariantInfo) SetVcf(v Vcf) {
	o.Vcf = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MyVariantInfo) GetVersion() int32 {
	if o == nil || isNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfo) GetVersionOk() (*int32, bool) {
	if o == nil || isNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MyVariantInfo) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *MyVariantInfo) SetVersion(v int32) {
	o.Version = &v
}

func (o MyVariantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyVariantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClinVar) {
		toSerialize["clinVar"] = o.ClinVar
	}
	if !isNil(o.Cosmic) {
		toSerialize["cosmic"] = o.Cosmic
	}
	if !isNil(o.Dbsnp) {
		toSerialize["dbsnp"] = o.Dbsnp
	}
	if !isNil(o.GnomadExome) {
		toSerialize["gnomadExome"] = o.GnomadExome
	}
	if !isNil(o.GnomadGenome) {
		toSerialize["gnomadGenome"] = o.GnomadGenome
	}
	if !isNil(o.Hgvs) {
		toSerialize["hgvs"] = o.Hgvs
	}
	if !isNil(o.Mutdb) {
		toSerialize["mutdb"] = o.Mutdb
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.Snpeff) {
		toSerialize["snpeff"] = o.Snpeff
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !isNil(o.Vcf) {
		toSerialize["vcf"] = o.Vcf
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableMyVariantInfo struct {
	value *MyVariantInfo
	isSet bool
}

func (v NullableMyVariantInfo) Get() *MyVariantInfo {
	return v.value
}

func (v *NullableMyVariantInfo) Set(val *MyVariantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMyVariantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMyVariantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyVariantInfo(val *MyVariantInfo) *NullableMyVariantInfo {
	return &NullableMyVariantInfo{value: val, isSet: true}
}

func (v NullableMyVariantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyVariantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


