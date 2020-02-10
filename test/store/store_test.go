package store_test

import (
	"github.com/ebladrocher/smarthome/system/config"
	"github.com/ebladrocher/smarthome/system/store/db"
	"os"
	"testing"
)

var (
	dataBaseURL string
)

func TestMain(m *testing.M)  {

	cfg := config.AppConfig{
		DbHost:"localhost",
		DbPort:"5432",
		DbName: "smarthome_test",
	}
	thisConfig := db.NewDbConfig(&cfg)

	if dataBaseURL == "" {
		dataBaseURL = thisConfig.ConnectionSting()
	}
	os.Exit(m.Run())

}