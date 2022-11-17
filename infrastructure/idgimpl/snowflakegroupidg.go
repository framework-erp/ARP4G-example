package idgimpl

import (
	"context"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeGroupIdGenerator struct {
	Node *snowflake.Node
}

func (sf *SnowflakeGroupIdGenerator) GenerateId(ctx context.Context) int64 {
	return sf.Node.Generate().Int64()
}
