package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// StorageSchema -
type StorageSchema struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	ItemID          string             `bson:"item_id, omitempty"`
	Name            string             `bson:"name, omitempty"`
	Type            string             `bson:"type, omitempty"`
	Sources         string             `bson:"sources, omitempty"`
	Amount          uint32             `bson:"amount, omitempty"`
	Price           uint16             `bson:"price, omitempty"`
	Description     string             `bson:"description, omitempty"`
	CreateTimestamp int64              `bson:"create_timestamp, omitempty"`
}
