package config

import (
	"github.com/gocql/gocql"
	"gopkg.in/yaml.v3"
	"os"
)

type CassandraSettings struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Keyspace string `yaml:"keyspace"`
}

func LoadCassandraSettings() (*CassandraSettings, error) {
	// Read all the contents of the settings file.
	data, err := os.ReadFile("settings.yml")
	if err != nil {
		return nil, err
	}

	// Decode the YAML data into a CassandraSettings struct.
	var settings CassandraSettings
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	// Return the CassandraSettings struct.
	return &settings, nil
}

func ConnectToCassandra(settings *CassandraSettings) (*gocql.Session, error) {
	cluster := gocql.NewCluster(settings.Host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: settings.Username,
		Password: settings.Password,
	}

	cluster.Keyspace = settings.Keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

//func main() {
//	// Load the Cassandra connection settings from the YAML file.
//	settings, err := LoadCassandraSettings()
//	if err != nil {
//		panic(err)
//	}
//
//	// Connect to the Cassandra database.
//	session, err := ConnectToCassandra(settings)
//	if err != nil {
//		panic(err)
//	}
//
//	// Close the Cassandra session when finished.
//	defer session.Close()
//
//	// Perform database operations using the `session` object.
//}
