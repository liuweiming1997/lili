package commands

import (
	. "github.com/bearyinnovative/lili/model"
)

type DingDingV2EX struct {
	*BaseV2EX
}

func NewDingDingV2EX() *DingDingV2EX {
	return &DingDingV2EX{
		&BaseV2EX{
			notifiers: LiliNotifiers,
			Query:     "钉钉",
		},
	}
}
