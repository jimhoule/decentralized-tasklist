// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tasklist

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TaskListInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskListInfo struct {
	Name        string
	Description string
	Email       string
}

// TaskListPartialTask is an auto generated low-level Go binding around an user-defined struct.
type TaskListPartialTask struct {
	Id     string
	Name   string
	IsDone bool
}

// TaskListTask is an auto generated low-level Go binding around an user-defined struct.
type TaskListTask struct {
	Id          string
	Name        string
	Description string
	IsDone      bool
}

// TasklistMetaData contains all meta data concerning the Tasklist contract.
var TasklistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"deleteTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"internalType\":\"structTaskList.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isDone\",\"type\":\"bool\"}],\"internalType\":\"structTaskList.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isDone\",\"type\":\"bool\"}],\"internalType\":\"structTaskList.PartialTask[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isDone\",\"type\":\"bool\"}],\"name\":\"updateTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161200f38038061200f833981810160405281019061003191906101f5565b335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826001908161007f91906104a6565b50816002908161008f91906104a6565b50806003908161009f91906104a6565b50505050610575565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610107826100c1565b810181811067ffffffffffffffff82111715610126576101256100d1565b5b80604052505050565b5f6101386100a8565b905061014482826100fe565b919050565b5f67ffffffffffffffff821115610163576101626100d1565b5b61016c826100c1565b9050602081019050919050565b8281835e5f83830152505050565b5f61019961019484610149565b61012f565b9050828152602081018484840111156101b5576101b46100bd565b5b6101c0848285610179565b509392505050565b5f82601f8301126101dc576101db6100b9565b5b81516101ec848260208601610187565b91505092915050565b5f805f6060848603121561020c5761020b6100b1565b5b5f84015167ffffffffffffffff811115610229576102286100b5565b5b610235868287016101c8565b935050602084015167ffffffffffffffff811115610256576102556100b5565b5b610262868287016101c8565b925050604084015167ffffffffffffffff811115610283576102826100b5565b5b61028f868287016101c8565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102e757607f821691505b6020821081036102fa576102f96102a3565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261035c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610321565b6103668683610321565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6103aa6103a56103a08461037e565b610387565b61037e565b9050919050565b5f819050919050565b6103c383610390565b6103d76103cf826103b1565b84845461032d565b825550505050565b5f90565b6103eb6103df565b6103f68184846103ba565b505050565b5b818110156104195761040e5f826103e3565b6001810190506103fc565b5050565b601f82111561045e5761042f81610300565b61043884610312565b81016020851015610447578190505b61045b61045385610312565b8301826103fb565b50505b505050565b5f82821c905092915050565b5f61047e5f1984600802610463565b1980831691505092915050565b5f610496838361046f565b9150826002028217905092915050565b6104af82610299565b67ffffffffffffffff8111156104c8576104c76100d1565b5b6104d282546102d0565b6104dd82828561041d565b5f60209050601f83116001811461050e575f84156104fc578287015190505b610506858261048b565b86555061056d565b601f19841661051c86610300565b5f5b828110156105435784890151825560018201915060208501945060208101905061051e565b86831015610560578489015161055c601f89168261046f565b8355505b6001600288020188555050505b505050505050565b611a8d806105825f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806304a87e38146100645780631272f20c146100825780634240803d1461009e578063517c0734146100ce5780635a9b0b89146100ea578063859c918614610108575b5f80fd5b61006c610124565b604051610079919061116b565b60405180910390f35b61009c600480360381019061009791906112f2565b61043a565b005b6100b860048036038101906100b391906113aa565b61053f565b6040516100c5919061145f565b60405180910390f35b6100e860048036038101906100e3919061147f565b61087e565b005b6100f2610985565b6040516100ff919061157e565b60405180910390f35b610122600480360381019061011d91906113aa565b610b9b565b005b60603373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461017c575f80fd5b5f60048054905067ffffffffffffffff81111561019c5761019b6111a4565b5b6040519080825280602002602001820160405280156101d557816020015b6101c2610f0e565b8152602001906001900390816101ba5790505b5090505f5b600480549050811015610432575f600482815481106101fc576101fb61159e565b5b905f5260205f2090600402016040518060800160405290815f82018054610222906115f8565b80601f016020809104026020016040519081016040528092919081815260200182805461024e906115f8565b80156102995780601f1061027057610100808354040283529160200191610299565b820191905f5260205f20905b81548152906001019060200180831161027c57829003601f168201915b505050505081526020016001820180546102b2906115f8565b80601f01602080910402602001604051908101604052809291908181526020018280546102de906115f8565b80156103295780601f1061030057610100808354040283529160200191610329565b820191905f5260205f20905b81548152906001019060200180831161030c57829003601f168201915b50505050508152602001600282018054610342906115f8565b80601f016020809104026020016040519081016040528092919081815260200182805461036e906115f8565b80156103b95780601f10610390576101008083540402835291602001916103b9565b820191905f5260205f20905b81548152906001019060200180831161039c57829003601f168201915b50505050508152602001600382015f9054906101000a900460ff16151515158152505090506040518060600160405280825f0151815260200182602001518152602001826060015115158152508383815181106104195761041861159e565b5b60200260200101819052505080806001019150506101da565b508091505090565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610490575f80fd5b5f61049a85610d4c565b905083600482815481106104b1576104b061159e565b5b905f5260205f20906004020160010190816104cc91906117ce565b5082600482815481106104e2576104e161159e565b5b905f5260205f20906004020160020190816104fd91906117ce565b5081600482815481106105135761051261159e565b5b905f5260205f2090600402016003015f6101000a81548160ff0219169083151502179055505050505050565b610547610f30565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461059d575f80fd5b6105a5610f30565b5f5b600480549050811015610874575f610667600483815481106105cc576105cb61159e565b5b905f5260205f2090600402015f0180546105e5906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610611906115f8565b801561065c5780601f106106335761010080835404028352916020019161065c565b820191905f5260205f20905b81548152906001019060200180831161063f57829003601f168201915b505050505086610e3c565b9050801561086657600482815481106106835761068261159e565b5b905f5260205f2090600402016040518060800160405290815f820180546106a9906115f8565b80601f01602080910402602001604051908101604052809291908181526020018280546106d5906115f8565b80156107205780601f106106f757610100808354040283529160200191610720565b820191905f5260205f20905b81548152906001019060200180831161070357829003601f168201915b50505050508152602001600182018054610739906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610765906115f8565b80156107b05780601f10610787576101008083540402835291602001916107b0565b820191905f5260205f20905b81548152906001019060200180831161079357829003601f168201915b505050505081526020016002820180546107c9906115f8565b80601f01602080910402602001604051908101604052809291908181526020018280546107f5906115f8565b80156108405780601f1061081757610100808354040283529160200191610840565b820191905f5260205f20905b81548152906001019060200180831161082357829003601f168201915b50505050508152602001600382015f9054906101000a900460ff16151515158152505092505b5080806001019150506105a7565b5080915050919050565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108d4575f80fd5b600460405180608001604052808581526020018481526020018381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f01908161093291906117ce565b50602082015181600101908161094891906117ce565b50604082015181600201908161095e91906117ce565b506060820151816003015f6101000a81548160ff0219169083151502179055505050505050565b61098d610f59565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146109e3575f80fd5b6040518060600160405280600180546109fb906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610a27906115f8565b8015610a725780601f10610a4957610100808354040283529160200191610a72565b820191905f5260205f20905b815481529060010190602001808311610a5557829003601f168201915b5050505050815260200160028054610a89906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab5906115f8565b8015610b005780601f10610ad757610100808354040283529160200191610b00565b820191905f5260205f20905b815481529060010190602001808311610ae357829003601f168201915b5050505050815260200160038054610b17906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610b43906115f8565b8015610b8e5780601f10610b6557610100808354040283529160200191610b8e565b820191905f5260205f20905b815481529060010190602001808311610b7157829003601f168201915b5050505050815250905090565b3373ffffffffffffffffffffffffffffffffffffffff165f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bf1575f80fd5b5f610bfb82610d4c565b90505f8190505b6001600480549050610c1491906118ca565b811015610cdf576004600182610c2a91906118fd565b81548110610c3b57610c3a61159e565b5b905f5260205f20906004020160048281548110610c5b57610c5a61159e565b5b905f5260205f2090600402015f8201815f019081610c799190611945565b5060018201816001019081610c8e9190611945565b5060028201816002019081610ca39190611945565b50600382015f9054906101000a900460ff16816003015f6101000a81548160ff0219169083151502179055509050508080600101915050610c02565b506004805480610cf257610cf1611a2a565b5b600190038181905f5260205f2090600402015f8082015f610d139190610f7a565b600182015f610d229190610f7a565b600282015f610d319190610f7a565b600382015f6101000a81549060ff0219169055505090555050565b5f805f90505f5b600480549050811015610e32575f610e1360048381548110610d7857610d7761159e565b5b905f5260205f2090600402015f018054610d91906115f8565b80601f0160208091040260200160405190810160405280929190818152602001828054610dbd906115f8565b8015610e085780601f10610ddf57610100808354040283529160200191610e08565b820191905f5260205f20905b815481529060010190602001808311610deb57829003601f168201915b505050505086610e3c565b90508015610e245781925050610e32565b508080600101915050610d53565b5080915050919050565b5f808390505f8390508051825114610e58575f92505050610f08565b5f5b8251811015610f0057818181518110610e7657610e7561159e565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916838281518110610eb657610eb561159e565b5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610ef3575f9350505050610f08565b8080600101915050610e5a565b506001925050505b92915050565b604051806060016040528060608152602001606081526020015f151581525090565b60405180608001604052806060815260200160608152602001606081526020015f151581525090565b60405180606001604052806060815260200160608152602001606081525090565b508054610f86906115f8565b5f825580601f10610f975750610fb4565b601f0160209004905f5260205f2090810190610fb39190610fb7565b5b50565b5b80821115610fce575f815f905550600101610fb8565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61103d82610ffb565b6110478185611005565b9350611057818560208601611015565b61106081611023565b840191505092915050565b5f8115159050919050565b61107f8161106b565b82525050565b5f606083015f8301518482035f86015261109f8282611033565b915050602083015184820360208601526110b98282611033565b91505060408301516110ce6040860182611076565b508091505092915050565b5f6110e48383611085565b905092915050565b5f602082019050919050565b5f61110282610fd2565b61110c8185610fdc565b93508360208202850161111e85610fec565b805f5b85811015611159578484038952815161113a85826110d9565b9450611145836110ec565b925060208a01995050600181019050611121565b50829750879550505050505092915050565b5f6020820190508181035f83015261118381846110f8565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6111da82611023565b810181811067ffffffffffffffff821117156111f9576111f86111a4565b5b80604052505050565b5f61120b61118b565b905061121782826111d1565b919050565b5f67ffffffffffffffff821115611236576112356111a4565b5b61123f82611023565b9050602081019050919050565b828183375f83830152505050565b5f61126c6112678461121c565b611202565b905082815260208101848484011115611288576112876111a0565b5b61129384828561124c565b509392505050565b5f82601f8301126112af576112ae61119c565b5b81356112bf84826020860161125a565b91505092915050565b6112d18161106b565b81146112db575f80fd5b50565b5f813590506112ec816112c8565b92915050565b5f805f806080858703121561130a57611309611194565b5b5f85013567ffffffffffffffff81111561132757611326611198565b5b6113338782880161129b565b945050602085013567ffffffffffffffff81111561135457611353611198565b5b6113608782880161129b565b935050604085013567ffffffffffffffff81111561138157611380611198565b5b61138d8782880161129b565b925050606061139e878288016112de565b91505092959194509250565b5f602082840312156113bf576113be611194565b5b5f82013567ffffffffffffffff8111156113dc576113db611198565b5b6113e88482850161129b565b91505092915050565b5f608083015f8301518482035f86015261140b8282611033565b915050602083015184820360208601526114258282611033565b9150506040830151848203604086015261143f8282611033565b91505060608301516114546060860182611076565b508091505092915050565b5f6020820190508181035f83015261147781846113f1565b905092915050565b5f805f6060848603121561149657611495611194565b5b5f84013567ffffffffffffffff8111156114b3576114b2611198565b5b6114bf8682870161129b565b935050602084013567ffffffffffffffff8111156114e0576114df611198565b5b6114ec8682870161129b565b925050604084013567ffffffffffffffff81111561150d5761150c611198565b5b6115198682870161129b565b9150509250925092565b5f606083015f8301518482035f86015261153d8282611033565b915050602083015184820360208601526115578282611033565b915050604083015184820360408601526115718282611033565b9150508091505092915050565b5f6020820190508181035f8301526115968184611523565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061160f57607f821691505b602082108103611622576116216115cb565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026116847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611649565b61168e8683611649565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6116d26116cd6116c8846116a6565b6116af565b6116a6565b9050919050565b5f819050919050565b6116eb836116b8565b6116ff6116f7826116d9565b848454611655565b825550505050565b5f90565b611713611707565b61171e8184846116e2565b505050565b5b81811015611741576117365f8261170b565b600181019050611724565b5050565b601f8211156117865761175781611628565b6117608461163a565b8101602085101561176f578190505b61178361177b8561163a565b830182611723565b50505b505050565b5f82821c905092915050565b5f6117a65f198460080261178b565b1980831691505092915050565b5f6117be8383611797565b9150826002028217905092915050565b6117d782610ffb565b67ffffffffffffffff8111156117f0576117ef6111a4565b5b6117fa82546115f8565b611805828285611745565b5f60209050601f831160018114611836575f8415611824578287015190505b61182e85826117b3565b865550611895565b601f19841661184486611628565b5f5b8281101561186b57848901518255600182019150602085019450602081019050611846565b868310156118885784890151611884601f891682611797565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118d4826116a6565b91506118df836116a6565b92508282039050818111156118f7576118f661189d565b5b92915050565b5f611907826116a6565b9150611912836116a6565b925082820190508082111561192a5761192961189d565b5b92915050565b5f8154905061193e816115f8565b9050919050565b818103611953575050611a28565b61195c82611930565b67ffffffffffffffff811115611975576119746111a4565b5b61197f82546115f8565b61198a828285611745565b5f601f8311600181146119b7575f84156119a5578287015490505b6119af85826117b3565b865550611a21565b601f1984166119c587611628565b96506119d086611628565b5f5b828110156119f7578489015482556001820191506001850194506020810190506119d2565b86831015611a145784890154611a10601f891682611797565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212206cc50410e96793966942509b363f91bfe868e73bdfb7d79035dbeff185ed5c2e64736f6c634300081a0033",
}

