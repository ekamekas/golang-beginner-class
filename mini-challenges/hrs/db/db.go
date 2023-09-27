package db

import "errors"

type Db struct {
    IndexById map[string]*PersonEntity
    IndexByName map[string]*PersonEntity
    database []PersonEntity
}

func (this *Db) Insert(entity PersonEntity) {
    this.database = append(this.database, entity);

    this.IndexById[entity.Id] = &this.database[len(this.database) - 1];
    this.IndexByName[entity.Name] = &this.database[len(this.database) - 1];
}

func (this *Db) FindByName(name string) (PersonEntity, error) {
    person, doesExist := this.IndexByName[name];

    if(!doesExist) {
        return PersonEntity{}, errors.New("Person does not found");
    }

    return *person, nil;
}

func (this *Db) FindById(id string) (PersonEntity, error) {
    person, doesExist := this.IndexById[id];

    if(!doesExist) {
        return PersonEntity{}, errors.New("Person does not found");
    }

    return *person, nil;
}
