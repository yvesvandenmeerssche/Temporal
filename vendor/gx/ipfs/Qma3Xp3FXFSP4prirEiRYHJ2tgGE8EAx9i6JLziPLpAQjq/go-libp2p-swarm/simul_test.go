package swarm_test

import (
	"context"
	"runtime"
	"sync"
	"testing"
	"time"

	ma "gx/ipfs/QmTZBfrPJmjWsCvHEtX5FE6KimVJhsJg5sBbqEFYf4UZtL/go-multiaddr"
	ci "gx/ipfs/QmWapVoHjtKhn4MhvKNoPTkJKADFGACfXPFnt7combwp5W/go-testutil/ci"
	peer "gx/ipfs/QmYVXrKrKHDC9FobgmcmshCDyWwdrfwfanNQN4oxJ9Fk3h/go-libp2p-peer"
	pstore "gx/ipfs/QmaCTz9RkrU13bm9kMB54f7atgqM4qkjDZpRwRoJiWXEqs/go-libp2p-peerstore"

	. "gx/ipfs/Qma3Xp3FXFSP4prirEiRYHJ2tgGE8EAx9i6JLziPLpAQjq/go-libp2p-swarm"
	swarmt "gx/ipfs/Qma3Xp3FXFSP4prirEiRYHJ2tgGE8EAx9i6JLziPLpAQjq/go-libp2p-swarm/testing"
)

func TestSimultOpen(t *testing.T) {

	t.Parallel()

	ctx := context.Background()
	swarms := makeSwarms(ctx, t, 2, swarmt.OptDisableReuseport)

	// connect everyone
	{
		var wg sync.WaitGroup
		connect := func(s *Swarm, dst peer.ID, addr ma.Multiaddr) {
			defer wg.Done()
			// copy for other peer
			log.Debugf("TestSimultOpen: connecting: %s --> %s (%s)", s.LocalPeer(), dst, addr)
			s.Peerstore().AddAddr(dst, addr, pstore.PermanentAddrTTL)
			if _, err := s.DialPeer(ctx, dst); err != nil {
				t.Error("error swarm dialing to peer", err)
			}
		}

		log.Info("Connecting swarms simultaneously.")
		wg.Add(2)
		go connect(swarms[0], swarms[1].LocalPeer(), swarms[1].ListenAddresses()[0])
		go connect(swarms[1], swarms[0].LocalPeer(), swarms[0].ListenAddresses()[0])
		wg.Wait()
	}

	for _, s := range swarms {
		s.Close()
	}
}

func TestSimultOpenMany(t *testing.T) {
	// t.Skip("very very slow")

	addrs := 20
	rounds := 10
	if ci.IsRunning() || runtime.GOOS == "darwin" {
		// osx has a limit of 256 file descriptors
		addrs = 10
		rounds = 5
	}
	SubtestSwarm(t, addrs, rounds)
}

func TestSimultOpenFewStress(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	// t.Skip("skipping for another test")
	t.Parallel()

	msgs := 40
	swarms := 2
	rounds := 10
	// rounds := 100

	for i := 0; i < rounds; i++ {
		SubtestSwarm(t, swarms, msgs)
		<-time.After(10 * time.Millisecond)
	}
}