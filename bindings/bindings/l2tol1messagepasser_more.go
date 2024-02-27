// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const L2ToL1MessagePasserStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var L2ToL1MessagePasserStorageLayout = new(solc.StorageLayout)

var L2ToL1MessagePasserDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610064575f3560e01c806389c09d381161004d57806389c09d38146100b1578063b58343bb146100b9578063d4b9f4fa146100c2575f80fd5b8063340735f714610068578063600a2e7714610090575b5f80fd5b61007b61007636600461049e565b6100cb565b60405190151581526020015b60405180910390f35b6100a361009e366004610532565b610194565b604051908152602001610087565b6100a361026e565b6100a360205481565b6100a360215481565b5f84815b6020811015610188578085901c600116600103610135578581602081106100f8576100f8610549565b602002015182604051602001610118929190918252602082015260400190565b604051602081830303815290604052805190602001209150610180565b8186826020811061014857610148610549565b6020020151604051602001610167929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b6001016100cf565b50909114949350505050565b5f3373530000000000000000000000000000000000000714610216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f6f6e6c79206d657373656e676572000000000000000000000000000000000000604482015260640160405180910390fd5b60208054604080519182529181018490527ffaa617c2d8ce12c62637dbce76efcc18dae60574aa95709bdcedce7e76071693910160405180910390a161025b8261035e565b61026361026e565b602181905592915050565b6020545f90819081805b6020811015610355578083901c6001166001036102d4575f81602081106102a1576102a1610549565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610301565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b6040805160208101849052908101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209150600101610278565b50919392505050565b80600161036d602060026106c3565b61037791906106d5565b602054106103b1576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60205f81546103c0906106e8565b918290555090505f5b6020811015610463578082901c6001166001036103fb57825f82602081106103f3576103f3610549565b015550505050565b5f816020811061040d5761040d610549565b01546040805160208101929092528101849052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012092506001016103c9565b5061046c61071f565b505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f8061046085870312156104b2575f80fd5b84359350602086603f8701126104c6575f80fd5b604051610400810181811067ffffffffffffffff821117156104ea576104ea610471565b604052806104208801898111156104ff575f80fd5b602089015b8181101561051b5780358352918401918401610504565b509699919850509435956104400135949350505050565b5f60208284031215610542575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b600181815b808511156105fc57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156105e2576105e2610576565b808516156105ef57918102915b93841c93908002906105a8565b509250929050565b5f82610612575060016106bd565b8161061e57505f6106bd565b8160018114610634576002811461063e5761065a565b60019150506106bd565b60ff84111561064f5761064f610576565b50506001821b6106bd565b5060208310610133831016604e8410600b841016171561067d575081810a6106bd565b61068783836105a3565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156106b9576106b9610576565b0290505b92915050565b5f6106ce8383610604565b9392505050565b818103818111156106bd576106bd610576565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361071857610718610576565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffdfea26469706673582212206c4baa4ca5b521e2d8063e55d878cbf838f5c71fbe10cf62886f40315f183ced64736f6c63430008180033"

func init() {
	if err := json.Unmarshal([]byte(L2ToL1MessagePasserStorageLayoutJSON), L2ToL1MessagePasserStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL1MessagePasser"] = L2ToL1MessagePasserStorageLayout
	deployedBytecodes["L2ToL1MessagePasser"] = L2ToL1MessagePasserDeployedBin
}
