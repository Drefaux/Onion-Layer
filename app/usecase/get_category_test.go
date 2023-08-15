package usecase

import (
	"OnionPractice/app/domain/model"
	mockRepo "OnionPractice/util/testhelper/mock/repository"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var ErrMessage = errors.New("error occurred")

func TestGetAllCategoriesUseCase_GetAllCategories(t *testing.T) {
	type fields struct {
		setCategoryRepository func(mock *mockRepo.MockCategoryRepository)
	}

	tests := []struct {
		name        string
		fields      fields
		expectError bool
	}{
		{name: "success",
			fields: fields{
				setCategoryRepository: func(mock *mockRepo.MockCategoryRepository) {
					mock.EXPECT().GetAllCategories(gomock.Any()).Return([]*model.Category{}, nil)
				},
			},
			expectError: false,
		},
		{name: "unsuccess",
			fields: fields{
				setCategoryRepository: func(mock *mockRepo.MockCategoryRepository) {
					mock.EXPECT().GetAllCategories(gomock.Any()).Return(nil, fmt.Errorf("failed to ParseInt startDate: %w", ErrMessage))
				},
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mockRepo.NewMockCategoryRepository(ctrl)
			test.fields.setCategoryRepository(categoryRepository)

			u := GetAllCategoriesUseCase{
				categoryRepository: categoryRepository,
			}

			res, err := u.GetAllCategoriesUseCase(context.Background())
			if test.expectError {
				assert.NotNil(t, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}
