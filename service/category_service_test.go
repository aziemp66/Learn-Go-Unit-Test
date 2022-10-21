package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
	categoryService    = CategoryService{Repository: categoryRepository}
)

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "0").Return(nil)
	category, err := categoryService.Get("0")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Handphone",
	}
	categoryRepository.Mock.On("FindById", "1").Return(category)

	result, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Id, category.Id)
	assert.Equal(t, result.Name, category.Name)
}
