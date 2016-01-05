package xenAPI

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"os"
	"testing"
)

func TestAuthentication(t *testing.T) {
	client := NewClient("http://localhost:40080")

	sessionRef, err := client.Session().LoginWithPassword("terraform", "testing", "1.0", "terraform")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	err = client.Session().Logout(sessionRef)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}
