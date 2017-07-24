package commands

import (
	. "github.com/bearyinnovative/lili/model"
)

type SlackV2EX struct {
	*BaseV2EX
}

func NewSlackV2EX() *SlackV2EX {
	return &SlackV2EX{
		&BaseV2EX{
			notifiers: LiliNotifiers,
			Query:     "Slack",
		},
	}
}
