package entities

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Metadata
}

func NewUser(id string, name string, email string, metadata Metadata) User {
	return User{
		Id:       id,
		Name:     name,
		Email:    email,
		Metadata: metadata,
	}
}

/* This, although works, is bad practice

type User struct {
	ID            string
	name          string
	email         string
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

func NewUser(ID string, name string, email string, createdAt time.Time, createdBy string, lastUpdatedAt time.Time, lastUpdatedBy string) *User {
	return &User{ID: ID, name: name, email: email, createdAt: createdAt, createdBy: createdBy, lastUpdatedAt: lastUpdatedAt, lastUpdatedBy: lastUpdatedBy}
}*/
