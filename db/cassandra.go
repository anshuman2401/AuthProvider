package db

import (
	"AuthProvider/conf"
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

func GetCassandraConnection() *gocql.Session{

	config := conf.GetConfig.Cassandra
	cluster := gocql.NewCluster(config.Cluster...)
	cluster.PoolConfig = gocql.PoolConfig{
		HostSelectionPolicy: gocql.DCAwareRoundRobinPolicy(config.DC),
	}
	cluster.Consistency = gocql.LocalOne
	cluster.SerialConsistency = gocql.LocalSerial
	cluster.Keyspace = config.Keyspace
	cluster.Timeout = config.Timeout * time.Millisecond
	cluster.ConnectTimeout = config.Timeout * time.Millisecond
	cluster.NumConns = config.NumConns
	cluster.SocketKeepalive = config.KeepAlive * time.Second

	session, err := cluster.CreateSession()

	if err != nil {
		fmt.Print("Error in connecting to Cassandra...")
		return nil
	}

	return session
}
