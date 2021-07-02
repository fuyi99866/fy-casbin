package sys

import (
	"fmt"
	"go_casbin/models/base"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", base.GetTablePrefix(),"sys_", name)
}