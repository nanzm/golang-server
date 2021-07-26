package elasticComponent

import (
	"dora/config"
	"dora/modules/logstore/core"
)

type elasticQuery struct {
	config config.Elastic
}

func NewElasticQuery() core.Query {
	return &elasticQuery{
		config: config.GetElastic(),
	}
}
