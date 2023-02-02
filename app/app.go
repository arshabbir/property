package app

import (
	"encoding/json"
	"fmt"
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
	// Check if the property is already exists
	prop, err := ctx.GetStub().GetState(property.Id)
	if err != nil {
		log.Println(" error while fetching the state .", err)
		return err
	}

	if prop != nil {
		log.Println("property already exits...")
		return fmt.Errorf("property already exists")
	}

	bytes, err := json.Marshal(&property)
	if err != nil {
		log.Println("error while marshaling ", err)
		return err
	}

	return ctx.GetStub().PutState(property.Id, bytes)

}

func (a *propertyApp) ReadAll(ctx contractapi.TransactionContextInterface) ([]models.Property, error) {
	propIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		log.Println("error while fetching the iterator ", err)
		return nil, err
	}
	defer propIterator.Close()
	var results []models.Property
	for propIterator.HasNext() {
		v, err := propIterator.Next()
		if err != nil {
			log.Println("error while navigating the iterator ", err)
			return nil, err
		}
		prop := models.Property{}
		if err = json.Unmarshal(v.Value, &prop); err != nil {
			return nil, err
		}
		results = append(results, prop)
	}
	return results, nil
}

func (a *propertyApp) ReadById(ctx contractapi.TransactionContextInterface, id string) (*models.Property, error) {
	prop, err := ctx.GetStub().GetState(id)
	if err != nil {
		log.Println("error while fetching the property ", err)
		return nil, err
	}
	if prop == nil {
		log.Println("No property exists with the ID  ", id)
		return nil, fmt.Errorf("no property exits with ID : %s", id)
	}

	p := &models.Property{}
	err = json.Unmarshal(prop, p)
	if err != nil {
		log.Println("error while unmarshaling ", err)
		return nil, err
	}

	return p, nil
}

func (a *propertyApp) TransferProperty(ctx contractapi.TransactionContextInterface, Id string, oldOwner string, newOwner string) error {
	prop, err := a.ReadById(ctx, Id)
	if err != nil {
		log.Println("error while fetching the property ", err)
		return err
	}
	prop.Owner = newOwner
	bytes, err := json.Marshal(&prop)
	if err != nil {
		log.Println("error while unmarshaling ", err)
		return err
	}
	return ctx.GetStub().PutState(Id, bytes)
}
