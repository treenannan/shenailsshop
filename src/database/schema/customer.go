package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// CustomerSchema - record customer id
type CustomerSchema struct {
	CustomerID  primitive.ObjectID `bson:"_id, omitempty"`
	FacbookID   string             `bson:"facebook_id, omitempty"`
	FacbookName string             `bson:"facebook_name, omitempty"`
}

// CustomizeProfileSchema - record customer profile
type CustomizeProfileSchema struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	CustomerID primitive.ObjectID `bson:"customer_id, omitempty"` // CustomerSchema.ID
	Name       string             `bson:"name, omitempty"`
	LastName   string             `bson:"lastname, omitempty"`
	Address    string             `bson:"address, omitempty"`
	City       string             `bson:"city, omitempty"`
	Zipcode    string             `bson:"zipcode, omitempty"`
	Phone      string             `bson:"phone, omitempty"`
}
