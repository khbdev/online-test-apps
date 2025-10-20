package client

import (
    "context"
    "fmt"
    "log"
    "time"

    testpb "github.com/khbdev/proto-online-test/proto/test"
    "google.golang.org/grpc"
)

type SectionClient struct {
    client testpb.TestServiceClient
    addr   string
}

// NewSectionClient – GRPC client yaratadi
func NewSectionClient(addr string) *SectionClient {
    conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    c := testpb.NewTestServiceClient(conn)
    return &SectionClient{client: c, addr: addr}
}


func (s *SectionClient) GetSection(sectionID uint64) (*testpb.Section, error) {
    var lastErr error
    for i := 0; i < 3; i++ { 
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        resp, err := s.client.GetFullSectionStructure(ctx, &testpb.GetFullSectionRequest{SectionId: sectionID})
        if err == nil {
            return resp.Section, nil
        }

        lastErr = err
        log.Printf("[Attempt %d] Failed to get section from %s: %v", i+1, s.addr, err)
        time.Sleep(1 * time.Second) 
    }
    return nil, fmt.Errorf("all retries failed: %w", lastErr)
}
