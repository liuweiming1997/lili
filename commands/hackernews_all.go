package commands

import (
	. "github.com/bearyinnovative/lili/model"
	. "github.com/bearyinnovative/lili/notifier"
)

type HackerNewsAll struct {
	*BaseHackerNews
}

func NewHackerNewsAll() *HackerNewsAll {
	return &HackerNewsAll{
		&BaseHackerNews{
			notifiers: []NotifierType{BCChannelNotifier("rocry_news")},
			name:      "rocry",
			shouldNotify: func(item *HNItem) bool {
				if item.Score < 50 || item.Comments < 5 {
					return false
				}

				return true
			},
		},
	}
}
