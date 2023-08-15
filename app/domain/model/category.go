package model

type Category struct {
	id        int
	companyID int
	name      string
}

const LrmCompanyID = 1

func (c *Category) ID() int {
	return c.id
}

func (c *Category) CompanyID() int {
	return c.companyID
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) IsLRMCategory() bool {
	return c.companyID == LrmCompanyID
}

func NewCategory(
	id int,
	companyID int,
	name string,
) Category {
	return Category{
		id:        UndefinedID,
		companyID: companyID,
		name:      name,
	}
}

func RestoreCategory(
	id int,
	companyID int,
	name string,
) Category {
	return Category{
		id:        id,
		companyID: companyID,
		name:      name,
	}
}
