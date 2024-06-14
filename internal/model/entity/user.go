/*
 * ----------------------------------------------------------------------
 * 版权声明 (Copyright Notice)
 * ----------------------------------------------------------------------
 * 项目名称: 竹监控「BambooDashboard」
 * 描述: 一个由 Go 编写的服务监控系统 (A service monitoring system written in Go)
 * 作者: 筱锋 (xiao_lfeng)
 *
 * 版权所有 © 2016-2024 筱锋(xiao_lfeng). 保留所有权利。
 * ----------------------------------------------------------------------
 * 许可声明:
 *   根据 MIT 许可协议，特此免费授予任何获得本软件和相关文档文件（“软件”）副本的人无限制地处
 *   理软件的权利，包括但不限于使用、复制、修改、合并、出版、分发、再许可和/或出售软件的副
 *   本，并允许软件提供给其使用的人这样做，但须符合以下条件：
 * 上述版权声明和本许可声明应包含在软件的所有副本或主要部分中。
 * ----------------------------------------------------------------------
 * 免责声明:
 *   本软件按“原样”提供，不提供任何明示或暗示的担保，包括但不限于对适销性、特定用途适用性和
 *   非侵权性的担保。在任何情况下，作者或版权持有人均不对因本软件或使用本软件而产生的任何索
 *   赔、损害或其他责任负责，无论是因合同、侵权或其他行为导致的。
 * ----------------------------------------------------------------------
 */

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Uuid            string      `json:"uuid"             orm:"uuid"             description:"用户uuid"`    // 用户uuid
	Username        string      `json:"username"         orm:"username"         description:"用户名"`       // 用户名
	Email           string      `json:"email"            orm:"email"            description:"邮箱"`        // 邮箱
	Phone           string      `json:"phone"            orm:"phone"            description:"手机号"`       // 手机号
	Nickname        string      `json:"nickname"         orm:"nickname"         description:"昵称"`        // 昵称
	Avatar          string      `json:"avatar"           orm:"avatar"           description:"头像地址"`      // 头像地址
	Password        string      `json:"password"         orm:"password"         description:"用户密码"`      // 用户密码
	Totp            string      `json:"totp"             orm:"totp"             description:"基于时间的二步验证"` // 基于时间的二步验证
	Role            string      `json:"role"             orm:"role"             description:"角色组"`       // 角色组
	AgentPermission *gjson.Json `json:"agent_permission" orm:"agent_permission" description:"Agent权限"`   // Agent权限
	CreatedAt       *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`      // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"修改时间"`      // 修改时间
}
