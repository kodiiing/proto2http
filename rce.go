package rce

import (
	"context"
	"encoding/json"
	"net/http"
)

type EmptyRequest struct {
}

type Runtimes struct {
	Runtime []Runtime `json:"runtime"`
}

type CodeRequest struct {
	Language string `json:"language"`
	Version string `json:"version"`
	Code string `json:"code"`
	CompileTimeout int32 `json:"compile_timeout"`
	RunTimeout int32 `json:"run_timeout"`
	MemoryLimit int32 `json:"memory_limit"`
}

type CodeResponse struct {
	Language string `json:"language"`
	Version string `json:"version"`
	Compile Output `json:"compile"`
	Runtime Output `json:"runtime"`
}

type PingResponse struct {
	Message string `json:"message"`
}

type Runtime struct {
	Language string `json:"language"`
	Version string `json:"version"`
	Aliases []string `json:"aliases"`
	Compiled bool `json:"compiled"`
}

type Output struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Output string `json:"output"`
	ExitCode int32 `json:"exitCode"`
}

type CodeExecutionEngineServiceServer interface {
	ListRuntimes(ctx context.Context, req *EmptyRequest) (*Runtimes, error)
	Execute(ctx context.Context, req *CodeRequest) (*CodeResponse, error)
	Ping(ctx context.Context, req *EmptyRequest) (*PingResponse, error)
}
func NewCodeExecutionEngineServiceServer(implementation CodeExecutionEngineServiceServer) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ListRuntimes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req EmptyRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := implementation.ListRuntimes(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/Execute", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req CodeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := implementation.Execute(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/Ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req EmptyRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp, err := implementation.Ping(r.Context(), &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return mux
}
