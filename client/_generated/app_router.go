package main

import (
	"context"

	pb "appRouter"
)

type appRouterImpl struct {
	pb.UnimplementedAppRouter
}

func (s *appRouterImpl) ById(
	ctx context.Context,
	req *pb.ByIdRequest,
) (*pb.ByIdResponse, error) {
	rsp := &pb.ByIdResponse{}
	return rsp, nil
}
