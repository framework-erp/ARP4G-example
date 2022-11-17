package repository

import (
	"context"
	"example/aggregate"
)

type GroupRepository interface {
	Find(ctx context.Context, id any) (entity *aggregate.Group, found bool)
	Take(ctx context.Context, id any) (entity *aggregate.Group, found bool)
	Put(ctx context.Context, id any, entity *aggregate.Group)
	Remove(ctx context.Context, id any) (removed *aggregate.Group, exists bool)
	GetAll(ctx context.Context) ([]*aggregate.Group, error)
	GetAllDeletedNotEmpty(ctx context.Context) ([]*aggregate.Group, error)
}
