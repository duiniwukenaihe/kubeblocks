[vttablet]
health_check_interval=1s
shard_sync_retry_delay=1s
remote_operation_timeout=1s
db_connect_timeout_ms=500
table_acl_config_mode=simple
enable_logs=true
enable_query_log=true
table_acl_config=
queryserver_config_strict_table_acl=false
table_acl_config_reload_interval=30s
enforce_tableacl_config=false

# OltpReadPool
queryserver-config-pool-size=30

# OlapReadPool
queryserver-config-stream-pool-size=30

# TxPool
queryserver-config-transaction-cap=50