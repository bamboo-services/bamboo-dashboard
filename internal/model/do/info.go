// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Info is the golang structure of table xf_info for DAO operations like Where/Data.
type Info struct {
	g.Meta     `orm:"table:xf_info, do:true"`
	SystemUuid interface{} // 系统主键
	Keyword    interface{} // 系统键
	Value      interface{} // 系统值
}
