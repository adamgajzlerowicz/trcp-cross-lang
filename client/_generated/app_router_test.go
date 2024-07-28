package main

import (
	"context"
	"reflect"
	"testing"

	pb "appRouter"

	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/appRouter/app_mock.go -package=appRouter -self_package=appRouter --source=stub/appRouter/app.trpc.go

func Test_appRouterImpl_ById(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	appRouterService := pb.NewMockAppRouterService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := appRouterService.EXPECT().ById(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.ByIdRequest) (*pb.ByIdResponse, error) {
		s := &appRouterImpl{}
		return s.ById(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.ByIdRequest
		rsp *pb.ByIdResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.ByIdResponse
			var err error
			if rsp, err = appRouterService.ById(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("appRouterImpl.ById() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("appRouterImpl.ById() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
