package mongo

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler[T any] struct {
    DB *DB
    Collection string
}

func (handler *Handler[T]) FindOne(ctx echo.Context, filter bson.M) (*T, error) {
    result := new(T)
    err := handler.DB.FindOne(ctx.Request().Context(), handler.Collection, result, filter)
    return result, err
}

func (handler *Handler[T]) FindBy(ctx echo.Context, field string, value any) (*T, error) {
    return handler.FindOne(ctx, bson.M{ field: value })
}

func (handler *Handler[T]) FindById(ctx echo.Context, value any) (*T, error) {
    return handler.FindOne(ctx, bson.M{ "_id": value })
}

func (handler *Handler[T]) FindAll(ctx echo.Context, filter bson.M) (result []T, err error) {
    err = handler.DB.FindAll(ctx.Request().Context(), handler.Collection, &result, filter)
    return result, err
}

func (handler *Handler[T]) UpdateBy(ctx echo.Context, field string, value any, update bson.M) error {
    return handler.DB.UpdateBy(ctx.Request().Context(), handler.Collection, field, value, update)
}

func (handler *Handler[T]) UpdateById(ctx echo.Context, id primitive.ObjectID, update bson.M) error {
    return handler.DB.UpdateBy(ctx.Request().Context(), handler.Collection, "_id", id, update)
}

func (handler *Handler[T]) Insert(ctx echo.Context, document any) (*primitive.ObjectID, error) {
    return handler.DB.Insert(ctx.Request().Context(), handler.Collection, document)
}

func (handler *Handler[T]) InsertMany(ctx echo.Context, documents []any) (ids []primitive.ObjectID, err error) {
    return handler.DB.InsertMany(ctx.Request().Context(), handler.Collection, documents)
}

func (handler *Handler[T]) DeleteById(ctx echo.Context, id primitive.ObjectID) error {
    return handler.DB.DeleteById(ctx.Request().Context(), handler.Collection, id)
}

func (handler *Handler[T]) Aggregate(ctx echo.Context, pipeline bson.A) (result []T, err error) {
    err = handler.DB.Aggregate(ctx.Request().Context(), handler.Collection, &result, pipeline)
    return result, err
}

func (handler *Handler[T]) AggregateOne(ctx echo.Context, pipeline bson.A) (result *T, err error) {
    results, err := handler.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    if len(results) == 0 {
        err = errors.New("Aggregate returned 0 results")
        return nil, err
    }
    return &results[0], err
}
