package mockrouting

import (
	"context"
	"time"

	"gx/ipfs/QmPuhRE325DR8ChNcFtgd6F1eANCHy1oohXZPpYop4xsK6/go-testutil"
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	routing "gx/ipfs/QmRASJXJUFygM5qU4YrH7k7jD6S4Hg8nJmgqJ4bYJvLatd/go-libp2p-routing"
	ropts "gx/ipfs/QmRASJXJUFygM5qU4YrH7k7jD6S4Hg8nJmgqJ4bYJvLatd/go-libp2p-routing/options"
	ma "gx/ipfs/QmRKLtwMw131aK7ugC3G7ybpumMz78YrJe5dzneyindvG1/go-multiaddr"
	peer "gx/ipfs/QmY5Grm8pJdiSSVsYxx4uNRgweY72EmYwuSDbRnbFok3iY/go-libp2p-peer"
	pstore "gx/ipfs/QmZ9zH2FnLcxv1xyzFeUpDUeo55xEhZQHgveZijcxr7TLj/go-libp2p-peerstore"
	logging "gx/ipfs/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"
)

var log = logging.Logger("mockrouter")

type client struct {
	vs     routing.ValueStore
	server server
	peer   testutil.Identity
}

// FIXME(brian): is this method meant to simulate putting a value into the network?
func (c *client) PutValue(ctx context.Context, key string, val []byte, opts ...ropts.Option) error {
	log.Debugf("PutValue: %s", key)
	return c.vs.PutValue(ctx, key, val, opts...)
}

// FIXME(brian): is this method meant to simulate getting a value from the network?
func (c *client) GetValue(ctx context.Context, key string, opts ...ropts.Option) ([]byte, error) {
	log.Debugf("GetValue: %s", key)
	return c.vs.GetValue(ctx, key, opts...)
}

func (c *client) SearchValue(ctx context.Context, key string, opts ...ropts.Option) (<-chan []byte, error) {
	log.Debugf("SearchValue: %s", key)
	return c.vs.SearchValue(ctx, key, opts...)
}

func (c *client) FindProviders(ctx context.Context, key cid.Cid) ([]pstore.PeerInfo, error) {
	return c.server.Providers(key), nil
}

func (c *client) FindPeer(ctx context.Context, pid peer.ID) (pstore.PeerInfo, error) {
	log.Debugf("FindPeer: %s", pid)
	return pstore.PeerInfo{}, nil
}

func (c *client) FindProvidersAsync(ctx context.Context, k cid.Cid, max int) <-chan pstore.PeerInfo {
	out := make(chan pstore.PeerInfo)
	go func() {
		defer close(out)
		for i, p := range c.server.Providers(k) {
			if max <= i {
				return
			}
			select {
			case out <- p:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Provide returns once the message is on the network. Value is not necessarily
// visible yet.
func (c *client) Provide(_ context.Context, key cid.Cid, brd bool) error {
	if !brd {
		return nil
	}
	info := pstore.PeerInfo{
		ID:    c.peer.ID(),
		Addrs: []ma.Multiaddr{c.peer.Address()},
	}
	return c.server.Announce(info, key)
}

func (c *client) Ping(ctx context.Context, p peer.ID) (time.Duration, error) {
	return 0, nil
}

func (c *client) Bootstrap(context.Context) error {
	return nil
}

var _ routing.IpfsRouting = &client{}