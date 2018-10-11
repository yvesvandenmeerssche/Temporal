package tns_test

import (
	"testing"

	"github.com/RTradeLtd/Temporal/tns"
)

// Issue with libp2p and being unable to run multiple tests one after another
// need to debug to figure out how we can avoid this

const (
	testZoneName  = "example.org"
	testIPAddress = "0.0.0.0"
	testPort      = "9999"
	testIPVersion = "ip4"
	testProtocol  = "tcp"
	testCfgPath   = "../test/config.json"
)

func TestTNS_Echo(t *testing.T) {
	manager, err := tns.GenerateTNSManager(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = manager.MakeHost(manager.PrivateKey, nil); err != nil {
		t.Fatal(err)
	}
	go manager.RunTNSDaemon()
	client, err := tns.GenerateTNSClient(true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = client.MakeHost(client.PrivateKey, nil); err != nil {
		t.Fatal(err)
	}
	addr, err := manager.ReachableAddress(0)
	if err != nil {
		t.Fatal(err)
	}
	pid, err := client.AddPeerToPeerStore(addr)
	if err != nil {
		t.Fatal(err)
	}
	if err = client.QueryTNS(pid, "echo"); err != nil {
		t.Fatal(err)
	}
}
func TestTNS_GenerateTNSClient(t *testing.T) {
	t.Skip()
	if _, err := tns.GenerateTNSClient(true, nil); err != nil {
		t.Fatal(err)
	}
}
func TestTNS_ClientMakeHost(t *testing.T) {
	t.Skip()
	c, err := tns.GenerateTNSClient(true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = c.MakeHost(c.PrivateKey, nil); err != nil {
		t.Fatal(err)
	}
}

func TestTNS_GenerateTNSManager(t *testing.T) {
	t.Skip()
	if _, err := tns.GenerateTNSManager(nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestTNS_ManagerMakeHost(t *testing.T) {
	t.Skip()
	m, err := tns.GenerateTNSManager(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = m.MakeHost(m.PrivateKey, nil); err != nil {
		t.Fatal(err)
	}
}
