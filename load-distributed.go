type Request struct {
    ID int
    Payload string
}

type Response struct {
    ID int
    Status string
}

type Server struct {
    Addr string
    Active bool
    Load int
}

func handleRequest(req Request) Response {
    return Response{ID: req.ID, Status: "Processed"}
}

func sendToServer(srv *Server, req Request) Response {
    if srv.Active {
        srv.Load++
        res := handleRequest(req)
        srv.Load--
        return res
    }
    return Response{ID: req.ID, Status: "Failed"}
}

func loadBalancer(req Request, servers []*Server) Response {
    minLoad := 9999
    var target *Server
    for _, srv := range servers {
        if srv.Active && srv.Load < minLoad {
            minLoad = srv.Load
            target = srv
        }
    }
    if target != nil {
        return sendToServer(target, req)
    }
    return Response{ID: req.ID, Status: "No Active Server"}
}

func monitor(servers []*Server) {
    for _, srv := range servers {
        srv.Active = ping(srv.Addr)
    }
}

func ping(addr string) bool {
    conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

func adaptiveScaling(servers []*Server) []*Server {
    avgLoad := 0
    activeCount := 0
    for _, srv := range servers {
        if srv.Active {
            avgLoad += srv.Load
            activeCount++
        }
    }
    if activeCount > 0 {
        avgLoad = avgLoad / activeCount
    }
    if avgLoad > 5 {
        servers = append(servers, &Server{Addr: fmt.Sprintf("10.0.0.%d:8080", rand.Intn(100)), Active: true})
    }
    return servers
}

func main() {
    servers := []*Server{
        {Addr: "*.*.*.*:8080", Active: true},
        {Addr: "*.*.*.*:8080", Active: true},
        {Addr: "*.*.*.*:8080", Active: true},
    }
    for i := 1; i <= 10; i++ {
        req := Request{ID: i, Payload: "Data"}
        monitor(servers)
        res := loadBalancer(req, servers)
        fmt.Println("Response:", res)
        servers = adaptiveScaling(servers)
    }
}
