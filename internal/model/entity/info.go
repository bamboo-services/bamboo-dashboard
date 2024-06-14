// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Info is the golang structure for table info.
type Info struct {
	SystemUuid string `json:"system_uuid" orm:"system_uuid" ` // 系统主键
	Keyword    string `json:"keyword"     orm:"keyword"     ` // 系统键
	Value      string `json:"value"       orm:"value"       ` // 系统值
}
