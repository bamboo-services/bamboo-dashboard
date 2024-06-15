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

package v1

import "github.com/gogf/gf/v2/frame/g"

// InitialConsoleUserReq
//
// # 初始化管理员
//
// 该接口为初始化管理员接口，用于初始化控制器的管理员账号；
// 该接口在初始化后一次后第二次启动将不再生效；
//
// # 请求
//   - Username:	用户名
//   - Password:	密码
//   - Email:		邮箱
//   - Phone:		手机号
type InitialConsoleUserReq struct {
	g.Meta   `path:"/console" method:"Post" tags:"初始化控制器" sm:"初始化管理员" dc:"该接口为初始化管理员接口，用于初始化控制器的管理员账号；该接口在初始化后一次后第二次启动将不再生效；"`
	Username string `json:"username" sm:"用户名" v:"required|length:4,30|regex:^[0-9A-Za-z_-]+$#请输入用户名|用户名长度为 4-30 位|用户名只能包含数字、字母、下划线和短横线"`
	Password string `json:"password" sm:"密码" v:"required|length:6,30|password#请输入密码|密码长度为 6-30 位|密码格式不正确"`
	Email    string `json:"email" sm:"邮箱" v:"required|email#请输入邮箱|邮箱格式不正确"`
	Phone    string `json:"phone" sm:"手机号" v:"required|phone#请输入手机号|手机号格式不正确"`
}

// InitialConsoleUserRes
//
// # 响应
//
// 初始化管理员信息相应结构体
type InitialConsoleUserRes struct {
	g.Meta `mime:"application/json"`
}
