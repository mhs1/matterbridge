// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MailSearchFolder undocumented
type MailSearchFolder struct {
	// MailFolder is the base model of MailSearchFolder
	MailFolder
	// IsSupported undocumented
	IsSupported *bool `json:"isSupported,omitempty"`
	// IncludeNestedFolders undocumented
	IncludeNestedFolders *bool `json:"includeNestedFolders,omitempty"`
	// SourceFolderIDs undocumented
	SourceFolderIDs []string `json:"sourceFolderIds,omitempty"`
	// FilterQuery undocumented
	FilterQuery *string `json:"filterQuery,omitempty"`
}