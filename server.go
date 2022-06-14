package wow

import (
	"errors"
	"github.com/denisskin/word-of-wisdom/netstat"
	"log"
	"net"
	"time"

	"github.com/denisskin/word-of-wisdom/pow"
)

// Server is "Word of Wisdom" tcp-server
type Server struct {
	db DB //

	// moving average of request count
	avgReq *netstat.MovingAverage
}

// anti-DDOS params and limits
const (
	// allowed number of requests per second
	limitRequests = 100

	// minimum number of hashes per request
	minDifficulty = 20e3
)

// StartServer creates and start "Word of Wisdom" server for givens tcp-address
func StartServer(addr string) {
	NewServer().Listen(addr)
}

// NewServer makes new "Word of Wisdom" server
func NewServer() *Server {
	return &Server{
		db: newDB(),

		avgReq: netstat.NewMovingAverage(time.Second, 15),
	}
}

// Listen announces on the local network address.
func (s *Server) Listen(addr string) {
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
	avgReq := s.avgReq.Add(1)

	// calculate difficulty as function of actual average number of requests per second
	// 		difficulty := ƒ(Avg-Requests)
	difficulty := uint64(avgReq / limitRequests * minDifficulty)
	if difficulty < minDifficulty {
		difficulty = minDifficulty
	}
	log.Printf("wow.Server> new request from [%s]. (avg-requests: %.1f/sec; difficulty: %d)", addr, avgReq, difficulty)

	//-- 1. request service
	_, err = readBytes(conn) // 'GET' - first hello-message
	if err != nil {
		return
	}

	//-- 2. send challenge
	token := pow.NewToken(difficulty)
	if err = writeBytes(conn, token); err != nil {
		return
	}

	//-- 3. read proof
	proof, err := readBytes(conn)
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
	return writeBytes(conn, resp)
}
