package channel

import (
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func init() {
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

	// Initialize the database.
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}
}

func TestChannel_CreateChannel(t *testing.T) {
	s := &channel{}
	createChannel, err := s.CreateChannel(context.Background(), &pb.CreateChannelRequest{
		PreferredUsername: "gateway",
		AccountId:         "746931987134185473",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(createChannel.Reply)
}

func TestChannel_GetChannelsByAccountID(t *testing.T) {
	s := &channel{}
	getChannelsByAccountID, err := s.GetChannelsByAccountID(context.Background(), &pb.GetChannelsByAccountIDRequest{
		AccountId: "746931987134185473",
	})
	if err != nil {
		return
	}
	fmt.Println(getChannelsByAccountID.Channels)
}

func TestChannel_DeleteChannel(t *testing.T) {
	s := &channel{}
	deleteChannel, err := s.DeleteChannel(context.Background(), &pb.DeleteChannelRequest{
		AccountId: "746931987134185473",
		ChannelId: "747178643349929985",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(deleteChannel.Reply)
}

func TestChannel_DeleteAllChannelsByAccountID(t *testing.T) {
	s := &channel{}
	deleteAllChannelsByAccountID, err := s.DeleteAllChannelsByAccountID(context.Background(), &pb.DeleteAllChannelsByAccountIDRequest{
		AccountId: "746588397237010433",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(deleteAllChannelsByAccountID.Reply)
}
