// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1024_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1023_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1023_storage\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1023_storage\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMessageSender\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_address\"},{\"astId\":1010,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"feeVault\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_address\"},{\"astId\":1012,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__rateLimiter\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1013,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_array(t_uint256)1022_storage\"},{\"astId\":1014,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"finalizedWithdrawals\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1015,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"messageSendTimestamp\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1016,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"isL1MessageDropped\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1017,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"rollup\",\"offset\":0,\"slot\":\"254\",\"type\":\"t_address\"},{\"astId\":1018,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"messageQueue\",\"offset\":0,\"slot\":\"255\",\"type\":\"t_address\"},{\"astId\":1019,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"maxReplayTimes\",\"offset\":0,\"slot\":\"256\",\"type\":\"t_uint256\"},{\"astId\":1020,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"replayStates\",\"offset\":0,\"slot\":\"257\",\"type\":\"t_mapping(t_bytes32,t_struct(ReplayState)1025_storage)\"},{\"astId\":1021,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"prevReplayIndex\",\"offset\":0,\"slot\":\"258\",\"type\":\"t_mapping(t_uint256,t_uint256)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1022_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\"},\"t_array(t_uint256)1023_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_struct(ReplayState)1025_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e struct L1CrossDomainMessenger.ReplayState)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_struct(ReplayState)1025_storage\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_struct(ReplayState)1025_storage\":{\"encoding\":\"inplace\",\"label\":\"struct L1CrossDomainMessenger.ReplayState\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x60806040526004361061019c575f3560e01c806386fa4b73116100dc578063bedb86fb11610087578063e70fc93b11610062578063e70fc93b1461052b578063ea7ec51414610556578063ecc7042814610582578063f2fde38b14610596575f80fd5b8063bedb86fb146104c1578063c0c53b8b146104e0578063cb23bcb5146104ff575f80fd5b8063a14238e7116100b7578063a14238e714610452578063b2267a7b14610480578063b604bf4c14610493575f80fd5b806386fa4b73146103e55780638da5cb5b14610404578063946130d81461042e575f80fd5b806355004105116101475780636e296e45116101225780636e296e45146102fa578063715018a614610326578063797594b01461033a578063846d4d7a14610366575f80fd5b806355004105146102bd5780635c975abb146102d05780635f7b1577146102e7575f80fd5b80633b70c18a116101775780633b70c18a14610221578063407c195514610272578063478222c214610291575f80fd5b806329907acd146101af5780632a6cccb2146101ce578063340735f7146101ed575f80fd5b366101ab576101a96105b5565b005b5f80fd5b3480156101ba575f80fd5b506101a96101c9366004612c27565b61063d565b3480156101d9575f80fd5b506101a96101e8366004612c94565b610a38565b3480156101f8575f80fd5b5061020c610207366004612ccb565b610ac7565b60405190151581526020015b60405180910390f35b34801561022c575f80fd5b5060ff5461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610218565b34801561027d575f80fd5b506101a961028c366004612d07565b610b90565b34801561029c575f80fd5b5060cb5461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b6101a96102cb366004612d1e565b610bd7565b3480156102db575f80fd5b5060655460ff1661020c565b6101a96102f5366004612db5565b6112d2565b348015610305575f80fd5b5060c95461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610331575f80fd5b506101a9611326565b348015610345575f80fd5b5060ca5461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610371575f80fd5b506103bc610380366004612d07565b6101016020525f90815260409020546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041682565b604080516fffffffffffffffffffffffffffffffff938416815292909116602083015201610218565b3480156103f0575f80fd5b506101a96103ff366004612e53565b611337565b34801561040f575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661024d565b348015610439575f80fd5b506104446101005481565b604051908152602001610218565b34801561045d575f80fd5b5061020c61046c366004612d07565b60fb6020525f908152604090205460ff1681565b6101a961048e366004612edc565b6118fe565b34801561049e575f80fd5b5061020c6104ad366004612d07565b60fd6020525f908152604090205460ff1681565b3480156104cc575f80fd5b506101a96104db366004612f44565b611919565b3480156104eb575f80fd5b506101a96104fa366004612f5f565b61193a565b34801561050a575f80fd5b5060fe5461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610536575f80fd5b50610444610545366004612d07565b60fc6020525f908152604090205481565b348015610561575f80fd5b50610444610570366004612d07565b6101026020525f908152604090205481565b34801561058d575f80fd5b50610444611c0a565b3480156105a1575f80fd5b506101a96105b0366004612c94565b611ca0565b60335473ffffffffffffffffffffffffffffffffffffffff16331461063b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b610645611d54565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead146106c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e006044820152606401610632565b60ff5473ffffffffffffffffffffffffffffffffffffffff165f6106ef8787878787611dc1565b90505f818051906020012090505f60fc5f8381526020019081526020015f20541161079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f71756575656400000000000000000000000000000000000000000000000000006064820152608401610632565b5f81815260fd602052604090205460ff1615610814576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d65737361676520616c72656164792064726f707065640000000000000000006044820152606401610632565b5f818152610101602052604081205470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16908190036108545750845b6040517f916524610000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff8516906391652461906024015f604051808303815f87803b1580156108b9575f80fd5b505af11580156108cb573d5f803e3d5ffd5b5050505f918252506101026020526040902054801561090b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610854565b5f82815260fd60205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560c980547fffffffffffffffffffffffff000000000000000000000000000000000000000016736f297c61b5c92ef107ffd30cd56affe5a273e841179055517f14298c5100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16906314298c519089906109d390899060040161300a565b5f604051808303818588803b1580156109ea575f80fd5b505af11580156109fc573d5f803e3d5ffd5b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055505050505050505050505050565b610a406105b5565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f591015b60405180910390a15050565b5f84815b6020811015610b84578085901c600116600103610b3157858160208110610af457610af461301c565b602002013582604051602001610b14929190918252602082015260400190565b604051602081830303815290604052805190602001209150610b7c565b81868260208110610b4457610b4461301c565b6020020135604051602001610b63929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b600101610acb565b50909114949350505050565b610b986105b5565b61010080549082905560408051828152602081018490527fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b9101610abb565b610bdf611d54565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead14610c62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e006044820152606401610632565b60ff5460ca5473ffffffffffffffffffffffffffffffffffffffff91821691165f610c908a8a8a8a8a611dc1565b90505f818051906020012090505f60fc5f8381526020019081526020015f205411610d3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f71756575656400000000000000000000000000000000000000000000000000006064820152608401610632565b5f81815260fd602052604090205460ff1615610db5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d65737361676520616c72656164792064726f707065640000000000000000006044820152606401610632565b6040517fd7704bae00000000000000000000000000000000000000000000000000000000815263ffffffff871660048201525f9073ffffffffffffffffffffffffffffffffffffffff86169063d7704bae90602401602060405180830381865afa158015610e25573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e499190613049565b905080341015610eb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e73756666696369656e74206d73672e76616c756520666f722066656500006044820152606401610632565b8015610f845760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114610f12576040519150601f19603f3d011682016040523d82523d5f602084013e610f17565b606091505b5050905080610f82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20646564756374207468652066656500000000000000006044820152606401610632565b505b5f8573ffffffffffffffffffffffffffffffffffffffff1663fd0ad31e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fce573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ff29190613049565b6040517f9b15978200000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff871690639b1597829061104b9088908c908990600401613060565b5f604051808303815f87803b158015611062575f80fd5b505af1158015611074573d5f803e3d5ffd5b5050505f848152610101602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416918101829052925090036110e4575f8281526101026020526040902060018c019055611115565b80602001516001016fffffffffffffffffffffffffffffffff166101025f8481526020019081526020015f20819055505b6fffffffffffffffffffffffffffffffff80831660208301526101005482519091161061119e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f457863656564206d6178696d756d207265706c61792074696d657300000000006044820152606401610632565b80516fffffffffffffffffffffffffffffffff600191909101811682525f8581526101016020908152604090912083519184015183167001000000000000000000000000000000000291909216179055348381039084146112c1575f8973ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f811461124f576040519150601f19603f3d011682016040523d82523d5f602084013e611254565b606091505b50509050806112bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20726566756e64207468652066656500000000000000006044820152606401610632565b505b505050505050505050505050505050565b6112da611d54565b61131e868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250879150611e5d9050565b505050505050565b61132e6105b5565b61063b5f612359565b61133f611d54565b60c95473ffffffffffffffffffffffffffffffffffffffff1661dead146113c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e006044820152606401610632565b60ff5473ffffffffffffffffffffffffffffffffffffffff9081169087160361146d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d657373656e6765723a20466f7262696420746f2063616c6c206d657373616760448201527f65207175657565000000000000000000000000000000000000000000000000006064820152608401610632565b611476866123cf565b60c95473ffffffffffffffffffffffffffffffffffffffff90811690881603611521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4d657373656e6765723a20496e76616c6964206d6573736167652073656e646560448201527f72000000000000000000000000000000000000000000000000000000000000006064820152608401610632565b5f61152f8888888888611dc1565b80516020918201205f81815260fb90925260409091205490915060ff16156115d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4d657373656e6765723a207769746864726177616c2068617320616c7265616460448201527f79206265656e2066696e616c697a6564000000000000000000000000000000006064820152608401610632565b60fe546040517f04d772150000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff909116905f9082906304d7721590602401602060405180830381865afa15801561164a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061166e91906130a3565b9050806116fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f4d657373656e6765723a207769746864726177616c526f6f74206e6f7420666960448201527f6e616c697a6564000000000000000000000000000000000000000000000000006064820152608401610632565b61170983868987610ac7565b611795576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d657373656e6765723a20696e76616c6964207769746864726177616c20696e60448201527f636c7573696f6e2070726f6f66000000000000000000000000000000000000006064820152608401610632565b5060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8b8116919091179091556040515f918a169089906117f09089906130be565b5f6040518083038185875af1925050503d805f811461182a576040519150601f19603f3d011682016040523d82523d5f602084013e61182f565b606091505b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156118c7575f83815260fb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26118f2565b60405183907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a25b50505050505050505050565b611906611d54565b6119138484848433611e5d565b50505050565b6119216105b5565b80156119325761192f61244e565b50565b61192f6124d3565b5f54610100900460ff161580801561195857505f54600160ff909116105b806119715750303b15801561197157505f5460ff166001145b6119fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610632565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611a59575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff83161580611a90575073ffffffffffffffffffffffffffffffffffffffff8216155b15611ac7576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ae57353000000000000000000000000000000000000018561252a565b60fe805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560ff80549285169282169290921790915560ca805490911673530000000000000000000000000000000000000717905560036101008190556040517fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b91611b9b915f9190918252602082015260400190565b60405180910390a18015611913575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b60ff54604080517ffd0ad31e00000000000000000000000000000000000000000000000000000000815290515f9273ffffffffffffffffffffffffffffffffffffffff169163fd0ad31e9160048083019260209291908290030181865afa158015611c77573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c9b9190613049565b905090565b611ca86105b5565b73ffffffffffffffffffffffffffffffffffffffff8116611d4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610632565b61192f81612359565b60655460ff161561063b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610632565b60608585858585604051602401611ddc9594939291906130d9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b611e65612676565b60ff5460ca54604080517ffd0ad31e000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff93841693909216915f91849163fd0ad31e916004808201926020929091908290030181865afa158015611ede573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f029190613049565b90505f611f12338a8a858b611dc1565b6040517fd7704bae000000000000000000000000000000000000000000000000000000008152600481018890529091505f9073ffffffffffffffffffffffffffffffffffffffff86169063d7704bae90602401602060405180830381865afa158015611f80573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fa49190613049565b9050611fb08982613128565b341015612019576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e74206d73672e76616c7565000000000000000000006044820152606401610632565b80156120e85760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114612076576040519150601f19603f3d011682016040523d82523d5f602084013e61207b565b606091505b50509050806120e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20646564756374207468652066656500000000000000006044820152606401610632565b505b6040517f9b15978200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff861690639b1597829061213e9087908b908790600401613160565b5f604051808303815f87803b158015612155575f80fd5b505af1158015612167573d5f803e3d5ffd5b505050505f8280519060200120905060fc5f8281526020019081526020015f20545f146121f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4475706c696361746564206d65737361676500000000000000000000000000006044820152606401610632565b5f81815260fc6020526040902042905573ffffffffffffffffffffffffffffffffffffffff8b163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8c878c8e6040516122639493929190613194565b60405180910390a3348290038a8103908b14612341575f8873ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146122cf576040519150601f19603f3d011682016040523d82523d5f602084013e6122d4565b606091505b505090508061233f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4661696c656420746f20726566756e64207468652066656500000000000000006044820152606401610632565b505b505050505050506123526001609755565b5050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b3073ffffffffffffffffffffffffffffffffffffffff82160361192f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d657373656e6765723a20466f7262696420746f2063616c6c2073656c6600006044820152606401610632565b612456611d54565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586124a93390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6124db6126f0565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336124a9565b5f54610100900460ff166125c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b6125c861275c565b6125d06127fa565b6125d8612898565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661dead1790915560ca805473ffffffffffffffffffffffffffffffffffffffff858116919093161790558116156126725760cb80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5050565b6002609754036126e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610632565b6002609755565b6001609755565b60655460ff1661063b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610632565b5f54610100900460ff166127f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b61063b612936565b5f54610100900460ff16612890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b61063b6129d5565b5f54610100900460ff1661292e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b61063b612a95565b5f54610100900460ff166129cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b61063b33612359565b5f54610100900460ff16612a6b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b5f54610100900460ff166126e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610632565b803573ffffffffffffffffffffffffffffffffffffffff81168114612b4e575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612b8f575f80fd5b813567ffffffffffffffff80821115612baa57612baa612b53565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612bf057612bf0612b53565b81604052838152866020858801011115612c08575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f60a08688031215612c3b575f80fd5b612c4486612b2b565b9450612c5260208701612b2b565b93506040860135925060608601359150608086013567ffffffffffffffff811115612c7b575f80fd5b612c8788828901612b80565b9150509295509295909350565b5f60208284031215612ca4575f80fd5b612cad82612b2b565b9392505050565b806104008101831015612cc5575f80fd5b92915050565b5f805f806104608587031215612cdf575f80fd5b84359350612cf08660208701612cb4565b939693955050505061042082013591610440013590565b5f60208284031215612d17575f80fd5b5035919050565b5f805f805f805f60e0888a031215612d34575f80fd5b612d3d88612b2b565b9650612d4b60208901612b2b565b95506040880135945060608801359350608088013567ffffffffffffffff811115612d74575f80fd5b612d808a828b01612b80565b93505060a088013563ffffffff81168114612d99575f80fd5b9150612da760c08901612b2b565b905092959891949750929550565b5f805f805f8060a08789031215612dca575f80fd5b612dd387612b2b565b955060208701359450604087013567ffffffffffffffff80821115612df6575f80fd5b818901915089601f830112612e09575f80fd5b813581811115612e17575f80fd5b8a6020828501011115612e28575f80fd5b60208301965080955050505060608701359150612e4760808801612b2b565b90509295509295509295565b5f805f805f805f6104c0888a031215612e6a575f80fd5b612e7388612b2b565b9650612e8160208901612b2b565b95506040880135945060608801359350608088013567ffffffffffffffff811115612eaa575f80fd5b612eb68a828b01612b80565b935050612ec68960a08a01612cb4565b91506104a0880135905092959891949750929550565b5f805f8060808587031215612eef575f80fd5b612ef885612b2b565b935060208501359250604085013567ffffffffffffffff811115612f1a575f80fd5b612f2687828801612b80565b949793965093946060013593505050565b801515811461192f575f80fd5b5f60208284031215612f54575f80fd5b8135612cad81612f37565b5f805f60608486031215612f71575f80fd5b612f7a84612b2b565b9250612f8860208501612b2b565b9150612f9660408501612b2b565b90509250925092565b5f5b83811015612fb9578181015183820152602001612fa1565b50505f910152565b5f8151808452612fd8816020860160208601612f9f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f612cad6020830184612fc1565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f60208284031215613059575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201525f61309a6060830184612fc1565b95945050505050565b5f602082840312156130b3575f80fd5b8151612cad81612f37565b5f82516130cf818460208701612f9f565b9190910192915050565b5f73ffffffffffffffffffffffffffffffffffffffff808816835280871660208401525084604083015283606083015260a0608083015261311d60a0830184612fc1565b979650505050505050565b80820180821115612cc5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f61309a6060830184612fc1565b848152836020820152826040820152608060608201525f6131b86080830184612fc1565b969550505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
