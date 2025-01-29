package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Profile struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
    Bio       string             `json:"bio,omitempty"`
    Avatar    string             `json:"avatar,omitempty"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}