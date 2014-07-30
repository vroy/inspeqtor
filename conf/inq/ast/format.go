package ast

import (
	"inspeqtor/conf/inq/token"
	"strconv"
	"strings"
)

type HostCheck struct {
	Rules []*Rule
}

type ProcessCheck struct {
	Name  string
	Rules RuleList
}

type RuleList []*Rule

type RuleMetric struct {
	Family string
	Name   string
}

type Rule struct {
	Metric     RuleMetric
	Operator   string
	Value      uint64
	Action     string
	CycleCount uint8
}

func NewProcessCheck(checkType interface{}, name interface{}, rules interface{}) *ProcessCheck {
	return &ProcessCheck{
		string(name.(*token.Token).Lit),
		rules.(RuleList),
	}
}
func NewHostCheck(rules interface{}) *HostCheck {
	return &HostCheck{
		rules.(RuleList),
	}
}

func NewRuleList(rule interface{}) RuleList {
	return RuleList{rule.(*Rule)}
}

func AppendRule(list interface{}, rule interface{}) RuleList {
	return append(list.(RuleList), rule.(*Rule))
}

func NewRule(metric interface{}, operator interface{}, value interface{}, action interface{}, cycleCount interface{}) *Rule {
	return &Rule{
		*metric.(*RuleMetric),
		string(operator.(*token.Token).Lit),
		value.(uint64),
		string(action.(*token.Token).Lit),
		uint8(cycleCount.(uint64)),
	}
}

func Metric(family interface{}, name interface{}) (*RuleMetric, error) {
	m := &RuleMetric{string(family.(*token.Token).Lit), ""}
	if name != nil {
		m.Name = string(name.(*token.Token).Lit)
	}
	return m, nil
}

func HumanAmount(digits interface{}, code interface{}) (uint64, error) {
	amt, err := strconv.ParseUint(string(digits.(*token.Token).Lit), 10, 64)
	if err != nil {
		return 0, err
	}

	if code != nil {
		sizecode := strings.ToLower(string(code.(*token.Token).Lit))
		if sizecode == "k" {
			amt *= 1024
		} else if sizecode == "m" {
			amt *= 1024 * 1024
		} else if sizecode == "g" {
			amt *= 1024 * 1024 * 1024
		} else if sizecode == "t" {
			amt *= 1024 * 1024 * 1024 * 1024
		} else if sizecode == "p" {
			amt *= 1024 * 1024 * 1024 * 1024 * 1024
		} else if sizecode == "%" {
			// nothing to do
		}
	}
	return amt, nil
}

func ToUint64(v interface{}) (uint64, error) {
	return strconv.ParseUint(string(v.(*token.Token).Lit), 10, 64)
}
