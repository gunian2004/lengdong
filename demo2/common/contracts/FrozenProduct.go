// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FrozenProductLogisticsRecord is an auto generated low-level Go binding around an user-defined struct.
type FrozenProductLogisticsRecord struct {
	Sku             string
	TransportMethod string
	Carrier         string
	Temperature     string
	Departure       string
	Destination     string
	Operator        common.Address
	OperatorName    string
	IsValid         bool
}

// FrozenProductProduct is an auto generated low-level Go binding around an user-defined struct.
type FrozenProductProduct struct {
	Sku               string
	ProductName       string
	Brand             string
	Specification     string
	ProductionDate    *big.Int
	ExpirationDate    *big.Int
	RawMaterialSource string
	ProcessingMethod  string
	LogisticsTemp     string
	StorageCondition  string
	QcReport          string
	ManufacturerId    string
	ImageUrl          string
	BatchNumber       string
	Quantity          *big.Int
	CurrentHolder     common.Address
	CurrentLocation   string
	IsValid           bool
}

// FrozenProductPurchaseRecord is an auto generated low-level Go binding around an user-defined struct.
type FrozenProductPurchaseRecord struct {
	PurchaseId   *big.Int
	Sku          string
	ProductName  string
	Quantity     *big.Int
	PurchaseDate *big.Int
	Status       string
	SellerId     common.Address
	BuyerId      common.Address
	IsValid      bool
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"logisticsId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"carrier\",\"type\":\"string\"}],\"name\":\"LogisticsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProductAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"purchaseId\",\"type\":\"uint256\"}],\"name\":\"PurchaseConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"purchaseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"PurchaseCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_transportMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_carrier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_temperature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_departure\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"}],\"name\":\"addLogisticsRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_current_location\",\"type\":\"string\"}],\"name\":\"addProduct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_purchaseId\",\"type\":\"uint256\"}],\"name\":\"confirmPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_product_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer_id\",\"type\":\"address\"}],\"name\":\"createPurchaseRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAvailableProducts\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"current_holder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"current_location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.Product[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllLogisticsIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllProductSkus\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"getBuyerCompletePurchaseHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"purchaseId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchase_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer_id\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.PurchaseRecord[]\",\"name\":\"purchases\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"current_holder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"current_location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.Product[]\",\"name\":\"products\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"getBuyerPurchaseIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"}],\"name\":\"getCompleteProductInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"current_holder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"current_location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.Product\",\"name\":\"product\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"carrier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"temperature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"departure\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.LogisticsRecord[]\",\"name\":\"logisticsRecords\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"}],\"name\":\"getLogisticsBySku\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"carrier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"temperature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"departure\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.LogisticsRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_logisticsId\",\"type\":\"uint256\"}],\"name\":\"getLogisticsRecord\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportMethod\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"carrier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"temperature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"departure\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"}],\"name\":\"getProductBySku\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"current_holder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"current_location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_purchaseId\",\"type\":\"uint256\"}],\"name\":\"getPurchaseRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"purchaseId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchase_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer_id\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getSellerCompleteSalesHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"purchaseId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"purchase_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller_id\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer_id\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.PurchaseRecord[]\",\"name\":\"purchases\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"product_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"brand\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"specification\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"production_date\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration_date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"raw_material_source\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"processing_method\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logistics_temp\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storage_condition\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"qc_report\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"manufacturer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"current_holder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"current_location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"internalType\":\"structFrozenProduct.Product[]\",\"name\":\"products\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getSellerPurchaseIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_newHolder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newLocation\",\"type\":\"string\"}],\"name\":\"updateProductHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346075573315605f5760008054336001600160a01b0319821681178355916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3612a12908161007b8239f35b631e4fbdf760e01b600052600060045260246000fd5b600080fdfe61012080604052600436101561001457600080fd5b60003560e01c9081630a67b45b14611af2575080631b0612d514611a7b5780634d24e902146119165780636ff8f2f7146117d0578063715018a6146117775780638da5cb5b1461174e578063966bd197146116f0578063ad8644231461148f578063bc8bf20a14611418578063c1cc2d2114610e9f578063c900500d14610c2d578063cbc9d8fa1461093d578063d433aba3146108ba578063da140376146107b4578063db10187b146106cb578063e64037fa14610576578063ecb5f4891461040b578063f2fde38b14610382578063f521c1e5146101e55763fad305a8146100fc57600080fd5b346101e05760203660031901126101e0576004356001600160401b0381116101e05761012f6101ce913690600401612012565b61013761220f565b506101dc6101ba6101b46101a4602060405161017660ff60118951938581818d019661016481838a611c85565b81016001815203019020015416612767565b60405182818951610188818387611c85565b8101600181520301902096604051938492839251928391611c85565b8101600381520301902093612406565b926126cc565b604051938493604085526040850190611ccd565b908382036020850152612068565b0390f35b600080fd5b346101e05760203660031901126101e0576001600160a01b03610206611c6f565b16600052600760205260406000206040518082602082945493848152019060005260206000209260005b81811061036957505061024592500382611ff1565b61024f815161217f565b9061025a815161229e565b9160005b825181101561035a5780610274600192856122ee565b516000526006602052604060002060ff60076040519261029384611fb9565b805484526102a286820161233c565b60208501526102b36002820161233c565b604085015260038101546060850152600481015460808501526102d86005820161233c565b60a0850152858060a01b0360068201541660c08501520154848060a01b03811660e084015260a01c16151561010082015261031382856122ee565b5261031e81846122ee565b5061033e610339602061033184876122ee565b5101516123e0565b612406565b61034882876122ee565b5261035381866122ee565b500161025e565b604051806101dc868583611e8c565b8454835260019485019486945060209093019201610230565b346101e05760203660031901126101e05761039b611c6f565b6103a36129b3565b6001600160a01b031680156103f557600080546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b600052600060045260246000fd5b346101e05760603660031901126101e0576004356001600160401b0381116101e05761043b903690600401612012565b602435906001600160a01b03821682036101e0576044356001600160401b0381116101e05761046e903690600401612012565b906040519061048f60ff60118351946020818187019761016481838b611c85565b604051600f82516104a1818487611c85565b600190830190815282900360200190912001546001600160a01b0316330361053157602060109261051f9261052f96600f604051858185516104e481838a611c85565b81016001815203019020019060018060a01b03166bffffffffffffffffffffffff60a01b825416179055604051938492839251928391611c85565b81016001815203019020016127ac565b005b60405162461bcd60e51b815260206004820152601e60248201527f4f6e6c792063757272656e7420686f6c6465722063616e2075706461746500006044820152606490fd5b346101e05760203660031901126101e05760043580600052600460205260ff600860406000200154161561068657600052600460205260406000206105ba90612628565b80516020820151916040810151906060810151608082015160a083015190600160a01b6001900360c0850151169260e08501519461010001511515956040519889986101208a526101208a0161060f91611ca8565b89810360208b015261062091611ca8565b88810360408a015261063191611ca8565b878103606089015261064291611ca8565b868103608088015261065391611ca8565b85810360a087015261066491611ca8565b9060c085015283810360e085015261067b91611ca8565b906101008301520390f35b60405162461bcd60e51b815260206004820152601f60248201527f4c6f67697374696373207265636f726420646f6573206e6f74206578697374006044820152606490fd5b346101e05760003660031901126101e0576002546106e881612168565b906106f66040519283611ff1565b80825260208201908160026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace6000915b83831061079757848660405191829160208301906020845251809152604083019060408160051b85010192916000905b82821061076857505050500390f35b919360019193955060206107878192603f198a82030186528851611ca8565b9601920192018594939192610759565b6001602081926107a68561233c565b815201920192019190610729565b346101e05760003660031901126101e0576002546000805b82811061086857506107dd9061229e565b906000805b8281106107ff57604051602080825281906101dc90820187611e33565b61081461033961080e8361294e565b5061259a565b6101c081015115158061085a575b610830575b506001016107e2565b826108539161084260019495886122ee565b5261084d81876122ee565b506129a4565b9190610827565b506102208101511515610822565b600e61087661080e8361294e565b01541515806108a1575b61088d575b6001016107cc565b906108996001916129a4565b919050610885565b5060ff60116108b261080e8461294e565b015416610880565b346101e05760203660031901126101e0576001600160a01b036108db611c6f565b166000526007602052604060002060405190816020825491828152019160005260206000209060005b818110610927576101dc8561091b81870382611ff1565b60405191829182611f7f565b8254845260209093019260019283019201610904565b346101e05760a03660031901126101e0576004356001600160401b0381116101e05761096d903690600401612012565b6024356001600160401b0381116101e05761098c903690600401612012565b60443591906064356001600160a01b03811691908290036101e0576084356001600160a01b03811692908390036101e05784600e6109f46020876109e360ff6011604051858181875197019661016481838a611c85565b604051809381928b51928391611c85565b81016001815203019020015410610bf0576009549260018401809411610bda578060016007604051610a2581611fb9565b87815285610acc8960208401998c8b52604085019081528d60608601908152610ab66080870192428452610aac60409e8f805197610a63828a611ff1565b8c89526670656e64696e6760c81b60208a015260a08c0198895260c08c019a8b5260e08c019e8f526101008c019d8e52600052600660205260002099518a555160018a016127ac565b51600288016127ac565b51600386015551600485015551600584016127ac565b516006820180546001600160a01b0319166001600160a01b0392831617905593519101805492516001600160a81b0319909316919093161790151560a01b60ff60a01b1617905560009081526007602052829020610b2b908490612981565b6000526008602052610b408282600020612981565b60095493600160401b851015610bc457610ba083610b878760017fe89d7cb7157508bfc3241a76bb061488782a9342e89e3ddb531745ef904e933999016009556009612969565b90919082549060031b91821b91600019901b1916179055565b610bbb82519485948552606060208601526060850190611ca8565b918301520390a1005b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60405162461bcd60e51b8152602060048201526015602482015274496e73756666696369656e74207175616e7469747960581b6044820152606490fd5b346101e05760203660031901126101e0576004356001600160401b0381116101e057610c5d903690600401612012565b60405190805191806020830193610c75818387611c85565b8101600181520360200190206011015460ff16610c9190612767565b60405180928192518092610ca492611c85565b810160018152036020019020610cb990612406565b805190602081015160e0526040810151610100526060810151608082015160a08301519060c084015160e08501516101008601516101208701519161014088015193610160890151956101808a0151976101a08b0151996101c08c01519c8c600160a01b60019003906101e00151169c6102008101516080526102200151151560a05260405160c05260c051610240905260c05161024001610d5a91611ca8565b60c051810360c0516020015260e05190610d7391611ca8565b60c051810360c051604001526101005190610d8d91611ca8565b60c051810360c05160600152610da291611ca8565b9160c0516080015260c05160a0015260c051810360c05160c00152610dc691611ca8565b60c051810360c05160e00152610ddb91611ca8565b60c051810360c0516101000152610df191611ca8565b60c051810360c0516101200152610e0791611ca8565b60c051810360c0516101400152610e1d91611ca8565b60c051810360c0516101600152610e3391611ca8565b60c051810360c0516101800152610e4991611ca8565b60c051810360c0516101a00152610e5f91611ca8565b9160c0516101c0015260c0516101e0015260c051810360c051610200015260805190610e8a91611ca8565b60a05160c051610220015260c051900360c051f35b346101e0576102003660031901126101e0576004356001600160401b0381116101e057610ed0903690600401612012565b6024356001600160401b0381116101e057610eef903690600401612012565b6044356001600160401b0381116101e057610f0e903690600401612012565b6064356001600160401b0381116101e057610f2d903690600401612012565b9060c4356001600160401b0381116101e057610f4d903690600401612012565b60e4356001600160401b0381116101e057610f6c903690600401612012565b610104356001600160401b0381116101e057610f8c903690600401612012565b610124356001600160401b0381116101e057610fac903690600401612012565b610144356001600160401b0381116101e057610fcc903690600401612012565b610164356001600160401b0381116101e057610fec903690600401612012565b91610184356001600160401b0381116101e05761100d903690600401612012565b936101a4356001600160401b0381116101e05761102e903690600401612012565b956101e435976001600160401b0389116101e05760ff60116110838f9b61105b6020913690600401612012565b9c6110646129b3565b61107081511515612725565b8160405193828580945193849201611c85565b810160018152030190200154166113c757608435156113825760843560a43511156113275761128c9a8d9a6040519b6110bb8d611fd5565b8c528d60208d015260408c015260608b015260843560808b015260a43560a08b015260c08a015260e08901526101008801526101208701526101408601526101608501526101808401526101a08301526101c4356101c0830152336101e08301526102008201526001610220820152601161022060405160208188516111448183858d01611c85565b810160018152030190209261115a8151856127ac565b61116b6020820151600186016127ac565b61117c6040820151600286016127ac565b61118d6060820151600386016127ac565b6080810151600485015560a081015160058501556111b260c0820151600686016127ac565b6111c360e0820151600786016127ac565b6111d5610100820151600886016127ac565b6111e7610120820151600986016127ac565b6111f9610140820151600a86016127ac565b61120b610160820151600b86016127ac565b61121d610180820151600c86016127ac565b61122f6101a0820151600d86016127ac565b6101c0810151600e8501556101e0810151600f850180546001600160a01b0319166001600160a01b039290921691909117905561020081015161127590601086016127ac565b0151151591019060ff801983541691151516179055565b600254600160401b811015610bc4578060016112ab920160025561294e565b929092611311576113066112f8926112e4837fcbebef04e1d7de535a91cb84796482718220565630fdbc1954fbf9d97f043453966127ac565b604051938493606085526060850190611ca8565b908382036020850152611ca8565b4260408301520390a1005b634e487b7160e01b600052600060045260246000fd5b60405162461bcd60e51b815260206004820152602d60248201527f45787069726174696f6e2064617465206d75737420626520616674657220707260448201526c6f64756374696f6e206461746560981b6064820152608490fd5b60405162461bcd60e51b815260206004820152601d60248201527f50726f64756374696f6e2064617465206d7573742062652076616c69640000006044820152606490fd5b60405162461bcd60e51b8152602060048201526024808201527f50726f647563742077697468207468697320534b5520616c72656164792065786044820152636973747360e01b6064820152608490fd5b346101e05760203660031901126101e0576001600160a01b03611439611c6f565b166000526008602052604060002060405190816020825491828152019160005260206000209060005b818110611479576101dc8561091b81870382611ff1565b8254845260209093019260019283019201611462565b346101e05760e03660031901126101e0576004356001600160401b0381116101e0576114bf903690600401612012565b6024356001600160401b0381116101e0576114de903690600401612012565b906044356001600160401b0381116101e0576114fe903690600401612012565b916064356001600160401b0381116101e05761151e903690600401612012565b906084356001600160401b0381116101e05761153e903690600401612012565b60a4356001600160401b0381116101e05761155d903690600401612012565b9260c4356001600160401b0381116101e05761157d903690600401612012565b61158986511515612725565b604051906115a960ff6011895194602081818d019761016481838b611c85565b6005549560018701809711610bda5761160c94602094604051976115cc89611fb9565b8a8952868901528a60408901526060880152608087015260a08601523360c086015260e08501526001610100850152604051809381928851928391611c85565b81016003815203019020908154600160401b811015610bc457600181018084558110156116da578161164e9161166194600052600960206000209102016128ac565b82600052600460205260406000206128ac565b600554600160401b811015610bc4577fc4f6b94487bd6f57ccd5df84506fd060f8c3c6e9b45cc0d50727219602e5d605936116ab83610b878460016116d596016005556005612969565b6116c76040519485948552606060208601526060850190611ca8565b908382036040850152611ca8565b0390a1005b634e487b7160e01b600052603260045260246000fd5b346101e05760203660031901126101e0576004356001600160401b0381116101e05761173a61172b60206110706101dc943690600401612012565b810160038152030190206126cc565b604051918291602083526020830190612068565b346101e05760003660031901126101e0576000546040516001600160a01b039091168152602090f35b346101e05760003660031901126101e0576117906129b3565b600080546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346101e05760203660031901126101e05760043580600052600660205261180460ff60076040600020015460a01c16612537565b600052600660205260406000206040519061181e82611fb9565b8054825261182e6001820161233c565b90602083019182526118426002820161233c565b92604081019384526003820154606082019081526119016004840154608084019081526118716005860161233c565b8060a0860152600760018060a01b03600688015416968760c08801520154936118e960ff60018060a01b038716968760e08a015260a01c1615159687610100820152519851995191519351916118db6040519b8c9b8c5261012060208d01526101208c0190611ca8565b908a820360408c0152611ca8565b926060890152608088015286820360a0880152611ca8565b9260c085015260e08401526101008301520390f35b346101e05760203660031901126101e05760043580600052600660205261194a60ff60076040600020015460a01c16612537565b60008181526006602052604090206007810180549092906001600160a01b03163303611a36576005820161197e8154612302565b601f8111611a14575b5060126818dbdb999a5c9b595960ba1b01905560036119a86001840161259a565b92015491600e81018054938403938411610bda57929092559154600f9190910180546001600160a01b0319166001600160a01b03929092169190911790556040519081527f2ce3b121b7afbc2724c6c53c061e65d64d4883db75d2f497a715b144a840427590602090a1005b611a309082600052601f6020600020910160051c810190612583565b84611987565b60405162461bcd60e51b815260206004820152601f60248201527f4f6e6c792062757965722063616e20636f6e6669726d207075726368617365006044820152606490fd5b346101e05760003660031901126101e05760405180602060055491828152019060056000527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db09060005b818110611adc576101dc8561091b81870382611ff1565b8254845260209093019260019283019201611ac5565b346101e05760203660031901126101e0576001600160a01b03611b13611c6f565b166000526008602052604060002081815491828252602082019060005260206000209260005b818110611c56575050611b4e92500382611ff1565b611b58815161217f565b90611b63815161229e565b9160005b825181101561035a5780611b7d600192856122ee565b516000526006602052604060002060ff600760405192611b9c84611fb9565b80548452611bab86820161233c565b6020850152611bbc6002820161233c565b60408501526003810154606085015260048101546080850152611be16005820161233c565b60a0850152858060a01b0360068201541660c08501520154848060a01b03811660e084015260a01c161515610100820152611c1c82856122ee565b52611c2781846122ee565b50611c3a610339602061033184876122ee565b611c4482876122ee565b52611c4f81866122ee565b5001611b67565b8454835260019485019486945060209093019201611b39565b600435906001600160a01b03821682036101e057565b60005b838110611c985750506000910152565b8181015183820152602001611c88565b90602091611cc181518092818552858086019101611c85565b601f01601f1916010190565b9061022080611e28611df4611de0611dcc611db8611da4611d90611d7c611d698b60c08e611d45611d33611d21611d0f84516102408852610240880190611ca8565b60208501518782036020890152611ca8565b60408401518682036040880152611ca8565b60608301518582036060870152611ca8565b926080820151608082015260a08083015191015201518d60c0818403910152611ca8565b60e08d01518c60e0818403910152611ca8565b6101008c01518b82036101008d0152611ca8565b6101208b01518a82036101208c0152611ca8565b6101408a01518982036101408b0152611ca8565b6101608901518882036101608a0152611ca8565b610180880151878203610180890152611ca8565b6101a08701518682036101a0880152611ca8565b6101c08601516101c086015260018060a01b036101e0870151166101e0860152610200860151858203610200870152611ca8565b930151151591015290565b9080602083519182815201916020808360051b8301019401926000915b838310611e5f57505050505090565b9091929394602080611e7d600193601f198682030187528951611ccd565b97019301930191939290611e50565b929160408401936040815282518095526060810194602060608260051b84010194019060005b818110611ed257505050611ecf9394506020818403910152611e33565b90565b909194602080600192605f19878203018b528851908151815261010080611f46611f20611f0e8787015161012089880152610120870190611ca8565b60408701518682036040880152611ca8565b606086015160608601526080860151608086015260a086015185820360a0870152611ca8565b93878060a01b0360c08201511660c0850152878060a01b0360e08201511660e08501520151151591015297019801910196919096611eb2565b602060408183019282815284518094520192019060005b818110611fa35750505090565b8251845260209384019390920191600101611f96565b61012081019081106001600160401b03821117610bc457604052565b61024081019081106001600160401b03821117610bc457604052565b90601f801991011681019081106001600160401b03821117610bc457604052565b81601f820112156101e0578035906001600160401b038211610bc45760405192612046601f8401601f191660200185611ff1565b828452602083830101116101e057816000926020809301838601378301015290565b9080602083519182815201916020808360051b8301019401926000915b83831061209457505050505090565b9091929394602080600192601f198582030186528851906101008061215161212e61211c61210a6120f86120e66120d68a516101208b526101208b0190611ca8565b8b8b01518a82038d8c0152611ca8565b60408a015189820360408b0152611ca8565b606089015188820360608a0152611ca8565b60808801518782036080890152611ca8565b60a087015186820360a0880152611ca8565b888060a01b0360c08701511660c086015260e086015185820360e0870152611ca8565b930151151591015297019301930191939290612085565b6001600160401b038111610bc45760051b60200190565b9061218982612168565b6121966040519182611ff1565b82815280926121a7601f1991612168565b019060005b8281106121b857505050565b6020906040516121c781611fb9565b60008152606083820152606060408201526000606082015260006080820152606060a0820152600060c0820152600060e08201526000610100820152828285010152016121ac565b6040519061221c82611fd5565b6000610220836060815260606020820152606060408201526060808201528260808201528260a0820152606060c0820152606060e08201526060610100820152606061012082015260606101408201526060610160820152606061018082015260606101a0820152826101c0820152826101e082015260606102008201520152565b906122a882612168565b6122b56040519182611ff1565b82815280926122c6601f1991612168565b019060005b8281106122d757505050565b6020906122e261220f565b828285010152016122cb565b80518210156116da5760209160051b010190565b90600182811c92168015612332575b602083101461231c57565b634e487b7160e01b600052602260045260246000fd5b91607f1691612311565b906040519182600082549261235084612302565b80845293600181169081156123be5750600114612377575b5061237592500383611ff1565b565b90506000929192526020600020906000915b8183106123a25750509060206123759282010138612368565b6020919350806001915483858901015201910190918492612389565b90506020925061237594915060ff191682840152151560051b82010138612368565b60206123f9918160405193828580945193849201611c85565b8101600181520301902090565b9060405161241381611fd5565b61022060ff601183956124258161233c565b85526124336001820161233c565b60208601526124446002820161233c565b60408601526124556003820161233c565b606086015260048101546080860152600581015460a086015261247a6006820161233c565b60c086015261248b6007820161233c565b60e086015261249c6008820161233c565b6101008601526124ae6009820161233c565b6101208601526124c0600a820161233c565b6101408601526124d2600b820161233c565b6101608601526124e4600c820161233c565b6101808601526124f6600d820161233c565b6101a0860152600e8101546101c0860152600f8101546001600160a01b03166101e08601526125276010820161233c565b6102008601520154161515910152565b1561253e57565b60405162461bcd60e51b815260206004820152601e60248201527f5075726368617365207265636f726420646f6573206e6f7420657869737400006044820152606490fd5b81811061258e575050565b60008155600101612583565b604051908160008254926125ad84612302565b936001811690811561260e57506001146125d2575b5060209250600181520301902090565b9150506000528160206000206000905b8382106125f65750506020918101386125c2565b602091925080600191548487015201910183916125e2565b60ff191684525060209380151502830191503890506125c2565b9060405161263581611fb9565b61010060ff600883956126478161233c565b85526126556001820161233c565b60208601526126666002820161233c565b60408601526126776003820161233c565b60608601526126886004820161233c565b60808601526126996005820161233c565b60a086015260068101546001600160a01b031660c08601526126bd6007820161233c565b60e08601520154161515910152565b9081546126d881612168565b926126e66040519485611ff1565b818452602084019060005260206000206000915b8383106127075750505050565b6009602060019261271785612628565b8152019201920191906126fa565b1561272c57565b60405162461bcd60e51b8152602060048201526013602482015272534b552063616e6e6f7420626520656d70747960681b6044820152606490fd5b1561276e57565b60405162461bcd60e51b8152602060048201526016602482015275141c9bd91d58dd08191bd95cc81b9bdd08195e1a5cdd60521b6044820152606490fd5b91909182516001600160401b038111610bc4576127c98254612302565b601f811161286f575b506020601f821160011461280d5781929394600092612802575b50508160011b916000199060031b1c1916179055565b0151905038806127ec565b601f1982169083600052806000209160005b8181106128575750958360019596971061283e575b505050811b019055565b015160001960f88460031b161c19169055388080612834565b9192602060018192868b01518155019401920161281f565b61289c90836000526020600020601f840160051c810191602085106128a2575b601f0160051c0190612583565b386127d2565b909150819061288f565b6008610100612375936128c08151856127ac565b6128d16020820151600186016127ac565b6128e26040820151600286016127ac565b6128f36060820151600386016127ac565b6129046080820151600486016127ac565b61291560a0820151600586016127ac565b60c08101516006850180546001600160a01b0319166001600160a01b039290921691909117905560e081015161127590600786016127ac565b6002548110156116da57600260005260206000200190600090565b80548210156116da5760005260206000200190600090565b90815491600160401b831015610bc45782610b8791600161237595018155612969565b6000198114610bda5760010190565b6000546001600160a01b031633036129c757565b63118cdaa760e01b6000523360045260246000fdfea2646970667358221220dcc6eef746bc78df010620948c1cea2f012e7eece08e9c7bf11e92f8e33a545364736f6c634300081c0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetAllAvailableProducts is a free data retrieval call binding the contract method 0xda140376.
//
// Solidity: function getAllAvailableProducts() view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[])
func (_Contracts *ContractsCaller) GetAllAvailableProducts(opts *bind.CallOpts) ([]FrozenProductProduct, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllAvailableProducts")

	if err != nil {
		return *new([]FrozenProductProduct), err
	}

	out0 := *abi.ConvertType(out[0], new([]FrozenProductProduct)).(*[]FrozenProductProduct)

	return out0, err

}

