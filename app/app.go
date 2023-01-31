package app

import (
	"log"

	"github.com/arshabbir/propertymod/models"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type propertyApp struct {
	contractapi.Contract
	cc *contractapi.ContractChaincode
}

type App interface {
	StartApp() error
}

func NewPropertyApp() App {
	pApp := new(propertyApp)
	cc, err := contractapi.NewChaincode(pApp)
	if err != nil {
		return nil
	}
	return &propertyApp{cc: cc}
}

func (a *propertyApp) StartApp() error {
	log.Println("Starting Property App...")
	if err := a.cc.Start(); err != nil {
		return err
	}
	return nil

}

func (a *propertyApp) AddProperty(ctx contractapi.TransactionContextInterface, property models.Property) error {
	// To be started
	return nil

}

func (a *propertyApp) ReadAll(ctx contractapi.TransactionContextInterface) ([]models.Property, error) {
	return nil, nil
}

func (a *propertyApp) ReadById(ctx contractapi.TransactionContextInterface, id string) (*models.Property, error) {
	return nil, nil
}

func (a *propertyApp) TransferProperty(ctx contractapi.TransactionContextInterface, Id string, oldOwner string, newOwner string) error {
	return nil
}
