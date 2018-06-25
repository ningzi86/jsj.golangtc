package tests

import (
    "testing"
    "net/rpc"
    "net/http"
    "log"
    "net"
    "time"
	"fmt"
)

//http://www.cnblogs.com/yjf512/archive/2013/02/28/2937261.html

type Args struct {
        A, B int
}
type Arith int
func (t *Arith) Multiply(args *Args, reply *([]string)) error {

        args.A = 123

        *reply = append(*reply, "test")
        fmt.Println(*t)
        return nil
}

func _Test_Rpc01(t *testing.T)  {
    
    var a Arith = 10

    args := &Args{7,8}
    reply := make([]string, 10)

    a.Multiply(args, &reply)
    fmt.Println(args)

}

func Test_Rpc(t *testing.T)  {

        arith := new(Arith)
        rpc.Register(arith)
        rpc.HandleHTTP()

        l, e := net.Listen("tcp", ":1234")
        if e != nil {
                log.Fatal("listen error:", e)
        }

        go http.Serve(l, nil)
        time.Sleep(1 * time.Second)

        client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
        if err != nil {
                log.Fatal("dialing:", err)
        }
         
        args := &Args{7,8}
        reply := make([]string, 10)
        err = client.Call("Arith.Multiply", args, &reply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        log.Println(reply)
        fmt.Println(reply)
        
        fmt.Println(args)


}

