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

package constants

// RedisConfig
//
// # Redis 配置
//
// 该结构体为 Redis 配置的结构体，用于存储 Redis 的配置信息；
// 用于存储 Redis 的配置信息；
//
// # 参数
//   - Host:		Redis 主机地址
//   - Port:		Redis 端口
//   - Password:	Redis 密码
//   - DB:		Redis 数据库
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// SystemConfig
//
// # 系统配置
//
// 该结构体为系统配置的结构体，用于存储系统的配置信息；
// 用于存储系统的配置信息；
//
// # 参数
//   - SystemName:		系统名称
//   - SystemVersion:	系统版本
//   - SystemAuthor:	系统作者
//   - SystemDescription:	系统描述
type SystemConfig struct {
	SystemName        string
	SystemVersion     string
	SystemAuthor      string
	SystemDescription string
}
