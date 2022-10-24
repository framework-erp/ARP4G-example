package idgenerator

import "context"

type ContactIdGenerator interface {
	GenerateId(ctx context.Context) int64
}

type GroupIdGenerator interface {
	GenerateId(ctx context.Context) int64
}
