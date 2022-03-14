package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/go-userMgmt-grpc/userMgmt"

	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	users_list := make(map[string]int32)
	var name string
	var age int32
	log.Println("Enter user name:")
	fmt.Scan(&name)
	log.Println("Enter user age:")
	fmt.Scan(&age)

	// users_list["Annie"] = 24
	// users_list["Bob"] = 38
	users_list[name] = int32(age)

	for name, age := range users_list {
		res, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})

		if err != nil {
			log.Fatalf("Failed to save new user: %v", err)
		}

		log.Printf("Created User Details: \nName: %v\nAge:%v\nId:%v", res.GetName(), res.GetAge(), res.GetId())
	}

	list, err := client.GetUsers(ctx, &pb.GetUsersParams{})

	if err != nil {
		log.Fatalf("Failed to get users list: %v", err)
	}

	fmt.Printf("\nUsers List: \n")
	for _, user := range list.GetUsers() {
		fmt.Println(user)
	}

	var userId int32
	log.Println("Enter user Id to get/update:")
	fmt.Scan(&userId)
	user, err := client.GetUserById(ctx, &pb.GetUserParams{UserId: userId})

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	log.Printf("Get User data for Id: %v\n%v\n", userId, user)

	updtUser := &pb.User{Id: userId, Name: "NEW name", Age: 16}
	
	res, err := client.UpdateUser(ctx, &pb.UpdateUserReq{User: updtUser})
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	log.Printf("Updated user Id: %v Rows updated: %v\n", userId, res.GetRowsAffected())

	user, err = client.GetUserById(ctx, &pb.GetUserParams{UserId: userId})

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	log.Printf("Get updated User data for Id: %v\n%v\n", userId, user)

	log.Println("Enter Id of user to delete:")
	fmt.Scan(&userId)
	res, err = client.DeleteUser(ctx, &pb.GetUserParams{UserId: userId})

	if err != nil {
		log.Fatalf("Delete user failed: %v", err)
	}

	log.Printf("Delete user Id: %v Rows Deleted: %v\n", userId, res.GetRowsAffected())
}
