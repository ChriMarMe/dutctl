package main

import (
	"github.com/BlindspotSoftware/dutctl/pkg/dut"
	"github.com/BlindspotSoftware/dutctl/pkg/module/dummy"
)

type devlist map[string]dut.Device

func (devs devlist) names() []string {
	names := make([]string, 0, len(devs))
	for d := range devs {
		names = append(names, d)
	}

	return names
}

func (devs devlist) cmds(device string) []string {
	dev, ok := devs[device]
	if !ok {
		return []string{}
	}

	cmds := make([]string, 0, len(dev.Cmds))
	for c := range dev.Cmds {
		cmds = append(cmds, c)
	}

	return cmds
}

// This is teporary test data

//nolint:gochecknoglobals
var testlist = devlist{
	"device1": {
		Desc: "Device 1",
		Cmds: map[string]dut.Command{
			"status": {
				Desc: "Report status",
				Modules: []dut.Module{
					{
						Name:    "dummy-status",
						Main:    true,
						Args:    nil,
						Options: nil,
						Module:  &dummy.Status{},
					},
				},
			},
		},
	},
	"device2": {
		Desc: "Device 2",
		Cmds: map[string]dut.Command{
			"status": {
				Desc: "Report status",
				Modules: []dut.Module{
					{
						Name:    "dummy-status",
						Main:    true,
						Args:    nil,
						Options: nil,
						Module:  &dummy.Status{},
					},
				},
			},
			"repeat": {
				Desc: "Repeat input",
				Modules: []dut.Module{
					{
						Name:    "dummy-repeat",
						Main:    true,
						Args:    nil,
						Options: nil,
						Module:  &dummy.Repeat{},
					},
				},
			},
		},
	},
	"device3": {
		Desc: "Device 3",
		Cmds: map[string]dut.Command{
			"status": {
				Desc: "Report status",
				Modules: []dut.Module{
					{
						Name:    "dummy-status",
						Main:    true,
						Args:    nil,
						Options: nil,
						Module:  &dummy.Status{},
					},
				},
			},
			"file-transfer": {
				Desc: "Transfer a file",
				Modules: []dut.Module{
					{
						Name:    "dummy-file-transfer",
						Main:    true,
						Args:    nil,
						Options: nil,
						Module:  &dummy.FT{},
					},
				},
			},
		},
	},
}