// Code generated by MockGen. DO NOT EDIT.
// Source: app/domain/repository/category_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "OnionPractice/app/domain/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockCategoryRepository) CreateCategory(ctx context.Context, category model.Category) (*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", ctx, category)
	ret0, _ := ret[0].(*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockCategoryRepositoryMockRecorder) CreateCategory(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockCategoryRepository)(nil).CreateCategory), ctx, category)
}

// GetAllCategories mocks base method.
func (m *MockCategoryRepository) GetAllCategories(ctx context.Context) ([]*model.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategories", ctx)
	ret0, _ := ret[0].([]*model.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategories indicates an expected call of GetAllCategories.
func (mr *MockCategoryRepositoryMockRecorder) GetAllCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategories", reflect.TypeOf((*MockCategoryRepository)(nil).GetAllCategories), ctx)
}