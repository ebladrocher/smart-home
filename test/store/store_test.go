package store_test

import (
	"github.com/ebladrocher/smarthome/system/store/postgrestore"
	"os"
	"testing"
)

var (
	dataBaseURL string
)

func TestMain(m *testing.M)  {

	cfg := postgrestore.ConfigDB{
		Host:"localhost",
		Port:"5432",
		Name: "smarthome_test",
	}

	if dataBaseURL == "" {
		dataBaseURL = cfg.ConnectionSting()
	}
	os.Exit(m.Run())

}