/*
 * Turing Minimal Openapi Spec for SDK
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EnsemblerInfraConfig struct for EnsemblerInfraConfig
type EnsemblerInfraConfig struct {
	ArtifactUri        *string                     `json:"artifact_uri,omitempty"`
	EnsemblerName      *string                     `json:"ensembler_name,omitempty"`
	ServiceAccountName *string                     `json:"service_account_name,omitempty" validate:"required"`
	Resources          NullableEnsemblingResources `json:"resources,omitempty"`
	RunId              *string                     `json:"run_id,omitempty"`
	Env                *[]EnvVar                   `json:"env,omitempty"`
}

// NewEnsemblerInfraConfig instantiates a new EnsemblerInfraConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnsemblerInfraConfig() *EnsemblerInfraConfig {
	this := EnsemblerInfraConfig{}
	return &this
}

// NewEnsemblerInfraConfigWithDefaults instantiates a new EnsemblerInfraConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnsemblerInfraConfigWithDefaults() *EnsemblerInfraConfig {
	this := EnsemblerInfraConfig{}
	return &this
}

// GetArtifactUri returns the ArtifactUri field value if set, zero value otherwise.
func (o *EnsemblerInfraConfig) GetArtifactUri() string {
	if o == nil || o.ArtifactUri == nil {
		var ret string
		return ret
	}
	return *o.ArtifactUri
}

// GetArtifactUriOk returns a tuple with the ArtifactUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerInfraConfig) GetArtifactUriOk() (*string, bool) {
	if o == nil || o.ArtifactUri == nil {
		return nil, false
	}
	return o.ArtifactUri, true
}

// HasArtifactUri returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasArtifactUri() bool {
	if o != nil && o.ArtifactUri != nil {
		return true
	}

	return false
}

// SetArtifactUri gets a reference to the given string and assigns it to the ArtifactUri field.
func (o *EnsemblerInfraConfig) SetArtifactUri(v string) {
	o.ArtifactUri = &v
}

// GetEnsemblerName returns the EnsemblerName field value if set, zero value otherwise.
func (o *EnsemblerInfraConfig) GetEnsemblerName() string {
	if o == nil || o.EnsemblerName == nil {
		var ret string
		return ret
	}
	return *o.EnsemblerName
}

// GetEnsemblerNameOk returns a tuple with the EnsemblerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerInfraConfig) GetEnsemblerNameOk() (*string, bool) {
	if o == nil || o.EnsemblerName == nil {
		return nil, false
	}
	return o.EnsemblerName, true
}

// HasEnsemblerName returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasEnsemblerName() bool {
	if o != nil && o.EnsemblerName != nil {
		return true
	}

	return false
}

// SetEnsemblerName gets a reference to the given string and assigns it to the EnsemblerName field.
func (o *EnsemblerInfraConfig) SetEnsemblerName(v string) {
	o.EnsemblerName = &v
}

// GetServiceAccountName returns the ServiceAccountName field value if set, zero value otherwise.
func (o *EnsemblerInfraConfig) GetServiceAccountName() string {
	if o == nil || o.ServiceAccountName == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerInfraConfig) GetServiceAccountNameOk() (*string, bool) {
	if o == nil || o.ServiceAccountName == nil {
		return nil, false
	}
	return o.ServiceAccountName, true
}

// HasServiceAccountName returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasServiceAccountName() bool {
	if o != nil && o.ServiceAccountName != nil {
		return true
	}

	return false
}

// SetServiceAccountName gets a reference to the given string and assigns it to the ServiceAccountName field.
func (o *EnsemblerInfraConfig) SetServiceAccountName(v string) {
	o.ServiceAccountName = &v
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnsemblerInfraConfig) GetResources() EnsemblingResources {
	if o == nil || o.Resources.Get() == nil {
		var ret EnsemblingResources
		return ret
	}
	return *o.Resources.Get()
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnsemblerInfraConfig) GetResourcesOk() (*EnsemblingResources, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources.Get(), o.Resources.IsSet()
}

// HasResources returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasResources() bool {
	if o != nil && o.Resources.IsSet() {
		return true
	}

	return false
}

// SetResources gets a reference to the given NullableEnsemblingResources and assigns it to the Resources field.
func (o *EnsemblerInfraConfig) SetResources(v EnsemblingResources) {
	o.Resources.Set(&v)
}

// SetResourcesNil sets the value for Resources to be an explicit nil
func (o *EnsemblerInfraConfig) SetResourcesNil() {
	o.Resources.Set(nil)
}

// UnsetResources ensures that no value is present for Resources, not even an explicit nil
func (o *EnsemblerInfraConfig) UnsetResources() {
	o.Resources.Unset()
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *EnsemblerInfraConfig) GetRunId() string {
	if o == nil || o.RunId == nil {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerInfraConfig) GetRunIdOk() (*string, bool) {
	if o == nil || o.RunId == nil {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasRunId() bool {
	if o != nil && o.RunId != nil {
		return true
	}

	return false
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *EnsemblerInfraConfig) SetRunId(v string) {
	o.RunId = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *EnsemblerInfraConfig) GetEnv() []EnvVar {
	if o == nil || o.Env == nil {
		var ret []EnvVar
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerInfraConfig) GetEnvOk() (*[]EnvVar, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *EnsemblerInfraConfig) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []EnvVar and assigns it to the Env field.
func (o *EnsemblerInfraConfig) SetEnv(v []EnvVar) {
	o.Env = &v
}

func (o EnsemblerInfraConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactUri != nil {
		toSerialize["artifact_uri"] = o.ArtifactUri
	}
	if o.EnsemblerName != nil {
		toSerialize["ensembler_name"] = o.EnsemblerName
	}
	if o.ServiceAccountName != nil {
		toSerialize["service_account_name"] = o.ServiceAccountName
	}
	if o.Resources.IsSet() {
		toSerialize["resources"] = o.Resources.Get()
	}
	if o.RunId != nil {
		toSerialize["run_id"] = o.RunId
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	return json.Marshal(toSerialize)
}

type NullableEnsemblerInfraConfig struct {
	value *EnsemblerInfraConfig
	isSet bool
}

func (v NullableEnsemblerInfraConfig) Get() *EnsemblerInfraConfig {
	return v.value
}

func (v *NullableEnsemblerInfraConfig) Set(val *EnsemblerInfraConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEnsemblerInfraConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEnsemblerInfraConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnsemblerInfraConfig(val *EnsemblerInfraConfig) *NullableEnsemblerInfraConfig {
	return &NullableEnsemblerInfraConfig{value: val, isSet: true}
}

func (v NullableEnsemblerInfraConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnsemblerInfraConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