// GetAllAvailableProducts is a free data retrieval call binding the contract method 0xda140376.
//
// Solidity: function getAllAvailableProducts() view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[])
func (_Contracts *ContractsSession) GetAllAvailableProducts() ([]FrozenProductProduct, error) {
	return _Contracts.Contract.GetAllAvailableProducts(&_Contracts.CallOpts)
}

// GetAllAvailableProducts is a free data retrieval call binding the contract method 0xda140376.
//
// Solidity: function getAllAvailableProducts() view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[])
func (_Contracts *ContractsCallerSession) GetAllAvailableProducts() ([]FrozenProductProduct, error) {
	return _Contracts.Contract.GetAllAvailableProducts(&_Contracts.CallOpts)
}

// GetAllLogisticsIds is a free data retrieval call binding the contract method 0x1b0612d5.
//
// Solidity: function getAllLogisticsIds() view returns(uint256[])
func (_Contracts *ContractsCaller) GetAllLogisticsIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllLogisticsIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllLogisticsIds is a free data retrieval call binding the contract method 0x1b0612d5.
//
// Solidity: function getAllLogisticsIds() view returns(uint256[])
func (_Contracts *ContractsSession) GetAllLogisticsIds() ([]*big.Int, error) {
	return _Contracts.Contract.GetAllLogisticsIds(&_Contracts.CallOpts)
}

