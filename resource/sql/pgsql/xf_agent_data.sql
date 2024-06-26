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

-- 探针数据表
create table "%xf_database%"
(
    aid         uuid                    not null
        constraint "%xf_database%_pk"
            primary key,
    agent_uuid  uuid                    not null
        constraint "%xf_database%_%xf_agent%_agent_uuid_fk"
            references "%xf_agent%"
            on update cascade on delete cascade,
    recorded_at timestamp default now() not null,
    data        jsonb                   not null
);

comment on table "%xf_database%" is '探针数据表';
comment on column "%xf_database%".aid is 'AgentDataUUID';
comment on column "%xf_database%".agent_uuid is '探针UUID';
comment on column "%xf_database%".recorded_at is '记录时间';
comment on column "%xf_database%".data is '探针数据';
