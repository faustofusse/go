package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) FindOne(ctx context.Context, collection string, target any, filter bson.M, opts ...*options.FindOneOptions) error {
    filter["deleted"] = bson.M{ "$exists": false }
    result := db.database.Collection(collection).FindOne(ctx, filter, opts...)
    if result.Err() != nil {
        return result.Err()
    }
    if target != nil {
        return result.Decode(target)
    }
    return nil
}

func (db *DB) FindAll(ctx context.Context, collection string, target any, filter bson.M) error {
    filter["deleted"] = bson.M{ "$exists": false }
    cursor, err := db.database.Collection(collection).Find(ctx, filter)
    if err != nil {
        return err
    }
    if target != nil {
        return cursor.All(ctx, target)
    }
    return nil
}

func (db *DB) UpdateBy(ctx context.Context, collection string, field string, value any, update bson.M) error {
    filter := bson.M{ field: value }
    // TODO: ineficiente que haga 2 updates
    res, err := db.database.Collection(collection).UpdateOne(ctx, filter, update)
    if err == nil {
        update = bson.M{ "$set": bson.M{ "updated": time.Now().Unix() }}
        _, err = db.database.Collection(collection).UpdateOne(ctx, filter, update)
    }
    if err == nil {
        if res.MatchedCount == 0 {
            return errors.New("not found")
        }
    }
    return err
}

func (db *DB) Insert(ctx context.Context, collection string, document any) (*primitive.ObjectID, error) {
    result, err := db.database.Collection(collection).InsertOne(ctx, document)
    if err != nil { return nil, err }
    id, ok := result.InsertedID.(primitive.ObjectID)
    if !ok {
        return nil, errors.New("Could not get id from object (weird shit)")
    }
    return &id, nil
}

func (db *DB) InsertMany(ctx context.Context, collection string, documents []any) (ids []primitive.ObjectID, err error) {
    // TODO: tiene que haber una mejor forma de solucionar el error (comentar las proximas 4 lineas)
    parsed := []interface{}{}
    for _, document := range documents {
        parsed = append(parsed, document)
    }
    result, err := db.database.Collection(collection).InsertMany(ctx, parsed)
    if err != nil { return nil, err }
    for _, insertedId := range result.InsertedIDs {
        if id, ok := insertedId.(primitive.ObjectID); !ok {
            return nil, errors.New("Could not get id from object (weird shit)")
        } else {
            ids = append(ids, id)
        }
    }
    return ids, err
}

func (db *DB) DeleteById(ctx context.Context, collection string, id primitive.ObjectID) error {
    result, err := db.database.Collection(collection).
        UpdateByID(ctx, id, bson.M{ "$set": bson.M{ "deleted": time.Now().Unix() } })
    if result.MatchedCount == 0 {
        return errors.New("Not found")
    }
    // } else if result.MatchedCount == 1 && result.ModifiedCount == 0 {
    //     return errors.New("Already deleted")
    // }
    return err
}

// TODO: add deleted filter
func (db *DB) Aggregate(ctx context.Context, collection string, target any, pipeline bson.A) error {
    cursor, err := db.database.Collection(collection).Aggregate(ctx, pipeline)
    if err != nil {
        return err
    }
    if target != nil {
        return cursor.All(ctx, target)
    }
    return nil
}

func (db *DB) Setup() {
    indexes := []mongo.IndexModel{
        {
            Keys: bson.M{ "professional.location": "2dsphere" },
            Options: options.Index().SetUnique(false),
        },
    }
    name, err := db.database.Collection("users").Indexes().CreateMany(context.Background(), indexes)
    if err != nil { panic(err) }
    fmt.Printf("Added mongodb index with name: %v\n", name)
}
