package repository

import (
	"gorm.io/gorm"
	"goutils/collections"
)

// CrudRepository defines generic CRUD operations akin to JpaRepository.
// T is the entity type, ID is the primary key type (e.g., uint, int64, string).
type CrudRepository[T comparable, ID comparable] interface {
	// Save persists the given entity.
	Save(entity *T) error

	SaveAll(entities []T) error

	// FindByID retrieves an entity by its primary key.
	FindByID(id ID) (*T, error)

	// FindAll returns all entities wrapped in a Collection.
	FindAll() (collections.Collection[T], error)

	// FindByQuery executes a custom query builder function and returns matching entities.
	FindByQuery(queryFn func(*gorm.DB) *gorm.DB) (collections.Collection[T], error)

	// Delete removes the given entity.
	Delete(entity *T) error

	// DeleteByID removes the entity with the given primary key.
	DeleteByID(id ID) error

	// Count returns the total number of entities of type T.
	Count() (int64, error)
}

// GormRepository is a GORM-based implementation of CrudRepository.
// Instantiate with NewGormRepository for any model type.
type GormRepository[T comparable, ID comparable] struct {
	db *gorm.DB
}

// Gorm constructs a new GormRepository for the given GORM DB instance.
func Gorm[T comparable, ID comparable](db *gorm.DB) *GormRepository[T, ID] {
	return &GormRepository[T, ID]{db: db}
}

func (r *GormRepository[T, ID]) Save(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GormRepository[T, ID]) SaveAll(entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	return r.db.Create(&entities).Error
}

func (r *GormRepository[T, ID]) FindByID(id ID) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GormRepository[T, ID]) FindAll() (collections.Collection[T], error) {
	var items []T
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return collections.List(items...), nil
}

func (r *GormRepository[T, ID]) FindByQuery(queryFn func(*gorm.DB) *gorm.DB) (collections.Collection[T], error) {
	var items []T
	q := queryFn(r.db)
	if err := q.Find(&items).Error; err != nil {
		return nil, err
	}
	return collections.List(items...), nil
}

func (r *GormRepository[T, ID]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *GormRepository[T, ID]) DeleteByID(id ID) error {
	return r.db.Delete(new(T), id).Error
}

func (r *GormRepository[T, ID]) Count() (int64, error) {
	var count int64
	if err := r.db.Model(new(T)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
