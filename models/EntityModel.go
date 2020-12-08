package models

// Entity Model
type Entity struct {
	ID     string `json:"id"`
	Value  string `json:"value"`
	Data   int64  `json:"data"`
	Result string `json:"result"`
}

type EntityDataStore struct {
	*Connection
}

// Creates a new Entity record
func (d *EntityDataStore) Insert(Entity *Entity) error {
	err := d.Connection.DB.Create(&Entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Getting a first Entity record by value
func (d *EntityDataStore) GetByValue(Value string) (*Entity, error) {
	entity := &Entity{}
	err := d.Connection.DB.Where("value = ?", Value).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Cleat Entities table
func (d *EntityDataStore) ClearTable() error {
	err := d.Connection.DB.Exec("delete from entities").Error
	if err != nil {
		return err
	}
	return nil
}

// Update Entity model
func (d *EntityDataStore) Update(Entity *Entity) error {
	err := d.Connection.DB.Model(&Entity).Save(&Entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Getting the all Entity records
func (d *EntityDataStore) Get() (*[]Entity, error) {
	entities := &[]Entity{}
	err := d.Connection.DB.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}
