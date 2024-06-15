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

package initial

import (
	v1 "bamboo-dashboard/api/initial/v1"
	"bamboo-dashboard/internal/dao"
	"bamboo-dashboard/internal/model/do"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
)

// ConsoleInitial
//
// # 控制台初始化
//
// 该方法为控制台初始化方法，用于初始化控制台用户；唯一超级管理员的注册；
//
// # 参数
//   - ctx:		上下文
//   - req:		请求参数
//
// # 返回
//   - err:		错误信息
func (s *sInitial) ConsoleInitial(ctx context.Context, req *v1.InitialConsoleUserReq) (consoleUUID *string, err error) {
	g.Log().Noticef(ctx, "[SERV] sInitial.ConsoleInitial | 控制台初始化")
	// 创建用户
	getConsoleUUID := butil.GenerateRandUUID().String()
	// TOTP 密钥
	totpSecret := butil.GenerateRandUUID()
	// 获取管理员角色
	record, err := dao.Role.Ctx(ctx).Fields(dao.Role.Columns().RoleUuid).Where(do.Role{RoleName: "admin"}).One()
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 创建用户
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Uuid:     getConsoleUUID,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: butil.PasswordEncode(req.Password),
		Totp:     totpSecret,
		Role:     record.GMap().Get(dao.Role.Columns().RoleUuid),
	}).Insert()
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return &getConsoleUUID, nil
}
