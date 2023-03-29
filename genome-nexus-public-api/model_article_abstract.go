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

// checks if the ArticleAbstract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleAbstract{}

// ArticleAbstract struct for ArticleAbstract
type ArticleAbstract struct {
	Abstract *string `json:"abstract,omitempty"`
	Link *string `json:"link,omitempty"`
}

// NewArticleAbstract instantiates a new ArticleAbstract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleAbstract() *ArticleAbstract {
	this := ArticleAbstract{}
	return &this
}

// NewArticleAbstractWithDefaults instantiates a new ArticleAbstract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleAbstractWithDefaults() *ArticleAbstract {
	this := ArticleAbstract{}
	return &this
}

// GetAbstract returns the Abstract field value if set, zero value otherwise.
func (o *ArticleAbstract) GetAbstract() string {
	if o == nil || isNil(o.Abstract) {
		var ret string
		return ret
	}
	return *o.Abstract
}

// GetAbstractOk returns a tuple with the Abstract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAbstract) GetAbstractOk() (*string, bool) {
	if o == nil || isNil(o.Abstract) {
		return nil, false
	}
	return o.Abstract, true
}

// HasAbstract returns a boolean if a field has been set.
func (o *ArticleAbstract) HasAbstract() bool {
	if o != nil && !isNil(o.Abstract) {
		return true
	}

	return false
}

// SetAbstract gets a reference to the given string and assigns it to the Abstract field.
func (o *ArticleAbstract) SetAbstract(v string) {
	o.Abstract = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ArticleAbstract) GetLink() string {
	if o == nil || isNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleAbstract) GetLinkOk() (*string, bool) {
	if o == nil || isNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ArticleAbstract) HasLink() bool {
	if o != nil && !isNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ArticleAbstract) SetLink(v string) {
	o.Link = &v
}

func (o ArticleAbstract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleAbstract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Abstract) {
		toSerialize["abstract"] = o.Abstract
	}
	if !isNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableArticleAbstract struct {
	value *ArticleAbstract
	isSet bool
}

func (v NullableArticleAbstract) Get() *ArticleAbstract {
	return v.value
}

func (v *NullableArticleAbstract) Set(val *ArticleAbstract) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleAbstract) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleAbstract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleAbstract(val *ArticleAbstract) *NullableArticleAbstract {
	return &NullableArticleAbstract{value: val, isSet: true}
}

func (v NullableArticleAbstract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleAbstract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


