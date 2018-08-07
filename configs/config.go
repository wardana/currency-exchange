package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"

	"gopkg.in/yaml.v2"
)

const (
	//FileConfig - path of configuration file
	fileConfig = "config.yaml"
)

var (
	//Configuration - application configuration purpose
	Configuration Config
	//Environtment - environtment variable
	Environtment Env
	//LogMode - logging mode
	LogMode bool
	//Deployment - deployment environment
	Deployment string
	//DefaultPort - Application port
	DefaultPort string
)

// Env - Set all environments that are needed
// Attach interface, so that we can mock it
type Env struct {
	MySQL *gorm.DB
}

// Config - System configuration
type Config struct {
	Port            string `yaml:"port"`
	WhitelistBranch string `yaml:"whitelist_branch"`
	MySQL
}

// MySQL - Configuration of MySQL database
type MySQL struct {
	Db       string `yaml:"db"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogMode  bool   `yaml:"log_mode"`
}

// ReadConfig - Read yaml configuration file
func ReadConfig(c *Config) {
	ymlFile, err := ioutil.ReadFile(fileConfig)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		log.Fatalln(err)
	}
}

// SetupEnvironment - Set environment system, eg: Database, Query
func SetupEnvironment(env *Env) {
	db, err := InitializeDatabase(&Configuration)
	if err != nil {
		log.Panic(err)
	}
	env.MySQL = db
}

// BindingPort - return binding port in string
// eg: :8080
func (c *Config) BindingPort() string {

	var ApplicationPort string

	ApplicationPort = c.Port
	if DefaultPort != "" {
		ApplicationPort = DefaultPort
	}

	return fmt.Sprintf(":%s", ApplicationPort)
}

func init() {
	logMode := flag.Bool("log-mode", false, "Log mode")
	deployment := flag.String("env", "development", "Environment")
	port := flag.String("port", "", "Port number")

	// Parse all command flags
	flag.Parse()

	LogMode = *logMode
	Deployment = *deployment
	DefaultPort = *port
	fmt.Println("------------------------------------")
	fmt.Println("Enable log mode		-", LogMode)
	fmt.Println("Environment		-", Deployment)
	fmt.Println("------------------------------------")
}
