package wow_server

import (
	"errors"
	"fmt"
	"github.com/denisskin/word-of-wisdom/common/netutils"
	"github.com/denisskin/word-of-wisdom/common/pow"
	"log"
	"net"
	"time"
)

// Server is "Word of Wisdom" tcp-server
type Server struct {
	db DB //

	// anti-DDoS options
	difficulty uint64 // minimal PoW Difficulty. (Number of hashes per request)
	reqLimit   uint64 // income Requests Limit. (Allowed number of requests per second)

	ma *netutils.MovingAverage // moving average of request count
}

// Start creates and start "Word of Wisdom" server for givens tcp-address
func Start(tcpPort uint, powDifficulty, requestsLimit uint64) {
	addr := fmt.Sprintf(":%d", tcpPort)
	New(powDifficulty, requestsLimit).Listen(addr)
}

// New makes new "Word of Wisdom" server
func New(difficulty, requestsLimit uint64) *Server {
	if difficulty <= 0 {
		difficulty = 10e3
	}
	return &Server{
		db:         newDB(),
		difficulty: difficulty,
		reqLimit:   requestsLimit,
		ma:         netutils.NewMovingAverage(time.Second, 15),
	}
}

// Listen announces on the local network address.
func (s *Server) Listen(addr string) {

	log.Printf("start server (PoW-difficulty:%d hash/sec; limit:%d reqs/sec).", s.difficulty, s.reqLimit)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("listening connection %s ...", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept-error: %v", err)
			continue
		}
		go func() {
			if err := s.handle(conn); err != nil {
				log.Printf("wow.Server> client[%s].handle-error: %v", conn.RemoteAddr(), err)
			}
		}()
	}
}

func (s *Server) handle(conn net.Conn) (err error) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	avgReq := s.ma.Add(1)

	// calculate difficulty as function of actual average number of requests per second
	// 		difficulty := ƒ(requestsPerSec)
	difficulty := uint64(avgReq / float64(s.reqLimit) * float64(s.difficulty))
	if difficulty < s.difficulty {
		difficulty = s.difficulty
	}
	log.Printf("wow.Server> new request from [%s]. (avg-requests: %.1f/sec; difficulty: %d)", addr, avgReq, difficulty)

	//-- 1. request service
	_, err = netutils.ReadBytes(conn) // 'GET' - first hello-message
	if err != nil {
		return
	}

	//-- 2. send challenge
	token := pow.NewToken(difficulty)
	if err = netutils.WriteBytes(conn, token); err != nil {
		return
	}

	//-- 3. read proof
	proof, err := netutils.ReadBytes(conn)
	if err != nil {
		return
	}

	//-- 4. verify proof
	if !pow.Verify(token, proof) {
		err = errors.New("invalid PoW-proof")
		return
	}

	//--- 5. send final response
	resp := s.db.randomWisdom()
	return netutils.WriteBytes(conn, resp)
}
