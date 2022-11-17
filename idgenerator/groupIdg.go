package idgenerator

import "context"

type GroupIdGenerator interface {
	GenerateId(ctx context.Context) int64
}
