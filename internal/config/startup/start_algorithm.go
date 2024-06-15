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
	"bamboo-dashboard/internal/dao"
	"bamboo-dashboard/internal/model/do"
	"bamboo-dashboard/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"os"
	"strings"
)

// databaseTablePrepare
//
// # 数据表准备
//
// 该方法为创建数据库表的方法，会判断数据库类型，然后执行对应的 SQL 文件。
//
// # 参数
//   - ctx:		上下文
//   - tableName:	表名
func databaseTablePrepare(ctx context.Context, tableName string) {
	getType, _ := g.Cfg().Get(ctx, "database.default.type")
	getResourcePath := ""
	switch getType.String() {
	case "pgsql":
		getResourcePath = "resource/sql/pgsql/xf_" + tableName + ".sql"
		break
	case "mysql":
		getResourcePath = "resource/sql/mysql/xf_" + tableName + ".sql"
		break
	default:
		// 如果数据库类型不支持，则直接返回
		g.Log().Errorf(ctx, "[STAR] 数据库类型不支持: %s", getType.String())
		os.Exit(502)
	}

	getContent := g.Res().GetContent(getResourcePath)
	if len(getContent) > 0 {
		// 获取表头
		getPrefix := g.DB().GetPrefix()
		result, _ := g.DB().Query(ctx, "SELECT * FROM information_schema.tables WHERE table_name = ?", getPrefix+tableName)
		if result.IsEmpty() {
			// 切分文件内容
			getSqlStatements := strings.Split(replaceStatement(tableName, string(getContent)), ";")
			err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				for _, sqlStatement := range getSqlStatements {
					if len(strings.TrimSpace(sqlStatement)) > 0 {
						_, err := tx.Exec(sqlStatement)
						if err != nil {
							return err
						}
					}
				}
				return nil
			})
			if err != nil {
				g.Log().Errorf(ctx, "[STAR] 数据表创建失败: %s", err.Error())
				os.Exit(502)
			} else {
				g.Log().Debugf(ctx, "[STAR] 数据表创建成功: %s", tableName)
			}
		} else {
			g.Log().Debugf(ctx, "[STAR] 数据表已存在: %s", tableName)
		}
	} else {
		g.Log().Errorf(ctx, "[STAR] 不存在数据表 %s", tableName)
		os.Exit(502)
	}
}

// replaceStatement
//
// # 替换语句
//
// 该方法为替换语句的方法，用于替换 SQL 文件中的变量名字；实现动态替换 SQL 文件中的变量名字；以实现前缀允许不相同的情况下建表；
// 例如：将 %xf_database% 替换为数据库表名等
//
// # 参数
//   - table:		表名
//   - statement:	SQL 语句
//
// # 返回
//   - string:		替换后的 SQL 语句
func replaceStatement(table, statement string) string {
	getPrefix := g.DB().GetPrefix()
	getReplace := strings.ReplaceAll(statement, "%xf_database%", getPrefix+table)
	getReplace = strings.ReplaceAll(getReplace, string(gres.GetContent("resource/license_sql.txt")), "")
	// 专项内容替换
	getReplace = strings.ReplaceAll(getReplace, "%xf_user%", getPrefix+"user")
	getReplace = strings.ReplaceAll(getReplace, "%xf_role%", getPrefix+"role")
	getReplace = strings.ReplaceAll(getReplace, "%xf_info%", getPrefix+"info")
	getReplace = strings.ReplaceAll(getReplace, "%xf_agent%", getPrefix+"agent")
	getReplace = strings.ReplaceAll(getReplace, "%xf_server%", getPrefix+"server")
	getReplace = strings.ReplaceAll(getReplace, "%xf_group%", getPrefix+"group")
	return getReplace
}

