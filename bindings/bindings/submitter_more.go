// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/morph-l2/bindings/solc"
)

const SubmitterStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"nextBatchIndex\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"nextBatchStartBlock\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"confirmedBatchs\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_uint256,t_struct(BatchInfo)1009_storage)\"},{\"astId\":1005,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"nextEpochStart\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"nextSubmitterIndex\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"calculatedEpochIndex\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/L2/submitter/Submitter.sol:Submitter\",\"label\":\"epochs\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_uint256,t_struct(EpochInfo)1010_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_uint256,t_struct(BatchInfo)1009_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct Types.BatchInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(BatchInfo)1009_storage\"},\"t_mapping(t_uint256,t_struct(EpochInfo)1010_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct Types.EpochInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(EpochInfo)1010_storage\"},\"t_struct(BatchInfo)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.BatchInfo\",\"numberOfBytes\":\"128\"},\"t_struct(EpochInfo)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct Types.EpochInfo\",\"numberOfBytes\":\"96\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SubmitterStorageLayout = new(solc.StorageLayout)

var SubmitterDeployedBin = "0x608060405234801561001057600080fd5b50600436106101815760003560e01c8063843e8a7b116100d8578063bddd8e731161008c578063cc0f858f11610066578063cc0f858f1461044b578063e8e3992514610454578063fe4b84df146104d957600080fd5b8063bddd8e73146103f4578063c58159c4146103fc578063c6b61e4c1461040557600080fd5b8063965fbb94116100bd578063965fbb941461036e578063a5af40d114610381578063bc0bc6ba146103a957600080fd5b8063843e8a7b1461030d578063927ede2d1461034757600080fd5b80633cb747bf1161013a5780636cb23707116101145780636cb23707146102d457806373790ab3146102fb5780637e4fa7001461030457600080fd5b80633cb747bf1461028257806354fd4d50146102a85780635c14c314146102bd57600080fd5b8063151232581161016b57806315123258146101fe57806316e2994a1461025a57806322caba241461026f57600080fd5b80628dbdb514610186578063047d0b6e146101d7575b600080fd5b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b61021161020c366004611868565b6104ec565b6040516101ce9190815173ffffffffffffffffffffffffffffffffffffffff16815260208083015190820152604080830151908201526060918201519181019190915260800190565b61026d610268366004611948565b610585565b005b61026d61027d3660046119e7565b6106cf565b7f00000000000000000000000000000000000000000000000000000000000000006101ad565b6102b0610a52565b6040516101ce9190611a53565b6102c660025481565b6040519081526020016101ce565b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6102c660045481565b6102c660015481565b610315610af5565b6040805173ffffffffffffffffffffffffffffffffffffffff90941684526020840192909252908201526060016101ce565b6101ad7f000000000000000000000000000000000000000000000000000000000000000081565b61026d61037c366004611868565b610ce0565b61039461038f366004611aa4565b610e5f565b604080519283526020830191909152016101ce565b6103bc6103b7366004611868565b611163565b60408051825173ffffffffffffffffffffffffffffffffffffffff1681526020808401519082015291810151908201526060016101ce565b61026d6111ec565b6102c660055481565b610315610413366004611868565b60076020526000908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff909116919083565b6102c660065481565b6104a2610462366004611868565b6003602081905260009182526040909120805460018201546002830154929093015473ffffffffffffffffffffffffffffffffffffffff90911692919084565b6040805173ffffffffffffffffffffffffffffffffffffffff909516855260208501939093529183015260608201526080016101ce565b61026d6104e7366004611868565b61135f565b61052d6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b506000908152600360208181526040928390208351608081018552815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600281015493820193909352910154606082015290565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c79206c322073657175656e63657220636f6e747261637400000000000060448201526064015b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ba9190611ac8565b600060055590506106cb8183611558565b5050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107ed57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d59190611ae1565b73ffffffffffffffffffffffffffffffffffffffff16145b610879576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f526f6c6c75703a2066756e6374696f6e2063616e206f6e6c792062652063616c60448201527f6c65642066726f6d20746865204c3120726f6c6c7570000000000000000000006064820152608401610620565b60015485146108e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206261746368496e64657800000000000000000000000000006044820152606401610620565b600254831461094f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c69642062617463685374617274426c6f636b0000000000000000006044820152606401610620565b604080516080808201835273ffffffffffffffffffffffffffffffffffffffff878116808452602080850189815285870189815260608088018a815260008f81526003808752908b902099518a547fffffffffffffffffffffffff000000000000000000000000000000000000000016981697909717895592516001890155905160028801559051959093019490945584518a8152938401529282018690529181018490529081018290527f516afe1b5719e7236e4c39aa8d6b5972e973d975aff7f724eeba95abd343664f9060a00160405180910390a160018054906000610a3783611b2d565b90915550610a489050826001611b65565b6002555050505050565b6060610a7d7f00000000000000000000000000000000000000000000000000000000000000006116c7565b610aa67f00000000000000000000000000000000000000000000000000000000000000006116c7565b610acf7f00000000000000000000000000000000000000000000000000000000000000006116c7565b604051602001610ae193929190611b78565b604051602081830303815290604052905090565b6040517fe597c19e0000000000000000000000000000000000000000000000000000000081526000600482018190529081908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e90602401600060405180830381865afa158015610b89573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610bcf9190810190611bee565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c629190611ac8565b825160045460055492935090915b42610c7b8584611b65565b11610ca95780610c8a81611b2d565b915050828103610c98575060005b610ca28483611b65565b9150610c70565b848181518110610cbb57610cbb611c7d565b6020026020010151828584610cd09190611b65565b9750975097505050505050909192565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f6f6e6c7920676f7620636f6e74726163740000000000000000000000000000006044820152606401610620565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e90602401600060405180830381865afa158015610e0d573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610e539190810190611bee565b90506106cb8282611558565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052908190819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e597c19e90602401600060405180830381865afa158015610ef1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610f379190810190611bee565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fca9190611ac8565b82519091506000805b8281101561104257848181518110610fed57610fed611c7d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16036110305760019150611042565b8061103a81611b2d565b915050610fd3565b816110a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f696e76616c6964207375626d69747465720000000000000000000000000000006044820152606401610620565b6004546005545b426110bb8784611b65565b116110e957806110ca81611b2d565b9150508481036110d8575060005b6110e28683611b65565b91506110b0565b80831115611126576000866110fe8386611cac565b6111089190611cbf565b9050806111158882611b65565b995099505050505050505050915091565b80831015611146576000868461113c8489611cac565b6110fe9190611b65565b6004546111538782611b65565b9850985050505050505050915091565b61119d6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b506000908152600760209081526040918290208251606081018452815473ffffffffffffffffffffffffffffffffffffffff168152600182015492810192909252600201549181019190915290565b6040517fe597c19e000000000000000000000000000000000000000000000000000000008152600060048201819052907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e597c19e90602401600060405180830381865afa15801561127a573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526112c09190810190611bee565b905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e5aec9956040518163ffffffff1660e01b8152600401602060405180830381865afa15801561132f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113539190611ac8565b90506106cb8183611558565b600054610100900460ff161580801561137f5750600054600160ff909116105b806113995750303b158015611399575060005460ff166001145b611425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610620565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561148357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b600082116114ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f696e76616c696420666972737445706f636853746172740000000000000000006044820152606401610620565b600482905580156106cb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b80515b428360045461156a9190611b65565b11611689576006805490600061157f83611b2d565b9190505550604051806060016040528083600554815181106115a3576115a3611c7d565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1681526020016004548152602001846004546115dc9190611b65565b90526006546000908152600760209081526040808320845181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161781559184015160018301559290920151600290920191909155600580549161165883611b2d565b9190505550806005540361166c5760006005555b826004600082825461167e9190611b65565b9091555061155b9050565b60408051848152602081018390527fabb37912485bfb13380247be2f4101619759991c9a13ef282eeb05108378b579910160405180910390a1505050565b606060006116d483611785565b600101905060008167ffffffffffffffff8111156116f4576116f4611881565b6040519080825280601f01601f19166020018201604052801561171e576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461172857509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106117ce577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106117fa576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061181857662386f26fc10000830492506010015b6305f5e1008310611830576305f5e100830492506008015b612710831061184457612710830492506004015b60648310611856576064830492506002015b600a8310611862576001015b92915050565b60006020828403121561187a57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156118f7576118f7611881565b604052919050565b600067ffffffffffffffff82111561191957611919611881565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461194557600080fd5b50565b6000602080838503121561195b57600080fd5b823567ffffffffffffffff81111561197257600080fd5b8301601f8101851361198357600080fd5b8035611996611991826118ff565b6118b0565b81815260059190911b820183019083810190878311156119b557600080fd5b928401925b828410156119dc5783356119cd81611923565b825292840192908401906119ba565b979650505050505050565b600080600080600060a086880312156119ff57600080fd5b853594506020860135611a1181611923565b94979496505050506040830135926060810135926080909101359150565b60005b83811015611a4a578181015183820152602001611a32565b50506000910152565b6020815260008251806020840152611a72816040850160208701611a2f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600060208284031215611ab657600080fd5b8135611ac181611923565b9392505050565b600060208284031215611ada57600080fd5b5051919050565b600060208284031215611af357600080fd5b8151611ac181611923565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b5e57611b5e611afe565b5060010190565b8082018082111561186257611862611afe565b60008451611b8a818460208901611a2f565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611bc6816001850160208a01611a2f565b60019201918201528351611be1816002840160208801611a2f565b0160020195945050505050565b60006020808385031215611c0157600080fd5b825167ffffffffffffffff811115611c1857600080fd5b8301601f81018513611c2957600080fd5b8051611c37611991826118ff565b81815260059190911b82018301908381019087831115611c5657600080fd5b928401925b828410156119dc578351611c6e81611923565b82529284019290840190611c5b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8181038181111561186257611862611afe565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611cf757611cf7611afe565b50029056fea164736f6c6343000810000a"

func init() {
	if err := json.Unmarshal([]byte(SubmitterStorageLayoutJSON), SubmitterStorageLayout); err != nil {
		panic(err)
	}

	layouts["Submitter"] = SubmitterStorageLayout
	deployedBytecodes["Submitter"] = SubmitterDeployedBin
}
