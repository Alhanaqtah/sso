package app

import grpcapp "sso/internal/app/grpc"

type App struct {
	GRPCSrv *grpcapp.App
}
