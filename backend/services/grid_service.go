package services

import (
	"context"
	"slices"

	"egm.com/backend/data"
	"egm.com/backend/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GridService struct {
	pb.UnimplementedGridServiceServer
}

func (s *GridService) GetGrid(ctx context.Context, req *emptypb.Empty) (*pb.GetGridResponse, error) {
	response := pb.GetGridResponse{Nodes: data.GridNodes, Links: data.Links}

	return &response, nil
}

func (s *GridService) GetNode(ctx context.Context, req *pb.GetNodeRequest) (*pb.GetNodeResponse, error) {
	id := req.GetId()
	i := slices.IndexFunc(data.GridNodes, func(node *pb.GridNode) bool { return node.Id == id })
	response := pb.GetNodeResponse{Node: data.GridNodes[i]}

	return &response, nil
}

func (s *GridService) UpdateNode(ctx context.Context, req *pb.UpdateNodeRequest) (*pb.UpdateNodeResponse, error) {
	id := req.Node.GetId()

	if id > 0 {
		// Update existing.
		i := slices.IndexFunc(data.GridNodes, func(node *pb.GridNode) bool { return node.Id == id })
		data.GridNodes[i] = req.Node
	} else {
		// Create new.
		id = int32(len(data.GridNodes) + 1)
		req.Node.Id = id
		data.GridNodes = append(data.GridNodes, req.Node)
	}

	response := pb.UpdateNodeResponse{Id: id}

	return &response, nil
}
