package idgimpl

import (
	"context"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeContactIdGenerator struct {
	Node *snowflake.Node
}

func (sf *SnowflakeContactIdGenerator) GenerateId(ctx context.Context) int64 {
	return sf.Node.Generate().Int64()
}
