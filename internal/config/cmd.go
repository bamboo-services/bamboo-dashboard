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

package config

import (
	"bamboo-dashboard/internal/config/startup"
	"bamboo-dashboard/internal/controller/initial"
	"context"
	"github.com/bamboo-services/bamboo-utils/bmiddle"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "竹监控「BambooDashboard」 - 一个由 Go 编写的服务监控系统(A service monitoring system written in Go)",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 系统初始化
			startup.Start(ctx)
			// 启动服务
			s := g.Server("bamboo")
			// 关闭路由映射输出
			s.SetDumpRouterMap(false)

			// 后端路由分组
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				// 初始化
				s.Group("/initial", func(group *ghttp.RouterGroup) {
					group.Middleware(bmiddle.BambooMiddleHandler)
					group.Bind(
						initial.NewV1(),
					)
				})
				s.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(bmiddle.BambooMiddleHandler)
					group.Bind()
				})
			})

			// 前端路由分组
			s.BindHandler("/**/**/**", func(r *ghttp.Request) {
				r.Response.ServeFile("bamboo-dashboard-web/dist/index.html")
			})

			// 加载静态资源
			s.AddStaticPath("/assets", "bamboo-dashboard-web/dist/assets")

			// 启动服务
			s.Run()

			// 服务结束
			g.Log().Warningf(ctx, "[SYST] 服务在 %s 结束运行", gtime.Now().String())
			return nil
		},
	}
)
