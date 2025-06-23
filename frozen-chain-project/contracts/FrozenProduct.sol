// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title FrozenProduct
 * @dev 冷链物流智能合约，整合冷冻品和物流信息管理
 */
contract FrozenProduct is Ownable {
    // 冷冻品结构体
    struct Product {
        string sku; // 唯一SKU码
        string product_name; // 产品名称
        string brand; // 品牌
        string specification; // 规格/包装
        uint256 production_date; // 生产日期
        uint256 expiration_date; // 保质期
        string raw_material_source; // 原材料来源
        string processing_method; // 加工方式
        string logistics_temp; // 运输温度要求
        string storage_condition; // 仓储条件
        string qc_report; // 质检报告
        string manufacturer_id; // 生产商ID
        string image_url; // 产品图片URL
        string batchNumber; // 批次号
        uint256 quantity; // 数量
        address current_holder; // 当前持有者
        string current_location; // 当前位置
        bool isValid; // 记录是否有效
    }

    // 物流记录结构体
    struct LogisticsRecord {
        string sku; // 关联的产品SKU
        string transportMethod; // 运输方式
        string carrier; // 运输公司
        string temperature; // 运输温度
        string departure; // 出发地
        string destination; // 目的地
        address operator; // 操作方地址
        string operatorName; // 操作方名称
        bool isValid; // 记录是否有效
    }

    // 采购记录结构体
    struct PurchaseRecord {
        uint256 purchaseId; // 采购ID
        string sku; // SKU码
        string product_name; // 产品名称
        uint256 quantity; // 采购数量
        uint256 purchase_date; // 采购日期
        string status; // 状态
        address seller_id; // 卖家ID
        address buyer_id; // 买家ID
        bool isValid; // 记录是否有效
    }

    // SKU到产品的映射
    mapping(string => Product) private productsBySku;
    // 所有有效产品的SKU列表
    string[] private allProductSkus;
    
    // SKU到物流记录数组的映射
    mapping(string => LogisticsRecord[]) private skuToLogistics;
    // 物流ID到物流记录的映射
    mapping(uint256 => LogisticsRecord) private logisticsById;
    // 所有物流记录ID
    uint256[] private allLogisticsIds;
    
    // 采购ID到采购记录的映射
    mapping(uint256 => PurchaseRecord) private purchaseRecordsById;
    // 买家地址到采购记录ID列表的映射
    mapping(address => uint256[]) private buyerToPurchaseIds;
    // 卖家地址到采购记录ID列表的映射
    mapping(address => uint256[]) private sellerToPurchaseIds;
    // 所有采购记录ID
    uint256[] private allPurchaseIds;

    // 事件
    event ProductAdded(string sku, string product_name, uint256 timestamp);
    event LogisticsAdded(uint256 logisticsId, string sku, string carrier);
    event PurchaseCreated(uint256 purchaseId, string sku, uint256 quantity);
    event PurchaseConfirmed(uint256 purchaseId);

    /**
     * @dev 构造函数
     */
    constructor() Ownable(msg.sender) {}

    /**
     * ====================== 冷冻品管理方法 ======================
     */

    /**
     * @dev 添加新的冷冻品
     */
    function addProduct(
        string memory _sku,
        string memory _product_name,
        string memory _brand,
        string memory _specification,
        uint256 _production_date,
        uint256 _expiration_date,
        string memory _raw_material_source,
        string memory _processing_method,
        string memory _logistics_temp,
        string memory _storage_condition,
        string memory _qc_report,
        string memory _manufacturer_id,
        string memory _image_url,
        string memory _batchNumber,
        uint256 _quantity,
        string memory _current_location
    ) public onlyOwner {
        // 确保SKU不为空且唯一
        require(bytes(_sku).length > 0, "SKU cannot be empty");
        require(!productsBySku[_sku].isValid, "Product with this SKU already exists");
        
        // 确保生产日期和保质期有效
        require(_production_date > 0, "Production date must be valid");
        require(_expiration_date > _production_date, "Expiration date must be after production date");

        // 创建新产品
        Product memory newProduct = Product({
            sku: _sku,
            product_name: _product_name,
            brand: _brand,
            specification: _specification,
            production_date: _production_date,
            expiration_date: _expiration_date,
            raw_material_source: _raw_material_source,
            processing_method: _processing_method,
            logistics_temp: _logistics_temp,
            storage_condition: _storage_condition,
            qc_report: _qc_report,
            manufacturer_id: _manufacturer_id,
            image_url: _image_url,
            batchNumber: _batchNumber,
            quantity: _quantity,
            current_holder: msg.sender,
            current_location: _current_location,
            isValid: true
        });

        // 添加到映射和数组中
        productsBySku[_sku] = newProduct;
        allProductSkus.push(_sku);

        // 触发事件
        emit ProductAdded(_sku, _product_name, block.timestamp);
    }

    /**
     * @dev 更新产品持有者和位置
     */
    function updateProductHolder(
        string memory _sku,
        address _newHolder,
        string memory _newLocation
    ) public {
        // 确保产品存在
        require(productsBySku[_sku].isValid, "Product does not exist");
        
        // 只有当前持有者可以更新
        require(productsBySku[_sku].current_holder == msg.sender, "Only current holder can update");
        
        // 更新持有者和位置
        productsBySku[_sku].current_holder = _newHolder;
        productsBySku[_sku].current_location = _newLocation;
    }

    /**
     * @dev 获取指定SKU的产品详情
     */
    function getProductBySku(
        string memory _sku
    ) public view returns (
        string memory sku,
        string memory product_name,
        string memory brand,
        string memory specification,
        uint256 production_date,
        uint256 expiration_date,
        string memory raw_material_source,
        string memory processing_method,
        string memory logistics_temp,
        string memory storage_condition,
        string memory qc_report,
        string memory manufacturer_id,
        string memory image_url,
        string memory batchNumber,
        uint256 quantity,
        address current_holder,
        string memory current_location,
        bool isValid
    ) {
        // 确保产品存在
        require(productsBySku[_sku].isValid, "Product does not exist");
        
        Product memory product = productsBySku[_sku];
        return (
            product.sku,
            product.product_name,
            product.brand,
            product.specification,
            product.production_date,
            product.expiration_date,
            product.raw_material_source,
            product.processing_method,
            product.logistics_temp,
            product.storage_condition,
            product.qc_report,
            product.manufacturer_id,
            product.image_url,
            product.batchNumber,
            product.quantity,
            product.current_holder,
            product.current_location,
            product.isValid
        );
    }

    /**
     * @dev 获取所有产品SKU
     */
    function getAllProductSkus() public view returns (string[] memory) {
        return allProductSkus;
    }

    /**
     * ====================== 物流管理方法 ======================
     */

    /**
     * @dev 添加新的物流记录
     */
    function addLogisticsRecord(
        string memory _sku,
        string memory _transportMethod,
        string memory _carrier,
        string memory _temperature,
        string memory _departure,
        string memory _destination,
        string memory _operatorName
    ) public {
        // 确保SKU不为空且产品存在
        require(bytes(_sku).length > 0, "SKU cannot be empty");
        require(productsBySku[_sku].isValid, "Product does not exist");

        uint256 logisticsId = allLogisticsIds.length + 1;

        // 创建新的物流记录
        LogisticsRecord memory newLogistics = LogisticsRecord({
            sku: _sku,
            transportMethod: _transportMethod,
            carrier: _carrier,
            temperature: _temperature,
            departure: _departure,
            destination: _destination,
            operator: msg.sender,
            operatorName: _operatorName,
            isValid: true
        });

        // 添加到映射和数组中
        skuToLogistics[_sku].push(newLogistics);
        logisticsById[logisticsId] = newLogistics;
        allLogisticsIds.push(logisticsId);

        // 触发事件
        emit LogisticsAdded(logisticsId, _sku, _carrier);
    }

    /**
     * @dev 获取指定SKU的所有物流记录
     */
    function getLogisticsBySku(
        string memory _sku
    ) public view returns (LogisticsRecord[] memory) {
        return skuToLogistics[_sku];
    }

    /**
     * @dev 获取指定物流记录详情
     */
    function getLogisticsRecord(
        uint256 _logisticsId
    ) 
        public 
        view 
        returns (
            string memory sku,
            string memory transportMethod,
            string memory carrier,
            string memory temperature,
            string memory departure,
            string memory destination,
            address operator,
            string memory operatorName,
            bool isValid
        ) 
    {
        // 确保物流记录存在
        require(logisticsById[_logisticsId].isValid, "Logistics record does not exist");
        
        LogisticsRecord memory logistics = logisticsById[_logisticsId];
        return (
            logistics.sku,
            logistics.transportMethod,
            logistics.carrier,
            logistics.temperature,
            logistics.departure,
            logistics.destination,
            logistics.operator,
            logistics.operatorName,
            logistics.isValid
        );
    }

    /**
     * @dev 获取所有物流记录ID
     */
    function getAllLogisticsIds() public view returns (uint256[] memory) {
        return allLogisticsIds;
    }

    /**
     * ====================== 采购管理方法 ======================
     */

    /**
     * @dev 创建采购记录
     */
    function createPurchaseRecord(
        string memory _sku,
        string memory _product_name,
        uint256 _quantity,
        address _seller_id,
        address _buyer_id
    ) public {
        // 确保产品存在且数量足够
        require(productsBySku[_sku].isValid, "Product does not exist");
        require(productsBySku[_sku].quantity >= _quantity, "Insufficient quantity");
        
        uint256 purchaseId = allPurchaseIds.length + 1;
        
        // 创建采购记录
        PurchaseRecord memory newPurchase = PurchaseRecord({
            purchaseId: purchaseId,
            sku: _sku,
            product_name: _product_name,
            quantity: _quantity,
            purchase_date: block.timestamp,
            status: "pending",
            seller_id: _seller_id,
            buyer_id: _buyer_id,
            isValid: true
        });
        
        // 添加到映射和数组中
        purchaseRecordsById[purchaseId] = newPurchase;
        buyerToPurchaseIds[_buyer_id].push(purchaseId);
        sellerToPurchaseIds[_seller_id].push(purchaseId);
        allPurchaseIds.push(purchaseId);
        
        // 触发事件
        emit PurchaseCreated(purchaseId, _sku, _quantity);
    }
    
    /**
     * @dev 确认采购
     */
    function confirmPurchase(uint256 _purchaseId) public {
        // 确保采购记录存在
        require(purchaseRecordsById[_purchaseId].isValid, "Purchase record does not exist");
        
        PurchaseRecord storage purchase = purchaseRecordsById[_purchaseId];
        
        // 只有买家可以确认
        require(purchase.buyer_id == msg.sender, "Only buyer can confirm purchase");
        
        // 更新状态
        purchase.status = "confirmed";
        
        // 更新产品信息
        Product storage product = productsBySku[purchase.sku];
        product.quantity -= purchase.quantity;
        product.current_holder = purchase.buyer_id;
        
        // 触发事件
        emit PurchaseConfirmed(_purchaseId);
    }
    
    /**
     * @dev 获取指定采购记录
     */
    function getPurchaseRecord(uint256 _purchaseId) public view returns (
        uint256 purchaseId,
        string memory sku,
        string memory product_name,
        uint256 quantity,
        uint256 purchase_date,
        string memory status,
        address seller_id,
        address buyer_id,
        bool isValid
    ) {
        // 确保采购记录存在
        require(purchaseRecordsById[_purchaseId].isValid, "Purchase record does not exist");
        
        PurchaseRecord memory purchase = purchaseRecordsById[_purchaseId];
        return (
            purchase.purchaseId,
            purchase.sku,
            purchase.product_name,
            purchase.quantity,
            purchase.purchase_date,
            purchase.status,
            purchase.seller_id,
            purchase.buyer_id,
            purchase.isValid
        );
    }
    
    /**
     * @dev 获取买家的所有采购记录ID
     */
    function getBuyerPurchaseIds(address _buyer) public view returns (uint256[] memory) {
        return buyerToPurchaseIds[_buyer];
    }
    
    /**
     * @dev 获取卖家的所有采购记录ID
     */
    function getSellerPurchaseIds(address _seller) public view returns (uint256[] memory) {
        return sellerToPurchaseIds[_seller];
    }

    /**
     * ====================== 统一查询方法 ======================
     */
    
    /**
     * @dev 获取指定SKU的完整信息（产品+物流）
     */
    function getCompleteProductInfo(string memory _sku) public view returns (
        Product memory product,
        LogisticsRecord[] memory logisticsRecords
    ) {
        // 确保产品存在
        require(productsBySku[_sku].isValid, "Product does not exist");
        
        return (productsBySku[_sku], skuToLogistics[_sku]);
    }
    
    /**
     * @dev 获取买家的完整采购历史（包含产品信息）
     */
    function getBuyerCompletePurchaseHistory(address _buyer) public view returns (
        PurchaseRecord[] memory purchases,
        Product[] memory products
    ) {
        uint256[] memory purchaseIds = buyerToPurchaseIds[_buyer];
        PurchaseRecord[] memory resultPurchases = new PurchaseRecord[](purchaseIds.length);
        Product[] memory resultProducts = new Product[](purchaseIds.length);
        
        for (uint256 i = 0; i < purchaseIds.length; i++) {
            resultPurchases[i] = purchaseRecordsById[purchaseIds[i]];
            resultProducts[i] = productsBySku[resultPurchases[i].sku];
        }
        
        return (resultPurchases, resultProducts);
    }
    
    /**
     * @dev 获取卖家的完整销售历史（包含产品信息）
     */
    function getSellerCompleteSalesHistory(address _seller) public view returns (
        PurchaseRecord[] memory purchases,
        Product[] memory products
    ) {
        uint256[] memory purchaseIds = sellerToPurchaseIds[_seller];
        PurchaseRecord[] memory resultPurchases = new PurchaseRecord[](purchaseIds.length);
        Product[] memory resultProducts = new Product[](purchaseIds.length);
        
        for (uint256 i = 0; i < purchaseIds.length; i++) {
            resultPurchases[i] = purchaseRecordsById[purchaseIds[i]];
            resultProducts[i] = productsBySku[resultPurchases[i].sku];
        }
        
        return (resultPurchases, resultProducts);
    }
    
    /**
     * @dev 获取所有可用产品（数量大于0）
     */
    function getAllAvailableProducts() public view returns (Product[] memory) {
        uint256 availableCount = 0;
        
        // 计算可用产品数量
        for (uint256 i = 0; i < allProductSkus.length; i++) {
            if (productsBySku[allProductSkus[i]].quantity > 0 && productsBySku[allProductSkus[i]].isValid) {
                availableCount++;
            }
        }
        
        // 创建结果数组
        Product[] memory availableProducts = new Product[](availableCount);
        uint256 index = 0;
        
        // 填充结果数组
        for (uint256 i = 0; i < allProductSkus.length; i++) {
            Product memory product = productsBySku[allProductSkus[i]];
            if (product.quantity > 0 && product.isValid) {
                availableProducts[index] = product;
                index++;
            }
        }
        
        return availableProducts;
    }
}