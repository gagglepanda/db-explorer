package clients

//goland:noinspection GoUnusedGlobalVariable
var (
	NewPostgresClient    = func() *PostgresClient { return &PostgresClient{} }
	NewMongoDBClient     = func() *MongoDBClient { return &MongoDBClient{} }
	NewRedisClient       = func() *RedisClient { return &RedisClient{} }
	NewRiakClient        = func() *RiakClient { return &RiakClient{} }
	NewCassandraClient   = func() *CassandraClient { return &CassandraClient{} }
	NewCockroachDBClient = func() *CockroachDBClient { return &CockroachDBClient{} }
	NewVoltDBClient      = func() *VoltDBClient { return &VoltDBClient{} }
	NewTiDBClient        = func() *TiDBClient { return &TiDBClient{} }
)
