## 嘉为蓝鲸SAP Hana数据库插件使用说明

## 使用说明

### 插件功能
从系统表/视图中收集所有 HANA 信息，SYS 模式下所有以 M_ 为前缀的表/视图都是监控表/视图，只从 SYS.M_* 表/视图中获取数据。  

### 版本支持

操作系统支持: linux, windows

是否支持arm: 支持

**组件支持版本：**

HanaData数据库版本: 2.x

**是否支持远程采集:**

是

### 参数说明


| **参数名**              | **含义**                | **是否必填** | **使用举例**       |
|----------------------|-----------------------|----------|----------------|
| HOST                 | 数据库服务IP(环境变量)         | 是        |                |
| PORT                 | 数据库服务端口(环境变量)         | 是        |                |
| USER                 | 数据库账户名(环境变量)          | 是        |                |
| PASS                 | 数据库密码(环境变量)           | 是        |                |
| TIMEOUT              | 连接超时时间(s), 默认5s(环境变量) | 否        | 5              |
| --web.listen-address | exporter监听id及端口地址     | 否        | 127.0.0.1:9601 |
| --log.level          | 日志级别                  | 否        | info           |

### 使用指引

1. 配置监控账户
用户需要拥有 SAP_INTERNAL_HANA_SUPPORT 角色才能访问 SYS 模式。如果没有 SAP_INTERNAL_HANA_SUPPORT 角色，这些信息只能由 SYSTEM 用户选择。  
注意：配置的用户至少需要对 SYS 模式具有 SELECT 权限，所有收集器将从该模式下的表/视图中收集信息。  

授予用户 SAP_INTERNAL_HANA_SUPPORT 角色：`GRANT SAP_INTERNAL_HANA_SUPPORT TO user;`

### 指标简介
| **指标ID**                                                          | **指标中文名**   | **维度ID**                                                                                         | **维度含义**                                   |
|-------------------------------------------------------------------|-------------|--------------------------------------------------------------------------------------------------|--------------------------------------------|
| hana_up                                                           | 插件运行状态      | -                                                                                                | -                                          |
| hana_info                                                         | hana信息      | db_name, db_version, sid                                                                         | 数据库名称, 版本, sid                             |
| hana_sys_m_cs_loads_count                                         | 列加载次数       | schema                                                                                           | 模式                                         |
| hana_sys_m_cs_tables_memory_size_in_total                         | 表总内存大小      | host, part_id, port, schema_name, table_name                                                     | 主机名, 分区id, 端口, 模式名称, 表名称                   |
| hana_sys_m_cs_tables_merge_count                                  | 表合并次数       | host, part_id, port, schema_name, table_name                                                     | 主机名, 分区id, 端口, 模式名称, 表名称                   |
| hana_sys_m_cs_tables_read_count                                   | 表读取次数       | host, part_id, port, schema_name, table_name                                                     | 主机名, 分区id, 端口, 模式名称, 表名称                   |
| hana_sys_m_cs_tables_record_count                                 | 表记录数        | host, part_id, port, schema_name, table_name                                                     | 主机名, 分区id, 端口, 模式名称, 表名称                   |
| hana_sys_m_cs_tables_write_count                                  | 表写入次数       | host, part_id, port, schema_name, table_name                                                     | 主机名, 分区id, 端口, 模式名称, 表名称                   |
| hana_sys_m_cs_unloads_count                                       | 列卸载次数       | schema                                                                                           | 模式                                         |
| hana_sys_m_disks_total_size                                       | 磁盘总大小       | host, path, usage_type                                                                           | 主机名, 路径, 类型                                |
| hana_sys_m_disks_used_size                                        | 磁盘已用大小      | host, path, usage_type                                                                           | 主机名, 路径, 类型                                |
| hana_sys_m_host_resource_utilization_free_physical_memory_bytes   | 空闲物理内存      | host                                                                                             | 主机名                                        |
| hana_sys_m_host_resource_utilization_used_physical_memory_bytes   | 已用物理内存      | host                                                                                             | 主机名                                        |
| hana_sys_m_rs_tables_total_allocated_size                         | 表总分配内存大小    | schema, table_name                                                                               | 模式, 表名称                                    |
| hana_sys_m_rs_tables_total_used_size                              | 表已用内存大小     | schema, table_name                                                                               | 模式, 表名称                                    |
| hana_sys_m_service_statistics_active_request_count                | 活跃请求数       | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_active_thread_count                 | 活跃线程数       | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_finished_non_internal_request_count | 完成的非内部请求数   | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_pending_request_count               | 待处理请求数      | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_physical_memory                     | 服务物理内存使用量   | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_process_cpu_time                    | 服务CPU时间     | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_process_physical_memory             | 服务进程物理内存使用量 | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_requests_per_sec                    | 每秒请求数       | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_response_time                       | 请求响应时间      | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_status                              | 服务状态统计      | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_status_duration_seconds             | 服务状态持续时间    | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_thread_count                        | 线程数         | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_total_cpu                           | 总CPU使用率     | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_service_statistics_total_cpu_time                      | 总CPU时间      | host, port, service_name, service_status                                                         | 主机名, 端口, 服务名, 服务状态                         |
| hana_sys_m_shared_memory_allocated_size                           | 分配的共享内存大小   | category, host, port                                                                             | 类目, 主机名, 端口                                |
| hana_sys_m_shared_memory_free_size                                | 空闲的共享内存大小   | category, host, port                                                                             | 类目, 主机名, 端口                                |
| hana_sys_m_shared_memory_used_size                                | 使用的共享内存大小   | category, host, port                                                                             | 类目, 主机名, 端口                                |
| hana_system_config_log_mode                                       | 日志模式        | -                                                                                                | -                                          |
| sys_m_system_replication_status                                   | 系统副本状态      | site_name, site_id, secondary_site_name, secondary_site_id, replication_mode,operation_mode,tier | 站点名称, 站点id, 备用站点名称, 备用站点id, 复制模式, 操作模式, 层级 |
| sys_m_service_replication_secondary_active_status                 | 服务从节点活动状态   | host, port, volume_id, secondary_host, secondary_port                                            | 主机名, 端口, 卷id, 从节点名称, 从节点端口                 |
| sys_m_service_replication_secondary_fully_recoverable             | 服务从节点是否全覆盖  | host, port, volume_id, secondary_host, secondary_port                                            | 主机名, 端口, 卷id, 从节点名称, 从节点端口                 |
| sys_m_service_replication_replication_status                      | 服务副本状态      | host, port, volume_id, secondary_host, secondary_port                                            | 主机名, 端口, 卷id, 从节点名称, 从节点端口                 |
| hana_exporter_collector_duration_seconds                          | 采集消耗时长      | collector                                                                                        | 采集类                                        |
| hana_exporter_last_scrape_error                                   | 最近一次采集状态    | -                                                                                                | -                                          |
| hana_exporter_scrape_errors_total                                 | 采集总错误数量     | collector                                                                                        | 采集类                                        |
| hana_exporter_scrapes_total                                       | 采集总次数       | -                                                                                                | -                                          |



### 版本日志

#### weops_hanadb_exporter 1.10.2

- weops调整

添加“小嘉”微信即可获取SAP hana数据库监控指标最佳实践礼包，其他更多问题欢迎咨询

<img src="https://wedoc.canway.net/imgs/img/小嘉.jpg" width="50%" height="50%">
