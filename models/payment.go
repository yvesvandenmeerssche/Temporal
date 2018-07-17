package models

import (
	"errors"
	"math/big"

	"github.com/jinzhu/gorm"
)

type PinPayment struct {
	gorm.Model
	Method          uint8
	Number          string
	ChargeAmount    string
	UploaderAddress string
	ContentHash     string
}

type PinPaymentManager struct {
	DB *gorm.DB
}

func NewPinPaymentManager(db *gorm.DB) *PinPaymentManager {
	return &PinPaymentManager{DB: db}
}

func (ppm *PinPaymentManager) NewPayment(method uint8, number, chargeAmount *big.Int, uploaderAddress, contentHash string) (*PinPayment, error) {
	pp := &PinPayment{
		Number:          number.String(),
		Method:          method,
		ChargeAmount:    chargeAmount.String(),
		UploaderAddress: uploaderAddress,
		ContentHash:     contentHash,
	}
	if check := ppm.DB.Create(pp); check.Error != nil {
		return nil, check.Error
	}
	return pp, nil
}

func (ppm *PinPaymentManager) RetrieveLatestPaymentNumber() (*big.Int, error) {
	pp := &PinPayment{}
	num := big.NewInt(0)
	check := ppm.DB.Table("pin_payments").Order("number desc").First(pp)
	if check.Error != nil && check.Error != gorm.ErrRecordNotFound {
		return nil, check.Error
	}
	if check.Error == gorm.ErrRecordNotFound {
		return num, nil
	}
	var valid bool
	num, valid = num.SetString(pp.Number, 10)
	if !valid {
		return nil, errors.New("failed to convert from string to big int")
	}
	return num, nil
}
