package gateway

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
	"github.com/projectriff/system/pkg/client/informers/externalversions/streaming/v1alpha1"
	"github.com/projectriff/system/pkg/gateway/liiklus"
	"github.com/projectriff/system/pkg/gateway/serialization"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const mimeTypeOctetStream = "application/octet-stream"

type Gateway struct {
	server         *http.Server
	liiklusClients map[string]liiklus.LiiklusServiceClient
	grpcConns      map[string]*grpc.ClientConn
	streamInformer v1alpha1.StreamInformer
	mutex sync.RWMutex
}

func NewGateway(informer v1alpha1.StreamInformer) *Gateway {
	g := Gateway{
		streamInformer: informer,
		liiklusClients: make(map[string]liiklus.LiiklusServiceClient),
		grpcConns:      make(map[string]*grpc.ClientConn),
	}

	m := http.NewServeMux()
	m.HandleFunc("/", g.ingest)

	g.server = &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	return &g
}

func (g *Gateway) Run(stopCh <-chan struct{}) error {

	err := g.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	} else {
		<-stopCh
		return nil
	}
}

func (g *Gateway) ingest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Only POSTs are accepted"))
		return
	}
	parts := strings.Split(r.RequestURI[1:], "/")
	if len(parts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Request URI should be of the form /<namespace>/<stream-name>"))
		return
	}

	ns := parts[0]
	streamName := parts[1]

	stream, err := g.streamInformer.Lister().Streams(ns).Get(streamName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf("Stream %s/%s not found", ns, streamName)))
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		contentType = mimeTypeOctetStream
	}

	// TODO: perform more loose type validation, eg supporting wildcards
	if stream.Spec.ContentType != contentType {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(fmt.Sprintf("Stream %s/%s only accepts %s", ns, streamName, stream.Spec.ContentType)))
	}

	m := serialization.Message{}
	m.ContentType = contentType
	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error reading request body"))
		return
	} else {
		m.Payload = bytes
	}
	_ = g.copyHeaders(r, m, stream)

	client, err := g.lookupClient(stream.Status.Address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error locating client for stream"))
		return
	}
	var value []byte
	if value, err = proto.Marshal(&m) ; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error serializing request"))
		return
	}
	request := liiklus.PublishRequest{
		Topic:streamName, // TODO: incorporate ns in name
		Value: value,
	}
	publishReply, err := client.Publish( /*TODO*/ context.Background(), &request)
	_ = publishReply
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error publishing message to stream"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	if err := g.server.Shutdown(ctx); err != nil {
		return err
	}
	var err error = nil
	for _, c := range g.grpcConns {
		if e := c.Close(); e != nil && err == nil {
			err = e
		}
	}
	return err
}

func (g *Gateway) copyHeaders(request *http.Request, message serialization.Message, stream *alpha1.Stream) error {
	return nil
}

func (g *Gateway) lookupClient(address alpha1.StreamAddress) (liiklus.LiiklusServiceClient, error) {
	g.mutex.RLock()
	if c, ok := g.liiklusClients[address.Gateway] ; ok {
		g.mutex.RUnlock()
		return c, nil
	}
	g.mutex.RUnlock()
	g.mutex.Lock()
	defer g.mutex.Unlock()

	timeout, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, err := grpc.DialContext(timeout, address.Gateway, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := liiklus.NewLiiklusServiceClient(conn)
	g.grpcConns[address.Gateway] = conn
	g.liiklusClients[address.Gateway] = client
	return client, nil
}
