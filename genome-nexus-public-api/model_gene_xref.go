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

// checks if the GeneXref type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneXref{}

// GeneXref struct for GeneXref
type GeneXref struct {
	// Database display name
	DbDisplayName string `json:"db_display_name"`
	// Database name
	Dbname string `json:"dbname"`
	// Description
	Description string `json:"description"`
	// Display id
	DisplayId string `json:"display_id"`
	EnsemblGeneId *string `json:"ensemblGeneId,omitempty"`
	// Database info text
	InfoText *string `json:"info_text,omitempty"`
	// Database info type
	InfoTypes *string `json:"info_types,omitempty"`
	// Primary id
	PrimaryId string `json:"primary_id"`
	// Synonyms
	Synonyms []string `json:"synonyms,omitempty"`
	// Version
	Version string `json:"version"`
}

// NewGeneXref instantiates a new GeneXref object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneXref(dbDisplayName string, dbname string, description string, displayId string, primaryId string, version string) *GeneXref {
	this := GeneXref{}
	this.DbDisplayName = dbDisplayName
	this.Dbname = dbname
	this.Description = description
	this.DisplayId = displayId
	this.PrimaryId = primaryId
	this.Version = version
	return &this
}

// NewGeneXrefWithDefaults instantiates a new GeneXref object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneXrefWithDefaults() *GeneXref {
	this := GeneXref{}
	return &this
}

// GetDbDisplayName returns the DbDisplayName field value
func (o *GeneXref) GetDbDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbDisplayName
}

// GetDbDisplayNameOk returns a tuple with the DbDisplayName field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetDbDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbDisplayName, true
}

// SetDbDisplayName sets field value
func (o *GeneXref) SetDbDisplayName(v string) {
	o.DbDisplayName = v
}

// GetDbname returns the Dbname field value
func (o *GeneXref) GetDbname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dbname
}

// GetDbnameOk returns a tuple with the Dbname field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetDbnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dbname, true
}

// SetDbname sets field value
func (o *GeneXref) SetDbname(v string) {
	o.Dbname = v
}

// GetDescription returns the Description field value
func (o *GeneXref) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GeneXref) SetDescription(v string) {
	o.Description = v
}

// GetDisplayId returns the DisplayId field value
func (o *GeneXref) GetDisplayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetDisplayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *GeneXref) SetDisplayId(v string) {
	o.DisplayId = v
}

// GetEnsemblGeneId returns the EnsemblGeneId field value if set, zero value otherwise.
func (o *GeneXref) GetEnsemblGeneId() string {
	if o == nil || isNil(o.EnsemblGeneId) {
		var ret string
		return ret
	}
	return *o.EnsemblGeneId
}

// GetEnsemblGeneIdOk returns a tuple with the EnsemblGeneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneXref) GetEnsemblGeneIdOk() (*string, bool) {
	if o == nil || isNil(o.EnsemblGeneId) {
		return nil, false
	}
	return o.EnsemblGeneId, true
}

// HasEnsemblGeneId returns a boolean if a field has been set.
func (o *GeneXref) HasEnsemblGeneId() bool {
	if o != nil && !isNil(o.EnsemblGeneId) {
		return true
	}

	return false
}

// SetEnsemblGeneId gets a reference to the given string and assigns it to the EnsemblGeneId field.
func (o *GeneXref) SetEnsemblGeneId(v string) {
	o.EnsemblGeneId = &v
}

// GetInfoText returns the InfoText field value if set, zero value otherwise.
func (o *GeneXref) GetInfoText() string {
	if o == nil || isNil(o.InfoText) {
		var ret string
		return ret
	}
	return *o.InfoText
}

// GetInfoTextOk returns a tuple with the InfoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneXref) GetInfoTextOk() (*string, bool) {
	if o == nil || isNil(o.InfoText) {
		return nil, false
	}
	return o.InfoText, true
}

// HasInfoText returns a boolean if a field has been set.
func (o *GeneXref) HasInfoText() bool {
	if o != nil && !isNil(o.InfoText) {
		return true
	}

	return false
}

// SetInfoText gets a reference to the given string and assigns it to the InfoText field.
func (o *GeneXref) SetInfoText(v string) {
	o.InfoText = &v
}

// GetInfoTypes returns the InfoTypes field value if set, zero value otherwise.
func (o *GeneXref) GetInfoTypes() string {
	if o == nil || isNil(o.InfoTypes) {
		var ret string
		return ret
	}
	return *o.InfoTypes
}

// GetInfoTypesOk returns a tuple with the InfoTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneXref) GetInfoTypesOk() (*string, bool) {
	if o == nil || isNil(o.InfoTypes) {
		return nil, false
	}
	return o.InfoTypes, true
}

// HasInfoTypes returns a boolean if a field has been set.
func (o *GeneXref) HasInfoTypes() bool {
	if o != nil && !isNil(o.InfoTypes) {
		return true
	}

	return false
}

// SetInfoTypes gets a reference to the given string and assigns it to the InfoTypes field.
func (o *GeneXref) SetInfoTypes(v string) {
	o.InfoTypes = &v
}

// GetPrimaryId returns the PrimaryId field value
func (o *GeneXref) GetPrimaryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryId
}

// GetPrimaryIdOk returns a tuple with the PrimaryId field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetPrimaryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryId, true
}

// SetPrimaryId sets field value
func (o *GeneXref) SetPrimaryId(v string) {
	o.PrimaryId = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *GeneXref) GetSynonyms() []string {
	if o == nil || isNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneXref) GetSynonymsOk() ([]string, bool) {
	if o == nil || isNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *GeneXref) HasSynonyms() bool {
	if o != nil && !isNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *GeneXref) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetVersion returns the Version field value
func (o *GeneXref) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GeneXref) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GeneXref) SetVersion(v string) {
	o.Version = v
}

func (o GeneXref) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneXref) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["db_display_name"] = o.DbDisplayName
	toSerialize["dbname"] = o.Dbname
	toSerialize["description"] = o.Description
	toSerialize["display_id"] = o.DisplayId
	if !isNil(o.EnsemblGeneId) {
		toSerialize["ensemblGeneId"] = o.EnsemblGeneId
	}
	if !isNil(o.InfoText) {
		toSerialize["info_text"] = o.InfoText
	}
	if !isNil(o.InfoTypes) {
		toSerialize["info_types"] = o.InfoTypes
	}
	toSerialize["primary_id"] = o.PrimaryId
	if !isNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableGeneXref struct {
	value *GeneXref
	isSet bool
}

func (v NullableGeneXref) Get() *GeneXref {
	return v.value
}

func (v *NullableGeneXref) Set(val *GeneXref) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneXref) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneXref) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneXref(val *GeneXref) *NullableGeneXref {
	return &NullableGeneXref{value: val, isSet: true}
}

func (v NullableGeneXref) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneXref) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


