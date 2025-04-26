package postgress

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/domain/entities"
	"github.com/Guilherme-DSGL/purchase_transaction_backend/internal/repository"
	"github.com/sirupsen/logrus"
)

type TransactionRepository struct {
	Conn *sql.DB
}

func NewTransactionRepository(conn *sql.DB) *TransactionRepository {
	return &TransactionRepository{conn}
}

func (m *TransactionRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []entities.Transaction, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]entities.Transaction, 0)
	for rows.Next() {
		t := entities.Transaction{}
		err = rows.Scan(
			&t.ID,
			&t.Description,
			&t.Date,
			&t.Value,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *TransactionRepository) Fetch(ctx context.Context, cursor string, num int64) (res []entities.Transaction, nextCursor string, err error) {
	query := `SELECT id, description, date, value, updated_at, created_at
  						FROM transaction WHERE created_at > ? ORDER BY created_at LIMIT ? `

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return res, nextCursor, nil
}

func (m *TransactionRepository) GetById(ctx context.Context, id string) (res entities.Transaction, err error) {
	query := `SELECT id, description, date, value, updated_at, created_at
  						FROM transaction WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return entities.Transaction{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return res, nil
}

func (m *TransactionRepository) Save(ctx context.Context, t *entities.Transaction) (err error) {
	query := `INSERT transaction SET id=?, description=?, date=?, value=?, updated_at=?, created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.ID, t.Description, t.Date, t.Value, t.UpdatedAt, t.CreatedAt)
	if err != nil {
		return err
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("atention! Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}

func (m *TransactionRepository) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM transaction WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("atention! Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}

func (m *TransactionRepository) Update(ctx context.Context, t *entities.Transaction) (err error) {
	query := `UPDATE transaction SET description=?, date=?, value=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.Description, t.Date, t.Value, t.UpdatedAt, t.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}
	return
}
