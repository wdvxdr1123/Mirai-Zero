package zero

import "github.com/wdvxdr1123/mirai-zero/types"

type Rule func(event types.IEvent) bool

type Matcher struct {
	Priority int64
	Block    bool
	Sessions []Session
	Rules    []Rule
}

// 所有匹配器列表  todo: 优先级,阻塞
var MatcherList []*Matcher

func On(priority int64, block bool, rules ...Rule) *Matcher {
	var matcher = &Matcher{
		Priority: priority,
		Block:    block,
		Rules:    rules,
	}
	if MatcherList != nil {
		MatcherList = []*Matcher{}
	}
	MatcherList = append(MatcherList, matcher)
	return matcher
}

func (m *Matcher)Run(event types.IEvent)  {
	panic("impl me")
}

func Emit(event types.IEvent)  {
	for _, matcher := range MatcherList {
		matcher.Run(event)
	}
}
