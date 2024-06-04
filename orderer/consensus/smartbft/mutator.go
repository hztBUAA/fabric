package smartbft

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hyperledger-labs/SmartBFT/smartbftprotos"
)
// 顶层Message的消息
// type Message struct {
// 	state         protoimpl.MessageState
// 	sizeCache     protoimpl.SizeCache
// 	unknownFields protoimpl.UnknownFields

// 	// Types that are assignable to Content:
// 	//	*Message_PrePrepare
// 	//	*Message_Prepare
// 	//	*Message_Commit
// 	//	*Message_ViewChange
// 	//	*Message_ViewData
// 	//	*Message_NewView
// 	//	*Message_HeartBeat
// 	//	*Message_HeartBeatResponse
// 	//	*Message_StateTransferRequest
// 	//	*Message_StateTransferResponse
// 	Content isMessage_Content `protobuf_oneof:"content"`
// }
// 顶层封装的消息
type Message = smartbftprotos.Message
type Message_PrePrepare = smartbftprotos.Message_PrePrepare
type Message_Prepare = smartbftprotos.Message_Prepare
type Message_Commit = smartbftprotos.Message_Commit
type Message_ViewChange = smartbftprotos.Message_ViewChange
type Message_ViewData = smartbftprotos.Message_ViewData//注意这里的ViewData里面的内容是SignedViewData 而不是ViewData
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
type ViewData = smartbftprotos.SignedViewData//attention！！！
type NewView = smartbftprotos.NewView
type HeartBeat = smartbftprotos.HeartBeat
type HeartBeatResponse = smartbftprotos.HeartBeatResponse
type StateTransferRequest = smartbftprotos.StateTransferRequest
type StateTransferResponse = smartbftprotos.StateTransferResponse


//内容的子字段
// type Proposal = smartbftprotos.Proposal

// go 1.20版本后  不再需要为了是的每次生成的数是随机的而费事设置time.Now  只有当希望是伪随机时才需要单独设置种子
// r := rand.New(rand.NewSource(1)) // 创建一个新的随机数生成器，种子值为 1
// fmt.Println(r.Int()) // 输出一个随机的整数

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func MutateSmartBFTMsg(msg *Message) *Message {
	switch content := msg.Content.(type) {
	case *Message_PrePrepare:
		// 变异 PrePrepare 消息
		msg.Content = &Message_PrePrepare{
			PrePrepare: MutatePrePrepare(content.PrePrepare),
		}
	case *Message_Prepare:
		// 变异 Prepare 消息
		msg.Content = &Message_Prepare{
			Prepare: MutatePrepare(content.Prepare),
		}
	case *Message_Commit:
		// 变异 Commit 消息
		msg.Content = &Message_Commit{
			Commit: MutateCommit(content.Commit),
		}
	case *Message_ViewChange:
		// 变异 ViewChange 消息
		msg.Content = &Message_ViewChange{
			ViewChange: MutateViewChange(content.ViewChange),
		}
	case *Message_ViewData:
		// 变异 ViewData 消息
		msg.Content = &Message_ViewData{
			ViewData: MutateSignedViewData(content.ViewData),//Message_ViewData实现了Message_content的接口  是一种message_content
		}
	case *Message_NewView:
		// 变异 NewView 消息
		msg.Content = &Message_NewView{
			NewView: MutateNewView(content.NewView),
		}
	case *Message_HeartBeat:
		// 变异 HeartBeat 消息
		msg.Content = &Message_HeartBeat{
			HeartBeat: MutateHeartBeat(content.HeartBeat),
		}
	case *Message_HeartBeatResponse:
		// 变异 HeartBeatResponse 消息
		msg.Content = &Message_HeartBeatResponse{
			HeartBeatResponse: MutateHeartBeatResponse(content.HeartBeatResponse),
		}
	case *Message_StateTransferRequest:
		// 变异 StateTransferRequest 消息
		msg.Content = &Message_StateTransferRequest{
			StateTransferRequest: MutateStateTransferRequest(content.StateTransferRequest),
		}
	case *Message_StateTransferResponse:
		// 变异 StateTransferResponse 消息
		msg.Content = &Message_StateTransferResponse{
			StateTransferResponse: MutateStateTransferResponse(content.StateTransferResponse),
		}
	default:
		// 不变异其他消息类型
		fmt.Println("----------------Dont mutate local msg---------------------")
	}

	return msg
}

