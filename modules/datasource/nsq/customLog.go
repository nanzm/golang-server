package nsq

import "dora/pkg/utils/logx"

type customLog struct {
}

func (c *customLog) Output(_ int, s string) error {
	logx.Warnf("nsq: %v", s)
	return nil
}

