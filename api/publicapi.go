package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/theQRL/zond/api/view"
	"github.com/theQRL/zond/chain"
	"github.com/theQRL/zond/config"
	"github.com/theQRL/zond/ntp"
	"github.com/theQRL/zond/p2p/messages"
	"github.com/theQRL/zond/protos"
	"net/http"
	"strconv"
	"time"
)

type GetHeightResponse struct {
	Height uint64 `json:"height"`
}

type GetEstimatedNetworkFeeResponse struct {
	Fee string `json:"fee"`
}

type GetVersionResponse struct {
	Version string `json:"version"`
}

type Response struct {
	Error        uint        `json:"error"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

type PublicAPIServer struct {
	chain                    *chain.Chain
	ntp                      ntp.NTPInterface
	config                   *config.Config
	visitors                 *visitors
	registerAndBroadcastChan chan *messages.RegisterMessage
}

func (p *PublicAPIServer) Start() error {
	c := config.GetConfig()

	router := mux.NewRouter()
	router.HandleFunc("/api/", p.RedirectToAPIDoc).Methods("GET")
	router.HandleFunc("/api/version", p.GetVersion).Methods("GET")
	router.HandleFunc("/api/block/{hash}", p.GetBlockByHash).Methods("GET")
	router.HandleFunc("/api/block/last", p.GetLastBlock).Methods("GET")
	router.HandleFunc("/api/address/{address}", p.GetAddressState).Methods("GET")
	router.HandleFunc("/api/balance/{address}", p.GetBalance).Methods("GET")
	router.HandleFunc("/api/fee", p.GetEstimatedNetworkFee).Methods("GET")
	router.HandleFunc("/api/broadcast/transfer", p.BroadcastTransferTx).Methods("POST")
	router.HandleFunc("/api/broadcast/stake", p.BroadcastStakeTx).Methods("POST")
	router.HandleFunc("/api/height", p.GetHeight).Methods("GET")
	//handler := cors.Default().Handler(router)
	co := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,//http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
			"post",
			"*",
			"*/*",
		},

		AllowedHeaders: []string{
			"*",//or you can your header key values which you are using in your application

		},
	})
	router.StrictSlash(false)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", c.User.API.PublicAPI.Host, c.User.API.PublicAPI.Port), co.Handler(router))
	if err != nil {

	}
	return nil
}

func (p *PublicAPIServer) prepareResponse(errorCode uint, errorMessage string, data interface{}) *Response {
	r := &Response{
		Error: errorCode,
		ErrorMessage: errorMessage,
		Data: data,
	}
	return r
}

func (p *PublicAPIServer) RedirectToAPIDoc(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	http.Redirect(w, r, "https://api.theqrl.org/", 301)
}

func (p *PublicAPIServer) GetVersion(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	fmt.Println("Get version called");
	getVersionResponse := &GetVersionResponse{
		Version: p.config.Dev.Version,
	}
	json.NewEncoder(w).Encode(p.prepareResponse(0, "", getVersionResponse))
}

func (p *PublicAPIServer) GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	vars := mux.Vars(r)
	hash, found := vars["hash"]
	if !found {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"block hash not provided",
			nil))
		return
	}
	headerHash, err := hex.DecodeString(hash)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error Decoding headerHash\n %s", err.Error()),
			nil))
		return
	}
	b, err := p.chain.GetBlock(headerHash)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error in GetBlock for headerHash %s\n %s", hash, err.Error()),
			nil))
		return
	}
	response := &view.PlainBlock{}
	response.BlockFromPBData(b.PBData(), b.HeaderHash())
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) GetLastBlock(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	b := p.chain.GetLastBlock()

	response := &view.PlainBlock{}
	response.BlockFromPBData(b.PBData(), b.HeaderHash())
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) GetAddressState(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	vars := mux.Vars(r)
	address := vars["address"]
	binAddress, err := hex.DecodeString(address)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error Decoding address %s\n %s", address, err.Error()),
			nil))
		return
	}
	addressState, err := p.chain.GetAddressState(binAddress)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error loading address State %s\n %s", address, err.Error()),
			nil))
		return
	}
	response := &view.PlainAddressState{}
	response.AddressStateFromPBData(addressState.PBData())

	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) GetBalance(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	vars := mux.Vars(r)
	address := vars["address"]
	binAddress, err := hex.DecodeString(address)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error Decoding address %s\n %s", address, err.Error()),
			nil))
		return
	}
	addressState, err := p.chain.GetAddressState(binAddress)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error loading address state %s\n %s", address, err.Error()),
			nil))
		return
	}
	response := view.PlainBalance{}
	response.Balance = strconv.FormatUint(addressState.Balance(), 10)
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) GetHeight(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	response := &GetHeightResponse{Height:p.chain.Height()}
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) GetNetworkStats(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
}

func (p *PublicAPIServer) GetEstimatedNetworkFee(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	// TODO: Fee needs to be calcuated by mean, median or mode
	response := &GetEstimatedNetworkFeeResponse{Fee:"1"}
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) BroadcastStakeTx(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var plainStakeTransaction view.PlainStakeTransaction
	err := decoder.Decode(&plainStakeTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error Decoding PlainStakeTransaction \n%s", err.Error()),
			nil))
		return
	}
	tx, err := plainStakeTransaction.ToStakeTransactionObject()
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error parsing ToStakeTransactionObject\n %s", err.Error()),
			nil))
		return
	}
	sc, err := p.chain.GetStateContext()
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Error getting state context for transaction validation",
			nil))
		return
	}
	if err := tx.SetAffectedAddress(sc); err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Error setting affected address into state context for transaction validation",
			nil))
		return
	}
	if !tx.Validate(sc) {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Transaction Validation Failed",
			nil))
		return
	}

	txHash := tx.TxHash(tx.GetSigningHash())
	err = p.chain.GetTransactionPool().Add(
		tx,
		txHash,
		p.chain.GetLastBlock().SlotNumber(),
		p.ntp.Time())
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Failed to Add Txn into txn pool \n %s", err.Error()),
			nil))
		return
	}

	registerMessage := &messages.RegisterMessage {
		Msg:&protos.LegacyMessage{
			Data: &protos.LegacyMessage_StData{
				StData:tx.PBData(),
			},
			FuncName:protos.LegacyMessage_ST,
		},
		MsgHash: hex.EncodeToString(txHash),
	}
	select {
	case p.registerAndBroadcastChan <- registerMessage:
	case <-time.After(10*time.Second):
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Transaction Broadcast Timeout",
			nil))
		return
	}

	response := &BroadcastTransactionResponse{
		TransactionHash: hex.EncodeToString(txHash),
	}
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func (p *PublicAPIServer) BroadcastTransferTx(w http.ResponseWriter, r *http.Request) {
	// Check Rate Limit
	if !p.visitors.isAllowed(r.RemoteAddr) {
		w.WriteHeader(429)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var plainTransferTransaction view.PlainTransferTransaction
	err := decoder.Decode(&plainTransferTransaction)
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error Decoding PlainTransferTransaction \n%s", err.Error()),
			nil))
		return
	}
	tx, err := plainTransferTransaction.ToTransferTransactionObject()
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Error parsing ToTransferTransactionObject\n %s", err.Error()),
			nil))
		return
	}
	sc, err := p.chain.GetStateContext()
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Error getting state context for transaction validation",
			nil))
		return
	}
	if err := tx.SetAffectedAddress(sc); err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Error setting affected address into state context for transaction validation",
			nil))
		return
	}
	if !tx.Validate(sc) {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Transaction Validation Failed",
			nil))
		return
	}

	txHash := tx.TxHash(tx.GetSigningHash())
	err = p.chain.GetTransactionPool().Add(
		tx,
		txHash,
		p.chain.GetLastBlock().SlotNumber(),
		p.ntp.Time())
	if err != nil {
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			fmt.Sprintf("Failed to Add Txn into txn pool \n %s", err.Error()),
			nil))
		return
	}

	registerMessage := &messages.RegisterMessage {
		Msg:&protos.LegacyMessage{
			Data: &protos.LegacyMessage_TtData{
				TtData:tx.PBData(),
			},
			FuncName:protos.LegacyMessage_TT,
		},
		MsgHash: hex.EncodeToString(txHash),
	}
	select {
	case p.registerAndBroadcastChan <- registerMessage:
	case <-time.After(10*time.Second):
		json.NewEncoder(w).Encode(p.prepareResponse(1,
			"Transaction Broadcast Timeout",
			nil))
		return
	}
	response := &BroadcastTransactionResponse{
		TransactionHash: hex.EncodeToString(txHash),
	}
	json.NewEncoder(w).Encode(p.prepareResponse(0,
		"",
		response))
}

func NewPublicAPIServer(c *chain.Chain, registerAndBroadcastChan chan *messages.RegisterMessage) *PublicAPIServer {
	return &PublicAPIServer{
		chain:                    c,
		ntp:                      ntp.GetNTP(),
		config:                   config.GetConfig(),
		visitors:                 newVisitors(),
		registerAndBroadcastChan: registerAndBroadcastChan,
	}
}