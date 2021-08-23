package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/ipfs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestInitChannelConfig(t *testing.T) {

	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err1 := viper.ReadInConfig(); err1 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	nd :=  cockroach.NewDb()
	if err2 := nd.InitDB(); err2 != nil {
		return
	}

	// If a configs file is found, read it in.
	if err3 := viper.ReadInConfig(); err3 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	ipfs.InitIPFS()

	cache.InitRedis(1)

}

func TestNewChannels(t *testing.T) {
	nc := NewChannels("ALICE HOUSE", "", "avatar", "bio", "alice", true)
	fmt.Println(nc.Id)


	nc1 := NewChannels("Hvxahv Chan", "", "avatar", "bio", "hvturingga",false)
	fmt.Println(nc1)
}

func TestNewChannels2(t *testing.T) {
	TestInitChannelConfig(t)
	nc := NewChannels("ALICE HOUSE", "", "avatar", "bio", "alice", true)
	code, s, id, err := nc.New()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("code: %v; message: %s; id: %s", code, s, id)

	nc2 := NewChannels("JSUT 4 FUN", "", "avatar", "bio", "hvturingga", false)
	i2, s2, _, _ := nc2.New()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("code: %v; message: %s", i2, s2)

}


func TestNewFindChannelByID(t *testing.T) {
	TestInitChannelConfig(t)
	nfc := NewChannelsByID("f6574uSSqGQ7CJX")
	ch := nfc.Find()
	fmt.Println(ch.Name)
}
