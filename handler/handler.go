package handler

import (
	"encoding/json"
	pb "github.com/micro/go-micro/v2/debug/service/proto"
	"github.com/micro/go-micro/v2/debug/stats"
	"net/http"
)

func StatusHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		rsp := &pb.StatsResponse{}
		sdata, err := stats.DefaultStats.Read()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(sdata) > 0 {
			// write the response values
			rsp.Timestamp = uint64(sdata[0].Timestamp)
			rsp.Started = uint64(sdata[0].Started)
			rsp.Uptime = uint64(sdata[0].Uptime)
			rsp.Memory = sdata[0].Memory
			rsp.Gc = sdata[0].GC
			rsp.Threads = sdata[0].Threads
			rsp.Requests = sdata[0].Requests
			rsp.Errors = sdata[0].Errors
		}
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(rsp); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func HealthHandler() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		rsp := &pb.HealthResponse{}
		rsp.Status = "ok"
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(rsp); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}
