package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	cards "cards/proto"
)

type Cards struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Cards) Call(ctx context.Context, req *cards.Request, rsp *cards.Response) error {
	log.Info("Received Cards.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Cards) Stream(ctx context.Context, req *cards.StreamingRequest, stream cards.Cards_StreamStream) error {
	log.Infof("Received Cards.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&cards.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Cards) PingPong(ctx context.Context, stream cards.Cards_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&cards.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
