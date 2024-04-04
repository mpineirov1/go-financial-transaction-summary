package repository

import (
	"context"
	"go-financial-transaction-summary/repository/entity"

	"github.com/stretchr/testify/mock"
)

type MockTransaction struct {
	*mock.Mock
	Transaction
}

func (m MockTransaction) Insert(ctx context.Context, transaction *entity.Transaction) (err error) {
	results := m.Called(ctx, transaction)
	return results.Error(0)
}
