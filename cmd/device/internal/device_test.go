/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"testing"

	"github.com/hvxahv/hvx/cockroach"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	if err := cockroach.NewRoach().Dial(); err != nil {
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

func TestDevices_Create(t *testing.T) {
	device, err := NewDevices(784092546599714817, "chrome").Create()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(device)
}

//
//func TestDevices_IsExistByHash(t *testing.T) {
//	d := NewDevicesHash("21ad23a6-1f20-48c1-abcb-9f2e03f6617f")
//	exist, err := d.IsExistByHash()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(exist)
//}
//
//func TestDevices_GetDeviceByHash(t *testing.T) {
//	d := NewDevicesHash("b2c9374c-ee29-40b9-9741-63c910f7eb1a")
//	device, err := d.GetByHash()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(device)
//}
//
//func TestDevices_IsExistById(t *testing.T) {
//	d := NewDevicesId(760036215582294017)
//	exist, err := d.IsExistById()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(exist)
//}
//
//func TestDevices_GetDeviceById(t *testing.T) {
//	d := NewDevicesId(763913016173002753)
//	device, err := d.GetById()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(device)
//}
//
//func TestDevices_CreateDevice(t *testing.T) {
//	d := NewDevices(763903590118359041, "chrome")
//	create, err := d.Create()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(create)
//}
//
//func TestDevices_GetDeviceByAccountId(t *testing.T) {
//	d := NewDevicesAccountID(763903590118359041)
//	devices, err := d.GetListByAccountId()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(devices)
//}
//
//func TestDevices_Delete(t *testing.T) {
//	d := NewDevicesDelete(763913016173002753, 763903590118359041)
//	if err := d.Delete(); err != nil {
//		t.Error(err)
//		return
//	}
//}
//
//func TestDevices_DeleteAccountDevices(t *testing.T) {
//	if err := NewDevicesAccountID(763903590118359041).DeleteAccountDevices(); err != nil {
//		t.Error(err)
//		return
//	}
//}
