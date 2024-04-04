package main

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go-financial-transaction-summary/models"
	"go-financial-transaction-summary/repository"
	"go-financial-transaction-summary/repository/entity"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	c := require.New(t)

	createdTime := time.Now()

	transactionData := models.Transaction{
		ID:          1,
		Date:        createdTime,
		Transaction: 10.50,
		CreatedAt:   createdTime,
	}

	m := mock.Mock{}
	mockTransaction := repository.MockTransaction{Mock: &m}

	transaction := entity.Transaction{
		ID:          transactionData.ID,
		Date:        transactionData.Date,
		Transaction: transactionData.Transaction,
		CreatedAt:   transactionData.CreatedAt,
	}

	ctx := context.Background()

	mockTransaction.On("Insert", ctx, &transaction).Return(nil)

	err := insertTransaction(ctx, mockTransaction, transactionData)
	c.NoError(err)

	m.AssertExpectations(t)
}
