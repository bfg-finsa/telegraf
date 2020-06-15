package all

import (
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/activemq"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/aerospike"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/amqp_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/apache"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/apcupsd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/aurora"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/azure_storage_queue"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/bcache"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/beanstalkd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/bind"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/bond"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/burrow"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cassandra"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ceph"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cgroup"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/chrony"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cisco_telemetry_gnmi"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cisco_telemetry_mdt"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/clickhouse"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cloud_pubsub"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cloud_pubsub_push"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cloudwatch"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/conntrack"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/consul"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/couchbase"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/couchdb"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/cpu"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/dcos"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/disk"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/diskio"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/disque"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/dmcache"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/dns_query"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/docker"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/docker_log"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/dovecot"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ecs"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/elasticsearch"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ethtool"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/eventhub_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/exec"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/execd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/fail2ban"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/fibaro"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/file"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/filecount"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/filestat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/fireboard"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/fluentd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/github"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/graylog"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/haproxy"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/hddtemp"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/http"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/http_listener_v2"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/http_response"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/httpjson"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/icinga2"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/infiniband"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/influxdb"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/internal"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/interrupts"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ipmi_sensor"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ipset"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/iptables"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ipvs"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/jenkins"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/jolokia"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/jolokia2"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/jti_openconfig_telemetry"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kafka_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kafka_consumer_legacy"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kapacitor"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kernel"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kernel_vmstat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kibana"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kinesis_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kube_inventory"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/kubernetes"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/lanz"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/leofs"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/linux_sysctl_fs"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/logparser"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/logstash"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/lustre2"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mailchimp"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/marklogic"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mcrouter"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mem"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/memcached"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mesos"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/minecraft"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/modbus"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mongodb"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/monit"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mqtt_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/multifile"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/mysql"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nats"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nats_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/neptune_apex"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/net"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/net_response"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nginx"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nginx_plus"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nginx_plus_api"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nginx_upstream_check"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nginx_vts"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nsq"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nsq_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nstat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ntpq"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/nvidia_smi"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/openldap"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/openntpd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/opensmtpd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/openweathermap"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/passenger"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/pf"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/pgbouncer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/phpfpm"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/ping"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/postfix"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/postgresql"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/postgresql_extensible"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/powerdns"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/powerdns_recursor"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/processes"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/procstat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/prometheus"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/puppetagent"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/rabbitmq"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/raindrops"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/redis"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/redis_consumer"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/rethinkdb"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/riak"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/salesforce"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/sensors"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/sflow"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/smart"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/snmp"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/snmp_legacy"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/snmp_trap"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/socket_listener"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/solr"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/sqlserver"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/stackdriver"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/statsd"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/suricata"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/swap"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/synproxy"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/syslog"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/sysstat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/system"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/systemd_units"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/tail"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/teamspeak"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/temp"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/tengine"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/tomcat"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/trig"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/twemproxy"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/udp_listener"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/unbound"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/uwsgi"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/varnish"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/vsphere"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/webhooks"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/win_perf_counters"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/win_services"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/wireguard"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/wireless"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/x509_cert"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/zfs"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/zipkin"
	_ "github.com/bfg-finsa/telegraf/plugins/inputs/zookeeper"
)
