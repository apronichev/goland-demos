package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "grpc-demo/proto"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users  map[int32]*pb.GetUserResponse
	nextID int32
}

// gRPC service methods remain the same
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	s.nextID++
	user := &pb.GetUserResponse{
		Id:    s.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[s.nextID] = user
	return &pb.CreateUserResponse{Id: s.nextID}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return &pb.UpdateUserResponse{Success: false}, status.Errorf(codes.NotFound, "user not found")
	}
	user.Name = req.Name
	user.Email = req.Email
	return &pb.UpdateUserResponse{Success: true}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, exists := s.users[req.Id]
	if !exists {
		return &pb.DeleteUserResponse{Success: false}, status.Errorf(codes.NotFound, "user not found")
	}
	delete(s.users, req.Id)
	return &pb.DeleteUserResponse{Success: true}, nil
}

// HTTP handlers
func (s *server) handleHTTP() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/grpc.yaml", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Failed to get working directory: %v", err)
			return
		}

		// Construct the full path to the YAML file
		yamlPath := filepath.Join(cwd, "grpc.yaml")

		// Check if file exists
		if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
			http.Error(w, "YAML specification file not found", http.StatusNotFound)
			log.Printf("YAML file not found at: %s", yamlPath)
			return
		}

		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, yamlPath)
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.handleGetUser(w, r)
		case http.MethodPut:
			s.handleUpdateUser(w, r)
		case http.MethodDelete:
			s.handleDeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			s.handleCreateUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}

func (s *server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := s.users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := s.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (s *server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req pb.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	req.Id = id

	resp, err := s.UpdateUser(r.Context(), &req)
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	resp, err := s.DeleteUser(r.Context(), &pb.DeleteUserRequest{Id: id})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func parseIDFromPath(path string) (int32, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path")
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

func main() {
	// Create server instance
	srv := &server{
		users:  make(map[int32]*pb.GetUserResponse),
		nextID: 0,
	}

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, srv)
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	log.Printf("HTTP server listening at :8080")
	if err := http.ListenAndServe(":8080", srv.handleHTTP()); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
