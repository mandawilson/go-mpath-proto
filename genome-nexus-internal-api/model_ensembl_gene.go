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

// checks if the EnsemblGene type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnsemblGene{}

// EnsemblGene struct for EnsemblGene
type EnsemblGene struct {
	// Ensembl gene id
	GeneId string `json:"geneId"`
	// Approved Hugo symbol
	HugoSymbol string `json:"hugoSymbol"`
	// Hugo symbol synonyms
	Synonyms []string `json:"synonyms,omitempty"`
	// Previous Hugo symbols
	PreviousSymbols []string `json:"previousSymbols,omitempty"`
	// Entrez Gene Id
	EntrezGeneId *string `json:"entrezGeneId,omitempty"`
}

// NewEnsemblGene instantiates a new EnsemblGene object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnsemblGene(geneId string, hugoSymbol string) *EnsemblGene {
	this := EnsemblGene{}
	this.GeneId = geneId
	this.HugoSymbol = hugoSymbol
	return &this
}

// NewEnsemblGeneWithDefaults instantiates a new EnsemblGene object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnsemblGeneWithDefaults() *EnsemblGene {
	this := EnsemblGene{}
	return &this
}

// GetGeneId returns the GeneId field value
func (o *EnsemblGene) GetGeneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneId
}

// GetGeneIdOk returns a tuple with the GeneId field value
// and a boolean to check if the value has been set.
func (o *EnsemblGene) GetGeneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneId, true
}

// SetGeneId sets field value
func (o *EnsemblGene) SetGeneId(v string) {
	o.GeneId = v
}

// GetHugoSymbol returns the HugoSymbol field value
func (o *EnsemblGene) GetHugoSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HugoSymbol
}

// GetHugoSymbolOk returns a tuple with the HugoSymbol field value
// and a boolean to check if the value has been set.
func (o *EnsemblGene) GetHugoSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HugoSymbol, true
}

// SetHugoSymbol sets field value
func (o *EnsemblGene) SetHugoSymbol(v string) {
	o.HugoSymbol = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *EnsemblGene) GetSynonyms() []string {
	if o == nil || isNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblGene) GetSynonymsOk() ([]string, bool) {
	if o == nil || isNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *EnsemblGene) HasSynonyms() bool {
	if o != nil && !isNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *EnsemblGene) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetPreviousSymbols returns the PreviousSymbols field value if set, zero value otherwise.
func (o *EnsemblGene) GetPreviousSymbols() []string {
	if o == nil || isNil(o.PreviousSymbols) {
		var ret []string
		return ret
	}
	return o.PreviousSymbols
}

// GetPreviousSymbolsOk returns a tuple with the PreviousSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblGene) GetPreviousSymbolsOk() ([]string, bool) {
	if o == nil || isNil(o.PreviousSymbols) {
		return nil, false
	}
	return o.PreviousSymbols, true
}

// HasPreviousSymbols returns a boolean if a field has been set.
func (o *EnsemblGene) HasPreviousSymbols() bool {
	if o != nil && !isNil(o.PreviousSymbols) {
		return true
	}

	return false
}

// SetPreviousSymbols gets a reference to the given []string and assigns it to the PreviousSymbols field.
func (o *EnsemblGene) SetPreviousSymbols(v []string) {
	o.PreviousSymbols = v
}

// GetEntrezGeneId returns the EntrezGeneId field value if set, zero value otherwise.
func (o *EnsemblGene) GetEntrezGeneId() string {
	if o == nil || isNil(o.EntrezGeneId) {
		var ret string
		return ret
	}
	return *o.EntrezGeneId
}

// GetEntrezGeneIdOk returns a tuple with the EntrezGeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblGene) GetEntrezGeneIdOk() (*string, bool) {
	if o == nil || isNil(o.EntrezGeneId) {
		return nil, false
	}
	return o.EntrezGeneId, true
}

// HasEntrezGeneId returns a boolean if a field has been set.
func (o *EnsemblGene) HasEntrezGeneId() bool {
	if o != nil && !isNil(o.EntrezGeneId) {
		return true
	}

	return false
}

// SetEntrezGeneId gets a reference to the given string and assigns it to the EntrezGeneId field.
func (o *EnsemblGene) SetEntrezGeneId(v string) {
	o.EntrezGeneId = &v
}

func (o EnsemblGene) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnsemblGene) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["geneId"] = o.GeneId
	toSerialize["hugoSymbol"] = o.HugoSymbol
	if !isNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !isNil(o.PreviousSymbols) {
		toSerialize["previousSymbols"] = o.PreviousSymbols
	}
	if !isNil(o.EntrezGeneId) {
		toSerialize["entrezGeneId"] = o.EntrezGeneId
	}
	return toSerialize, nil
}

type NullableEnsemblGene struct {
	value *EnsemblGene
	isSet bool
}

func (v NullableEnsemblGene) Get() *EnsemblGene {
	return v.value
}

func (v *NullableEnsemblGene) Set(val *EnsemblGene) {
	v.value = val
	v.isSet = true
}

func (v NullableEnsemblGene) IsSet() bool {
	return v.isSet
}

func (v *NullableEnsemblGene) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnsemblGene(val *EnsemblGene) *NullableEnsemblGene {
	return &NullableEnsemblGene{value: val, isSet: true}
}

func (v NullableEnsemblGene) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnsemblGene) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


