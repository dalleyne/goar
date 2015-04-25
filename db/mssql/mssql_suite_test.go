package mssql

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMsSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MsSql Suite")
}

//var Migration *MsSqlMigration = &MsSqlMigration{}
//var migrationDbName string = "migration_test_db"

var _ = BeforeSuite(func() {
	Client()
	//dbName := "goar_test"
	//SetDbName(dbName)

	//// drop databases from prior test(s)
	//err := Migration.DropDb(DbName())
	//Migration.DropDb(migrationDbName)

	//// prep for current test(s)
	//err = Migration.CreateDb(DbName())
	//Expect(err).NotTo(HaveOccurred())

	//err = Migration.CreateTable("rethink_db_automobiles")
	//Expect(err).NotTo(HaveOccurred())

	//err = Migration.CreateTable("callback_error_models")
	//Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	//Session().Close()
})
