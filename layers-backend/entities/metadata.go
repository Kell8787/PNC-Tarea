package entities

type Metadata struct {
	CreatedAt string `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	DeleteAt  string `json:"delete,omitempty"`
}
