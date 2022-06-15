package wow_client

import (
	"errors"
	"github.com/denisskin/word-of-wisdom/common/netutils"
	"github.com/denisskin/word-of-wisdom/common/pow"
	"net"
)

type Client struct {
	addr string
}

// New returns new "Word of Wisdom" client
func New(addr string) *Client {
	return &Client{addr}
}

// Get requests new "Word of Wisdom"
func (c *Client) Get() (_ string, err error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return
	}
	defer conn.Close()

	//--- 1. send request
	if err = netutils.WriteBytes(conn, []byte("GET")); err != nil {
		return
	}

	//--- 2. PoW-challenge. read token
	token, err := netutils.ReadBytes(conn)
	if err != nil {
		return
	}

	//--- 3. solve PoW-challenge
	solution := pow.Solve(token)

	//--- 4. send PoW-solution
	if err = netutils.WriteBytes(conn, solution); err != nil {
		return
	}

	//--- 5. read final response
	resp, err := netutils.ReadBytes(conn)
	if err == nil && len(resp) == 0 {
		err = errors.New("wow.Client: empty response")
	}
	if err != nil {
		return
	}
	return string(resp), err
}
