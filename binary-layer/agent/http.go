package agent

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-collections/go-datastructures/queue"
)

const DefaultSDKHttpListenerPort = "6723"

type LogsApiHttpListener struct {
	sdkHttpServer *http.Server
	queue         *queue.Queue
}

func NewLogsApiHttpListener(queue *queue.Queue) (*LogsApiHttpListener, error) {

	return &LogsApiHttpListener{
		sdkHttpServer: nil,
		queue:         queue,
	}, nil
}

func SdkListenOnAddress() string {
	env_aws_local, ok := os.LookupEnv("SLS_TEST_EXTENSION")
	if ok && env_aws_local == "1" {
		return "127.0.0.1:" + DefaultSDKHttpListenerPort
	}
	return "localhost:" + DefaultSDKHttpListenerPort
}

func (s *LogsApiHttpListener) Start() (bool, error) {
	sdkAddress := SdkListenOnAddress()
	sdkMux := http.NewServeMux()
	sdkMux.HandleFunc("/async", s.http_handler)
	s.sdkHttpServer = &http.Server{Addr: sdkAddress, Handler: sdkMux}

	go func() {
		err := s.sdkHttpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			s.Shutdown()
		}
	}()

	return true, nil
}

func (h *LogsApiHttpListener) http_handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	logSet := string(body)
	fmt.Println(logSet)
	h.queue.Put(logSet)
}

func (s *LogsApiHttpListener) Shutdown() {
	if s.sdkHttpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
		err := s.sdkHttpServer.Shutdown(ctx)
		if err != nil {
		} else {
			s.sdkHttpServer = nil
		}
	}
}

type HttpAgent struct {
	listener *LogsApiHttpListener
}

func NewHttpAgent(queue *queue.Queue) (*HttpAgent, error) {
	logsApiListener, err := NewLogsApiHttpListener(queue)
	if err != nil {
		return nil, err
	}

	return &HttpAgent{
		listener: logsApiListener,
	}, nil
}

func (a HttpAgent) Init() {
	a.listener.Start()
}

func (a *HttpAgent) Shutdown() {
	a.listener.Shutdown()
}
