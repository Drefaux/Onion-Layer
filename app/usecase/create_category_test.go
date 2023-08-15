package usecase

import (
	"OnionPractice/app/domain/model"
	mockRepo "OnionPractice/util/testhelper/mock/repository"
	"context"
	"fmt"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryUseCase_CreateCategory(t *testing.T) {
	category := model.NewCategory(1, 1, "Test")

	type fields struct {
		setCategoryRepository func(mock *mockRepo.MockCategoryRepository)
	}

	type args struct {
		input CreateCategoryUseCaseInput
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		expectErr bool
	}{
		{
			name: "success",
			fields: fields{
				setCategoryRepository: func(mock *mockRepo.MockCategoryRepository) {
					mock.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(&category, nil)
				},
			},
			args: args{
				input: CreateCategoryUseCaseInput{
					ID:        1,
					CompanyID: 1,
					Name:      "Test",
				},
			},
			expectErr: false,
		},
		{
			name: "unsuccess",
			fields: fields{
				setCategoryRepository: func(mock *mockRepo.MockCategoryRepository) {
					mock.EXPECT().CreateCategory(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("some error occured %w", ErrMessage))
				},
			},
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			categoryRepository := mockRepo.NewMockCategoryRepository(ctrl)
			test.fields.setCategoryRepository(categoryRepository)

			u := CreateCategoryUseCase{
				categoryRepository: categoryRepository,
			}

			res, err := u.CreateCategoryUseCase(context.Background(), test.args.input)
			if test.expectErr {
				assert.NotNil(t, err)
				assert.Nil(t, res)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, res)
			}
		})
	}
}
