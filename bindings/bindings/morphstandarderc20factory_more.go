// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const MorphStandardERC20FactoryStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var MorphStandardERC20FactoryStorageLayout = new(solc.StorageLayout)

var MorphStandardERC20FactoryDeployedBin = "0x608060405234801561000f575f80fd5b506004361061006f575f3560e01c80637bdbcbbf1161004d5780637bdbcbbf146100d95780638da5cb5b146100ec578063f2fde38b14610109575f80fd5b80635c60da1b1461007357806361e98ca1146100bc578063715018a6146100cf575b5f80fd5b6001546100939073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100936100ca3660046105b1565b61011c565b6100d7610158565b005b6100936100e73660046105b1565b61016b565b5f5473ffffffffffffffffffffffffffffffffffffffff16610093565b6100d76101173660046105e2565b61020b565b5f8061012884846102c7565b60015490915061014e9073ffffffffffffffffffffffffffffffffffffffff1682610372565b9150505b92915050565b6101606103d4565b6101695f610454565b565b5f6101746103d4565b5f61017f84846102c7565b6001549091505f906101a79073ffffffffffffffffffffffffffffffffffffffff16836104c8565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f07ab516ad4f19b4465f15fa7c2dbc064f18e734a0846d6e0932da244aa3d8a7160405160405180910390a3949350505050565b6102136103d4565b73ffffffffffffffffffffffffffffffffffffffff81166102bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c481610454565b50565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606083901b1660208201525f9083906034016040516020818303038152906040528051906020012060405160200161035492919060609290921b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000168252601482015260340190565b60405160208183030381529060405280519060200120905092915050565b6040513060388201526f5af43d82803e903d91602b57fd5bf3ff602482015260148101839052733d602d80600a3d3981f3363d3d373d3d3d363d738152605881018290526037600c820120607882015260556043909101205f905b9392505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610169576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b2565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b1760205281603760095ff5905073ffffffffffffffffffffffffffffffffffffffff8116610152576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f455243313136373a2063726561746532206661696c656400000000000000000060448201526064016102b2565b803573ffffffffffffffffffffffffffffffffffffffff811681146105ac575f80fd5b919050565b5f80604083850312156105c2575f80fd5b6105cb83610589565b91506105d960208401610589565b90509250929050565b5f602082840312156105f2575f80fd5b6103cd8261058956fea26469706673582212204b00fba924982df4bb20a965559191acffef58ba7b2a61e3fec619c4a45a6edb64736f6c63430008180033"

func init() {
	if err := json.Unmarshal([]byte(MorphStandardERC20FactoryStorageLayoutJSON), MorphStandardERC20FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["MorphStandardERC20Factory"] = MorphStandardERC20FactoryStorageLayout
	deployedBytecodes["MorphStandardERC20Factory"] = MorphStandardERC20FactoryDeployedBin
}
