// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x6080604052600436106100f2575f3560e01c80638da5cb5b11610087578063c4d66de811610057578063c4d66de8146102c3578063e70fc93b146102e2578063ecc704281461031b578063f2fde38b1461032f575f80fd5b80638da5cb5b146102485780638ef1332e14610272578063b2267a7b14610291578063bedb86fb146102a4575f80fd5b80635f7b1577116100c25780635f7b1577146101c95780636e296e45146101dc578063715018a614610208578063797594b01461021c575f80fd5b806302345b50146100fd5780632a6cccb214610140578063478222c2146101615780635c975abb146101b2575f80fd5b366100f957005b5f80fd5b348015610108575f80fd5b5061012b6101173660046115ea565b60fc6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561014b575f80fd5b5061015f61015a366004611629565b61034e565b005b34801561016c575f80fd5b5060cb5461018d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610137565b3480156101bd575f80fd5b5060655460ff1661012b565b61015f6101d7366004611649565b6103dd565b3480156101e7575f80fd5b5060c95461018d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610213575f80fd5b5061015f61042f565b348015610227575f80fd5b5060ca5461018d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610253575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661018d565b34801561027d575f80fd5b5061015f61028c3660046117bb565b610442565b61015f61029f366004611828565b6105be565b3480156102af575f80fd5b5061015f6102be366004611883565b6105d8565b3480156102ce575f80fd5b5061015f6102dd366004611629565b6105f9565b3480156102ed575f80fd5b5061030d6102fc3660046115ea565b60fb6020525f908152604090205481565b604051908152602001610137565b348015610326575f80fd5b5061030d6107b7565b34801561033a575f80fd5b5061015f610349366004611629565b61083e565b6103566108d8565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f591015b60405180910390a15050565b6103e561093f565b610427868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250610992915050565b505050505050565b6104376108d8565b6104405f610bfd565b565b61044a61093f565b60ca5473ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff16146105135760405162461bcd60e51b8152602060048201526024808201527f43616c6c6572206973206e6f74204c3143726f7373446f6d61696e4d6573736560448201527f6e6765720000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b5f6105218686868686610c73565b80516020918201205f81815260fc90925260409091205490915060ff16156105b15760405162461bcd60e51b815260206004820152602960248201527f4d6573736167652077617320616c7265616479207375636365737366756c6c7960448201527f2065786563757465640000000000000000000000000000000000000000000000606482015260840161050a565b6104278686868585610d0f565b6105c661093f565b6105d284848484610992565b50505050565b6105e06108d8565b80156105f1576105ee610f91565b50565b6105ee611016565b5f54610100900460ff161580801561061757505f54600160ff909116105b806106305750303b15801561063057505f5460ff166001145b6106a25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161050a565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156106fe575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff821661074b576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610755825f61106d565b80156107b3575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016103d1565b5050565b5f73530000000000000000000000000000000000000173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610815573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083991906118a2565b905090565b6108466108d8565b73ffffffffffffffffffffffffffffffffffffffff81166108cf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161050a565b6105ee81610bfd565b60335473ffffffffffffffffffffffffffffffffffffffff1633146104405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161050a565b60655460ff16156104405760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161050a565b61099a6111a0565b8234146109e95760405162461bcd60e51b815260206004820152601260248201527f6d73672e76616c7565206d69736d617463680000000000000000000000000000604482015260640161050a565b5f73530000000000000000000000000000000000000190505f8173ffffffffffffffffffffffffffffffffffffffff1663b58343bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a4b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a6f91906118a2565b90505f610a7f3388888589610c73565b80516020918201205f81815260fb90925260409091205490915015610ae65760405162461bcd60e51b815260206004820152601260248201527f4475706c696361746564206d6573736167650000000000000000000000000000604482015260640161050a565b5f81815260fb602052604090819020429055517f600a2e770000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff84169063600a2e77906024016020604051808303815f875af1158015610b60573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b8491906118a2565b5073ffffffffffffffffffffffffffffffffffffffff87163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8885888a604051610be89493929190611924565b60405180910390a35050506105d26001609755565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60608585858585604051602401610c8e959493929190611952565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b7fffffffffffffffffffffffffacffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff851601610dba5760405162461bcd60e51b815260206004820152602660248201527f466f7262696420746f2063616c6c206c3220746f206c31206d6573736167652060448201527f7061737365720000000000000000000000000000000000000000000000000000606482015260840161050a565b610dc384611200565b60c95473ffffffffffffffffffffffffffffffffffffffff90811690861603610e2e5760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6573736167652073656e64657200000000000000000000604482015260640161050a565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff878116919091179091556040515f918616908590610e889086906119a1565b5f6040518083038185875af1925050503d805f8114610ec2576040519150601f19603f3d011682016040523d82523d5f602084013e610ec7565b606091505b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015610f5f575f82815260fc602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2610427565b60405182907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a2505050505050565b610f9961093f565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610fec3390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b61101e611265565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33610fec565b5f54610100900460ff166110e95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b6110f16112b7565b6110f961133b565b6111016113bf565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661dead1790915560ca805473ffffffffffffffffffffffffffffffffffffffff858116919093161790558116156107b35760cb805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790555050565b6002609754036111f25760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161050a565b6002609755565b6001609755565b3073ffffffffffffffffffffffffffffffffffffffff8216036105ee5760405162461bcd60e51b815260206004820152601e60248201527f4d657373656e6765723a20466f7262696420746f2063616c6c2073656c660000604482015260640161050a565b60655460ff166104405760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161050a565b5f54610100900460ff166113335760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b610440611443565b5f54610100900460ff166113b75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b6104406114c8565b5f54610100900460ff1661143b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b61044061156e565b5f54610100900460ff166114bf5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b61044033610bfd565b5f54610100900460ff166115445760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b5f54610100900460ff166111f95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161050a565b5f602082840312156115fa575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611624575f80fd5b919050565b5f60208284031215611639575f80fd5b61164282611601565b9392505050565b5f805f805f8060a0878903121561165e575f80fd5b61166787611601565b955060208701359450604087013567ffffffffffffffff8082111561168a575f80fd5b818901915089601f83011261169d575f80fd5b8135818111156116ab575f80fd5b8a60208285010111156116bc575f80fd5b602083019650809550505050606087013591506116db60808801611601565b90509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611723575f80fd5b813567ffffffffffffffff8082111561173e5761173e6116e7565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611784576117846116e7565b8160405283815286602085880101111561179c575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a086880312156117cf575f80fd5b6117d886611601565b94506117e660208701611601565b93506040860135925060608601359150608086013567ffffffffffffffff81111561180f575f80fd5b61181b88828901611714565b9150509295509295909350565b5f805f806080858703121561183b575f80fd5b61184485611601565b935060208501359250604085013567ffffffffffffffff811115611866575f80fd5b61187287828801611714565b949793965093946060013593505050565b5f60208284031215611893575f80fd5b81358015158114611642575f80fd5b5f602082840312156118b2575f80fd5b5051919050565b5f5b838110156118d35781810151838201526020016118bb565b50505f910152565b5f81518084526118f28160208601602086016118b9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b848152836020820152826040820152608060608201525f61194860808301846118db565b9695505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015283606083015260a0608083015261199660a08301846118db565b979650505050505050565b5f82516119b28184602087016118b9565b919091019291505056fea2646970667358221220096a0602d9334ed30477f78291328baca386e8dd8db2c744d816da6e457563b164736f6c63430008180033"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