// GetAllLogisticsIds is a free data retrieval call binding the contract method 0x1b0612d5.
//
// Solidity: function getAllLogisticsIds() view returns(uint256[])
func (_Contracts *ContractsCallerSession) GetAllLogisticsIds() ([]*big.Int, error) {
	return _Contracts.Contract.GetAllLogisticsIds(&_Contracts.CallOpts)
}

// GetAllProductSkus is a free data retrieval call binding the contract method 0xdb10187b.
//
// Solidity: function getAllProductSkus() view returns(string[])
func (_Contracts *ContractsCaller) GetAllProductSkus(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllProductSkus")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllProductSkus is a free data retrieval call binding the contract method 0xdb10187b.
//
// Solidity: function getAllProductSkus() view returns(string[])
func (_Contracts *ContractsSession) GetAllProductSkus() ([]string, error) {
	return _Contracts.Contract.GetAllProductSkus(&_Contracts.CallOpts)
}

// GetAllProductSkus is a free data retrieval call binding the contract method 0xdb10187b.
//
// Solidity: function getAllProductSkus() view returns(string[])
func (_Contracts *ContractsCallerSession) GetAllProductSkus() ([]string, error) {
	return _Contracts.Contract.GetAllProductSkus(&_Contracts.CallOpts)
}

// GetBuyerCompletePurchaseHistory is a free data retrieval call binding the contract method 0xf521c1e5.
//
// Solidity: function getBuyerCompletePurchaseHistory(address _buyer) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsCaller) GetBuyerCompletePurchaseHistory(opts *bind.CallOpts, _buyer common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBuyerCompletePurchaseHistory", _buyer)

	outstruct := new(struct {
		Purchases []FrozenProductPurchaseRecord
		Products  []FrozenProductProduct
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Purchases = *abi.ConvertType(out[0], new([]FrozenProductPurchaseRecord)).(*[]FrozenProductPurchaseRecord)
	outstruct.Products = *abi.ConvertType(out[1], new([]FrozenProductProduct)).(*[]FrozenProductProduct)

	return *outstruct, err

}

// GetBuyerCompletePurchaseHistory is a free data retrieval call binding the contract method 0xf521c1e5.
//
// Solidity: function getBuyerCompletePurchaseHistory(address _buyer) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsSession) GetBuyerCompletePurchaseHistory(_buyer common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	return _Contracts.Contract.GetBuyerCompletePurchaseHistory(&_Contracts.CallOpts, _buyer)
}

// GetBuyerCompletePurchaseHistory is a free data retrieval call binding the contract method 0xf521c1e5.
//
// Solidity: function getBuyerCompletePurchaseHistory(address _buyer) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsCallerSession) GetBuyerCompletePurchaseHistory(_buyer common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	return _Contracts.Contract.GetBuyerCompletePurchaseHistory(&_Contracts.CallOpts, _buyer)
}

