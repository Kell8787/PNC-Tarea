package data

import (
	"layersapi/entities"
	"time"
)

var meta = entities.Metadata{
	CreatedAt: time.Now().String(),
	UpdatedAt: time.Now().String(),
	CreatedBy: "webapp",
	UpdatedBy: "webapp",
	DeleteAt:  time.Now().String(),
}

var Data = []entities.User{
	entities.NewUser("25dedd65-ca82-4271-b147-2ad36c04adcc", "Miguel", "miguel@gmail.com", meta),
	entities.NewUser("bb62e33c-1a8a-4530-bd28-b6c0824e85dc", "Leonardo", "leonardo@gmail.com", meta),
	entities.NewUser("848905a8-956a-4c36-9715-cefd0f7648b8", "Rafael", "rafael@gmail.com", meta),
}
