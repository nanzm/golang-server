package elasticComponent

import (
	"dora/internal/config"
	"dora/internal/logstore/core"
)

type elasticQuery struct {
	config config.Elastic
}

func NewElasticQuery() core.Query {
	return &elasticQuery{
		config: config.GetElastic(),
	}
}
