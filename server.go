package main
import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "githum.com/baadjis/fipay/transaction-service/pb"

	

)
//_"githum.com/baadjis/fipay/transaction-service/pb"
const (
	port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.Customer
	savedTransactions []*pb.Transaction
}


func main(){
   fmt.Println("hello")
}