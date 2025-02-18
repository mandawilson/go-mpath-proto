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

// checks if the Clinvar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Clinvar{}

// Clinvar struct for Clinvar
type Clinvar struct {
	AlternateAllele *string `json:"alternateAllele,omitempty"`
	Chromosome *string `json:"chromosome,omitempty"`
	ClinicalSignificance *string `json:"clinicalSignificance,omitempty"`
	ClinvarId *int32 `json:"clinvarId,omitempty"`
	ConflictingClinicalSignificance *string `json:"conflictingClinicalSignificance,omitempty"`
	EndPosition *int32 `json:"endPosition,omitempty"`
	ReferenceAllele *string `json:"referenceAllele,omitempty"`
	StartPosition *int32 `json:"startPosition,omitempty"`
}

// NewClinvar instantiates a new Clinvar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClinvar() *Clinvar {
	this := Clinvar{}
	return &this
}

// NewClinvarWithDefaults instantiates a new Clinvar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClinvarWithDefaults() *Clinvar {
	this := Clinvar{}
	return &this
}

// GetAlternateAllele returns the AlternateAllele field value if set, zero value otherwise.
func (o *Clinvar) GetAlternateAllele() string {
	if o == nil || isNil(o.AlternateAllele) {
		var ret string
		return ret
	}
	return *o.AlternateAllele
}

// GetAlternateAlleleOk returns a tuple with the AlternateAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetAlternateAlleleOk() (*string, bool) {
	if o == nil || isNil(o.AlternateAllele) {
		return nil, false
	}
	return o.AlternateAllele, true
}

// HasAlternateAllele returns a boolean if a field has been set.
func (o *Clinvar) HasAlternateAllele() bool {
	if o != nil && !isNil(o.AlternateAllele) {
		return true
	}

	return false
}

// SetAlternateAllele gets a reference to the given string and assigns it to the AlternateAllele field.
func (o *Clinvar) SetAlternateAllele(v string) {
	o.AlternateAllele = &v
}

// GetChromosome returns the Chromosome field value if set, zero value otherwise.
func (o *Clinvar) GetChromosome() string {
	if o == nil || isNil(o.Chromosome) {
		var ret string
		return ret
	}
	return *o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetChromosomeOk() (*string, bool) {
	if o == nil || isNil(o.Chromosome) {
		return nil, false
	}
	return o.Chromosome, true
}

// HasChromosome returns a boolean if a field has been set.
func (o *Clinvar) HasChromosome() bool {
	if o != nil && !isNil(o.Chromosome) {
		return true
	}

	return false
}

// SetChromosome gets a reference to the given string and assigns it to the Chromosome field.
func (o *Clinvar) SetChromosome(v string) {
	o.Chromosome = &v
}

// GetClinicalSignificance returns the ClinicalSignificance field value if set, zero value otherwise.
func (o *Clinvar) GetClinicalSignificance() string {
	if o == nil || isNil(o.ClinicalSignificance) {
		var ret string
		return ret
	}
	return *o.ClinicalSignificance
}

// GetClinicalSignificanceOk returns a tuple with the ClinicalSignificance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetClinicalSignificanceOk() (*string, bool) {
	if o == nil || isNil(o.ClinicalSignificance) {
		return nil, false
	}
	return o.ClinicalSignificance, true
}

// HasClinicalSignificance returns a boolean if a field has been set.
func (o *Clinvar) HasClinicalSignificance() bool {
	if o != nil && !isNil(o.ClinicalSignificance) {
		return true
	}

	return false
}

// SetClinicalSignificance gets a reference to the given string and assigns it to the ClinicalSignificance field.
func (o *Clinvar) SetClinicalSignificance(v string) {
	o.ClinicalSignificance = &v
}

// GetClinvarId returns the ClinvarId field value if set, zero value otherwise.
func (o *Clinvar) GetClinvarId() int32 {
	if o == nil || isNil(o.ClinvarId) {
		var ret int32
		return ret
	}
	return *o.ClinvarId
}

// GetClinvarIdOk returns a tuple with the ClinvarId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetClinvarIdOk() (*int32, bool) {
	if o == nil || isNil(o.ClinvarId) {
		return nil, false
	}
	return o.ClinvarId, true
}

// HasClinvarId returns a boolean if a field has been set.
func (o *Clinvar) HasClinvarId() bool {
	if o != nil && !isNil(o.ClinvarId) {
		return true
	}

	return false
}

