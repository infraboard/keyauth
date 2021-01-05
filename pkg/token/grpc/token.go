package grpc

// HelloServiceImpl todo
type HelloServiceImpl struct{}

// Config todo
func (p *HelloServiceImpl) Config() error {
	return nil
}

// Registry todo
// func (p *HelloServiceImpl) Registry(server *grpc.Server) {
// 	pb.RegisterTokenServiceServer(server, p)
// }

// // Hello todo
// func (p *HelloServiceImpl) Hello(
// 	ctx context.Context, args *pb.String,
// ) (*pb.String, error) {
// 	reply := &pb.String{Value: "hello:" + args.GetValue()}
// 	return reply, nil
// }

// func init() {
// 	pkg.RegistryGRPCV1("token", &HelloServiceImpl{})
// }
