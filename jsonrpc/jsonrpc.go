package jsonrpc

import (
	"encoding/json"
	"io"
)

type Request struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type Response struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  any    `json:"result"`
}

func Read(data []byte) (req Request, err error) {
	err = json.Unmarshal(data, &req)
	return
}

func Write(writer io.Writer, id int, result any) error {
	res, err := json.Marshal(Response{
		JsonRPC: "2.0",
		ID:      id,
		Result:  result,
	})

	if err == nil {
		_, err = writer.Write(append(res, '\n'))
	}

	return err
}
