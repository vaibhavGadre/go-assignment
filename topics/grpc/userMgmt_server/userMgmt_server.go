package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "example.com/go-userMgmt-grpc/userMgmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

const PORT = ":50051"

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	*pgx.Conn
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

func (server *UserManagementServer) Run() error {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		//register
		err := pb.RegisterUserManagementHandlerServer(context.Background(), mux, server)

		if err != nil {
			log.Fatalf("http listen err %v", err)
		}

		//listen
		// Start HTTP server (and proxy calls to gRPC server endpoint)
		http.ListenAndServe("localhost:3000", mux)
	}()

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)
	// err = pb.RegisterUserManagementHandlerFromEndpoint(context.Background(), mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
	log.Printf("Server is listening on port: %v", lis.Addr())

	return s.Serve(lis)
}

func (server *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received name: %v", in.GetName())
	// var user_list *pb.UserList = &pb.UserList{}
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge()}
	createSql := `
	CREATE TABLE IF NOT EXISTS USERS (
		id SERIAL PRIMARY KEY,
		name text,
		age int
	)
	`
	_, err := server.Conn.Exec(context.Background(), createSql)

	if err != nil {
		log.Fatalf("Error during users table creation: %v", err)
	}

	tx, err := server.Conn.Begin(context.Background())

	if err != nil {
		log.Fatalf("transaction begin error: %v", err)
	}

	_, err = tx.Exec(context.Background(), "insert into users(name, age) values ($1, $2)", created_user.Name, created_user.Age)

	if err != nil {
		log.Fatalf("tx exec error: %v", err)
	}

	tx.Commit(context.Background())
	return created_user, nil
}

func (server *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {

	var user_list *pb.UserList = &pb.UserList{}
	rows, err := server.Conn.Query(context.Background(), "select * from users")
	if err != nil {
		log.Fatalf("get users query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := pb.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Fatalf("query row scan error: %v", err)
		}
		user_list.Users = append(user_list.Users, &user)
	}
	return user_list, err
}

func (server *UserManagementServer) GetUserById(ctx context.Context, in *pb.GetUserParams) (*pb.User, error) {
	user := pb.User{}
	err := server.Conn.QueryRow(context.Background(), "select * from users where id=$1", in.GetUserId()).Scan(&user.Id, &user.Name, &user.Age)

	if err != nil {
		log.Fatalf("Get user for id: %v failed with err: %v", in.GetUserId(), err)
	}

	return &user, nil
}

func (server *UserManagementServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	var res = pb.UpdateUserRes{}
	x, err := server.Conn.Exec(context.Background(), "update users set name=$1, age=$2 where id=$3;", in.GetUser().GetName(), in.GetUser().GetAge(), in.GetUser().GetId())
	if err != nil {
		log.Fatalf("Update user err: %v", err)
	}
	res.RowsAffected = x.RowsAffected()
	return &res, nil
}

func (server *UserManagementServer) DeleteUser(ctx context.Context, in *pb.GetUserParams) (*pb.UpdateUserRes, error) {
	var res = pb.UpdateUserRes{}
	x, err := server.Conn.Exec(context.Background(), "Delete from users where id=$1;", in.GetUserId())
	if err != nil {
		log.Fatalf("Delete user err: %v", err)
	}
	res.RowsAffected = x.RowsAffected()
	return &res, nil
}

func main() {
	var user_mgmt_server *UserManagementServer = NewUserManagementServer()

	connString := "postgresql://postgres:password@localhost:5432/test_db"
	conn, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		log.Fatalf("Unable to connect to db: %v", err)
	}

	defer conn.Close(context.Background())

	user_mgmt_server.Conn = conn
	if err := user_mgmt_server.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
