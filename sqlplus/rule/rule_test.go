package rule

import (
	"context"
	"testing"
)

func mySite(ctx context.Context,
	field, v string, args map[string]interface{}) (valueArgs []interface{}, err error) {
	valueArgs = append(valueArgs, 1, 2, 3)
	return
}

func currentUserID(ctx context.Context,
	field, v string, args map[string]interface{}) (valueArgs []interface{}, err error) {
	valueArgs = append(valueArgs, 1)
	return
}

func TestBuildSQL(t *testing.T) {
	ctx := context.Background()
	valueFuncs := map[string]ValueFunc{
		"my_site":         mySite,
		"current_user_id": currentUserID,
	}
	rules := []*DataRule{
		{
			Name:       "test1",
			As:         "star_ai",
			ValueFuncs: valueFuncs,
			Expressions: []*DataRuleExpression{
				{
					Connector: "",
					Field:     "site_id",
					Operator:  "in",
					ValueType: "dynamic",
					Value:     "my_site",
				},
				{
					Connector: "or",
					Field:     "creator",
					Operator:  "eq",
					ValueType: "dynamic",
					Value:     "current_user_id",
				},
			},
		},
	}
	rsql, rargs, err := BuildSQL(ctx, rules)
	if err != nil {
		t.Error(err)
	}
	t.Logf("rule sql: %s, rule args: %v\n", rsql, rargs)

	rules = append(rules, &DataRule{
		Name:       "test2",
		As:         "t2",
		ValueFuncs: valueFuncs,
		Expressions: []*DataRuleExpression{
			{
				Connector: "",
				Field:     "site_id",
				Operator:  "in",
				ValueType: "dynamic",
				Value:     "my_site",
			},
			{
				Connector: "or",
				Field:     "creator",
				Operator:  "eq",
				ValueType: "dynamic",
				Value:     "current_user_id",
			},
		},
	})
	rsql, rargs, err = BuildSQL(ctx, rules)
	if err != nil {
		t.Error(err)
	}
	t.Logf("rule sql: %s, rule args: %v\n", rsql, rargs)
}
