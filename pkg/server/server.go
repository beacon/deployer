package server

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"github.com/beacon/deployer/pkg/config"
	pb "github.com/beacon/deployer/pkg/proto"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/proto.proto

// Server for grpc
type Server struct {
	srv    *http.Server
	rpcSrv *grpc.Server

	restful *gin.Engine
}

func (s *Server) UpdateDeployStatus(ctx context.Context, status *pb.DeployStatus) (*pb.Reply, error) {
	return &pb.Reply{
		Code:    200,
		Message: "OK",
	}, nil
}

func New(cfg *config.Config) *Server {
	rpcSrv := grpc.NewServer()
	s := &Server{
		rpcSrv:  rpcSrv,
		restful: gin.New(),
	}
	if cfg.TLS == nil {
		h2Srv := &http2.Server{}

		s.srv = &http.Server{
			Addr:    cfg.Addr,
			Handler: h2c.NewHandler(s, h2Srv),
		}
	} else {
		s.srv = &http.Server{
			Addr:    cfg.Addr,
			Handler: s,
		}
	}

	pb.RegisterServerServer(rpcSrv, s)

	s.routeRestful()
	return s
}

func (s *Server) ListenAndServe(cfg *config.Config) error {
	if cfg.TLS == nil {
		log.Println("NextProto:", s.srv.TLSNextProto)
		log.Println("TLS:", s.srv.TLSConfig)
		return s.srv.ListenAndServe()
	} else {
		return s.srv.ListenAndServeTLS(cfg.TLS.CertFile, cfg.TLS.KeyFile)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request, uri=", r.RequestURI, "method=", r.Method)
	if r.ProtoMajor == 2 && strings.HasPrefix(
		r.Header.Get("Content-Type"), "application/grpc") {
		log.Println("Request handled by grpc")
		s.rpcSrv.ServeHTTP(w, r)
	} else {
		log.Println("Request handled by restful")
		s.restful.ServeHTTP(w, r)
	}
}

func (s *Server) routeRestful() {
	{
		g := s.restful.Group("/actions")
		g.POST("", s.postAction)
	}
}

func (s *Server) postAction(c *gin.Context) {
	c.Status(200)
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("Error shutting down:", err)
	}
}
