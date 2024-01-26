// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const AddressManagerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/deployment/AddressManager.sol:AddressManager\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/deployment/AddressManager.sol:AddressManager\",\"label\":\"addresses\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_address)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_address\"}}}"

var AddressManagerStorageLayout = new(solc.StorageLayout)

var AddressManagerDeployedBin = "0x608060405234801561001057600080fd5b50600436106100675760003560e01c80639b2ea4bd116100505780639b2ea4bd146100b9578063bf40fac1146100cc578063f2fde38b146100df57600080fd5b8063715018a61461006c5780638da5cb5b14610076575b600080fd5b6100746100f2565b005b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100746100c73660046104fa565b610106565b6100906100da366004610548565b6101d9565b6100746100ed366004610585565b610215565b6100fa6102d1565b6101046000610352565b565b61010e6102d1565b6000610119836103c7565b60008181526001602052604090819020805473ffffffffffffffffffffffffffffffffffffffff8681167fffffffffffffffffffffffff00000000000000000000000000000000000000008316179092559151929350169061017c9085906105a7565b6040805191829003822073ffffffffffffffffffffffffffffffffffffffff808716845284166020840152917f9416a153a346f93d95f94b064ae3f148b6460473c6e82b3f9fc2521b873fcd6c910160405180910390a250505050565b6000600160006101e8846103c7565b815260208101919091526040016000205473ffffffffffffffffffffffffffffffffffffffff1692915050565b61021d6102d1565b73ffffffffffffffffffffffffffffffffffffffff81166102c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102ce81610352565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610104576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102bc565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000816040516020016103da91906105a7565b604051602081830303815290604052805190602001209050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261043757600080fd5b813567ffffffffffffffff80821115610452576104526103f7565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610498576104986103f7565b816040528381528660208588010111156104b157600080fd5b836020870160208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146104f557600080fd5b919050565b6000806040838503121561050d57600080fd5b823567ffffffffffffffff81111561052457600080fd5b61053085828601610426565b92505061053f602084016104d1565b90509250929050565b60006020828403121561055a57600080fd5b813567ffffffffffffffff81111561057157600080fd5b61057d84828501610426565b949350505050565b60006020828403121561059757600080fd5b6105a0826104d1565b9392505050565b6000825160005b818110156105c857602081860181015185830152016105ae565b50600092019182525091905056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(AddressManagerStorageLayoutJSON), AddressManagerStorageLayout); err != nil {
		panic(err)
	}

	layouts["AddressManager"] = AddressManagerStorageLayout
	deployedBytecodes["AddressManager"] = AddressManagerDeployedBin
}
