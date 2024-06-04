package smartbft

import (
	"fmt"
	"math/rand"

	"github.com/hyperledger-labs/SmartBFT/smartbftprotos"
)

// 顶层封装的消息
type Message = smartbftprotos.Message
type Message_PrePrepare = smartbftprotos.Message_PrePrepare
type Message_Prepare = smartbftprotos.Message_Prepare
type Message_Commit = smartbftprotos.Message_Commit
type Message_ViewChange = smartbftprotos.Message_ViewChange
type Message_ViewData = smartbftprotos.Message_ViewData
type Message_NewView = smartbftprotos.Message_NewView
type Message_HeartBeat = smartbftprotos.Message_HeartBeat
type Message_HeartBeatResponse = smartbftprotos.Message_HeartBeatResponse
type Message_StateTransferRequest = smartbftprotos.Message_StateTransferRequest
type Message_StateTransferResponse = smartbftprotos.Message_StateTransferResponse

// 对应的实际内容
type PrePrepare = smartbftprotos.PrePrepare
type Prepare = smartbftprotos.Prepare
type Commit = smartbftprotos.Commit
type ViewChange = smartbftprotos.ViewChange
type ViewData = smartbftprotos.SignedViewData
type NewView = smartbftprotos.NewView
type HeartBeat = smartbftprotos.HeartBeat
type HeartBeatResponse = smartbftprotos.HeartBeatResponse
type StateTransferRequest = smartbftprotos.StateTransferRequest
type StateTransferResponse = smartbftprotos.StateTransferResponse

// go 1.20版本后  不再需要为了是的每次生成的数是随机的而费事设置time.Now  只有当希望是伪随机时才需要单独设置种子
// r := rand.New(rand.NewSource(1)) // 创建一个新的随机数生成器，种子值为 1
// fmt.Println(r.Int()) // 输出一个随机的整数

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func MutateSmartBFTMsg(msg *Message) *Message {
	mutatedMsg := *msg

	switch content := msg.Content.(type) {
	case *Message_PrePrepare:
		// 变异 PrePrepare 消息
		mutatedMsg.Content = &Message_PrePrepare{
			PrePrepare: MutatePrePrepare(content.PrePrepare),
		}
	case *Message_Prepare:
		// 变异 Prepare 消息
		mutatedMsg.Content = &Message_Prepare{
			Prepare: MutatePrepare(content.Prepare),
		}
	case *Message_Commit:
		// 变异 Commit 消息
		mutatedMsg.Content = &Message_Commit{
			Commit: MutateCommit(content.Commit),
		}
	case *Message_ViewChange:
		// 变异 ViewChange 消息
		mutatedMsg.Content = &Message_ViewChange{
			ViewChange: MutateViewChange(content.ViewChange),
		}
	case *Message_ViewData:
		// 变异 ViewData 消息
		mutatedMsg.Content = &Message_ViewData{
			ViewData: MutateViewData(content.ViewData),
		}
	case *Message_NewView:
		// 变异 NewView 消息
		mutatedMsg.Content = &Message_NewView{
			NewView: MutateNewView(content.NewView),
		}
	case *Message_HeartBeat:
		// 变异 HeartBeat 消息
		mutatedMsg.Content = &Message_HeartBeat{
			HeartBeat: MutateHeartBeat(content.HeartBeat),
		}
	case *Message_HeartBeatResponse:
		// 变异 HeartBeatResponse 消息
		mutatedMsg.Content = &Message_HeartBeatResponse{
			HeartBeatResponse: MutateHeartBeatResponse(content.HeartBeatResponse),
		}
	case *Message_StateTransferRequest:
		// 变异 StateTransferRequest 消息
		mutatedMsg.Content = &Message_StateTransferRequest{
			StateTransferRequest: MutateStateTransferRequest(content.StateTransferRequest),
		}
	case *Message_StateTransferResponse:
		// 变异 StateTransferResponse 消息
		mutatedMsg.Content = &Message_StateTransferResponse{
			StateTransferResponse: MutateStateTransferResponse(content.StateTransferResponse),
		}
	default:
		// 不变异其他消息类型
		fmt.Println("----------------Dont mutate local msg---------------------")
	}

	return &mutatedMsg
}

func MutatePrePrepare(prePrepare *PrePrepare) *PrePrepare {
	// 变异逻辑
	mutatedPrePare := *prePrepare
	rand.Int()
	mutatedPrePare.View = uint64(rand.Int63())
	mutatedPrePare.Seq = uint64(rand.Int63())
	return &mutatedPrePare
}

func MutatePrepare(prepare *Prepare) *Prepare {
	// 变异逻辑
	return prepare
}

func MutateCommit(commit *Commit) *Commit {
	// 变异逻辑
	return commit
}

func MutateViewChange(viewChange *ViewChange) *ViewChange {
	// 变异逻辑
	return viewChange
}

func MutateViewData(viewData *ViewData) *ViewData {
	// 变异逻辑
	return viewData
}

func MutateNewView(newView *NewView) *NewView {
	// 变异逻辑
	return newView
}

func MutateHeartBeat(heartBeat *HeartBeat) *HeartBeat {
	// 变异逻辑
	return heartBeat
}

func MutateHeartBeatResponse(heartBeatResponse *HeartBeatResponse) *HeartBeatResponse {
	// 变异逻辑
	return heartBeatResponse
}

func MutateStateTransferRequest(stateTransferRequest *StateTransferRequest) *StateTransferRequest {
	// 变异逻辑
	return stateTransferRequest
}

func MutateStateTransferResponse(stateTransferResponse *StateTransferResponse) *StateTransferResponse {
	// 变异逻辑
	return stateTransferResponse
}
