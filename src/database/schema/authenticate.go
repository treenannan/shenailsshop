package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthenticateSchema - define authority rule for edit data
type AuthenticateSchema struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `bson:"username, omitempty"`
	Password string             `bson:"password, omitempty"`
	ShotID   int16              `bson:"shotId, omitempty"`
}