// informationDataPrepare
//
// # 信息表初始化
//
// 该方法为初始化信息表的方法，检查信息表是否有此数据，若没有则初始化
//
// # 参数
//   - ctx:		上下文
//   - key:		键
//   - value:	值
func informationDataPrepare(ctx context.Context, key, value string) {
	// 检查数据是否存在
	record, _ := dao.Info.Ctx(ctx).Where(do.Info{Keyword: key}).One()
	if record.IsEmpty() {
		_, err := dao.Info.Ctx(ctx).Data(entity.Info{
			SystemUuid: butil.GenerateRandUUID().String(),
			Keyword:    key,
			Value:      value,
		}).OnConflict("keyword").Save()
		if err != nil {
			g.Log().Errorf(ctx, "[STAR] 信息表初始化失败: %s", err.Error())
			os.Exit(502)
		}
		if len(value) > 50 {
			g.Log().Debugf(ctx, "[STAR] 初始化信息表键 %s 值 %s 成功", key, strings.TrimSpace(
				strings.ReplaceAll(value[0:50], "\n", "")+"..."),
			)
		} else {
			g.Log().Debugf(ctx, "[STAR] 初始化信息表键 %s 值 %s 成功", key, value)
		}
	}
}

// mailReplaceLicense
//
// # 替换许可
//
// 该方法为替换许可的方法，用于替换 SQL 文件中的许可声明；实现动态替换 SQL 文件中的许可声明；
//
// # 参数
//   - mailTemplateName:	邮件模板名
//
// # 返回
//   - string:			替换后的邮件模板
func mailReplaceLicense(mailTemplateName string) string {
	// 获取邮件模板
	mailTemplate := gres.GetContent("resource/template/mail/mail_" + mailTemplateName + ".html")
	if len(mailTemplate) > 0 {
		getReplace := strings.Replace(
			string(mailTemplate),
			string(gres.GetContent("resource/license_html.txt")),
			"",
			-1,
		)
		return getReplace
	} else {
		g.Log().Errorf(context.Background(), "[STAR] 不存在邮件模板 %s", mailTemplateName)
		os.Exit(502)
	}
	return ""
}

// roleDataPrepare
//
// # 角色初始化
//
// 该方法为初始化角色的方法，检查角色是否有此数据，若没有则初始化; 该方法会检查角色是否存在，若不存在则初始化角色；
//
// # 参数
//   - ctx:			上下文
//   - roleName:		角色名
//   - displayName:	显示名
//   - description:	描述
func roleDataPrepare(ctx context.Context, roleName, displayName, description string) {
	// 检查数据是否存在
	record, _ := dao.Role.Ctx(ctx).Where(do.Role{RoleName: roleName}).One()
	if record.IsEmpty() {
		_, err := dao.Role.Ctx(ctx).Data(do.Role{
			RoleUuid:    butil.GenerateRandUUID().String(),
			RoleName:    roleName,
			DisplayName: displayName,
			Description: description,
		}).OnConflict("role_name").Save()
		if err != nil {
			g.Log().Errorf(ctx, "[STAR] 角色初始化失败: %s", err.Error())
			os.Exit(502)
		}
		g.Log().Debugf(ctx, "[STAR] 初始化角色 %s 成功", roleName)
	}
}

// getValueWithInfoTable
//
// # 信息表获取
//
// 该方法为获取信息表的方法，用于获取信息表中的数据；若数据不存在则直接退出程序；请注意，数据是一定存在的，如果不存在则说明人为操作过数据库；
// 并非原始数据库导致的，原始数据库不会删除 info 表的内容信息；
//
// # 参数
//   - ctx:		上下文
//   - key:		键
//
// # 返回
//   - interface{}:	值
func getValueWithInfoTable(ctx context.Context, key string) interface{} {
	var getValue *entity.Info
	err := dao.Info.Ctx(ctx).Where(do.Info{Keyword: key}).Scan(&getValue)
	if err != nil {
		g.Log().Errorf(ctx, "[STAR] 数据 %s 获取失败: %s", key, err.Error())
		os.Exit(502)
		return nil
	} else {
		return getValue.Value
	}
}
