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
	"bamboo-dashboard/internal/constants"
	"bamboo-dashboard/internal/dao"
	"bamboo-dashboard/internal/model/do"
	"bamboo-dashboard/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-dashboard/api/initial/v1"
)

// InitialConsoleUser
//
// # 初始化控制台用户
//
// 该方法为初始化控制台用户的方法，用于初始化控制台用户; 只有系统处于初始化状态时才能调用此方法;
// 方法将在执行成功后拒绝所有对该方法的调用; 并且在第二次重启服务器时注销此接口的访问权;
//
// # 参数
//   - ctx:		上下文(context.Context)
//   - req:		请求(*v1.InitConsoleUserReq)
//
// # 返回
//   - res:		响应(*v1.InitConsoleUserRes)
//   - err:		错误(error)
func (c *ControllerV1) InitialConsoleUser(
	ctx context.Context,
	req *v1.InitialConsoleUserReq,
) (res *v1.InitialConsoleUserRes, err error) {
	g.Log().Noticef(ctx, "[CONT] InitConsoleUser | 初始化控制台用户")
	if constants.HasInitialMode {
		return nil, berror.NewError(bcode.OperationFailed, "系统已初始化，无法再次初始化管理员")
	} else {
		// 获取数据，进行数据信息初始化
		consoleUUID, err := service.Initial().ConsoleInitial(ctx, req)
		if err != nil {
			return nil, err
		}
		// 对配置文件添加超级管理员的 UUID 信息
		_, err = dao.Info.Ctx(ctx).
			Where(do.Info{Keyword: "system_admin_user"}).
			Data(do.Info{Value: *consoleUUID}).
			Update()
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		constants.GetConsoleAdminUUID = *consoleUUID
		// 初始化完毕后，将系统初始化状态置为 true 状态
		_, err = dao.Info.Ctx(ctx).
			Where(do.Info{Keyword: "system_initial"}).
			Data(do.Info{Value: "true"}).
			Update()
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		return nil, nil
	}
}
