package mssql_test

import (
	"testing"

	. "github.com/obieq/goar"
	. "github.com/obieq/goar/db/mssql"
	. "github.com/obieq/goar/tests/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type MsSqlAutomobile struct {
	ArMsSql
	Automobile
	SafetyRating int
}

func (m *MsSqlAutomobile) Validate() {
	m.Validation.Required("Year", m.Year)
	m.Validation.Required("Make", m.Make)
	m.Validation.Required("Model", m.Model)
}

func (model MsSqlAutomobile) ToActiveRecord() *MsSqlAutomobile {
	return ToAR(&model).(*MsSqlAutomobile)
}

func (dbModel MsSqlAutomobile) AssertDbPropertyMappings(model MsSqlAutomobile, isDbUpdate bool) {
	Ω(dbModel.ID).Should(Equal(model.ID))
	Ω(dbModel.Year).Should(Equal(model.Year))
	Ω(dbModel.Make).Should(Equal(model.Make))
	Ω(dbModel.Model).Should(Equal(model.Model))
	Ω(dbModel.SafetyRating).Should(Equal(model.SafetyRating))

	Ω(dbModel.CreatedAt).ShouldNot(BeNil())
	if isDbUpdate {
		Ω(dbModel.UpdatedAt).ShouldNot(BeNil())
	} else {
		Ω(dbModel.UpdatedAt).Should(BeNil())
	}
}

func TestMsSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MsSql Suite")
}

var _ = BeforeSuite(func() {
	client := Client()

	// clean up previous test data
	client.DropTable(&MsSqlAutomobile{})
	client.CreateTable(&MsSqlAutomobile{})
})

var _ = AfterSuite(func() {
})
