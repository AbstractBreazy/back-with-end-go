package managers

import (
	"back-with-end-go/models"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math"
	"strconv"
)

type EntityManager struct {
	EntityDataStore models.EntityDataStore
}

func InitEntityManager(connection *models.Connection) EntityManager {
	return EntityManager{EntityDataStore: models.EntityDataStore{connection}}
}

func (d *EntityManager) PostEntity(value string, data float64) (*models.Entity, error) {
	// Generating new UUID
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	newEntity := &models.Entity{
		ID:    newID.String(),
		Value: value,
		Data:  int64(data),
	}
	// Creates newEntity record
	err = d.EntityDataStore.Insert(newEntity)
	if err != nil {
		return nil, err
	}
	return newEntity, nil
}

func (d *EntityManager) ProcessEntity(value string, data float64) ([]float64, error) {
	// Getting a first Entity record by value
	fetchedEntity, err := d.EntityDataStore.GetByValue(value)
	if err != nil {
		return nil, err
	}
	// Calculate result string
	r := data / float64(4)
	integerPoint, _ := math.Modf(r)
	r = integerPoint * float64(fetchedEntity.Data) * 8
	dataString := fmt.Sprintf("%v,%v,%v,%v", r, r, r, r)
	fmt.Printf("fetchedString: %v\n", dataString)
	resultSlice := []float64{float64(r), float64(r), float64(r), float64(r)}
	fetchedEntity.Result = dataString
	// updating the result field in a Entity record
	err = d.EntityDataStore.Update(fetchedEntity)
	if err != nil {
		return nil, err
	}
	return resultSlice, nil
}

func (d *EntityManager) IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
