package common

import "github.com/casbin/casbin"

const (
	PrefixUserID = "u"
	PrefixRoleID = "r"
)

var Enforcer *casbin.Enforcer


