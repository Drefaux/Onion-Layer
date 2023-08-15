package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory_IsLRM(t *testing.T) {
	type fields struct {
		companyID int
	}

	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			name: "is LRM category",
			fields: fields{
				companyID: 1,
			},
			expected: true,
		},
		{
			name: "is not LRM category",
			fields: fields{
				companyID: 0,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := &Category{
				companyID: test.fields.companyID,
			}
			actual := e.IsLRMCategory()

			assert.Equal(t, test.expected, actual, "Expected IsLRMCategory() result to be %v, but got %v", test.expected, actual)
		})
	}
}
