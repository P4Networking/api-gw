package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/P4Networking/api-gw/proto/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"sync"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50050", "gRPC server endpoint")
)

type Server struct {
	api.UnimplementedPiscServiceServer
}

func StartGRPCServer(wg *sync.WaitGroup) {
	defer wg.Done()
	addr := ":50050"
	fmt.Printf("Service %v started! Listen at %v\n", "GRPC Server", addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	api.RegisterPiscServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

func StartHttpReverseProxyServer(wg *sync.WaitGroup) {
	defer wg.Done()
	addr := ":8080"

	fmt.Printf("Service %v started! Listen at %v\n", "HTTP Reverse Proxy Server", addr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterPiscServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		fmt.Println(err)
	}

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
	}
}

func (s *Server) GetSystemInformation(ctx context.Context, req *emptypb.Empty) (*api.SystemInformationResponse, error) {
	return &api.SystemInformationResponse{
		ProductName:               "Hygea",
		ProductSerialNumber:       "AI12001995",
		SystemManufacturer:        "Edge-core Networks",
		SystemManufacturingDate:   "05-11-20",
		SoftwareVersion:           "v1.0",
		FirmwareVersion:           "v1.0",
		HardwareVersion:           "R0A",
		HardwareManufacturingDate: "05-10-19",
		LocalMac:                  "3C:2C:99:6F:9D:06",
		CpuCore:                   4,
		RamSpaceTotalGb:           8,
		DiskSpaceTotalGb:          128,
	}, nil
}

func (s *Server) GetSystemCPUInformation(ctx context.Context, in *emptypb.Empty) (*api.SystemCPUInformationResponse, error) {
	cpuNum := 4
	cpuUsage := make([]uint32, 0)

	// Notice the order is dependent with core number
	for i := 0; i < cpuNum; i++ {
		cpuUsage = append(cpuUsage, uint32(rand.Intn(100)))
	}

	return &api.SystemCPUInformationResponse{
		CpuUsage: cpuUsage,
	}, nil
}
func (s *Server) GetSystemRAMInformation(ctx context.Context, in *emptypb.Empty) (*api.SystemRAMInformationResponse, error) {
	return &api.SystemRAMInformationResponse{
		RamUsage: uint32(rand.Intn(100)),
	}, nil
}
func (s *Server) GetSystemDiskInformation(ctx context.Context, in *emptypb.Empty) (*api.SystemDiskInformationResponse, error) {
	return &api.SystemDiskInformationResponse{
		DiskSpaceUsage: uint32(rand.Intn(100)),
	}, nil
}

func (s *Server) GetFeature(ctx context.Context, req *api.NameMessage) (*api.FeatureResponse, error) {
	// Debug code
	debugMsg("Feature", req)

	switch strings.ToLower(req.Name) {
	case "driver":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "Driver", Description: "Driver system", Version: "v0.99", Enabled: true})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "bgp":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "BGP", Description: "Route BGP Protocol", Version: "v0.1", Enabled: true})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "device":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "Device", Description: "Device Manager", Version: "v0.2", Enabled: false})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "flow":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "Flow", Description: "Flow Manager", Version: "v0.3", Enabled: false})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "host":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "Host", Description: "Host Manager", Version: "v0.3", Enabled: false})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "vlan":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "VLAN", Description: "IEEE 802.1Q", Version: "v0.5", Enabled: true})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	case "":
		lists := make([]*api.Feature, 0)
		lists = append(lists, &api.Feature{Name: "BGP", Description: "Route BGP Protocol", Version: "v0.1", Enabled: true})
		lists = append(lists, &api.Feature{Name: "Device", Description: "Device Manager", Version: "v0.2", Enabled: false})
		lists = append(lists, &api.Feature{Name: "Flow", Description: "Flow Manager", Version: "v0.3", Enabled: false})
		lists = append(lists, &api.Feature{Name: "Host", Description: "Host Manager", Version: "v0.3", Enabled: false})
		lists = append(lists, &api.Feature{Name: "VLAN", Description: "IEEE 802.1Q", Version: "v0.5", Enabled: true})
		lists = append(lists, &api.Feature{Name: "Driver", Description: "Driver system", Version: "v0.99", Enabled: true})
		return &api.FeatureResponse{
			FeatureLists: lists,
		}, nil
	default:
		return nil, status.Errorf(codes.NotFound, "not found")
	}
}

