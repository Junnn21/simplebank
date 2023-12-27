// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt sql.NullTime
}

type Entry struct {
	ID        int64
	AccountID int64
	// can be negative or positive
	Amount    int64
	CreatedAt sql.NullTime
}

type Session struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type Transfer struct {
	ID            int64
	FromAccountID int64
	ToAccountID   int64
	// must be positive
	Amount    int64
	CreatedAt sql.NullTime
}

type User struct {
	Username          string
	HashedPassword    string
	FullName          string
	Email             string
	PasswordChangedAt time.Time
	CreatedAt         sql.NullTime
	IsEmailVerified   bool
}

type VerifyEmail struct {
	ID         int64
	Username   string
	Email      string
	SecretCode string
	IsUsed     bool
	CreatedAt  time.Time
	ExpiredAt  time.Time
}
