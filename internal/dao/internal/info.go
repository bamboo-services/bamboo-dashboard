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

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// InfoDao is the data access object for table info.
type InfoDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns InfoColumns // columns contains all the column names of Table for convenient usage.
}

// InfoColumns defines and stores column names for table info.
type InfoColumns struct {
	SystemUuid string // 系统主键
	Keyword    string // 系统键
	Value      string // 系统值
}

// infoColumns holds the columns for table info.
var infoColumns = InfoColumns{
	SystemUuid: "system_uuid",
	Keyword:    "keyword",
	Value:      "value",
}

// NewInfoDao creates and returns a new DAO object for table data access.
func NewInfoDao() *InfoDao {
	return &InfoDao{
		group:   "default",
		table:   "info",
		columns: infoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *InfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *InfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *InfoDao) Columns() InfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *InfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *InfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *InfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