func (s *Server) GetPort(ctx context.Context, req *api.IdMessage) (*api.PortResponse, error) {
	// Debug code
	debugMsg("Port", req)

	var data = make(map[uint64]*api.PortResponse_Port)
	data[1] = &api.PortResponse_Port{
		Id:        1,
		Enabled:   true,
		TxPackets: 23, RxPackets: 20, TxBytes: 568, RxBytes: 512, TxDrop: 1, RxDrop: 3,
		Type:       api.PortResponse_Port_FIBER,
		MacAddress: "aa:bb:cc:dd:ee:01",
		VlanMode:   api.PortResponse_Port_ACCESS,
		VlanId:     []uint32{11},
		IpAddress:  []string{"192.168.1.1"}}
	data[2] = &api.PortResponse_Port{
		Id:        2,
		Enabled:   true,
		TxPackets: 44, RxPackets: 2, TxBytes: 1002, RxBytes: 4, TxDrop: 0, RxDrop: 0,
		Type:       api.PortResponse_Port_FIBER,
		MacAddress: "aa:bb:cc:dd:ee:02",
		VlanMode:   api.PortResponse_Port_ACCESS,
		VlanId:     []uint32{12},
		IpAddress:  []string{"192.168.2.2"}}
	data[3] = &api.PortResponse_Port{
		Id:        3,
		Enabled:   false,
		TxPackets: 0, RxPackets: 0, TxBytes: 0, RxBytes: 0, TxDrop: 0, RxDrop: 0,
		Type:       api.PortResponse_Port_FIBER,
		MacAddress: "aa:bb:cc:dd:ee:03",
		VlanMode:   api.PortResponse_Port_TRUNK,
		VlanId:     []uint32{13, 14, 15},
		IpAddress:  []string{"192.168.3.3", "192.168.4.4", "192.168.5.5"}}

	rsp := &api.PortResponse{PortLists: make([]*api.PortResponse_Port, 0)}

	if req.Id == 0 {
		for _, v := range data {
			rsp.PortLists = append(rsp.PortLists, v)
		}
		//fmt.Printf("rsp: %v, lists: %v, data: %v, rsp.list pointer: %p, list pointer: %p\n", rsp, lists, data,rsp.PortLists,lists)
		return rsp, nil
	} else {
		if p, found := data[req.Id]; found {
			rsp.PortLists = append(rsp.PortLists, p)
			return rsp, nil
		}
		return nil, status.Errorf(codes.NotFound, "not found")
	}
}

func (s *Server) GetLicense(ctx context.Context, req *emptypb.Empty) (*api.LicenseResponse, error) {
	return &api.LicenseResponse{
		LicenseLevel: api.LicenseResponse_EVALUATION,
	}, nil
}

func (s *Server) LoginAuth(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {

	debugMsg("Login", req)
	username := req.Username
	password := req.Password
	rsp := &api.LoginResponse{Username: username, IsSuccess: false, UserRole: api.LoginResponse_UNKNOWN}

	if username == "admin" && password == "admin" {
		rsp.IsSuccess = true
		rsp.UserRole = api.LoginResponse_ADMIN
	}

	if username == "light_admin" && password == "light_admin" {
		rsp.IsSuccess = true
		rsp.UserRole = api.LoginResponse_LIGHT_ADMIN
	}

	if username == "user" && password == "user" {
		rsp.IsSuccess = true
		rsp.UserRole = api.LoginResponse_USER
	}

	if username == "guest" && password == "guest" {
		rsp.IsSuccess = true
		rsp.UserRole = api.LoginResponse_GUEST
	}

	return rsp, nil
}

func debugMsg(name string, req interface{}) {
	fmt.Printf("Component: %v, Request: %v\n", name, req)
}
