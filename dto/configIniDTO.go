package dto

const ConfigIniFile = `
plugin = eosio::producer_plugin
plugin = eosio::net_plugin
plugin = eosio::net_api_plugin
plugin = eosio::chain_api_plugin
plugin = eosio::account_history_plugin
plugin = eosio::account_history_api_plugin
plugin = eosio::wallet_api_plugin

get-transactions-time-limit = 3
genesis-json = "genesis.json"
block-log-dir = "blocks"
max-reversible-block-time = -1
max-pending-transaction-time = -1
max-deferred-transaction-time = 20
access-control-allow-credentials = false
mongodb-queue-size = 256
allowed-connection = any
log-level-net-plugin = info
max-clients = 25
connection-cleanup-period = 30
network-version-match = 0
sync-fetch-span = 1000
enable-stale-production = true
required-participation = 33
wallet-dir = "."



private-key = ["{{.PublicKey}}","{{.PrivateKey}}"]
http-server-address = 0.0.0.0:{{.HttpPort}}
p2p-listen-endpoint = 0.0.0.0:{{.P2pPort}}
p2p-server-address = 0.0.0.0:{{.P2pPort}}
producer-name = {{.ProducerName}}
agent-name = {{.ProducerName}}


###########################
p2p-peer-address = node0.eosio.sg:9800
###########################
`

const Genesis=`{
  "initial_key": "EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV",
  "initial_timestamp": "2018-04-20T00:00:00",
  "initial_parameters": {
    "maintenance_interval": 86400,
    "maintenance_skip_slots": 3,
    "maximum_transaction_size": 2048,
    "maximum_block_size": 2048000000,
    "maximum_time_until_expiration": 86400,
    "maximum_producer_count": 1001
  },
  "immutable_parameters": {
    "min_producer_count": 21
  },
  "initial_chain_id": "0000000000000000000000000000000000000000000000000000000000000001"
}`

const Run=`mkdir -p logs;
nohup nodeos --data-dir ./data-dir --config-dir ./config-dir  > ./logs/eos.log 2>&1 &
echo $! > eos.pid`

const Kill="kill `cat eos.pid`\nrm eos.pid"