func MutatePrePrepare(prePrepare *PrePrepare) *PrePrepare {
	// 变异逻辑
	mutatedPrePare := *prePrepare
	mutatedPrePare.View = uint64(rand.Int63())
	mutatedPrePare.Seq = uint64(rand.Int63())
	return &mutatedPrePare
}

func MutatePrepare(prepare *Prepare) *Prepare {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择变异的字段
	switch rand.Intn(4) {
	case 0:
		// 变异 View 字段
		prepare.View = rand.Uint64()
	case 1:
		// 变异 Seq 字段
		prepare.Seq = rand.Uint64()
	case 2:
		// 变异 Digest 字段
		prepare.Digest = randomString(32) // 假设 Digest 字段是一个32字符长度的字符串
	case 3:
		// 变异 Assist 字段
		prepare.Assist = !prepare.Assist
	}

	return prepare
}


func MutateCommit(commit *Commit) *Commit {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择变异的字段
	switch rand.Intn(5) {
	case 0:
		// 变异 View 字段
		commit.View = rand.Uint64()
	case 1:
		// 变异 Seq 字段
		commit.Seq = rand.Uint64()
	case 2:
		// 变异 Digest 字段
		commit.Digest = randomString(32) // 假设 Digest 字段是一个32字符长度的字符串
	case 3:
		// 变异 Signature 字段
		commit.Signature = mutateSignature(commit.Signature)
	case 4:
		// 变异 Assist 字段
		commit.Assist = !commit.Assist
	}

	return commit
}

// mutateSignature 变异 Signature 字段
func mutateSignature(sig *smartbftprotos.Signature) *smartbftprotos.Signature {
    if sig == nil {
        sig = &smartbftprotos.Signature{}
    }
    switch rand.Intn(3) {
    case 0:
        // 变异 Signer
        sig.Signer = rand.Uint64()
    case 1:
        // 变异 Value
        sig.Value = randomBytes(64)
    case 2:
        // 变异 Msg
        sig.Msg = randomBytes(128)
    }
    return sig
}

// MutateSignedViewData 变异 SignedViewData 字段
func MutateSignedViewData(signedViewData *ViewData) *ViewData {
	switch rand.Intn(3) {
	case 0:
		signedViewData.RawViewData = randomBytes(256) // 假设 RawViewData 是 256 字节
	case 1:
		signedViewData.Signer = rand.Uint64()
	case 2:
		signedViewData.Signature = randomBytes(64)
	}
	return signedViewData
}


// randomString 生成指定长度的随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// randomBytes 生成指定长度的随机字节数组
func randomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}

// MutateViewChange 变异函数
func MutateViewChange(viewChange *ViewChange) *ViewChange {
	switch rand.Intn(2) {
	case 0:
		viewChange.NextView = rand.Uint64()
	case 1:
		viewChange.Reason = randomString(32)
	}
	return viewChange
}



// // mutateProposal 变异 Proposal 实例  只有ViewData才用到 Proposal  目前还不是很清楚 ViewData的作用  和signedViewData的作用
// func mutateProposal() *Proposal {
// 	return &Proposal{
// 		ID:      randomString(16),
// 		Content: randomString(64),
// 	}
// }

// MutateNewView 变异函数
func MutateNewView(newView *NewView) *NewView {
	for _, signedViewData := range newView.SignedViewData {
		// 定义 SignedViewData 变异逻辑
		signedViewData = MutateSignedViewData(signedViewData)
	}
	return newView
}

func MutateHeartBeat(heartBeat *HeartBeat) *HeartBeat {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(2) {
	case 0:
		heartBeat.View = rand.Uint64()
	case 1:
		heartBeat.Seq = rand.Uint64()
	}
	return heartBeat
}

func MutateHeartBeatResponse(heartBeatResponse *HeartBeatResponse) *HeartBeatResponse {
	rand.Seed(time.Now().UnixNano())
	heartBeatResponse.View = rand.Uint64()
	return heartBeatResponse
}

func MutateStateTransferRequest(stateTransferRequest *StateTransferRequest) *StateTransferRequest {
	// StateTransferRequest 消息无需变异  因为没有可变异字段
	return stateTransferRequest
}

func MutateStateTransferResponse(stateTransferResponse *StateTransferResponse) *StateTransferResponse {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(2) {
	case 0:
		stateTransferResponse.ViewNum = rand.Uint64()
	case 1:
		stateTransferResponse.Sequence = rand.Uint64()
	}
	return stateTransferResponse
}