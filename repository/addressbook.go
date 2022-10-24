package repository

import (
	"context"
	"example/aggregate"
)

type ContactRepository interface {
	Find(ctx context.Context, id any) (entity aggregate.Contact, found bool)
	Take(ctx context.Context, id any) (entity aggregate.Contact, found bool)
	Put(ctx context.Context, id any, entity aggregate.Contact)
	Remove(ctx context.Context, id any) (removed aggregate.Contact, exists bool)
}

type GroupRepository interface {
	Find(ctx context.Context, id any) (entity aggregate.Group, found bool)
	Take(ctx context.Context, id any) (entity aggregate.Group, found bool)
	Put(ctx context.Context, id any, entity aggregate.Group)
}
