package main

import (
	"context"
	// "log"
	// "net"
	"reflect"
	"testing"

	// "google.golang.org/grpc"
	// "google.golang.org/grpc/test/bufconn"

	pb "example.com/go-userMgmt-grpc/userMgmt"
	"github.com/jackc/pgx/v4"
)

// func dialer() func(context.Context, string) (net.Conn, error) {
// 	listener := bufconn.Listen(1024 * 1024)

// 	server := grpc.NewServer()

// 	pb.RegisterUserManagementServer(server, &UserManagementServer{})

// 	go func() {
// 		if err := server.Serve(listener); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	return func(context.Context, string) (net.Conn, error) {
// 		return listener.Dial()
// 	}
// }

func TestUserManagementServer_CreateNewUser(t *testing.T) {
	connString := "postgresql://postgres:password@localhost:5432/test_db"
	conn, err := pgx.Connect(context.Background(), connString)
	type fields struct {
		UnimplementedUserManagementServer pb.UnimplementedUserManagementServer
		Conn                              *pgx.Conn
	}
	type args struct {
		ctx context.Context
		in  *pb.NewUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		// insert cases here
		{
			name:    "Create User test",
			fields:  fields{
				UnimplementedUserManagementServer: pb.UnimplementedUserManagementServer{},
				Conn:                              conn,
			},
			args:    args{},
			want:    &pb.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &UserManagementServer{
				UnimplementedUserManagementServer: tt.fields.UnimplementedUserManagementServer,
				Conn:                              tt.fields.Conn,
			}
			got, err := server.CreateNewUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserManagementServer.CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserManagementServer.CreateNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
