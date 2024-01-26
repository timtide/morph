// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1WETHGatewayStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"_status\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_address\"},{\"astId\":1006,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1014_storage\"},{\"astId\":1007,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1008,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"router\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_array(t_uint256)1013_storage\"},{\"astId\":1012,\"contract\":\"contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_array(t_uint256)1015_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1WETHGatewayStorageLayout = new(solc.StorageLayout)

var L1WETHGatewayDeployedBin = "0x6080604052600436106100ec5760003560e01c8063885586871161008a578063c676ad2911610059578063c676ad2914610333578063f219fa6614610373578063f2fde38b14610386578063f887ea40146103a657600080fd5b806388558687146102805780638da5cb5b146102b4578063ad5c4648146102df578063c0c53b8b1461031357600080fd5b80633cb747bf116100c65780633cb747bf146101d5578063715018a61461022b578063797594b01461024057806384bd13b01461026d57600080fd5b80630aea8c261461019c57806314298c51146101af57806321425ee0146101c257600080fd5b3661019757337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610195576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6f6e6c792057455448000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b005b600080fd5b6101956101aa3660046122d4565b6103d3565b6101956101bd366004612391565b6103e7565b6101956101d03660046123d3565b610714565b3480156101e157600080fd5b506099546102029073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561023757600080fd5b5061019561074e565b34801561024c57600080fd5b506097546102029073ffffffffffffffffffffffffffffffffffffffff1681565b61019561027b366004612408565b610762565b34801561028c57600080fd5b506102027f000000000000000000000000000000000000000000000000000000000000000081565b3480156102c057600080fd5b5060655473ffffffffffffffffffffffffffffffffffffffff16610202565b3480156102eb57600080fd5b506102027f000000000000000000000000000000000000000000000000000000000000000081565b34801561031f57600080fd5b5061019561032e3660046124a0565b6109e4565b34801561033f57600080fd5b5061020261034e3660046124eb565b507f000000000000000000000000000000000000000000000000000000000000000090565b61019561038136600461250f565b610bf9565b34801561039257600080fd5b506101956103a13660046124eb565b610c06565b3480156103b257600080fd5b506098546102029073ffffffffffffffffffffffffffffffffffffffff1681565b6103e08585858585610cbd565b5050505050565b60995473ffffffffffffffffffffffffffffffffffffffff16338114610469576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161018c565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d89190612555565b73ffffffffffffffffffffffffffffffffffffffff16736f297c61b5c92ef107ffd30cd56affe5a273e84173ffffffffffffffffffffffffffffffffffffffff1614610580576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f6f6e6c792063616c6c656420696e2064726f7020636f6e746578740000000000604482015260640161018c565b61058861108e565b7f8431f5c1000000000000000000000000000000000000000000000000000000006105b7600460008587612572565b6105c09161259c565b7fffffffff000000000000000000000000000000000000000000000000000000001614610649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f696e76616c69642073656c6563746f7200000000000000000000000000000000604482015260640161018c565b6000808061065a8560048189612572565b81019061066791906125e4565b50945050935050925061067b838383611101565b61069c73ffffffffffffffffffffffffffffffffffffffff8416838361127b565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fdbdf8eb487847e4c0f22847f5dac07f2d3690f96f581a6ae4b102769917645a8836040516106fb91815260200190565b60405180910390a350505061070f60018055565b505050565b61070f83338460005b6040519080825280601f01601f191660200182016040528015610747576020820181803683370190505b5085610cbd565b610756611355565b61076060006113d6565b565b60995473ffffffffffffffffffffffffffffffffffffffff163381146107e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f6f6e6c79206d657373656e6765722063616e2063616c6c000000000000000000604482015260640161018c565b8073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561082f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108539190612555565b60975473ffffffffffffffffffffffffffffffffffffffff9081169116146108d7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792063616c6c20627920636f756e746572706172740000000000000000604482015260640161018c565b6108df61108e565b6108ee8888888888888861144d565b61090f73ffffffffffffffffffffffffffffffffffffffff8916868661127b565b61094f8584848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061168a92505050565b8573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff167fc6f985873b37805705f6bce756dce3d1ff4b603e298d506288cce499926846a7888888886040516109c99493929190612674565b60405180910390a46109da60018055565b5050505050505050565b600054610100900460ff1615808015610a045750600054600160ff909116105b80610a1e5750303b158015610a1e575060005460ff166001145b610aaa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161018c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610b0857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610b85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f7a65726f20726f75746572206164647265737300000000000000000000000000604482015260640161018c565b610b90848484611740565b8015610bf357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610bf3848484600061071d565b610c0e611355565b73ffffffffffffffffffffffffffffffffffffffff8116610cb1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161018c565b610cba816113d6565b50565b610cc561108e565b60008311610d2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e7400000000000000000000000000604482015260640161018c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610de4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6f6e6c79205745544820697320616c6c6f776564000000000000000000000000604482015260640161018c565b6000610df18685856118eb565b6040517f2e1a7d4d00000000000000000000000000000000000000000000000000000000815260048101839052919650945090915073ffffffffffffffffffffffffffffffffffffffff871690632e1a7d4d90602401600060405180830381600087803b158015610e6157600080fd5b505af1158015610e75573d6000803e3d6000fd5b505050506000867f000000000000000000000000000000000000000000000000000000000000000083888888604051602401610eb696959493929190612753565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8431f5c10000000000000000000000000000000000000000000000000000000017905260995490915073ffffffffffffffffffffffffffffffffffffffff16635f7b1577610f5734886127dd565b6097546040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152610fb19173ffffffffffffffffffffffffffffffffffffffff16908a9087908a908a906004016127f6565b6000604051808303818588803b158015610fca57600080fd5b505af1158015610fde573d6000803e3d6000fd5b50505050508173ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f31cd3b976e4d654022bf95c68a2ce53f1d5d94afabe0454d2832208eeb40af2589898960405161107b93929190612846565b60405180910390a450506103e060018055565b6002600154036110fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161018c565b6002600155565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146111b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f746f6b656e206e6f742057455448000000000000000000000000000000000000604482015260640161018c565b34811461121f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d73672e76616c7565206d69736d617463680000000000000000000000000000604482015260640161018c565b8273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561126757600080fd5b505af11580156109da573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff831660248201526044810182905261070f9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611ba7565b60018055565b60655473ffffffffffffffffffffffffffffffffffffffff163314610760576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161018c565b6065805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614611502576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3120746f6b656e206e6f742057455448000000000000000000000000000000604482015260640161018c565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146115b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6c3220746f6b656e206e6f742057455448000000000000000000000000000000604482015260640161018c565b348314611620576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f6d73672e76616c7565206d69736d617463680000000000000000000000000000604482015260640161018c565b8673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561166857600080fd5b505af115801561167c573d6000803e3d6000fd5b505050505050505050505050565b600081511180156116b2575060008273ffffffffffffffffffffffffffffffffffffffff163b115b1561173c576040517f444b281f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063444b281f90611709908490600401612884565b600060405180830381600087803b15801561172357600080fd5b505af1158015611737573d6000803e3d6000fd5b505050505b5050565b73ffffffffffffffffffffffffffffffffffffffff83166117bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f7a65726f20636f756e7465727061727420616464726573730000000000000000604482015260640161018c565b73ffffffffffffffffffffffffffffffffffffffff811661183a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f7a65726f206d657373656e676572206164647265737300000000000000000000604482015260640161018c565b611842611cb6565b61184a611d55565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560998054848416921691909117905582161561070f576098805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b60985460009081906060903390819073ffffffffffffffffffffffffffffffffffffffff168190036119d9578580602001905181019061192b9190612897565b6040517fc52a3bbc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301528b81166024830152604482018b905291985091925083169063c52a3bbc906064016020604051808303816000875af11580156119ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119d29190612924565b9650611b30565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8a16906370a0823190602401602060405180830381865afa158015611a46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a6a9190612924565b9050611a8e73ffffffffffffffffffffffffffffffffffffffff8a1683308b611df4565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8b16906370a0823190602401602060405180830381865afa158015611afb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b1f9190612924565b9050611b2b828261293d565b985050505b60008711611b9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f6465706f736974207a65726f20616d6f756e7400000000000000000000000000604482015260640161018c565b9795965093949350505050565b6000611c09826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611e529092919063ffffffff16565b9050805160001480611c2a575080806020019051810190611c2a9190612950565b61070f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161018c565b600054610100900460ff16611d4d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161018c565b610760611e69565b600054610100900460ff16611dec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161018c565b610760611f00565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610bf39085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016112cd565b6060611e618484600085611fa0565b949350505050565b600054610100900460ff1661134f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161018c565b600054610100900460ff16611f97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161018c565b610760336113d6565b606082471015612032576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161018c565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161205b9190612972565b60006040518083038185875af1925050503d8060008114612098576040519150601f19603f3d011682016040523d82523d6000602084013e61209d565b606091505b50915091506120ae878383876120b9565b979650505050505050565b6060831561214f5782516000036121485773ffffffffffffffffffffffffffffffffffffffff85163b612148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161018c565b5081611e61565b611e6183838151156121645781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018c9190612884565b73ffffffffffffffffffffffffffffffffffffffff81168114610cba57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612230576122306121ba565b604052919050565b600067ffffffffffffffff821115612252576122526121ba565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261228f57600080fd5b81356122a261229d82612238565b6121e9565b8181528460208386010111156122b757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a086880312156122ec57600080fd5b85356122f781612198565b9450602086013561230781612198565b935060408601359250606086013567ffffffffffffffff81111561232a57600080fd5b6123368882890161227e565b95989497509295608001359392505050565b60008083601f84011261235a57600080fd5b50813567ffffffffffffffff81111561237257600080fd5b60208301915083602082850101111561238a57600080fd5b9250929050565b600080602083850312156123a457600080fd5b823567ffffffffffffffff8111156123bb57600080fd5b6123c785828601612348565b90969095509350505050565b6000806000606084860312156123e857600080fd5b83356123f381612198565b95602085013595506040909401359392505050565b600080600080600080600060c0888a03121561242357600080fd5b873561242e81612198565b9650602088013561243e81612198565b9550604088013561244e81612198565b9450606088013561245e81612198565b93506080880135925060a088013567ffffffffffffffff81111561248157600080fd5b61248d8a828b01612348565b989b979a50959850939692959293505050565b6000806000606084860312156124b557600080fd5b83356124c081612198565b925060208401356124d081612198565b915060408401356124e081612198565b809150509250925092565b6000602082840312156124fd57600080fd5b813561250881612198565b9392505050565b6000806000806080858703121561252557600080fd5b843561253081612198565b9350602085013561254081612198565b93969395505050506040820135916060013590565b60006020828403121561256757600080fd5b815161250881612198565b6000808585111561258257600080fd5b8386111561258f57600080fd5b5050820193919092039150565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156125dc5780818660040360031b1b83161692505b505092915050565b60008060008060008060c087890312156125fd57600080fd5b863561260881612198565b9550602087013561261881612198565b9450604087013561262881612198565b9350606087013561263881612198565b92506080870135915060a087013567ffffffffffffffff81111561265b57600080fd5b61266789828a0161227e565b9150509295509295509295565b73ffffffffffffffffffffffffffffffffffffffff8516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01601019392505050565b60005b838110156127005781810151838201526020016126e8565b50506000910152565b600081518084526127218160208601602086016126e5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a08301526127a260c0830184612709565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156127f0576127f06127ae565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015260a0604084015261282c60a0840187612709565b606084019590955292909216608090910152509392505050565b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061287b6060830184612709565b95945050505050565b6020815260006125086020830184612709565b600080604083850312156128aa57600080fd5b82516128b581612198565b602084015190925067ffffffffffffffff8111156128d257600080fd5b8301601f810185136128e357600080fd5b80516128f161229d82612238565b81815286602083850101111561290657600080fd5b6129178260208301602086016126e5565b8093505050509250929050565b60006020828403121561293657600080fd5b5051919050565b818103818111156127f0576127f06127ae565b60006020828403121561296257600080fd5b8151801515811461250857600080fd5b600082516129848184602087016126e5565b919091019291505056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(L1WETHGatewayStorageLayoutJSON), L1WETHGatewayStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1WETHGateway"] = L1WETHGatewayStorageLayout
	deployedBytecodes["L1WETHGateway"] = L1WETHGatewayDeployedBin
}
