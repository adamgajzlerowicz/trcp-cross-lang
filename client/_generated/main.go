package main

import (
	pb "appRouter"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterAppRouterService(s.Service("appRouter.AppRouter"), &appRouterImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
