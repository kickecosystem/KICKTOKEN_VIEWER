package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/schollz/progressbar/v2"
	"github.com/shopspring/decimal"
)

type TxFromBc struct {
	AddressFirst  string
	AddressSecond string
	Amount        decimal.Decimal
	BlockNumber   uint64
	EventName     string
}

type TxFromDb struct {
	AddressFirst  string
	AddressSecond string
	Amount        decimal.Decimal
	EventName     string
	Method        string
}

func (client *Client) InsertTxs(txFromBc map[string][]TxFromBc) error {

	tx, err := client.Begin()
	if err != nil {
		client.Log.Error().Caller().Err(err).Msg("")
		return err
	}

	defer func() {
		err := tx.Rollback()
		if err != nil && err != sql.ErrTxDone {
			client.Log.Fatal().Caller().Err(err).Msg("ошибка при Rollback транзакции")
		}
	}()

	for txHex, txLogs := range txFromBc {

		_, err = tx.Exec(`DELETE
		FROM t_eth_logs
		WHERE hash = $1`, txHex)
		if err != nil && err != sql.ErrNoRows {
			client.Log.Error().Caller().Err(err).Msg("")
			return err
		}

		for _, tr := range txLogs {
			_, err := tx.Exec(`INSERT INTO t_eth_logs 
				(hash, address_first, address_second, block_number, kick_amount, event, ts_updated)
				VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)`,
				txHex, tr.AddressFirst, tr.AddressSecond, tr.BlockNumber, tr.Amount, tr.EventName)
			if err != nil {
				client.Log.Error().Caller().Err(err).Msg("")
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		client.Log.Error().Caller().Err(err).Msg("")
		return err
	}

	return nil
}

func (client *Client) UpdateTxMethod(hash string, method string) error {
	_, err := client.Exec(`UPDATE t_eth_logs 
		SET method = $1, ts_updated = CURRENT_TIMESTAMP 
		WHERE hash = $2 `, method, hash)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (client *Client) GetEthBcTxEmptyMethod() ([]string, error) {

	var txs []string
	rows, err := client.Query(`SELECT DISTINCT hash
		FROM t_eth_logs
		WHERE method IS NULL`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer func() {
		if rows != nil {
			err := rows.Close()
			if err != nil {
				client.Log.Error().Caller().Err(err).Msg("")
			}
		}
	}()

	if rows != nil {
		for rows.Next() {
			var txHash sql.NullString
			err = rows.Scan(&txHash)
			if err != nil {
				return nil, err
			}
			txs = append(txs, txHash.String)
		}
	}

	return txs, nil
}

func (client *Client) GetTxCount() (int64, error) {
	var count sql.NullInt64
	err := client.QueryRow(`SELECT count(*) count 
		FROM t_eth_logs`).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count.Int64, nil
}

func (client *Client) GetAllTransactions(txAmount int) ([]TxFromDb, error) {

	var txs []TxFromDb
	rows, err := client.Query(`SELECT address_first, address_second, event, method, kick_amount
		FROM t_eth_logs`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer func() {
		if rows != nil {
			err := rows.Close()
			if err != nil {
				client.Log.Error().Caller().Err(err).Msg("")
			}
		}
	}()

	bar := progressbar.New(txAmount)

	if rows != nil {
		for rows.Next() {
			var addressFirst sql.NullString
			var addressSecond sql.NullString
			var event sql.NullString
			var method sql.NullString
			var kickAmountStr sql.NullString
			err = rows.Scan(&addressFirst, &addressSecond, &event, &method, &kickAmountStr)
			if err != nil {
				return nil, err
			}

			kickAmountDecimal, err := decimal.NewFromString(kickAmountStr.String)
			if err != nil {
				return nil, err
			}

			txs = append(txs, TxFromDb{
				AddressFirst:  addressFirst.String,
				AddressSecond: addressSecond.String,
				Amount:        kickAmountDecimal,
				EventName:     event.String,
				Method:        method.String,
			})

			err = bar.Add(1)
			if err != nil {
				return nil, err
			}
		}
	}

	err = bar.Finish()
	if err != nil {
		return nil, err
	}
	fmt.Println("")

	return txs, nil
}
