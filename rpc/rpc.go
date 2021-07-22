package rpc

import (
	"encoding/json"

	"main.go/api"
	"main.go/config"
)

type RPCClient struct {
	Url string
}

func NewRPCClient(cfg config.Config) {

}

type RPCResponse struct {
	Id     *json.RawMessage       `json:"id"`
	Result *json.RawMessage       `json:"result"`
	Error  map[string]interface{} `json:"error"`
}

type Work struct {
	Header string // current block header pow-hash
	Seed   string // the seed hash used for the DAG.
	Target string // boundary condition
}

// GetWork() makes POSt call to RPC, returns work
func (r *RPCClient) GetWork() (*Work, error) {
	// jsonRpc and id are static in docs
	httpResp, err := api.Post(r.Url, "2.0", "eth_getWork", nil, 0)
	if err != nil {
		return nil, err
	}

	var resp *RPCResponse
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	var work *Work
	err = json.Unmarshal(*resp.Result, &work)

	return work, err
}
