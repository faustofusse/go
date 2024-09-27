package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
    ID *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" form:"_id,omitempty"`
    Created *int64 `bson:"created,omitempty" json:"created,omitempty" form:"created,omitempty"`
    Updated *int64 `bson:"updated,omitempty" json:"updated,omitempty" form:"updated,omitempty"`
    Deleted *int64 `bson:"deleted,omitempty" json:"deleted,omitempty" form:"deleted,omitempty"`
}

func (document Document) GetID() string {
    return document.ID.Hex()
}
