/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeploymentObservation struct {
	AllowsUpload *bool `json:"allowsUpload,omitempty" tf:"allows_upload,omitempty"`
}

type DeploymentParameters struct {

	// Any comment on the specific deployment.
	// +kubebuilder:validation:Optional
	Annotation *string `json:"annotation,omitempty" tf:"annotation,omitempty"`

	// Policy to determine how the TTL on imgix images is set.
	// +kubebuilder:validation:Optional
	CacheTTLBehavior *string `json:"cacheTtlBehavior,omitempty" tf:"cache_ttl_behavior,omitempty"`

	// TTL (in seconds) for any error image served when unable to fetch a file from origin.
	// +kubebuilder:validation:Optional
	CacheTTLError *float64 `json:"cacheTtlError,omitempty" tf:"cache_ttl_error,omitempty"`

	// TTL (in seconds) for any error image served when unable to fetch a file from origin.
	// +kubebuilder:validation:Optional
	CacheTTLValue *float64 `json:"cacheTtlValue,omitempty" tf:"cache_ttl_value,omitempty"`

	// Whether this Source should serve a Cross-Domain Policy file if requested.
	// +kubebuilder:validation:Optional
	CrossdomainXMLEnabled *bool `json:"crossdomainXmlEnabled,omitempty" tf:"crossdomain_xml_enabled,omitempty"`

	// Non-imgix.net domains you want to use to access your images. Custom domains must be unique across all Sources and must be valid domains.
	// +kubebuilder:validation:Optional
	CustomDomains []*string `json:"customDomains,omitempty" tf:"custom_domains,omitempty"`

	// Parameters that should be set on all requests to this Source.
	// +kubebuilder:validation:Optional
	DefaultParams map[string]*string `json:"defaultParams,omitempty" tf:"default_params,omitempty"`

	// GCS Access Key ID.
	// +kubebuilder:validation:Optional
	GcsAccessKey *string `json:"gcsAccessKey,omitempty" tf:"gcs_access_key,omitempty"`

	// GCS bucket name.
	// +kubebuilder:validation:Optional
	GcsBucket *string `json:"gcsBucket,omitempty" tf:"gcs_bucket,omitempty"`

	// The folder prefix prepended to the image path before resolving the image in GCS.
	// +kubebuilder:validation:Optional
	GcsPrefix *string `json:"gcsPrefix,omitempty" tf:"gcs_prefix,omitempty"`

	// GCS Secret Access Key.
	// +kubebuilder:validation:Optional
	GcsSecretKeySecretRef *v1.SecretKeySelector `json:"gcsSecretKeySecretRef,omitempty" tf:"-"`

	// Image URL imgix should serve instead when a request results in an error.
	// +kubebuilder:validation:Optional
	ImageError *string `json:"imageError,omitempty" tf:"image_error,omitempty"`

	// Whether imgix should pass the parameters on the request that received an error to the URL described in image_error.
	// +kubebuilder:validation:Optional
	ImageErrorAppendQs *bool `json:"imageErrorAppendQs,omitempty" tf:"image_error_append_qs,omitempty"`

	// Image URL imgix should serve instead when a request results in a missing image.
	// +kubebuilder:validation:Optional
	ImageMissing *string `json:"imageMissing,omitempty" tf:"image_missing,omitempty"`

	// Whether imgix should pass the parameters on the request that resulted in a missing image to the URL described in image_missing.
	// +kubebuilder:validation:Optional
	ImageMissingAppendQs *bool `json:"imageMissingAppendQs,omitempty" tf:"image_missing_append_qs,omitempty"`

	// Subdomain you want to use on *.imgix.net to access your images.
	// +kubebuilder:validation:Required
	ImgixSubdomains []*string `json:"imgixSubdomains" tf:"imgix_subdomains,omitempty"`

	// AWS Access Key ID.
	// +kubebuilder:validation:Optional
	S3AccessKey *string `json:"s3AccessKey,omitempty" tf:"s3_access_key,omitempty"`

	// AWS S3 bucket name.
	// +kubebuilder:validation:Optional
	S3Bucket *string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`

	// The folder prefix prepended to the image path before resolving the image in S3.
	// +kubebuilder:validation:Optional
	S3Prefix *string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`

	// AWS S3 Secret Access Key.
	// +kubebuilder:validation:Optional
	S3SecretKeySecretRef *v1.SecretKeySelector `json:"s3SecretKeySecretRef,omitempty" tf:"-"`

	// Whether requests must be signed with the secure_url_token to be considered valid.
	// +kubebuilder:validation:Optional
	SecureURLEnabled *bool `json:"secureUrlEnabled,omitempty" tf:"secure_url_enabled,omitempty"`

	// Type of the deployment.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SourceObservation struct {
	DateDeployed *float64 `json:"dateDeployed,omitempty" tf:"date_deployed,omitempty"`

	DeploymentStatus *string `json:"deploymentStatus,omitempty" tf:"deployment_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SecureURLToken *string `json:"secureUrlToken,omitempty" tf:"secure_url_token,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	Deployment []DeploymentParameters `json:"deployment" tf:"deployment,omitempty"`

	// Whether or not a Source is enabled and capable of serving traffic.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Determines if Terraform should wait for deployed status after any change.
	// +kubebuilder:validation:Optional
	WaitForDeployed *bool `json:"waitForDeployed,omitempty" tf:"wait_for_deployed,omitempty"`
}

// SourceSpec defines the desired state of Source
type SourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SourceParameters `json:"forProvider"`
}

// SourceStatus defines the observed state of Source.
type SourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Source is the Schema for the Sources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,imgixjet}
type Source struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourceSpec   `json:"spec"`
	Status            SourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourceList contains a list of Sources
type SourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Source `json:"items"`
}

// Repository type metadata.
var (
	Source_Kind             = "Source"
	Source_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Source_Kind}.String()
	Source_KindAPIVersion   = Source_Kind + "." + CRDGroupVersion.String()
	Source_GroupVersionKind = CRDGroupVersion.WithKind(Source_Kind)
)

func init() {
	SchemeBuilder.Register(&Source{}, &SourceList{})
}
