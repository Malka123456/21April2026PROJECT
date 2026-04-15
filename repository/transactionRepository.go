package repository

// import (
// 	"learning-backend/models"
// 	// "learning-backend/dto"
// 	"gorm.io/gorm"
// )

// type TransactionRepository interface {
// 	CreatePayment(payment *models.Payment) error
// 	FindInitialPayment(uId uint) (*models.Payment, error)
// 	UpdatePayment(payment *models.Payment) error
// 	FindOrders(uId uint) ([]models.OrderItem, error)
// 	FindOrderById(uId uint, id uint) (dto.SellerOrderDetails, error)
// }

// type transactionStorage struct {
// 	db *gorm.DB
// }

// func (t transactionStorage) UpdatePayment(payment *models.Payment) error {
// 	return t.db.Save(payment).Error
// }

// func (t transactionStorage) FindInitialPayment(uId uint) (*models.Payment, error) {
// 	var payment *models.Payment
// 	err := t.db.First(&payment, "user_id=? AND status=?", uId, "initial").Order("created_at desc").Error
// 	return payment, err
// }

// func (t transactionStorage) CreatePayment(payment *models.Payment) error {
// 	return t.db.Create(payment).Error
// }

// func (t transactionStorage) FindOrders(uId uint) ([]models.OrderItem, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (t transactionStorage) FindOrderById(uId uint, id uint) (dto.SellerOrderDetails, error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func NewTransactionRepository(db *gorm.DB) TransactionRepository {
// 	return &transactionStorage{db: db}
// }
