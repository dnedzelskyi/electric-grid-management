package data

import "egm.com/backend/pb"

var (
	GridNodes = []*pb.GridNode{
		{Id: 1, Name: "Power Plant #1", Type: pb.GridNodeTypeEnum_POWER_STATION, Power: 800.0},
		{Id: 2, Name: "Transmission Tower #1", Type: pb.GridNodeTypeEnum_TRANSMISSION, Power: -10.0},
		{Id: 3, Name: "Transmission Tower #2", Type: pb.GridNodeTypeEnum_TRANSMISSION, Power: -10.0},
		{Id: 4, Name: "Substation #1", Type: pb.GridNodeTypeEnum_SUBSTATION, Power: 10.0},
		{Id: 5, Name: "Substation #2", Type: pb.GridNodeTypeEnum_SUBSTATION, Power: 10.0},
		{Id: 6, Name: "Office #1", Type: pb.GridNodeTypeEnum_CONSUMER, Power: -100.0},
		{Id: 7, Name: "Office #2", Type: pb.GridNodeTypeEnum_SUBSTATION, Power: -50.0},
		{Id: 8, Name: "Office #3", Type: pb.GridNodeTypeEnum_SUBSTATION, Power: -65.0},
		{Id: 9, Name: "Office #4", Type: pb.GridNodeTypeEnum_SUBSTATION, Power: -50.0},
	}
	Links = []*pb.NodeLink{
		{Source: 1, Target: 2},
		{Source: 1, Target: 3},
		{Source: 2, Target: 4},
		{Source: 3, Target: 5},
		{Source: 4, Target: 6},
		{Source: 4, Target: 7},
		{Source: 5, Target: 8},
		{Source: 5, Target: 9},
	}
)
