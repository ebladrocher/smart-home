package store

import (
	"fmt"
	"github.com/ebladrocher/smarthome/system/config"
	"strings"
	"testing"
)

func TestStore(t *testing.T) (*Store, func( ... string)) {
	t.Helper()
	//conf, _ :=config.ReadConfig()
	cfg := config.AppConfig{
		//ServerHost:conf.ServerHost,
		//ServerPort:conf.ServerPort,
		//Mode:conf.Mode,
		DbHost:"localhost",
		DbPort:"5432",
		DbName: "smarthome_test",
	}
	thisConfig := NewDbConfig(&cfg)
	s := InitStore(thisConfig)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ... string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
