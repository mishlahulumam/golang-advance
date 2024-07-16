package postgres_gorm

import (
	"context"
	"errors"
	"golang-advance/assignment-2/wallet-service/entity"
	"golang-advance/assignment-2/wallet-service/service"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}
type walletRepository struct {
	db GormDBIface
}

func NewWalletRepository(db GormDBIface) service.IWalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) CreateWallet(ctx context.Context, userID int32) (*entity.Wallet, error) {
	wallet := &entity.Wallet{
		UserID:  uint(userID),
		Balance: 0.0,
	}

	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		log.Printf("Error creating wallet: %v\n", err)
		return nil, err
	}
	return wallet, nil
}

func (r *walletRepository) GetWallet(ctx context.Context, userID int32) (*entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Printf("Error getting wallet for user %d: %v\n", userID, err)
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) TopUp(ctx context.Context, userID int32, amount float32) (*entity.Wallet, error) {
	wallet, err := r.GetWallet(ctx, userID)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	wallet.Balance += amount
	if err := r.db.WithContext(ctx).Save(wallet).Error; err != nil {
		log.Printf("Error topping up wallet for user %d: %v\n", userID, err)
		return nil, err
	}
	return wallet, nil
}

func (r *walletRepository) Transfer(ctx context.Context, fromUserID, toUserID int32, amount float32) (*entity.Wallet, error) {
	fromWallet, err := r.GetWallet(ctx, fromUserID)
	if err != nil {
		return nil, err
	}
	if fromWallet == nil {
		return nil, errors.New("from user's wallet not found")
	}

	toWallet, err := r.GetWallet(ctx, toUserID)
	if err != nil {
		return nil, err
	}
	if toWallet == nil {
		return nil, errors.New("to user's wallet not found")
	}

	if fromWallet.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	fromWallet.Balance -= amount
	toWallet.Balance += amount

	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(fromWallet).Error; err != nil {
		tx.Rollback()
		log.Printf("Error updating from user's wallet: %v\n", err)
		return nil, err
	}

	if err := tx.Save(toWallet).Error; err != nil {
		tx.Rollback()
		log.Printf("Error updating to user's wallet: %v\n", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return nil, err
	}

	return fromWallet, nil
}

func (r *walletRepository) GetTransactions(ctx context.Context, userID int32) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transactions, nil
		}
		log.Printf("Error getting transactions for user %d: %v\n", userID, err)
		return nil, err
	}
	return transactions, nil
}
