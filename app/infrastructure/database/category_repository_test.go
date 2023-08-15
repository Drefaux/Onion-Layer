package database

import (
	"OnionPractice/app/domain/model"
	"OnionPractice/app/infrastructure/database/dbmodel"
	"OnionPractice/db"
	"OnionPractice/util/testhelper"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func newCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{db: db.DB}
}

func getAllCategories() []dbmodel.Category {
	var categories []dbmodel.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		panic(err)
	}

	return categories
}

func insertCategoryTestData(category model.Category) int {
	dbModel := CategoryRepositoryImpl{}.domainModelToDBModel(category)
	if err := db.DB.Create(&dbModel).Error; err != nil {
		panic(err)
	}

	return dbModel.ID
}

func truncateCategoryTable() {
	// Disable foreign key constraint check
	if err := db.DB.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}

	// Truncate the table
	if err := db.DB.Exec("TRUNCATE TABLE category").Error; err != nil {
		panic(err)
	}

	// Re-enable foreign key constraint check
	if err := db.DB.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
		panic(err)
	}
}

func TestCategoryRepositoryImp_Create(t *testing.T) {
	category := model.NewCategory(1, 1, "category1")
	type args struct {
		category model.Category
	}
	tests := []struct {
		name   string
		args   args
		expect *model.Category
	}{
		{
			name: "success",
			args: args{
				category: category,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := newCategoryRepositoryImpl()
			_, err := repo.CreateCategory(context.Background(), test.args.category)
			assert.Nil(t, err)
			stored := getAllCategories()

			if diff := cmp.Diff(repo.dbModelToDomainModel(stored[0]), test.args.category, testhelper.ToDoCmpOptions...); diff != "" {
				t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
			}
		})
	}
}

func TestCategoryRepositoryImp_Get(t *testing.T) {
	truncateCategoryTable()

	categories := []model.Category{
		model.NewCategory(1, 1, "category1"),
		model.NewCategory(2, 1, "category2"),
		model.NewCategory(3, 1, "category3"),
	}

	for _, category := range categories {
		insertCategoryTestData(category)
	}

	tests := []struct {
		name     string
		expected []*model.Category
	}{
		{
			name: "success",
			expected: []*model.Category{
				&categories[0],
				&categories[1],
				&categories[2],
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := newCategoryRepositoryImpl()
			actual, err := repo.GetAllCategories(context.Background())
			assert.Nil(t, err)
			if diff := cmp.Diff(actual, test.expected, testhelper.ToDoCmpOptions...); diff != "" {
				t.Errorf("Compare value is mismatch (-v1 +v2):%s\n", diff)
			}
		})
	}
}
