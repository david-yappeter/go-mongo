package model

//User User Struct
type User struct {
	ID        string  `bson:"_id,omitempty"`
	Name      string  `bson:"name"`
	CreatedAt *string `bson:"created_at"`
	UpdatedAt *string `bson:"updated_at"`
}
