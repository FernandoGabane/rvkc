package repositories

import (
	"rvkc/config"
)


type Repository[T any] interface {
	Create(entity *T) error
	GetAll() ([]T, error)
	GetByID(id uint) (*T, error)
	GetBy(query interface{}, args ...interface{})(*T, error)
	Update(entity *T) error
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


func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := config.DB.Find(&entities).Error
	return entities, err
}


func (r *GenericRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := config.DB.First(&entity, id).Error
	return &entity, err
}

func (r *GenericRepository[T]) GetBy(query interface{}, args ...interface{}) (*T, error) {
	var entity T
	err := config.DB.Where(query, args...).First(&entity).Error
    return &entity, err
}


func (r *GenericRepository[T]) Update(entity *T) error {
	return config.DB.Save(entity).Error
}


func (r *GenericRepository[T]) Delete(id uint) error {
	return config.DB.Delete(new(T), id).Error
}