// TasklistABI is the input ABI used to generate the binding from.
// Deprecated: Use TasklistMetaData.ABI instead.
var TasklistABI = TasklistMetaData.ABI

// TasklistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TasklistMetaData.Bin instead.
var TasklistBin = TasklistMetaData.Bin

// DeployTasklist deploys a new Ethereum contract, binding an instance of Tasklist to it.
func DeployTasklist(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _description string, _email string) (common.Address, *types.Transaction, *Tasklist, error) {
	parsed, err := TasklistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TasklistBin), backend, _name, _description, _email)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tasklist{TasklistCaller: TasklistCaller{contract: contract}, TasklistTransactor: TasklistTransactor{contract: contract}, TasklistFilterer: TasklistFilterer{contract: contract}}, nil
}

// Tasklist is an auto generated Go binding around an Ethereum contract.
type Tasklist struct {
	TasklistCaller     // Read-only binding to the contract
	TasklistTransactor // Write-only binding to the contract
	TasklistFilterer   // Log filterer for contract events
}

// TasklistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TasklistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TasklistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TasklistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TasklistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TasklistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TasklistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TasklistSession struct {
	Contract     *Tasklist         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TasklistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TasklistCallerSession struct {
	Contract *TasklistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TasklistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TasklistTransactorSession struct {
	Contract     *TasklistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TasklistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TasklistRaw struct {
	Contract *Tasklist // Generic contract binding to access the raw methods on
}

// TasklistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TasklistCallerRaw struct {
	Contract *TasklistCaller // Generic read-only contract binding to access the raw methods on
}

// TasklistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TasklistTransactorRaw struct {
	Contract *TasklistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTasklist creates a new instance of Tasklist, bound to a specific deployed contract.
func NewTasklist(address common.Address, backend bind.ContractBackend) (*Tasklist, error) {
	contract, err := bindTasklist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tasklist{TasklistCaller: TasklistCaller{contract: contract}, TasklistTransactor: TasklistTransactor{contract: contract}, TasklistFilterer: TasklistFilterer{contract: contract}}, nil
}

// NewTasklistCaller creates a new read-only instance of Tasklist, bound to a specific deployed contract.
func NewTasklistCaller(address common.Address, caller bind.ContractCaller) (*TasklistCaller, error) {
	contract, err := bindTasklist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TasklistCaller{contract: contract}, nil
}

// NewTasklistTransactor creates a new write-only instance of Tasklist, bound to a specific deployed contract.
func NewTasklistTransactor(address common.Address, transactor bind.ContractTransactor) (*TasklistTransactor, error) {
	contract, err := bindTasklist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TasklistTransactor{contract: contract}, nil
}

// NewTasklistFilterer creates a new log filterer instance of Tasklist, bound to a specific deployed contract.
func NewTasklistFilterer(address common.Address, filterer bind.ContractFilterer) (*TasklistFilterer, error) {
	contract, err := bindTasklist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TasklistFilterer{contract: contract}, nil
}

// bindTasklist binds a generic wrapper to an already deployed contract.
func bindTasklist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TasklistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tasklist *TasklistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tasklist.Contract.TasklistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tasklist *TasklistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tasklist.Contract.TasklistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tasklist *TasklistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tasklist.Contract.TasklistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tasklist *TasklistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tasklist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tasklist *TasklistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tasklist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tasklist *TasklistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tasklist.Contract.contract.Transact(opts, method, params...)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((string,string,string))
func (_Tasklist *TasklistCaller) GetInfo(opts *bind.CallOpts) (TaskListInfo, error) {
	var out []interface{}
	err := _Tasklist.contract.Call(opts, &out, "getInfo")

	if err != nil {
		return *new(TaskListInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskListInfo)).(*TaskListInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((string,string,string))
func (_Tasklist *TasklistSession) GetInfo() (TaskListInfo, error) {
	return _Tasklist.Contract.GetInfo(&_Tasklist.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns((string,string,string))
func (_Tasklist *TasklistCallerSession) GetInfo() (TaskListInfo, error) {
	return _Tasklist.Contract.GetInfo(&_Tasklist.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0x4240803d.
//
// Solidity: function getTask(string _id) view returns((string,string,string,bool))
func (_Tasklist *TasklistCaller) GetTask(opts *bind.CallOpts, _id string) (TaskListTask, error) {
	var out []interface{}
	err := _Tasklist.contract.Call(opts, &out, "getTask", _id)

	if err != nil {
		return *new(TaskListTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskListTask)).(*TaskListTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x4240803d.
//
// Solidity: function getTask(string _id) view returns((string,string,string,bool))
func (_Tasklist *TasklistSession) GetTask(_id string) (TaskListTask, error) {
	return _Tasklist.Contract.GetTask(&_Tasklist.CallOpts, _id)
}

// GetTask is a free data retrieval call binding the contract method 0x4240803d.
//
// Solidity: function getTask(string _id) view returns((string,string,string,bool))
func (_Tasklist *TasklistCallerSession) GetTask(_id string) (TaskListTask, error) {
	return _Tasklist.Contract.GetTask(&_Tasklist.CallOpts, _id)
}

// GetTasks is a free data retrieval call binding the contract method 0x04a87e38.
//
// Solidity: function getTasks() view returns((string,string,bool)[])
func (_Tasklist *TasklistCaller) GetTasks(opts *bind.CallOpts) ([]TaskListPartialTask, error) {
	var out []interface{}
	err := _Tasklist.contract.Call(opts, &out, "getTasks")

	if err != nil {
		return *new([]TaskListPartialTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TaskListPartialTask)).(*[]TaskListPartialTask)

	return out0, err

}

// GetTasks is a free data retrieval call binding the contract method 0x04a87e38.
//
// Solidity: function getTasks() view returns((string,string,bool)[])
func (_Tasklist *TasklistSession) GetTasks() ([]TaskListPartialTask, error) {
	return _Tasklist.Contract.GetTasks(&_Tasklist.CallOpts)
}

// GetTasks is a free data retrieval call binding the contract method 0x04a87e38.
//
// Solidity: function getTasks() view returns((string,string,bool)[])
func (_Tasklist *TasklistCallerSession) GetTasks() ([]TaskListPartialTask, error) {
	return _Tasklist.Contract.GetTasks(&_Tasklist.CallOpts)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(string _id, string _name, string _description) returns()
func (_Tasklist *TasklistTransactor) AddTask(opts *bind.TransactOpts, _id string, _name string, _description string) (*types.Transaction, error) {
	return _Tasklist.contract.Transact(opts, "addTask", _id, _name, _description)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(string _id, string _name, string _description) returns()
func (_Tasklist *TasklistSession) AddTask(_id string, _name string, _description string) (*types.Transaction, error) {
	return _Tasklist.Contract.AddTask(&_Tasklist.TransactOpts, _id, _name, _description)
}

// AddTask is a paid mutator transaction binding the contract method 0x517c0734.
//
// Solidity: function addTask(string _id, string _name, string _description) returns()
func (_Tasklist *TasklistTransactorSession) AddTask(_id string, _name string, _description string) (*types.Transaction, error) {
	return _Tasklist.Contract.AddTask(&_Tasklist.TransactOpts, _id, _name, _description)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x859c9186.
//
// Solidity: function deleteTask(string _id) returns()
func (_Tasklist *TasklistTransactor) DeleteTask(opts *bind.TransactOpts, _id string) (*types.Transaction, error) {
	return _Tasklist.contract.Transact(opts, "deleteTask", _id)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x859c9186.
//
// Solidity: function deleteTask(string _id) returns()
func (_Tasklist *TasklistSession) DeleteTask(_id string) (*types.Transaction, error) {
	return _Tasklist.Contract.DeleteTask(&_Tasklist.TransactOpts, _id)
}

// DeleteTask is a paid mutator transaction binding the contract method 0x859c9186.
//
// Solidity: function deleteTask(string _id) returns()
func (_Tasklist *TasklistTransactorSession) DeleteTask(_id string) (*types.Transaction, error) {
	return _Tasklist.Contract.DeleteTask(&_Tasklist.TransactOpts, _id)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x1272f20c.
//
// Solidity: function updateTask(string _id, string _name, string _description, bool _isDone) returns()
func (_Tasklist *TasklistTransactor) UpdateTask(opts *bind.TransactOpts, _id string, _name string, _description string, _isDone bool) (*types.Transaction, error) {
	return _Tasklist.contract.Transact(opts, "updateTask", _id, _name, _description, _isDone)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x1272f20c.
//
// Solidity: function updateTask(string _id, string _name, string _description, bool _isDone) returns()
func (_Tasklist *TasklistSession) UpdateTask(_id string, _name string, _description string, _isDone bool) (*types.Transaction, error) {
	return _Tasklist.Contract.UpdateTask(&_Tasklist.TransactOpts, _id, _name, _description, _isDone)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x1272f20c.
//
// Solidity: function updateTask(string _id, string _name, string _description, bool _isDone) returns()
func (_Tasklist *TasklistTransactorSession) UpdateTask(_id string, _name string, _description string, _isDone bool) (*types.Transaction, error) {
	return _Tasklist.Contract.UpdateTask(&_Tasklist.TransactOpts, _id, _name, _description, _isDone)
}
