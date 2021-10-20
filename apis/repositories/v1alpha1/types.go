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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// RepositoryParameters represents a GitHub repository
// type RepositoryParameters struct {
// }
type RepositoryParameters struct {
	ID                  *int64          `json:"id,omitempty"`
	NodeID              *string         `json:"node_id,omitempty"`
	Owner               *string         `json:"owner,omitempty"`
	Name                *string         `json:"name,omitempty"`
	FullName            *string         `json:"fullName,omitempty"`
	Description         *string         `json:"description,omitempty"`
	Homepage            *string         `json:"homepage,omitempty"`
	DefaultBranch       *string         `json:"defaultBranch,omitempty"`
	MasterBranch        *string         `json:"masterBranch,omitempty"`
	CreatedAt           *string         `json:"createdAt,omitempty"`
	PushedAt            *string         `json:"pushedAt,omitempty"`
	UpdatedAt           *string         `json:"updatedAt,omitempty"`
	HTMLURL             *string         `json:"htmlUrl,omitempty"`
	CloneURL            *string         `json:"cloneUrl,omitempty"`
	GitURL              *string         `json:"gitUrl,omitempty"`
	MirrorURL           *string         `json:"mirrorUrl,omitempty"`
	SSHURL              *string         `json:"sshUrl,omitempty"`
	SVNURL              *string         `json:"svnUrl,omitempty"`
	Language            *string         `json:"language,omitempty"`
	Fork                *bool           `json:"fork,omitempty"`
	ForksCount          *int            `json:"forksCount,omitempty"`
	NetworkCount        *int            `json:"networkCount,omitempty"`
	OpenIssuesCount     *int            `json:"openIssuesCount,omitempty"`
	StargazersCount     *int            `json:"stargazersCount,omitempty"`
	SubscribersCount    *int            `json:"subscribersCount,omitempty"`
	Size                *int            `json:"size,omitempty"`
	AutoInit            *bool           `json:"autoInit,omitempty"`
	Parent              *string         `json:"parent,omitempty"`
	Source              *string         `json:"source,omitempty"`
	TemplateRepository  *string         `json:"templateRepository,omitempty"`
	Organization        *string         `json:"organization,omitempty"`
	Permissions         map[string]bool `json:"permissions,omitempty"`
	AllowRebaseMerge    *bool           `json:"allowRebaseMerge,omitempty"`
	AllowSquashMerge    *bool           `json:"allowSquashMerge,omitempty"`
	AllowMergeCommit    *bool           `json:"allowMergeCommit,omitempty"`
	AllowAutoMerge      *bool           `json:"allowAutoMerge,omitempty"`
	DeleteBranchOnMerge *bool           `json:"deleteBranchOnMerge,omitempty"`
	Topics              []string        `json:"topics,omitempty"`
	Archived            *bool           `json:"archived,omitempty"`
	Disabled            *bool           `json:"disabled,omitempty"`

	// Only provided when using RepositoriesService.Get while in preview
	License *string `json:"license,omitempty"`

	// Additional mutable fields when creating and editing a repository
	Private           *bool   `json:"private,omitempty"`
	HasIssues         *bool   `json:"hasIssues,omitempty"`
	HasWiki           *bool   `json:"hasWiki,omitempty"`
	HasPages          *bool   `json:"hasPages,omitempty"`
	HasProjects       *bool   `json:"hasProjects,omitempty"`
	HasDownloads      *bool   `json:"hasDownloads,omitempty"`
	IsTemplate        *bool   `json:"isTemplate,omitempty"`
	LicenseTemplate   *string `json:"licenseTemplate,omitempty"`
	GitignoreTemplate *string `json:"gitignoreTemplate,omitempty"`

	// Creating an organization repository. Required for non-owners.
	TeamID *int64 `json:"teamId,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archiveUrl,omitempty"`
	AssigneesURL     *string `json:"assigneesUrl,omitempty"`
	BlobsURL         *string `json:"blobsUrl,omitempty"`
	BranchesURL      *string `json:"branchesUrl,omitempty"`
	CollaboratorsURL *string `json:"collaboratorsUrl,omitempty"`
	CommentsURL      *string `json:"commentsUrl,omitempty"`
	CommitsURL       *string `json:"commitsUrl,omitempty"`
	CompareURL       *string `json:"compareUrl,omitempty"`
	ContentsURL      *string `json:"contentsUrl,omitempty"`
	ContributorsURL  *string `json:"contributorsUrl,omitempty"`
	DeploymentsURL   *string `json:"deploymentsUrl,omitempty"`
	DownloadsURL     *string `json:"downloadsUrl,omitempty"`
	EventsURL        *string `json:"eventsUrl,omitempty"`
	ForksURL         *string `json:"forksUrl,omitempty"`
	GitCommitsURL    *string `json:"gitCommitsUrl,omitempty"`
	GitRefsURL       *string `json:"git_refsUrl,omitempty"`
	GitTagsURL       *string `json:"git_tagsUrl,omitempty"`
	HooksURL         *string `json:"hooksUrl,omitempty"`
	IssueCommentURL  *string `json:"issue_commentUrl,omitempty"`
	IssueEventsURL   *string `json:"issue_eventsUrl,omitempty"`
	IssuesURL        *string `json:"issuesUrl,omitempty"`
	KeysURL          *string `json:"keysUrl,omitempty"`
	LabelsURL        *string `json:"labelsUrl,omitempty"`
	LanguagesURL     *string `json:"languagesUrl,omitempty"`
	MergesURL        *string `json:"mergesUrl,omitempty"`
	MilestonesURL    *string `json:"milestonesUrl,omitempty"`
	NotificationsURL *string `json:"notificationsUrl,omitempty"`
	PullsURL         *string `json:"pullsUrl,omitempty"`
	ReleasesURL      *string `json:"releasesUrl,omitempty"`
	StargazersURL    *string `json:"stargazersUrl,omitempty"`
	StatusesURL      *string `json:"statusesUrl,omitempty"`
	SubscribersURL   *string `json:"subscribersUrl,omitempty"`
	SubscriptionURL  *string `json:"subscriptionUrl,omitempty"`
	TagsURL          *string `json:"tagsUrl,omitempty"`
	TreesURL         *string `json:"treesUrl,omitempty"`
	TeamsURL         *string `json:"teamsUrl,omitempty"`

	// Visibility is only used for Create and Edit endpoints. The visibility field
	// overrides the field parameter when both are used.
	// Can be one of public, private or internal.
	Visibility *string `json:"visibility,omitempty"`
}

// MembershipSpec defines the desired state of a Membership.
type RepositorySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RepositoryParameters `json:"forProvider"`
}

// RepositoryObservation is the representation of the current state that is observed
type RepositoryObservation struct {
	URL *string `json:"url,omitempty"`

	// State is the repository's state.
	// Possible values are: "Active", "Archived", "Disabled"
	State *string `json:"state,omitempty"`
}

// RepositoryStatus represents the observed state of a Repository.
type RepositoryStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RepositoryObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// A Repository is a managed resource that represents a GitHub repository
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RepositorySpec   `json:"spec"`
	Status RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repository
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}
