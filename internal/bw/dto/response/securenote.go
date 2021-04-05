package response

type SecureNote struct {
	CollectionIds  []string `json:"collectionIds"`
	Favorite       bool     `json:"favorite"`
	FolderID       string   `json:"folderId"`
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Notes          string   `json:"notes"`
	Object         string   `json:"object"`
	OrganizationID string   `json:"organizationId"`
	RevisionDate   string   `json:"revisionDate"`
	SecureNote     struct {
		Type int64 `json:"type"`
	} `json:"secureNote"`
	Type int64 `json:"type"`
}
