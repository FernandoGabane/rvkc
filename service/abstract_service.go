package service

import (
	"rvkc/repositories"
)


type GenericService[T any] struct {
	repo repositories.Repository[T]
}


func NewGenericService[T any](repo repositories.Repository[T]) *GenericService[T] {
	return &GenericService[T]{repo: repo}
}


func (s *GenericService[T]) Create(entity *T) error {
	return s.repo.Create(entity)
}


func (s *GenericService[T]) GetAll() ([]T, error) {
	return s.repo.GetAll()
}


func (s *GenericService[T]) GetByID(id uint) (*T, error) {
	return s.repo.GetByID(id)
}

func (s *GenericService[T]) GetBy(query interface{}, args ...interface{}) (*T, error) {
	return s.repo.GetBy(query, args...)
}

func (s *GenericService[T]) Update(entity *T) error {
	return s.repo.Update(entity)
}


func (s *GenericService[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}