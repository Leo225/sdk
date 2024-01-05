package rule

import (
	"bytes"
	"context"
	"fmt"

	"github.com/Leo225/sdk/sqlplus"
)

type DataRule struct {
	Name        string
	As          string // Databases table name(option)
	ValueFuncs  map[string]ValueFunc
	Expressions []*DataRuleExpression
	Args        map[string]interface{}
}

type DataRuleExpression struct {
	Connector string
	Field     string
	Operator  string
	ValueType string
	Value     string
}

type ValueFunc func(ctx context.Context,
	filed, k string, args map[string]interface{}) (valueArgs []interface{}, err error)

const (
	ConstantKey = "constant"
	DynamicKey  = "dynamic"
)

var (
	ConnectorMap = map[string]string{
		"and": "AND",
		"or":  "OR",
	}
	OperatorMap = map[string]string{
		"eq":     "=",
		"ne":     "!=",
		"lt":     "<",
		"le":     "<=",
		"gt":     ">",
		"ge":     ">=",
		"in":     "IN",
		"not-in": "NOT IN",
	}
)

func BuildSQL(ctx context.Context, rules []*DataRule) (ruleSQL string, ruleArgs []interface{}, err error) {
	var sqlBuffer bytes.Buffer
	if len(rules) > 0 {
		sqlBuffer.WriteString("(")
	}

	for i := 0; i < len(rules); i++ {
		l := len(rules[i].Expressions)
		if i > 0 && l > 0 {
			sqlBuffer.WriteString(" OR ")
		}

		as := rules[i].As
		valueFuncs := rules[i].ValueFuncs
		args := rules[i].Args
		if l > 0 {
			sqlBuffer.WriteString("(")
		}

		for j := 0; j < l; j++ {
			if j > 0 {
				connector := rules[i].Expressions[j].Connector
				sqlBuffer.WriteString(fmt.Sprintf(" %s ", ConnectorMap[connector]))
			}

			field := rules[i].Expressions[j].Field
			operator := rules[i].Expressions[j].Operator
			valueType := rules[i].Expressions[j].ValueType
			value := rules[i].Expressions[j].Value
			var valueArgs []interface{}
			if valueType == ConstantKey {
				valueArgs = append(valueArgs, value)
			} else if valueType == DynamicKey {
				f, ok := valueFuncs[value]
				if !ok {
					err = fmt.Errorf("data rule [%s] not in engine [%s]", rules[i].Name, value)
				}
				if valueArgs, err = f(ctx, field, value, args); err != nil {
					return
				}
			}

			if len(as) > 0 {
				field = fmt.Sprintf("%s.%s", as, field)
			}
			sqlBuffer.WriteString(fmt.Sprintf("%s %s ", field, OperatorMap[operator]))
			placeholders := sqlplus.SQLPlaceholders(len(valueArgs))
			if operator == "in" || operator == "not-in" {
				sqlBuffer.WriteString(fmt.Sprintf("(%s)", placeholders))
			} else {
				sqlBuffer.WriteString(placeholders)
			}
			ruleArgs = append(ruleArgs, valueArgs...)
		} // End for j
		if l > 0 {
			sqlBuffer.WriteString(")")
		}
	} // End for i
	if len(rules) > 1 {
		sqlBuffer.WriteString(")")
	}
	ruleSQL = sqlBuffer.String()
	return
}
