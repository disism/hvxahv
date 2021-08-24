package cockroach

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}


	n := NewDBAddr()
	if err := n.InitDB(); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}

func TestCreateDB(t *testing.T) {

	n := NewDBAddr()
	name := "hvxahv"
	if err := n.New(name); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}
