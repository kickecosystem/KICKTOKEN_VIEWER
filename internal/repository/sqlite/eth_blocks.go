package sqlite

import (
	"database/sql"
)

func (client *Client) SyncLastBlock(defaultLastBlock int64) error {

	lastBlockNumber, err := client.GetLastBlock()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		err := client.InsertLastBlock(defaultLastBlock)
		if err != nil {
			return err
		}
		return nil
	}

	if lastBlockNumber < defaultLastBlock {
		err := client.UpdateLastBlock(defaultLastBlock)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (client *Client) GetLastBlock() (int64, error) {
	var lastBlockNumber sql.NullInt64
	err := client.QueryRow(`SELECT last_block_number 
		FROM t_last_blocks `).Scan(&lastBlockNumber)
	if err != nil {
		return 0, err
	}

	return lastBlockNumber.Int64, nil
}

func (client *Client) InsertLastBlock(lastBlock int64) error {
	_, err := client.Exec(`INSERT INTO t_last_blocks 
		(last_block_number, ts_updated)
		VALUES ($1, CURRENT_TIMESTAMP)`, lastBlock)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) UpdateLastBlock(lastBlock int64) error {
	_, err := client.Exec(`UPDATE t_last_blocks 
		SET last_block_number = $1, ts_updated = CURRENT_TIMESTAMP`, lastBlock)
	if err != nil {
		return err
	}
	return nil
}
