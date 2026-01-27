package sdk_rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	ErrInvalidResponse = errors.New("invalid JSON-RPC response")
	ErrHTTPError       = errors.New("http error")
)

type Response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// Call
// ────────────────────────────────────────────────
// Core calling method (internal)
// ────────────────────────────────────────────────
func (r *RpcClient) Call(ctx context.Context, method string, params []interface{}, result interface{}) error {
	p := struct {
		JSONRPC string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params,omitempty"`
		Id      int         `json:"id"`
	}{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
	}
	body, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("marshal request body failed: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, r.host, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "tronrpc-go/0.1")

	resp, err := r.client.Do(request)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer func() {
		_, _ = io.Copy(ioutil.Discard, request.Body)
		request.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%w: %d %s", ErrHTTPError, resp.StatusCode, string(body))
	}

	var rpcResp Response
	if err = json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return fmt.Errorf("decode json-rpc response failed: %w", err)
	}
	if rpcResp.Error != nil {
		return fmt.Errorf("tron error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	if rpcResp.Result == nil {
		return ErrInvalidResponse
	}

	if err = json.Unmarshal(rpcResp.Result, result); err != nil {
		return fmt.Errorf("unmarshal json-rpc response failed: %w", err)
	}

	return nil
}
