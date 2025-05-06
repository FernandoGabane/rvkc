package service

import (
	"rvkc/repositories"

	"gorm.io/gorm"
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


func (s *GenericService[T]) CreateBatch(entities *[]T) error {
	return s.repo.CreateBatch(entities)
}


func (s *GenericService[T]) CreateWithAssociations(entity *T, associations map[string][]any) error {
	return s.repo.CreateWithAssociations(entity, associations)
}


func (s *GenericService[T]) GetAll() ([]*T, error) {
	return s.repo.GetAll()
}


func (s *GenericService[T]) GetByID(id uint) (*T, error) {
	return s.repo.GetByID(id)
}


func (s *GenericService[T]) GetBy(query interface{}, args ...interface{}) (*T, error) {
	return s.repo.GetBy(query, args...)
}


func (s *GenericService[T]) FindBy(query interface{}, args ...interface{}) (*T, error) {
	return s.repo.FindBy(query, args...)
}


func (s *GenericService[T]) FindAllBy(query interface{}, args ...interface{}) ([]*T, error) {
	return s.repo.FindAllBy(query, args...)
}


func (s *GenericService[T]) SubQuery(selectClause string, groupBy string) *gorm.DB {
	return s.repo.SubQuery(selectClause, groupBy)
}


func (s *GenericService[T]) FindWithPreloads(preloads []string, whereClause interface{}, args ...interface{}) ([]*T, error) {
	return s.repo.FindWithPreloads(preloads, whereClause, args...)
}


func (s *GenericService[T]) FindWithJoinsAndFilters(joins []string, whereClause interface{}, args ...interface{}) ([]*T, error) {
	return s.repo.FindWithJoinsAndFilters(joins, whereClause, args...)
}


func (s *GenericService[T]) FindWithJoinsSubqueryAndPreloads(joins []interface{}, preloads []string,	whereClause interface{}, args ...interface{}) ([]*T, error) {
	return s.repo.FindWithJoinsSubqueryAndPreloads(joins, preloads, whereClause, args...)
}


func (s *GenericService[T]) Update(entity *T) error {
	return s.repo.Update(entity)
}


func (s *GenericService[T]) UpdateWithAssociations(entity *T, associations map[string][]any) error {
	return s.repo.UpdateWithAssociations(entity, associations)
}


func (s *GenericService[T]) Delete(id uint) error {
	return s.repo.Delete(id)
}