// GetBuyerPurchaseIds is a free data retrieval call binding the contract method 0xd433aba3.
//
// Solidity: function getBuyerPurchaseIds(address _buyer) view returns(uint256[])
func (_Contracts *ContractsCaller) GetBuyerPurchaseIds(opts *bind.CallOpts, _buyer common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBuyerPurchaseIds", _buyer)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBuyerPurchaseIds is a free data retrieval call binding the contract method 0xd433aba3.
//
// Solidity: function getBuyerPurchaseIds(address _buyer) view returns(uint256[])
func (_Contracts *ContractsSession) GetBuyerPurchaseIds(_buyer common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetBuyerPurchaseIds(&_Contracts.CallOpts, _buyer)
}

// GetBuyerPurchaseIds is a free data retrieval call binding the contract method 0xd433aba3.
//
// Solidity: function getBuyerPurchaseIds(address _buyer) view returns(uint256[])
func (_Contracts *ContractsCallerSession) GetBuyerPurchaseIds(_buyer common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetBuyerPurchaseIds(&_Contracts.CallOpts, _buyer)
}

// GetCompleteProductInfo is a free data retrieval call binding the contract method 0xfad305a8.
//
// Solidity: function getCompleteProductInfo(string _sku) view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool) product, (string,string,string,string,string,string,address,string,bool)[] logisticsRecords)
func (_Contracts *ContractsCaller) GetCompleteProductInfo(opts *bind.CallOpts, _sku string) (struct {
	Product          FrozenProductProduct
	LogisticsRecords []FrozenProductLogisticsRecord
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCompleteProductInfo", _sku)

	outstruct := new(struct {
		Product          FrozenProductProduct
		LogisticsRecords []FrozenProductLogisticsRecord
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Product = *abi.ConvertType(out[0], new(FrozenProductProduct)).(*FrozenProductProduct)
	outstruct.LogisticsRecords = *abi.ConvertType(out[1], new([]FrozenProductLogisticsRecord)).(*[]FrozenProductLogisticsRecord)

	return *outstruct, err

}

// GetCompleteProductInfo is a free data retrieval call binding the contract method 0xfad305a8.
//
// Solidity: function getCompleteProductInfo(string _sku) view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool) product, (string,string,string,string,string,string,address,string,bool)[] logisticsRecords)
func (_Contracts *ContractsSession) GetCompleteProductInfo(_sku string) (struct {
	Product          FrozenProductProduct
	LogisticsRecords []FrozenProductLogisticsRecord
}, error) {
	return _Contracts.Contract.GetCompleteProductInfo(&_Contracts.CallOpts, _sku)
}

// GetCompleteProductInfo is a free data retrieval call binding the contract method 0xfad305a8.
//
// Solidity: function getCompleteProductInfo(string _sku) view returns((string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool) product, (string,string,string,string,string,string,address,string,bool)[] logisticsRecords)
func (_Contracts *ContractsCallerSession) GetCompleteProductInfo(_sku string) (struct {
	Product          FrozenProductProduct
	LogisticsRecords []FrozenProductLogisticsRecord
}, error) {
	return _Contracts.Contract.GetCompleteProductInfo(&_Contracts.CallOpts, _sku)
}

// GetLogisticsBySku is a free data retrieval call binding the contract method 0x966bd197.
//
// Solidity: function getLogisticsBySku(string _sku) view returns((string,string,string,string,string,string,address,string,bool)[])
func (_Contracts *ContractsCaller) GetLogisticsBySku(opts *bind.CallOpts, _sku string) ([]FrozenProductLogisticsRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLogisticsBySku", _sku)

	if err != nil {
		return *new([]FrozenProductLogisticsRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FrozenProductLogisticsRecord)).(*[]FrozenProductLogisticsRecord)

	return out0, err

}

// GetLogisticsBySku is a free data retrieval call binding the contract method 0x966bd197.
//
// Solidity: function getLogisticsBySku(string _sku) view returns((string,string,string,string,string,string,address,string,bool)[])
func (_Contracts *ContractsSession) GetLogisticsBySku(_sku string) ([]FrozenProductLogisticsRecord, error) {
	return _Contracts.Contract.GetLogisticsBySku(&_Contracts.CallOpts, _sku)
}

// GetLogisticsBySku is a free data retrieval call binding the contract method 0x966bd197.
//
// Solidity: function getLogisticsBySku(string _sku) view returns((string,string,string,string,string,string,address,string,bool)[])
func (_Contracts *ContractsCallerSession) GetLogisticsBySku(_sku string) ([]FrozenProductLogisticsRecord, error) {
	return _Contracts.Contract.GetLogisticsBySku(&_Contracts.CallOpts, _sku)
}

// GetLogisticsRecord is a free data retrieval call binding the contract method 0xe64037fa.
//
// Solidity: function getLogisticsRecord(uint256 _logisticsId) view returns(string sku, string transportMethod, string carrier, string temperature, string departure, string destination, address operator, string operatorName, bool isValid)
func (_Contracts *ContractsCaller) GetLogisticsRecord(opts *bind.CallOpts, _logisticsId *big.Int) (struct {
	Sku             string
	TransportMethod string
	Carrier         string
	Temperature     string
	Departure       string
	Destination     string
	Operator        common.Address
	OperatorName    string
	IsValid         bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLogisticsRecord", _logisticsId)

	outstruct := new(struct {
		Sku             string
		TransportMethod string
		Carrier         string
		Temperature     string
		Departure       string
		Destination     string
		Operator        common.Address
		OperatorName    string
		IsValid         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sku = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TransportMethod = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Carrier = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Temperature = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Departure = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Destination = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Operator = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.OperatorName = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.IsValid = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetLogisticsRecord is a free data retrieval call binding the contract method 0xe64037fa.
//
// Solidity: function getLogisticsRecord(uint256 _logisticsId) view returns(string sku, string transportMethod, string carrier, string temperature, string departure, string destination, address operator, string operatorName, bool isValid)
func (_Contracts *ContractsSession) GetLogisticsRecord(_logisticsId *big.Int) (struct {
	Sku             string
	TransportMethod string
	Carrier         string
	Temperature     string
	Departure       string
	Destination     string
	Operator        common.Address
	OperatorName    string
	IsValid         bool
}, error) {
	return _Contracts.Contract.GetLogisticsRecord(&_Contracts.CallOpts, _logisticsId)
}

// GetLogisticsRecord is a free data retrieval call binding the contract method 0xe64037fa.
//
// Solidity: function getLogisticsRecord(uint256 _logisticsId) view returns(string sku, string transportMethod, string carrier, string temperature, string departure, string destination, address operator, string operatorName, bool isValid)
func (_Contracts *ContractsCallerSession) GetLogisticsRecord(_logisticsId *big.Int) (struct {
	Sku             string
	TransportMethod string
	Carrier         string
	Temperature     string
	Departure       string
	Destination     string
	Operator        common.Address
	OperatorName    string
	IsValid         bool
}, error) {
	return _Contracts.Contract.GetLogisticsRecord(&_Contracts.CallOpts, _logisticsId)
}

// GetProductBySku is a free data retrieval call binding the contract method 0xc900500d.
//
// Solidity: function getProductBySku(string _sku) view returns(string sku, string product_name, string brand, string specification, uint256 production_date, uint256 expiration_date, string raw_material_source, string processing_method, string logistics_temp, string storage_condition, string qc_report, string manufacturer_id, string image_url, string batchNumber, uint256 quantity, address current_holder, string current_location, bool isValid)
func (_Contracts *ContractsCaller) GetProductBySku(opts *bind.CallOpts, _sku string) (struct {
	Sku               string
	ProductName       string
	Brand             string
	Specification     string
	ProductionDate    *big.Int
	ExpirationDate    *big.Int
	RawMaterialSource string
	ProcessingMethod  string
	LogisticsTemp     string
	StorageCondition  string
	QcReport          string
	ManufacturerId    string
	ImageUrl          string
	BatchNumber       string
	Quantity          *big.Int
	CurrentHolder     common.Address
	CurrentLocation   string
	IsValid           bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProductBySku", _sku)

	outstruct := new(struct {
		Sku               string
		ProductName       string
		Brand             string
		Specification     string
		ProductionDate    *big.Int
		ExpirationDate    *big.Int
		RawMaterialSource string
		ProcessingMethod  string
		LogisticsTemp     string
		StorageCondition  string
		QcReport          string
		ManufacturerId    string
		ImageUrl          string
		BatchNumber       string
		Quantity          *big.Int
		CurrentHolder     common.Address
		CurrentLocation   string
		IsValid           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sku = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ProductName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Brand = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Specification = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ProductionDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ExpirationDate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RawMaterialSource = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.ProcessingMethod = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.LogisticsTemp = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.StorageCondition = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.QcReport = *abi.ConvertType(out[10], new(string)).(*string)
	outstruct.ManufacturerId = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.ImageUrl = *abi.ConvertType(out[12], new(string)).(*string)
	outstruct.BatchNumber = *abi.ConvertType(out[13], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)
	outstruct.CurrentHolder = *abi.ConvertType(out[15], new(common.Address)).(*common.Address)
	outstruct.CurrentLocation = *abi.ConvertType(out[16], new(string)).(*string)
	outstruct.IsValid = *abi.ConvertType(out[17], new(bool)).(*bool)

	return *outstruct, err

}

// GetProductBySku is a free data retrieval call binding the contract method 0xc900500d.
//
// Solidity: function getProductBySku(string _sku) view returns(string sku, string product_name, string brand, string specification, uint256 production_date, uint256 expiration_date, string raw_material_source, string processing_method, string logistics_temp, string storage_condition, string qc_report, string manufacturer_id, string image_url, string batchNumber, uint256 quantity, address current_holder, string current_location, bool isValid)
func (_Contracts *ContractsSession) GetProductBySku(_sku string) (struct {
	Sku               string
	ProductName       string
	Brand             string
	Specification     string
	ProductionDate    *big.Int
	ExpirationDate    *big.Int
	RawMaterialSource string
	ProcessingMethod  string
	LogisticsTemp     string
	StorageCondition  string
	QcReport          string
	ManufacturerId    string
	ImageUrl          string
	BatchNumber       string
	Quantity          *big.Int
	CurrentHolder     common.Address
	CurrentLocation   string
	IsValid           bool
}, error) {
	return _Contracts.Contract.GetProductBySku(&_Contracts.CallOpts, _sku)
}

// GetProductBySku is a free data retrieval call binding the contract method 0xc900500d.
//
// Solidity: function getProductBySku(string _sku) view returns(string sku, string product_name, string brand, string specification, uint256 production_date, uint256 expiration_date, string raw_material_source, string processing_method, string logistics_temp, string storage_condition, string qc_report, string manufacturer_id, string image_url, string batchNumber, uint256 quantity, address current_holder, string current_location, bool isValid)
func (_Contracts *ContractsCallerSession) GetProductBySku(_sku string) (struct {
	Sku               string
	ProductName       string
	Brand             string
	Specification     string
	ProductionDate    *big.Int
	ExpirationDate    *big.Int
	RawMaterialSource string
	ProcessingMethod  string
	LogisticsTemp     string
	StorageCondition  string
	QcReport          string
	ManufacturerId    string
	ImageUrl          string
	BatchNumber       string
	Quantity          *big.Int
	CurrentHolder     common.Address
	CurrentLocation   string
	IsValid           bool
}, error) {
	return _Contracts.Contract.GetProductBySku(&_Contracts.CallOpts, _sku)
}

// GetPurchaseRecord is a free data retrieval call binding the contract method 0x6ff8f2f7.
//
// Solidity: function getPurchaseRecord(uint256 _purchaseId) view returns(uint256 purchaseId, string sku, string product_name, uint256 quantity, uint256 purchase_date, string status, address seller_id, address buyer_id, bool isValid)
func (_Contracts *ContractsCaller) GetPurchaseRecord(opts *bind.CallOpts, _purchaseId *big.Int) (struct {
	PurchaseId   *big.Int
	Sku          string
	ProductName  string
	Quantity     *big.Int
	PurchaseDate *big.Int
	Status       string
	SellerId     common.Address
	BuyerId      common.Address
	IsValid      bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPurchaseRecord", _purchaseId)

	outstruct := new(struct {
		PurchaseId   *big.Int
		Sku          string
		ProductName  string
		Quantity     *big.Int
		PurchaseDate *big.Int
		Status       string
		SellerId     common.Address
		BuyerId      common.Address
		IsValid      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PurchaseId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Sku = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ProductName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Quantity = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PurchaseDate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.SellerId = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.BuyerId = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.IsValid = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetPurchaseRecord is a free data retrieval call binding the contract method 0x6ff8f2f7.
//
// Solidity: function getPurchaseRecord(uint256 _purchaseId) view returns(uint256 purchaseId, string sku, string product_name, uint256 quantity, uint256 purchase_date, string status, address seller_id, address buyer_id, bool isValid)
func (_Contracts *ContractsSession) GetPurchaseRecord(_purchaseId *big.Int) (struct {
	PurchaseId   *big.Int
	Sku          string
	ProductName  string
	Quantity     *big.Int
	PurchaseDate *big.Int
	Status       string
	SellerId     common.Address
	BuyerId      common.Address
	IsValid      bool
}, error) {
	return _Contracts.Contract.GetPurchaseRecord(&_Contracts.CallOpts, _purchaseId)
}

// GetPurchaseRecord is a free data retrieval call binding the contract method 0x6ff8f2f7.
//
// Solidity: function getPurchaseRecord(uint256 _purchaseId) view returns(uint256 purchaseId, string sku, string product_name, uint256 quantity, uint256 purchase_date, string status, address seller_id, address buyer_id, bool isValid)
func (_Contracts *ContractsCallerSession) GetPurchaseRecord(_purchaseId *big.Int) (struct {
	PurchaseId   *big.Int
	Sku          string
	ProductName  string
	Quantity     *big.Int
	PurchaseDate *big.Int
	Status       string
	SellerId     common.Address
	BuyerId      common.Address
	IsValid      bool
}, error) {
	return _Contracts.Contract.GetPurchaseRecord(&_Contracts.CallOpts, _purchaseId)
}

// GetSellerCompleteSalesHistory is a free data retrieval call binding the contract method 0x0a67b45b.
//
// Solidity: function getSellerCompleteSalesHistory(address _seller) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsCaller) GetSellerCompleteSalesHistory(opts *bind.CallOpts, _seller common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSellerCompleteSalesHistory", _seller)

	outstruct := new(struct {
		Purchases []FrozenProductPurchaseRecord
		Products  []FrozenProductProduct
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Purchases = *abi.ConvertType(out[0], new([]FrozenProductPurchaseRecord)).(*[]FrozenProductPurchaseRecord)
	outstruct.Products = *abi.ConvertType(out[1], new([]FrozenProductProduct)).(*[]FrozenProductProduct)

	return *outstruct, err

}

// GetSellerCompleteSalesHistory is a free data retrieval call binding the contract method 0x0a67b45b.
//
// Solidity: function getSellerCompleteSalesHistory(address _seller) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsSession) GetSellerCompleteSalesHistory(_seller common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	return _Contracts.Contract.GetSellerCompleteSalesHistory(&_Contracts.CallOpts, _seller)
}

// GetSellerCompleteSalesHistory is a free data retrieval call binding the contract method 0x0a67b45b.
//
// Solidity: function getSellerCompleteSalesHistory(address _seller) view returns((uint256,string,string,uint256,uint256,string,address,address,bool)[] purchases, (string,string,string,string,uint256,uint256,string,string,string,string,string,string,string,string,uint256,address,string,bool)[] products)
func (_Contracts *ContractsCallerSession) GetSellerCompleteSalesHistory(_seller common.Address) (struct {
	Purchases []FrozenProductPurchaseRecord
	Products  []FrozenProductProduct
}, error) {
	return _Contracts.Contract.GetSellerCompleteSalesHistory(&_Contracts.CallOpts, _seller)
}

// GetSellerPurchaseIds is a free data retrieval call binding the contract method 0xbc8bf20a.
//
// Solidity: function getSellerPurchaseIds(address _seller) view returns(uint256[])
func (_Contracts *ContractsCaller) GetSellerPurchaseIds(opts *bind.CallOpts, _seller common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSellerPurchaseIds", _seller)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSellerPurchaseIds is a free data retrieval call binding the contract method 0xbc8bf20a.
//
// Solidity: function getSellerPurchaseIds(address _seller) view returns(uint256[])
func (_Contracts *ContractsSession) GetSellerPurchaseIds(_seller common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetSellerPurchaseIds(&_Contracts.CallOpts, _seller)
}

// GetSellerPurchaseIds is a free data retrieval call binding the contract method 0xbc8bf20a.
//
// Solidity: function getSellerPurchaseIds(address _seller) view returns(uint256[])
func (_Contracts *ContractsCallerSession) GetSellerPurchaseIds(_seller common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetSellerPurchaseIds(&_Contracts.CallOpts, _seller)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// AddLogisticsRecord is a paid mutator transaction binding the contract method 0xad864423.
//
// Solidity: function addLogisticsRecord(string _sku, string _transportMethod, string _carrier, string _temperature, string _departure, string _destination, string _operatorName) returns()
func (_Contracts *ContractsTransactor) AddLogisticsRecord(opts *bind.TransactOpts, _sku string, _transportMethod string, _carrier string, _temperature string, _departure string, _destination string, _operatorName string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addLogisticsRecord", _sku, _transportMethod, _carrier, _temperature, _departure, _destination, _operatorName)
}

// AddLogisticsRecord is a paid mutator transaction binding the contract method 0xad864423.
//
// Solidity: function addLogisticsRecord(string _sku, string _transportMethod, string _carrier, string _temperature, string _departure, string _destination, string _operatorName) returns()
func (_Contracts *ContractsSession) AddLogisticsRecord(_sku string, _transportMethod string, _carrier string, _temperature string, _departure string, _destination string, _operatorName string) (*types.Transaction, error) {
	return _Contracts.Contract.AddLogisticsRecord(&_Contracts.TransactOpts, _sku, _transportMethod, _carrier, _temperature, _departure, _destination, _operatorName)
}

// AddLogisticsRecord is a paid mutator transaction binding the contract method 0xad864423.
//
// Solidity: function addLogisticsRecord(string _sku, string _transportMethod, string _carrier, string _temperature, string _departure, string _destination, string _operatorName) returns()
func (_Contracts *ContractsTransactorSession) AddLogisticsRecord(_sku string, _transportMethod string, _carrier string, _temperature string, _departure string, _destination string, _operatorName string) (*types.Transaction, error) {
	return _Contracts.Contract.AddLogisticsRecord(&_Contracts.TransactOpts, _sku, _transportMethod, _carrier, _temperature, _departure, _destination, _operatorName)
}

// AddProduct is a paid mutator transaction binding the contract method 0xc1cc2d21.
//
// Solidity: function addProduct(string _sku, string _product_name, string _brand, string _specification, uint256 _production_date, uint256 _expiration_date, string _raw_material_source, string _processing_method, string _logistics_temp, string _storage_condition, string _qc_report, string _manufacturer_id, string _image_url, string _batchNumber, uint256 _quantity, string _current_location) returns()
func (_Contracts *ContractsTransactor) AddProduct(opts *bind.TransactOpts, _sku string, _product_name string, _brand string, _specification string, _production_date *big.Int, _expiration_date *big.Int, _raw_material_source string, _processing_method string, _logistics_temp string, _storage_condition string, _qc_report string, _manufacturer_id string, _image_url string, _batchNumber string, _quantity *big.Int, _current_location string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addProduct", _sku, _product_name, _brand, _specification, _production_date, _expiration_date, _raw_material_source, _processing_method, _logistics_temp, _storage_condition, _qc_report, _manufacturer_id, _image_url, _batchNumber, _quantity, _current_location)
}

// AddProduct is a paid mutator transaction binding the contract method 0xc1cc2d21.
//
// Solidity: function addProduct(string _sku, string _product_name, string _brand, string _specification, uint256 _production_date, uint256 _expiration_date, string _raw_material_source, string _processing_method, string _logistics_temp, string _storage_condition, string _qc_report, string _manufacturer_id, string _image_url, string _batchNumber, uint256 _quantity, string _current_location) returns()
func (_Contracts *ContractsSession) AddProduct(_sku string, _product_name string, _brand string, _specification string, _production_date *big.Int, _expiration_date *big.Int, _raw_material_source string, _processing_method string, _logistics_temp string, _storage_condition string, _qc_report string, _manufacturer_id string, _image_url string, _batchNumber string, _quantity *big.Int, _current_location string) (*types.Transaction, error) {
	return _Contracts.Contract.AddProduct(&_Contracts.TransactOpts, _sku, _product_name, _brand, _specification, _production_date, _expiration_date, _raw_material_source, _processing_method, _logistics_temp, _storage_condition, _qc_report, _manufacturer_id, _image_url, _batchNumber, _quantity, _current_location)
}

// AddProduct is a paid mutator transaction binding the contract method 0xc1cc2d21.
//
// Solidity: function addProduct(string _sku, string _product_name, string _brand, string _specification, uint256 _production_date, uint256 _expiration_date, string _raw_material_source, string _processing_method, string _logistics_temp, string _storage_condition, string _qc_report, string _manufacturer_id, string _image_url, string _batchNumber, uint256 _quantity, string _current_location) returns()
func (_Contracts *ContractsTransactorSession) AddProduct(_sku string, _product_name string, _brand string, _specification string, _production_date *big.Int, _expiration_date *big.Int, _raw_material_source string, _processing_method string, _logistics_temp string, _storage_condition string, _qc_report string, _manufacturer_id string, _image_url string, _batchNumber string, _quantity *big.Int, _current_location string) (*types.Transaction, error) {
	return _Contracts.Contract.AddProduct(&_Contracts.TransactOpts, _sku, _product_name, _brand, _specification, _production_date, _expiration_date, _raw_material_source, _processing_method, _logistics_temp, _storage_condition, _qc_report, _manufacturer_id, _image_url, _batchNumber, _quantity, _current_location)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _purchaseId) returns()
func (_Contracts *ContractsTransactor) ConfirmPurchase(opts *bind.TransactOpts, _purchaseId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "confirmPurchase", _purchaseId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _purchaseId) returns()
func (_Contracts *ContractsSession) ConfirmPurchase(_purchaseId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfirmPurchase(&_Contracts.TransactOpts, _purchaseId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _purchaseId) returns()
func (_Contracts *ContractsTransactorSession) ConfirmPurchase(_purchaseId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ConfirmPurchase(&_Contracts.TransactOpts, _purchaseId)
}

// CreatePurchaseRecord is a paid mutator transaction binding the contract method 0xcbc9d8fa.
//
// Solidity: function createPurchaseRecord(string _sku, string _product_name, uint256 _quantity, address _seller_id, address _buyer_id) returns()
func (_Contracts *ContractsTransactor) CreatePurchaseRecord(opts *bind.TransactOpts, _sku string, _product_name string, _quantity *big.Int, _seller_id common.Address, _buyer_id common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createPurchaseRecord", _sku, _product_name, _quantity, _seller_id, _buyer_id)
}

// CreatePurchaseRecord is a paid mutator transaction binding the contract method 0xcbc9d8fa.
//
// Solidity: function createPurchaseRecord(string _sku, string _product_name, uint256 _quantity, address _seller_id, address _buyer_id) returns()
func (_Contracts *ContractsSession) CreatePurchaseRecord(_sku string, _product_name string, _quantity *big.Int, _seller_id common.Address, _buyer_id common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePurchaseRecord(&_Contracts.TransactOpts, _sku, _product_name, _quantity, _seller_id, _buyer_id)
}

// CreatePurchaseRecord is a paid mutator transaction binding the contract method 0xcbc9d8fa.
//
// Solidity: function createPurchaseRecord(string _sku, string _product_name, uint256 _quantity, address _seller_id, address _buyer_id) returns()
func (_Contracts *ContractsTransactorSession) CreatePurchaseRecord(_sku string, _product_name string, _quantity *big.Int, _seller_id common.Address, _buyer_id common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CreatePurchaseRecord(&_Contracts.TransactOpts, _sku, _product_name, _quantity, _seller_id, _buyer_id)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// UpdateProductHolder is a paid mutator transaction binding the contract method 0xecb5f489.
//
// Solidity: function updateProductHolder(string _sku, address _newHolder, string _newLocation) returns()
func (_Contracts *ContractsTransactor) UpdateProductHolder(opts *bind.TransactOpts, _sku string, _newHolder common.Address, _newLocation string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateProductHolder", _sku, _newHolder, _newLocation)
}

// UpdateProductHolder is a paid mutator transaction binding the contract method 0xecb5f489.
//
// Solidity: function updateProductHolder(string _sku, address _newHolder, string _newLocation) returns()
func (_Contracts *ContractsSession) UpdateProductHolder(_sku string, _newHolder common.Address, _newLocation string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProductHolder(&_Contracts.TransactOpts, _sku, _newHolder, _newLocation)
}

// UpdateProductHolder is a paid mutator transaction binding the contract method 0xecb5f489.
//
// Solidity: function updateProductHolder(string _sku, address _newHolder, string _newLocation) returns()
func (_Contracts *ContractsTransactorSession) UpdateProductHolder(_sku string, _newHolder common.Address, _newLocation string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProductHolder(&_Contracts.TransactOpts, _sku, _newHolder, _newLocation)
}

// ContractsLogisticsAddedIterator is returned from FilterLogisticsAdded and is used to iterate over the raw logs and unpacked data for LogisticsAdded events raised by the Contracts contract.
type ContractsLogisticsAddedIterator struct {
	Event *ContractsLogisticsAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogisticsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogisticsAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogisticsAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogisticsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogisticsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogisticsAdded represents a LogisticsAdded event raised by the Contracts contract.
type ContractsLogisticsAdded struct {
	LogisticsId *big.Int
	Sku         string
	Carrier     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogisticsAdded is a free log retrieval operation binding the contract event 0xc4f6b94487bd6f57ccd5df84506fd060f8c3c6e9b45cc0d50727219602e5d605.
//
// Solidity: event LogisticsAdded(uint256 logisticsId, string sku, string carrier)
func (_Contracts *ContractsFilterer) FilterLogisticsAdded(opts *bind.FilterOpts) (*ContractsLogisticsAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogisticsAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsLogisticsAddedIterator{contract: _Contracts.contract, event: "LogisticsAdded", logs: logs, sub: sub}, nil
}

// WatchLogisticsAdded is a free log subscription operation binding the contract event 0xc4f6b94487bd6f57ccd5df84506fd060f8c3c6e9b45cc0d50727219602e5d605.
//
// Solidity: event LogisticsAdded(uint256 logisticsId, string sku, string carrier)
func (_Contracts *ContractsFilterer) WatchLogisticsAdded(opts *bind.WatchOpts, sink chan<- *ContractsLogisticsAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogisticsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogisticsAdded)
				if err := _Contracts.contract.UnpackLog(event, "LogisticsAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogisticsAdded is a log parse operation binding the contract event 0xc4f6b94487bd6f57ccd5df84506fd060f8c3c6e9b45cc0d50727219602e5d605.
//
// Solidity: event LogisticsAdded(uint256 logisticsId, string sku, string carrier)
func (_Contracts *ContractsFilterer) ParseLogisticsAdded(log types.Log) (*ContractsLogisticsAdded, error) {
	event := new(ContractsLogisticsAdded)
	if err := _Contracts.contract.UnpackLog(event, "LogisticsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProductAddedIterator is returned from FilterProductAdded and is used to iterate over the raw logs and unpacked data for ProductAdded events raised by the Contracts contract.
type ContractsProductAddedIterator struct {
	Event *ContractsProductAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsProductAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsProductAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsProductAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsProductAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProductAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsProductAdded represents a ProductAdded event raised by the Contracts contract.
type ContractsProductAdded struct {
	Sku         string
	ProductName string
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProductAdded is a free log retrieval operation binding the contract event 0xcbebef04e1d7de535a91cb84796482718220565630fdbc1954fbf9d97f043453.
//
// Solidity: event ProductAdded(string sku, string product_name, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterProductAdded(opts *bind.FilterOpts) (*ContractsProductAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ProductAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsProductAddedIterator{contract: _Contracts.contract, event: "ProductAdded", logs: logs, sub: sub}, nil
}

// WatchProductAdded is a free log subscription operation binding the contract event 0xcbebef04e1d7de535a91cb84796482718220565630fdbc1954fbf9d97f043453.
//
// Solidity: event ProductAdded(string sku, string product_name, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchProductAdded(opts *bind.WatchOpts, sink chan<- *ContractsProductAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ProductAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsProductAdded)
				if err := _Contracts.contract.UnpackLog(event, "ProductAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProductAdded is a log parse operation binding the contract event 0xcbebef04e1d7de535a91cb84796482718220565630fdbc1954fbf9d97f043453.
//
// Solidity: event ProductAdded(string sku, string product_name, uint256 timestamp)
func (_Contracts *ContractsFilterer) ParseProductAdded(log types.Log) (*ContractsProductAdded, error) {
	event := new(ContractsProductAdded)
	if err := _Contracts.contract.UnpackLog(event, "ProductAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPurchaseConfirmedIterator is returned from FilterPurchaseConfirmed and is used to iterate over the raw logs and unpacked data for PurchaseConfirmed events raised by the Contracts contract.
type ContractsPurchaseConfirmedIterator struct {
	Event *ContractsPurchaseConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPurchaseConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPurchaseConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPurchaseConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPurchaseConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPurchaseConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPurchaseConfirmed represents a PurchaseConfirmed event raised by the Contracts contract.
type ContractsPurchaseConfirmed struct {
	PurchaseId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchaseConfirmed is a free log retrieval operation binding the contract event 0x2ce3b121b7afbc2724c6c53c061e65d64d4883db75d2f497a715b144a8404275.
//
// Solidity: event PurchaseConfirmed(uint256 purchaseId)
func (_Contracts *ContractsFilterer) FilterPurchaseConfirmed(opts *bind.FilterOpts) (*ContractsPurchaseConfirmedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return &ContractsPurchaseConfirmedIterator{contract: _Contracts.contract, event: "PurchaseConfirmed", logs: logs, sub: sub}, nil
}

// WatchPurchaseConfirmed is a free log subscription operation binding the contract event 0x2ce3b121b7afbc2724c6c53c061e65d64d4883db75d2f497a715b144a8404275.
//
// Solidity: event PurchaseConfirmed(uint256 purchaseId)
func (_Contracts *ContractsFilterer) WatchPurchaseConfirmed(opts *bind.WatchOpts, sink chan<- *ContractsPurchaseConfirmed) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PurchaseConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPurchaseConfirmed)
				if err := _Contracts.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePurchaseConfirmed is a log parse operation binding the contract event 0x2ce3b121b7afbc2724c6c53c061e65d64d4883db75d2f497a715b144a8404275.
//
// Solidity: event PurchaseConfirmed(uint256 purchaseId)
func (_Contracts *ContractsFilterer) ParsePurchaseConfirmed(log types.Log) (*ContractsPurchaseConfirmed, error) {
	event := new(ContractsPurchaseConfirmed)
	if err := _Contracts.contract.UnpackLog(event, "PurchaseConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPurchaseCreatedIterator is returned from FilterPurchaseCreated and is used to iterate over the raw logs and unpacked data for PurchaseCreated events raised by the Contracts contract.
type ContractsPurchaseCreatedIterator struct {
	Event *ContractsPurchaseCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsPurchaseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPurchaseCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsPurchaseCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsPurchaseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPurchaseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPurchaseCreated represents a PurchaseCreated event raised by the Contracts contract.
type ContractsPurchaseCreated struct {
	PurchaseId *big.Int
	Sku        string
	Quantity   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCreated is a free log retrieval operation binding the contract event 0xe89d7cb7157508bfc3241a76bb061488782a9342e89e3ddb531745ef904e9339.
//
// Solidity: event PurchaseCreated(uint256 purchaseId, string sku, uint256 quantity)
func (_Contracts *ContractsFilterer) FilterPurchaseCreated(opts *bind.FilterOpts) (*ContractsPurchaseCreatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PurchaseCreated")
	if err != nil {
		return nil, err
	}
	return &ContractsPurchaseCreatedIterator{contract: _Contracts.contract, event: "PurchaseCreated", logs: logs, sub: sub}, nil
}

// WatchPurchaseCreated is a free log subscription operation binding the contract event 0xe89d7cb7157508bfc3241a76bb061488782a9342e89e3ddb531745ef904e9339.
//
// Solidity: event PurchaseCreated(uint256 purchaseId, string sku, uint256 quantity)
func (_Contracts *ContractsFilterer) WatchPurchaseCreated(opts *bind.WatchOpts, sink chan<- *ContractsPurchaseCreated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PurchaseCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPurchaseCreated)
				if err := _Contracts.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePurchaseCreated is a log parse operation binding the contract event 0xe89d7cb7157508bfc3241a76bb061488782a9342e89e3ddb531745ef904e9339.
//
// Solidity: event PurchaseCreated(uint256 purchaseId, string sku, uint256 quantity)
func (_Contracts *ContractsFilterer) ParsePurchaseCreated(log types.Log) (*ContractsPurchaseCreated, error) {
	event := new(ContractsPurchaseCreated)
	if err := _Contracts.contract.UnpackLog(event, "PurchaseCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
