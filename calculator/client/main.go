package main
import (
	"log"

	"google.golang.org/grpc"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	
)
var addr string = "localhost:50051"
func main() {
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	doSum(c)

}