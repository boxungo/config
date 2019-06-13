package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ghodss/yaml"
)

func TestConfigParsingFlags(t *testing.T) {
	args := []string{
		"-name=testname",
	}

	cfg := NewConfig()
	err := cfg.parse(args)
	if err != nil {
		t.Fatal(err)
	}

	validateFlags(t, cfg)
}

func TestConfigParseFromFile(t *testing.T) {
	c := struct {
		Name string `json:"name"`
	}{
		"testname",
	}
	b, err := yaml.Marshal(&c)
	if err != nil {
		t.Fatal(err)
	}

	tmpFile := createTempConfigFile(t, b)
	defer os.Remove(tmpFile.Name())

	args := []string{fmt.Sprintf("--config-file=%s", tmpFile.Name())}

	cfg := NewConfig()
	err = cfg.parse(args)
	if err != nil {
		t.Fatal(err)
	}

	err = cfg.configFromFile(cfg.configFile)
	if err != nil {
		t.Fatal(err)
	}

	validateFlags(t, cfg)
}

func validateFlags(t *testing.T, cfg *config) {
	if cfg.Name != "testname" {
		t.Errorf("Name = %v, want %v", cfg.Name, "testname")
	}
}

func createTempConfigFile(t *testing.T, b []byte) *os.File {
	tmpfile, err := ioutil.TempFile("", "config")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tmpfile.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	err = tmpfile.Close()
	if err != nil {
		t.Fatal(err)
	}

	return tmpfile
}
