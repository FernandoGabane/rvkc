package repositories

import (
	"fmt"
	"rvkc/config"

	"gorm.io/gorm"
)


type Repository[T any] interface {
	Create(entity *T) error
	CreateBatch(entities *[]T) error
	CreateWithAssociations(entity *T, associations map[string][]any) error
	GetAll() ([]*T, error)
	GetByID(id uint) (*T, error)
	GetBy(query interface{}, args ...interface{}) (*T, error)
	FindBy(query interface{}, args ...interface{}) (*T, error)
	FindAllBy(query interface{}, args ...interface{}) ([]*T, error)
	SubQuery(selectClause string, groupBy string) *gorm.DB
	FindWithPreloads(preloads []string, whereClause interface{}, args ...interface{}) ([]*T, error)
	FindWithJoinsAndFilters(joins []string, whereClause interface{}, args ...interface{}) ([]*T, error)
	FindWithJoinsSubqueryAndPreloads(joins []interface{}, preloads []string,	whereClause interface{}, args ...interface{}) ([]*T, error)
	Update(entity *T) error
	UpdateWithAssociations(entity *T, associations map[string][]any) error
	Delete(id uint) error
}


type GenericRepository[T any] struct {
	
}


func NewGenericRepository[T any]() *GenericRepository[T] {
	return &GenericRepository[T]{}
}


func (r *GenericRepository[T]) Create(entity *T) error {
	return config.DB.Create(entity).Error
}


func (r *GenericRepository[T]) CreateBatch(entities *[]T) error {
	if len(*entities) == 0 {
		return nil
	}

	tx := config.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(&entities).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}



func (r *GenericRepository[T]) CreateWithAssociations(entity *T, associations map[string][]any) error {
	tx := config.DB.Begin()

	if err := tx.Create(entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	for assocName, values := range associations {
		if len(values) == 0 {
			continue
		}
		// usa slice genérico como referência
		if err := tx.Model(entity).Association(assocName).Append(values...); err != nil {
			tx.Rollback()
			return fmt.Errorf("erro ao associar '%s': %w", assocName, err)
		}
	}

	return tx.Commit().Error
}


func (r *GenericRepository[T]) GetAll() ([]*T, error) {
	var entities []*T
	err := config.DB.Find(&entities).Error
	return entities, err
}


func (r *GenericRepository[T]) GetByID(id uint) (*T, error) {
	var entity *T
	err := config.DB.First(&entity, id).Error
	return entity, err
}


func (r *GenericRepository[T]) GetBy(query interface{}, args ...interface{}) (*T, error) {
	var entity *T
	err := config.DB.Where(query, args...).First(&entity).Error
    return entity, err
}


func (r *GenericRepository[T]) FindBy(query interface{}, args ...interface{}) (*T, error) {
	var entity *T
	err := config.DB.Where(query, args...).Find(&entity).Error
    return entity, err
}


func (r *GenericRepository[T]) FindAllBy(query interface{}, args ...interface{}) ([]*T, error) {
	var entity []*T
	err := config.DB.Where(query, args...).Find(&entity).Error
    return entity, err
}


func (r *GenericRepository[T]) FindWithJoinsAndFilters(joins []string, whereClause interface{},	args ...interface{}) ([]*T, error) {
	var entities []*T

	db := config.DB

	for _, join := range joins {
		db = db.Joins(join)
	} 

	if whereClause != nil {
		db = db.Where(whereClause, args...)
	}

	err := db.Find(&entities).Error
	return entities, err
}


func (s *GenericRepository[T]) SubQuery(selectClause string, groupBy string) *gorm.DB {
	var entity T
	return config.DB.
		Model(&entity).
		Select(selectClause).
		Group(groupBy)
}


func(s *GenericRepository[T]) Preload(preloads []string) *gorm.DB {
	var entity T
	db := config.DB.Model(&entity)

	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	return db
}

func (r *GenericRepository[T]) FindWithJoinsSubqueryAndPreloads(joins []interface{}, preloads []string,	whereClause interface{}, args ...interface{}) ([]*T, error) {
	var entities []*T

	db := config.DB.Model(&entities)

	for _, join := range joins {
		switch j := join.(type) {
		case string:
			db = db.Joins(j)
		case []interface{}:
			if len(j) == 2 {
				db = db.Joins(j[0].(string), j[1])
			}
		}
	}

	db = r.applyPreloads(db, preloads)

	if whereClause != nil {
		db = db.Where(whereClause, args...)
	}

	err := db.Find(&entities).Error
	return entities, err
}


func (r *GenericRepository[T]) FindWithPreloads(preloads []string, whereClause interface{}, args ...interface{}) ([]*T, error) {
	var entities []*T
	db := config.DB.Model(&entities)

	db = r.applyPreloads(db, preloads)

	if whereClause != nil {
		db = db.Where(whereClause, args...)
	}

	err := db.Find(&entities).Error
	return entities, err
}


func (r *GenericRepository[T]) Update(entity *T) error {
	return config.DB.Save(entity).Error
}


func (r *GenericRepository[T]) UpdateWithAssociations(entity *T, associations map[string][]any) error {
	tx := config.DB.Begin()

	if err := tx.Save(entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	for assocName, values := range associations {
		if err := tx.Model(entity).Association(assocName).Replace(values...); err != nil {
			tx.Rollback()
			return fmt.Errorf("erro ao atualizar associação '%s': %w", assocName, err)
		}
	}

	return tx.Commit().Error
}


func (r *GenericRepository[T]) Delete(id uint) error {
	return config.DB.Delete(new(T), id).Error
}


func (r *GenericRepository[T]) applyPreloads(db *gorm.DB, preloads []string) *gorm.DB {
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db
}
