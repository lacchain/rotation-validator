package service


import(
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/lacchain/rotation-validator/rpc"
	"github.com/lacchain/rotation-validator/audit"
)

func vote(rpcURL string, id string, address string, kind bool) *rpc.JsonrpcMessage{
	data := fmt.Sprintf(`{"jsonrpc":"2.0","method":"ibft_proposeValidatorVote",
	"params":["%s",%t], "id":"%s"}`,address,kind,id)

	requestBody := []byte(data)

	timeout := time.Duration(5*time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type","application/json")

	if err != nil {
		handleError(id,err)
	}

	response, err := client.Do(request)
	if err != nil {
		handleError(id,err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		handleError(id,err)
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))

	var rpcMessage rpc.JsonrpcMessage

	err = json.NewDecoder(rdr1).Decode(&rpcMessage)
	if err != nil {
		handleError(id,err)
	}

	return &rpcMessage
}

//HandleError
func handleError(id string, err error)(*rpc.JsonrpcMessage){
	audit.GeneralLogger.Println(err.Error())
	result := new(rpc.JsonrpcMessage)
	result.ID = json.RawMessage(id)
	return result.ErrorResponse(err)
}