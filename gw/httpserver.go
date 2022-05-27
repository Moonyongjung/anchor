package gw

import (
	"encoding/json"
	"reflect"
	"net/http"
	"strings"

	"github.com/Moonyongjung/anchor/manage"
	ctypes "github.com/Moonyongjung/anchor/types"
	"github.com/Moonyongjung/anchor/util"
	"github.com/rs/cors"

)

//-- HTTPServer operates for sending or invoking transaction when user call 
func RunHttpServer(channel ctypes.ChannelStruct) {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		util.LogHttpServer("Client IP addr : " + r.RemoteAddr)
		util.LogHttpServer("Request API : " + r.URL.Path)

		apiVersionCheck := strings.Split(r.URL.Path, "/")

		if !(apiVersionCheck[1]=="api" && apiVersionCheck[2]=="v1") {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "application/json;charset=utf-8")
			w.Write([]byte("Wrong API version"))

		} else if !(apiVersionCheck[3] == "call" || apiVersionCheck[3] == "invoke"){
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("content-type", "application/json;charset=utf-8")
			w.Write([]byte("Wrong API version"))

		} else {
			var result interface{}

			doTransactionbyType(r, apiVersionCheck[3:], channel)	

			result = <- channel.CallContractResultChan
			w.WriteHeader(http.StatusOK)

			resultJsonByte, err := returnBlockListJson(result)			
			if err != nil {
				util.LogHttpServer(err)
			}
			w.Write(resultJsonByte)
		}
	})

	handler := cors.Default().Handler(mux)
	util.LogHttpServer("Server Listen...")	

	err := http.ListenAndServe(":"+gatewayServerPort, handler)
	if err != nil {
		util.LogHttpServer(err)	}
}

func doTransactionbyType(request *http.Request, transactionType []string, channel ctypes.ChannelStruct) {
	if transactionType[0] == "call" {
		channel.CallContractChan <- ctypes.CallContractChanStruct{
			InitBool: false,
			CallName: transactionType[1],
		}
	} else if transactionType[0] == "invoke" {
		//-- Not using "invoke" through HTTPServer

		// var invokeContractStruct ctypes.InvokeContractChanStruct
		// body, err := ioutil.ReadAll(request.Body)
		// storeInputData := util.JsonUnmarshalData(invokeContractStruct, body)
		// storeInput := storeInputData.(map[string]interface{})["storeInput"].(string)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		
		// channel.InvokeContractChan <- ctypes.InvokeContractChanStruct{
		// 	InvokeName: transactionType[1],
		// 	StoreInput: storeInput,
		// }
	}
}

func returnBlockListJson(result interface{}) ([]byte, error){
	var interfaceSlice []interface{}
	resultType := reflect.TypeOf(result)	

	if resultType == reflect.TypeOf(interfaceSlice) {
		resultSlice := result.([]interface{})

		blockList := manage.BlockList {
			resultSlice[2].(string),
			resultSlice[1].(string),
			resultSlice[0].(string),
		}
		b, err := json.Marshal(blockList)

		return b, err

	} else {
		util.LogHttpServer("Response type is not []interface{}")		
		return nil, nil
	}	
}