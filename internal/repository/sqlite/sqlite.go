package sqlite

import (
	"database/sql"
	"github.com/rs/zerolog"
	"os"
)

type Client struct {
	*sql.DB
	Log *zerolog.Logger
}

func InitSqliteClient(log *zerolog.Logger) (*Client, error) {

	dbPath := "kicktoken_cache.sqlite"
	var fileDb *os.File

	_, err := os.Stat(dbPath)
	if err == nil {
		fileDb, err = os.Open(dbPath)
		if err != nil {
			return nil, err
		}
	} else {
		fileDb, err = os.Create(dbPath)
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	defer func() {
		err := fileDb.Close()
		if err != nil {
			log.Error().Caller().Err(err).Msg("")
		}
	}()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Client{
		DB:  db,
		Log: log,
	}, nil
}

func (client *Client) MigrateUp() error {
	_, err := client.Exec("CREATE TABLE IF NOT EXISTS t_eth_logs (hash text,event text, address_first text, address_second text, kick_amount text, block_number numeric, method text, ts_created timestamp default CURRENT_TIMESTAMP, ts_updated timestamp);" +
		"CREATE TABLE IF NOT EXISTS t_last_blocks (last_block_number bigint, ts_created timestamp default CURRENT_TIMESTAMP, ts_updated timestamp);" +
		"CREATE INDEX IF NOT EXISTS t_eth_logs_hash_index on t_eth_logs (hash);" +
		"CREATE INDEX IF NOT EXISTS t_eth_logs_address_first_index on t_eth_logs (address_first);" +
		"CREATE INDEX IF NOT EXISTS t_eth_logs_address_second_index on t_eth_logs (address_second);" +
		"CREATE INDEX IF NOT EXISTS t_eth_logs_event_index on t_eth_logs (event);" +
		"CREATE INDEX IF NOT EXISTS t_eth_logs_method_index on t_eth_logs (method);")
	if err != nil {
		return nil
	}
	return err
}
