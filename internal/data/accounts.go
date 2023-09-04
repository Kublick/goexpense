package data

import (
	"database/sql"
	"time"

	"github.com/kublick/goexpense/internal/validator"
)

type Account struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	AccountType string    `json:"accountType"`
	Budget      bool      `json:"budget"`
	UserID      int64     `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AccountModel struct {
	DB *sql.DB
}

func ValidateAccount(v *validator.Validator, account *Account) {
	v.Check(account.Name != "", "name", "must be provided")
	v.Check(account.AccountType != "", "accountType", "must be provided")
	v.Check(account.Balance != 0, "balance", "must be provided")
	v.Check(account.UserID != 0, "userId", "must be provided")
}

func (m AccountModel) Insert(a *Account) error {
	query := `
	INSERT INTO accounts (name, initial_balance, current_balance, account_type, budget, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, created_at, updated_at`

	args := []interface{}{a.Name, a.Balance, a.AccountType, a.Budget, a.UserID, a.CreatedAt, a.UpdatedAt}

	return m.DB.QueryRow(query, args...).Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt)
}

func (m AccountModel) GetAll(userID int64) ([]*Account, error) {
	query := `
	SELECT id, name, initial_balance, current_balance, account_type, budget, user_id, created_at, updated_at
	FROM accounts
	WHERE user_id = $1`

	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []*Account{}

	for rows.Next() {
		account := &Account{}
		err := rows.Scan(&account.ID, &account.Name, &account.Balance, &account.AccountType, &account.Budget, &account.UserID, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
