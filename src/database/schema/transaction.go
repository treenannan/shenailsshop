package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

// TransactionSchema -
type TransactionSchema struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	CustomerID      primitive.ObjectID `bson:"customer_id, omitempty"`
	ItemID          string             `bson:"item_id, omitempty"`
	Cost            uint16             `bson:"cost, omitempty"`
	Status          uint8              `bson:"status, omitempty"`
	PaymentType     uint8              `bson:"payment_type, omitempty"`
	TimestampCreate int64              `bson:"timestamp_create, omitempty"`
	TimestampUpdate int64              `bson:"imestamp_update, omitempty"`
}
