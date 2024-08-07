package clients

import (
	"context"
	"database/sql"
	"github.com/VoltDB/voltdb-client-go/voltdbclient"
	_ "github.com/go-sql-driver/mysql" // MySQL driver for TiDB
	"github.com/gocql/gocql"
	"github.com/jackc/pgx/v5"
	"github.com/rqlite/gorqlite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/redis.v5"
)

type DatabaseClient interface {
	Connect() error
	Close() error
}

type PostgresClient struct {
	conn *pgx.Conn
}

func (c *PostgresClient) Connect() error {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:admin@localhost:5432/testdb")
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *PostgresClient) Close() error {
	return c.conn.Close(context.Background())
}

type MongoDBClient struct {
	client *mongo.Client
}

func (c *MongoDBClient) Connect() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *MongoDBClient) Close() error {
	return c.client.Disconnect(context.Background())
}

type RedisClient struct {
	client *redis.Client
}

func (c *RedisClient) Connect() error {
	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)
	c.client = client
	return nil
}

func (c *RedisClient) Close() error { return c.client.Close() }

type RiakClient struct{ client *gorqlite.Connection }

func (c *RiakClient) Connect() error {
	client, err := gorqlite.Open("http://localhost:8087")
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *RiakClient) Close() error { return nil }

type CassandraClient struct{ session *gocql.Session }

func (c *CassandraClient) Connect() error {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	c.session = session
	return nil
}

func (c *CassandraClient) Close() error {
	c.session.Close()
	return nil
}

type CockroachDBClient struct{ conn *pgx.Conn }

func (c *CockroachDBClient) Connect() error {
	//goland:noinspection SpellCheckingInspection
	conn, err := pgx.Connect(context.Background(), "postgres://root@localhost:26257/defaultdb?sslmode=disable")
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *CockroachDBClient) Close() error { return c.conn.Close(context.Background()) }

type VoltDBClient struct{ client *voltdbclient.Conn }

func (c *VoltDBClient) Connect() error {
	client, err := voltdbclient.OpenConn("localhost:21212")
	if err != nil {
		return err
	}
	c.client = client
	return nil
}

func (c *VoltDBClient) Close() error { return c.client.Close() }

type TiDBClient struct{ db *sql.DB }

func (c *TiDBClient) Connect() error {
	db, err := sql.Open("mysql", "root:@tcp(localhost:4000)/test")
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *TiDBClient) Close() error { return c.db.Close() }