// SetClinvarId gets a reference to the given int32 and assigns it to the ClinvarId field.
func (o *Clinvar) SetClinvarId(v int32) {
	o.ClinvarId = &v
}

// GetConflictingClinicalSignificance returns the ConflictingClinicalSignificance field value if set, zero value otherwise.
func (o *Clinvar) GetConflictingClinicalSignificance() string {
	if o == nil || isNil(o.ConflictingClinicalSignificance) {
		var ret string
		return ret
	}
	return *o.ConflictingClinicalSignificance
}

// GetConflictingClinicalSignificanceOk returns a tuple with the ConflictingClinicalSignificance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetConflictingClinicalSignificanceOk() (*string, bool) {
	if o == nil || isNil(o.ConflictingClinicalSignificance) {
		return nil, false
	}
	return o.ConflictingClinicalSignificance, true
}

// HasConflictingClinicalSignificance returns a boolean if a field has been set.
func (o *Clinvar) HasConflictingClinicalSignificance() bool {
	if o != nil && !isNil(o.ConflictingClinicalSignificance) {
		return true
	}

	return false
}

// SetConflictingClinicalSignificance gets a reference to the given string and assigns it to the ConflictingClinicalSignificance field.
func (o *Clinvar) SetConflictingClinicalSignificance(v string) {
	o.ConflictingClinicalSignificance = &v
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *Clinvar) GetEndPosition() int32 {
	if o == nil || isNil(o.EndPosition) {
		var ret int32
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetEndPositionOk() (*int32, bool) {
	if o == nil || isNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *Clinvar) HasEndPosition() bool {
	if o != nil && !isNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given int32 and assigns it to the EndPosition field.
func (o *Clinvar) SetEndPosition(v int32) {
	o.EndPosition = &v
}

// GetReferenceAllele returns the ReferenceAllele field value if set, zero value otherwise.
func (o *Clinvar) GetReferenceAllele() string {
	if o == nil || isNil(o.ReferenceAllele) {
		var ret string
		return ret
	}
	return *o.ReferenceAllele
}

// GetReferenceAlleleOk returns a tuple with the ReferenceAllele field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetReferenceAlleleOk() (*string, bool) {
	if o == nil || isNil(o.ReferenceAllele) {
		return nil, false
	}
	return o.ReferenceAllele, true
}

// HasReferenceAllele returns a boolean if a field has been set.
func (o *Clinvar) HasReferenceAllele() bool {
	if o != nil && !isNil(o.ReferenceAllele) {
		return true
	}

	return false
}

// SetReferenceAllele gets a reference to the given string and assigns it to the ReferenceAllele field.
func (o *Clinvar) SetReferenceAllele(v string) {
	o.ReferenceAllele = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *Clinvar) GetStartPosition() int32 {
	if o == nil || isNil(o.StartPosition) {
		var ret int32
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Clinvar) GetStartPositionOk() (*int32, bool) {
	if o == nil || isNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *Clinvar) HasStartPosition() bool {
	if o != nil && !isNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given int32 and assigns it to the StartPosition field.
func (o *Clinvar) SetStartPosition(v int32) {
	o.StartPosition = &v
}

func (o Clinvar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Clinvar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AlternateAllele) {
		toSerialize["alternateAllele"] = o.AlternateAllele
	}
	if !isNil(o.Chromosome) {
		toSerialize["chromosome"] = o.Chromosome
	}
	if !isNil(o.ClinicalSignificance) {
		toSerialize["clinicalSignificance"] = o.ClinicalSignificance
	}
	if !isNil(o.ClinvarId) {
		toSerialize["clinvarId"] = o.ClinvarId
	}
	if !isNil(o.ConflictingClinicalSignificance) {
		toSerialize["conflictingClinicalSignificance"] = o.ConflictingClinicalSignificance
	}
	if !isNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !isNil(o.ReferenceAllele) {
		toSerialize["referenceAllele"] = o.ReferenceAllele
	}
	if !isNil(o.StartPosition) {
		toSerialize["startPosition"] = o.StartPosition
	}
	return toSerialize, nil
}

type NullableClinvar struct {
	value *Clinvar
	isSet bool
}

func (v NullableClinvar) Get() *Clinvar {
	return v.value
}

func (v *NullableClinvar) Set(val *Clinvar) {
	v.value = val
	v.isSet = true
}

func (v NullableClinvar) IsSet() bool {
	return v.isSet
}

func (v *NullableClinvar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClinvar(val *Clinvar) *NullableClinvar {
	return &NullableClinvar{value: val, isSet: true}
}

func (v NullableClinvar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClinvar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


