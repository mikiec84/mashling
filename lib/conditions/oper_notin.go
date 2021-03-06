/*
* Copyright © 2017. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package condition

import (
	"strings"
)

/**
specify the exact name of the operator such that the operator can be
used directly in an expression in mashling event-links. the operator
must be preceded by a space (' ') and succeeded by a space (' ') when
used in an expression.

an operator can have more than one names. make sure that the names do not
collide with other operators.

this is a string (alphanumeric) operator that evaluates if a given string belongs to a collection of values.

e.g. ${trigger.content.country notin (USA,IND,CHN,JPN)}
*/
var infoNotIn = &OperatorInfo{
	Names:       []string{"notin"},
	Description: `Support for 'not in' operation to be used in the conditions`,
}

func init() {
	OperatorRegistry.RegisterOperator(&NotIn{info: infoNotIn})
}

type NotIn struct {
	info *OperatorInfo
}

func (o *NotIn) OperatorInfo() *OperatorInfo {
	return o.info
}

// Eval implementation of condition.Operator.Eval
func (o *NotIn) Eval(lhs string, rhs string) bool {
	//RHS will be starting with '(' and ending with ')' and the values will be separated by a comma ','
	rhs = strings.TrimPrefix(rhs, "(")
	rhs = strings.TrimSuffix(rhs, ")")
	values := strings.Split(rhs, ",")
	for _, value := range values {
		if strings.TrimSpace(value) == lhs {
			return false
		}
	}
	return true
}
