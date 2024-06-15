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

package startup

import (
	"bamboo-dashboard/internal/constants"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// startDatabase
//
// # 数据表准备
//
// 该方法为创建数据库表的方法，在 GoFrame 的 cmd.go 文件使用；
// 会检查数据库是否存在此方法，若不存在
//
// # 参数
//   - ctx:		上下文
func startDatabase(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 数据表初始化")

	// 系统表
	databaseTablePrepare(ctx, "info")
	// 角色表
	databaseTablePrepare(ctx, "role")
	// 用户表
	databaseTablePrepare(ctx, "user")
	// 服务器表
	databaseTablePrepare(ctx, "server")
	// 探针分组表
	databaseTablePrepare(ctx, "group")
	// 探针表
	databaseTablePrepare(ctx, "agent")
	// 探针数据表
	databaseTablePrepare(ctx, "agent_data")
}

// startInformation
//
// # 信息表初始化
//
// 该方法为初始化信息表的方法，检查信息表是否有此数据，若没有则初始化
//
// # 参数
//   - ctx:		上下文
func startInformation(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 信息表初始化")

	// 系统相关信息
	informationDataPrepare(ctx, "system_name", "竹监控")
	informationDataPrepare(ctx, "system_version", "1.0.0")
	informationDataPrepare(ctx, "system_author", "筱锋")
	informationDataPrepare(ctx, "system_license", "MIT")
	informationDataPrepare(ctx, "system_license_link", "https://opensource.org/license/MIT")
	informationDataPrepare(ctx, "system_description", "一个由 Go 编写的服务监控系统")
	informationDataPrepare(ctx, "system_initial", "false")
	informationDataPrepare(ctx, "system_admin_user", "")
	// 配置相关信息
	informationDataPrepare(ctx, "config_redis", "false")
	informationDataPrepare(ctx, "config_redis_host", "localhost")
	informationDataPrepare(ctx, "config_redis_port", "6379")
	informationDataPrepare(ctx, "config_redis_password", "")
	informationDataPrepare(ctx, "config_redis_db", "0")
	informationDataPrepare(ctx, "config_smtp", "false")
	informationDataPrepare(ctx, "config_smtp_host", "smtp.example.com")
	informationDataPrepare(ctx, "config_smtp_port", "25")
	informationDataPrepare(ctx, "config_smtp_user", "bamboo@x-lf.com")
	informationDataPrepare(ctx, "config_smtp_password", "password")
	// 网站相关信息
	informationDataPrepare(ctx, "web_title", "竹控")
	informationDataPrepare(ctx, "web_description", "一个由 Go 编写的服务监控系统")
	informationDataPrepare(ctx, "web_keywords", "Go, 竹监控, 服务监控, 系统监控")
	// 邮件模板配置信息
	informationDataPrepare(
		ctx,
		"mail_template_verification_code",
		mailReplaceLicense("verification_code"),
	)
}

// startRole
//
// # 角色初始化
//
// 该方法为初始化角色的方法，检查角色是否有此数据，若没有则初始化；
//
// # 参数
//   - ctx:		上下文
func startRole(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 角色初始化")

	// 系统管理员
	roleDataPrepare(ctx, "admin", "系统管理员", "系统级别的管理员，具有最高权限")
	// 普通用户
	roleDataPrepare(ctx, "user", "普通用户", "普通用户，具有一般权限，无法进行系统级别的操作")
}

// startGlobalVariable
//
// # 全局变量初始化
//
// 该方法为全局变量初始化的方法，用于初始化全局变量；
//
// # 参数
//   - ctx:		上下文
func startGlobalVariable(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 全局变量初始化")

	// 系统初始化时间
	constants.HasInitialMode = gconv.Bool(getValueWithInfoTable(ctx, "system_initial"))
	// 系统管理员 UUID
	constants.GetConsoleAdminUUID = gconv.String(getValueWithInfoTable(ctx, "system_admin_user"))
	// 是否启用 Redis
	constants.HasEnableRedis = gconv.Bool(getValueWithInfoTable(ctx, "config_redis"))
	// Redis 配置
	constants.Redis = &constants.RedisConfig{
		Host:     gconv.String(getValueWithInfoTable(ctx, "config_redis_host")),
		Port:     gconv.Int(getValueWithInfoTable(ctx, "config_redis_port")),
		Password: gconv.String(getValueWithInfoTable(ctx, "config_redis_password")),
		DB:       gconv.Int(getValueWithInfoTable(ctx, "config_redis_db")),
	}
	// 系统名称
	constants.System = &constants.SystemConfig{
		SystemName:        gconv.String(getValueWithInfoTable(ctx, "system_name")),
		SystemVersion:     gconv.String(getValueWithInfoTable(ctx, "system_version")),
		SystemAuthor:      gconv.String(getValueWithInfoTable(ctx, "system_author")),
		SystemDescription: gconv.String(getValueWithInfoTable(ctx, "system_description")),
	}
}
