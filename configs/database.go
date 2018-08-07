package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//get postgres gorm extension
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitializeDatabase - Instantiate and check connection to Postgre database
func InitializeDatabase(c *Config) (*gorm.DB, error) {
	var err error
	con, err := gorm.Open("mysql", c.GetDatasource())

	if err != nil {
		fmt.Println("[ERROR] Failed to connect to MySQL. Config= " + Configuration.MySQL.Host)
		return nil, err
	}

	con.LogMode(LogMode)
	con.SingularTable(true)

	fmt.Println("[INFO] Connected to MySQL. Config => " + Configuration.MySQL.Host + ", LogMode => " + fmt.Sprintf("%v", LogMode))
	return con, nil
}

// GetDatasource - return datasource name database in used
func (c *Config) GetDatasource() string {
	// Mysql Host Address
	var datasource = c.MySQL.Username + ":" + c.MySQL.Password + "@tcp(" + c.MySQL.Host + ":" + c.MySQL.Port + ")/" + c.MySQL.Db + "?charset=utf8&parseTime=True&loc=Local"
	return datasource
}
