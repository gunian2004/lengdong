import * as ethers from "ethers";

// 导入合约ABI
const FrozenProductABI = {
  abi: [
    // 获取所有产品SKU
    {
      inputs: [],
      name: "getAllProductSkus",
      outputs: [{ internalType: "string[]", name: "", type: "string[]" }],
      stateMutability: "view",
      type: "function",
    },
    // 获取产品详情
    {
      inputs: [{ internalType: "string", name: "_sku", type: "string" }],
      name: "getProductBySku",
      outputs: [
        { internalType: "string", name: "sku", type: "string" },
        { internalType: "string", name: "product_name", type: "string" },
        { internalType: "string", name: "brand", type: "string" },
        { internalType: "string", name: "specification", type: "string" },
        { internalType: "uint256", name: "production_date", type: "uint256" },
        { internalType: "uint256", name: "expiration_date", type: "uint256" },
        { internalType: "string", name: "raw_material_source", type: "string" },
        { internalType: "string", name: "processing_method", type: "string" },
        { internalType: "string", name: "logistics_temp", type: "string" },
        { internalType: "string", name: "storage_condition", type: "string" },
        { internalType: "string", name: "qc_report", type: "string" },
        { internalType: "string", name: "manufacturer_id", type: "string" },
        { internalType: "string", name: "image_url", type: "string" },
        { internalType: "string", name: "batchNumber", type: "string" },
        { internalType: "uint256", name: "quantity", type: "uint256" },
        { internalType: "address", name: "current_holder", type: "address" },
        { internalType: "string", name: "current_location", type: "string" },
        { internalType: "bool", name: "isValid", type: "bool" },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 添加产品
    {
      inputs: [
        { internalType: "string", name: "_sku", type: "string" },
        { internalType: "string", name: "_product_name", type: "string" },
        { internalType: "string", name: "_brand", type: "string" },
        { internalType: "string", name: "_specification", type: "string" },
        { internalType: "uint256", name: "_production_date", type: "uint256" },
        { internalType: "uint256", name: "_expiration_date", type: "uint256" },
        { internalType: "string", name: "_raw_material_source", type: "string" },
        { internalType: "string", name: "_processing_method", type: "string" },
        { internalType: "string", name: "_logistics_temp", type: "string" },
        { internalType: "string", name: "_storage_condition", type: "string" },
        { internalType: "string", name: "_qc_report", type: "string" },
        { internalType: "string", name: "_manufacturer_id", type: "string" },
        { internalType: "string", name: "_image_url", type: "string" },
        { internalType: "string", name: "_batchNumber", type: "string" },
        { internalType: "uint256", name: "_quantity", type: "uint256" },
        { internalType: "string", name: "_current_location", type: "string" },
      ],
      name: "addProduct",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    // 更新产品持有者
    {
      inputs: [
        { internalType: "string", name: "_sku", type: "string" },
        { internalType: "address", name: "_newHolder", type: "address" },
        { internalType: "string", name: "_newLocation", type: "string" },
      ],
      name: "updateProductHolder",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    // 添加物流记录
    {
      inputs: [
        { internalType: "string", name: "_sku", type: "string" },
        { internalType: "string", name: "_transportMethod", type: "string" },
        { internalType: "string", name: "_carrier", type: "string" },
        { internalType: "string", name: "_temperature", type: "string" },
        { internalType: "string", name: "_departure", type: "string" },
        { internalType: "string", name: "_destination", type: "string" },
        { internalType: "string", name: "_operatorName", type: "string" },
      ],
      name: "addLogisticsRecord",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    // 获取指定SKU的物流记录
    {
      inputs: [{ internalType: "string", name: "_sku", type: "string" }],
      name: "getLogisticsBySku",
      outputs: [
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "transportMethod", type: "string" },
            { internalType: "string", name: "carrier", type: "string" },
            { internalType: "string", name: "temperature", type: "string" },
            { internalType: "string", name: "departure", type: "string" },
            { internalType: "string", name: "destination", type: "string" },
            { internalType: "address", name: "operator", type: "address" },
            { internalType: "string", name: "operatorName", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.LogisticsRecord[]",
          name: "",
          type: "tuple[]",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 获取完整产品信息（产品+物流）
    {
      inputs: [{ internalType: "string", name: "_sku", type: "string" }],
      name: "getCompleteProductInfo",
      outputs: [
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "string", name: "brand", type: "string" },
            { internalType: "string", name: "specification", type: "string" },
            { internalType: "uint256", name: "production_date", type: "uint256" },
            { internalType: "uint256", name: "expiration_date", type: "uint256" },
            { internalType: "string", name: "raw_material_source", type: "string" },
            { internalType: "string", name: "processing_method", type: "string" },
            { internalType: "string", name: "logistics_temp", type: "string" },
            { internalType: "string", name: "storage_condition", type: "string" },
            { internalType: "string", name: "qc_report", type: "string" },
            { internalType: "string", name: "manufacturer_id", type: "string" },
            { internalType: "string", name: "image_url", type: "string" },
            { internalType: "string", name: "batchNumber", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "address", name: "current_holder", type: "address" },
            { internalType: "string", name: "current_location", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.Product",
          name: "product",
          type: "tuple",
        },
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "transportMethod", type: "string" },
            { internalType: "string", name: "carrier", type: "string" },
            { internalType: "string", name: "temperature", type: "string" },
            { internalType: "string", name: "departure", type: "string" },
            { internalType: "string", name: "destination", type: "string" },
            { internalType: "address", name: "operator", type: "address" },
            { internalType: "string", name: "operatorName", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.LogisticsRecord[]",
          name: "logisticsRecords",
          type: "tuple[]",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 创建采购记录
    {
      inputs: [
        { internalType: "string", name: "_sku", type: "string" },
        { internalType: "string", name: "_product_name", type: "string" },
        { internalType: "uint256", name: "_quantity", type: "uint256" },
        { internalType: "address", name: "_seller_id", type: "address" },
        { internalType: "address", name: "_buyer_id", type: "address" },
      ],
      name: "createPurchaseRecord",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    // 确认采购
    {
      inputs: [{ internalType: "uint256", name: "_purchaseId", type: "uint256" }],
      name: "confirmPurchase",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    // 获取买家的所有采购记录ID
    {
      inputs: [{ internalType: "address", name: "_buyer", type: "address" }],
      name: "getBuyerPurchaseIds",
      outputs: [{ internalType: "uint256[]", name: "", type: "uint256[]" }],
      stateMutability: "view",
      type: "function",
    },
    // 获取卖家的所有采购记录ID
    {
      inputs: [{ internalType: "address", name: "_seller", type: "address" }],
      name: "getSellerPurchaseIds",
      outputs: [{ internalType: "uint256[]", name: "", type: "uint256[]" }],
      stateMutability: "view",
      type: "function",
    },
    // 获取指定采购记录
    {
      inputs: [{ internalType: "uint256", name: "_purchaseId", type: "uint256" }],
      name: "getPurchaseRecord",
      outputs: [
        { internalType: "uint256", name: "purchaseId", type: "uint256" },
        { internalType: "string", name: "sku", type: "string" },
        { internalType: "string", name: "product_name", type: "string" },
        { internalType: "uint256", name: "quantity", type: "uint256" },
        { internalType: "uint256", name: "purchase_date", type: "uint256" },
        { internalType: "string", name: "status", type: "string" },
        { internalType: "address", name: "seller_id", type: "address" },
        { internalType: "address", name: "buyer_id", type: "address" },
        { internalType: "bool", name: "isValid", type: "bool" },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 获取买家的完整采购历史
    {
      inputs: [{ internalType: "address", name: "_buyer", type: "address" }],
      name: "getBuyerCompletePurchaseHistory",
      outputs: [
        {
          components: [
            { internalType: "uint256", name: "purchaseId", type: "uint256" },
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "uint256", name: "purchase_date", type: "uint256" },
            { internalType: "string", name: "status", type: "string" },
            { internalType: "address", name: "seller_id", type: "address" },
            { internalType: "address", name: "buyer_id", type: "address" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.PurchaseRecord[]",
          name: "purchases",
          type: "tuple[]",
        },
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "string", name: "brand", type: "string" },
            { internalType: "string", name: "specification", type: "string" },
            { internalType: "uint256", name: "production_date", type: "uint256" },
            { internalType: "uint256", name: "expiration_date", type: "uint256" },
            { internalType: "string", name: "raw_material_source", type: "string" },
            { internalType: "string", name: "processing_method", type: "string" },
            { internalType: "string", name: "logistics_temp", type: "string" },
            { internalType: "string", name: "storage_condition", type: "string" },
            { internalType: "string", name: "qc_report", type: "string" },
            { internalType: "string", name: "manufacturer_id", type: "string" },
            { internalType: "string", name: "image_url", type: "string" },
            { internalType: "string", name: "batchNumber", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "address", name: "current_holder", type: "address" },
            { internalType: "string", name: "current_location", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.Product[]",
          name: "products",
          type: "tuple[]",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 获取卖家的完整销售历史
    {
      inputs: [{ internalType: "address", name: "_seller", type: "address" }],
      name: "getSellerCompleteSalesHistory",
      outputs: [
        {
          components: [
            { internalType: "uint256", name: "purchaseId", type: "uint256" },
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "uint256", name: "purchase_date", type: "uint256" },
            { internalType: "string", name: "status", type: "string" },
            { internalType: "address", name: "seller_id", type: "address" },
            { internalType: "address", name: "buyer_id", type: "address" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.PurchaseRecord[]",
          name: "purchases",
          type: "tuple[]",
        },
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "string", name: "brand", type: "string" },
            { internalType: "string", name: "specification", type: "string" },
            { internalType: "uint256", name: "production_date", type: "uint256" },
            { internalType: "uint256", name: "expiration_date", type: "uint256" },
            { internalType: "string", name: "raw_material_source", type: "string" },
            { internalType: "string", name: "processing_method", type: "string" },
            { internalType: "string", name: "logistics_temp", type: "string" },
            { internalType: "string", name: "storage_condition", type: "string" },
            { internalType: "string", name: "qc_report", type: "string" },
            { internalType: "string", name: "manufacturer_id", type: "string" },
            { internalType: "string", name: "image_url", type: "string" },
            { internalType: "string", name: "batchNumber", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "address", name: "current_holder", type: "address" },
            { internalType: "string", name: "current_location", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.Product[]",
          name: "products",
          type: "tuple[]",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    // 获取所有可用产品
    {
      inputs: [],
      name: "getAllAvailableProducts",
      outputs: [
        {
          components: [
            { internalType: "string", name: "sku", type: "string" },
            { internalType: "string", name: "product_name", type: "string" },
            { internalType: "string", name: "brand", type: "string" },
            { internalType: "string", name: "specification", type: "string" },
            { internalType: "uint256", name: "production_date", type: "uint256" },
            { internalType: "uint256", name: "expiration_date", type: "uint256" },
            { internalType: "string", name: "raw_material_source", type: "string" },
            { internalType: "string", name: "processing_method", type: "string" },
            { internalType: "string", name: "logistics_temp", type: "string" },
            { internalType: "string", name: "storage_condition", type: "string" },
            { internalType: "string", name: "qc_report", type: "string" },
            { internalType: "string", name: "manufacturer_id", type: "string" },
            { internalType: "string", name: "image_url", type: "string" },
            { internalType: "string", name: "batchNumber", type: "string" },
            { internalType: "uint256", name: "quantity", type: "uint256" },
            { internalType: "address", name: "current_holder", type: "address" },
            { internalType: "string", name: "current_location", type: "string" },
            { internalType: "bool", name: "isValid", type: "bool" },
          ],
          internalType: "struct FrozenProduct.Product[]",
          name: "",
          type: "tuple[]",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
  ],
};

class BlockchainService {
  constructor() {
    this.provider = null;
    this.contract = null;
    this.initialized = false;
    this.accountAddress = null;
    // 合约地址（需替换为实际部署的地址）
    this.contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
  }

  // 初始化区块链连接
  async init() {
    try {
      if (window.ethereum) {
        // 创建Provider
        this.provider = new ethers.BrowserProvider(window.ethereum);

        // 获取签名者
        const signer = await this.provider.getSigner();

        // 创建合约实例
        this.contract = new ethers.Contract(
            this.contractAddress,
            FrozenProductABI.abi,
            signer
        );

        this.initialized = true;
        return true;
      } else {
        console.error("请安装MetaMask钱包");
        return false;
      }
    } catch (error) {
      console.error("初始化区块链服务失败:", error);
      return false;
    }
  }

  // 确保服务已初始化
  async ensureInitialized() {
    if (!this.initialized) {
      await this.init();
    }
    return this.initialized;
  }

  // 获取所有产品SKU
  async getAllProductSkus() {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const skus = await this.contract.getAllProductSkus();
      return skus;
    } catch (error) {
      console.error("获取所有产品SKU失败:", error);
      return [];
    }
  }

  // 根据SKU获取产品详情
  async getProductBySku(sku) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const productData = await this.contract.getProductBySku(sku);

      // 转换时间戳为可读格式
      const productionDate = new Date(
          Number(productData.production_date) * 1000
      ).toISOString();
      const expirationDate = new Date(
          Number(productData.expiration_date) * 1000
      ).toISOString();

      return {
        sku: productData.sku,
        productName: productData.product_name,
        brand: productData.brand,
        specification: productData.specification,
        productionDate,
        expirationDate,
        rawMaterialSource: productData.raw_material_source,
        processingMethod: productData.processing_method,
        logisticsTemp: productData.logistics_temp,
        storageCondition: productData.storage_condition,
        qcReport: productData.qc_report,
        manufacturerID: productData.manufacturer_id,
        imageURL: productData.image_url,
        batchNumber: productData.batchNumber,
        quantity: Number(productData.quantity),
        currentHolder: productData.current_holder,
        currentLocation: productData.current_location,
        isValid: productData.isValid,
      };
    } catch (error) {
      console.error(`获取产品 ${sku} 失败:`, error);
      throw error;
    }
  }

  // 添加产品
  async addProduct(productData) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      // 构造参数（时间戳需转换为uint256）
      const productionDate = Math.floor(new Date(productData.productionDate).getTime() / 1000);
      const expirationDate = Math.floor(new Date(productData.expirationDate).getTime() / 1000);

      // 调用合约方法（需钱包签名）
      const tx = await this.contract.addProduct(
          productData.sku,
          productData.productName,
          productData.brand,
          productData.specification,
          productionDate,
          expirationDate,
          productData.rawMaterialSource,
          productData.processingMethod,
          productData.logisticsTemp,
          productData.storageCondition,
          productData.qcReport,
          productData.manufacturerID,
          productData.imageURL,
          productData.batchNumber,
          productData.quantity,
          productData.currentLocation
      );

      // 等待交易确认
      const receipt = await tx.wait();
      console.log("产品添加成功，交易哈希:", receipt.hash);
      return receipt;
    } catch (error) {
      console.error("添加产品失败:", error);
      throw error;
    }
  }

  // 更新产品持有者
  async updateProductHolder(sku, newHolder, newLocation) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const tx = await this.contract.updateProductHolder(
          sku,
          newHolder,
          newLocation
      );

      const receipt = await tx.wait();
      console.log("产品持有者更新成功，交易哈希:", receipt.hash);
      return receipt;
    } catch (error) {
      console.error("更新产品持有者失败:", error);
      throw error;
    }
  }

  // 添加物流记录
  async addLogisticsRecord(logisticsData) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const tx = await this.contract.addLogisticsRecord(
          logisticsData.sku,
          logisticsData.transportMethod,
          logisticsData.carrier,
          logisticsData.temperature,
          logisticsData.departure,
          logisticsData.destination,
          logisticsData.operatorName
      );

      const receipt = await tx.wait();
      console.log("物流记录添加成功，交易哈希:", receipt.hash);
      return receipt;
    } catch (error) {
      console.error("添加物流记录失败:", error);
      throw error;
    }
  }

  // 获取指定SKU的物流记录
  async fetchLogisticsRecords() {
    try {
      this.loading = true;

      // 确保区块链服务已初始化
      if (!this.blockchainConnected) {
        await this.connectBlockchain();
      }

      // 获取当前用户地址
      const address = await blockchainService.getAccountAddress();

      // 获取该用户的采购记录
      const { purchases } = await blockchainService.getBuyerCompletePurchaseHistory(address);

      // 清空当前物流记录
      this.logisticsRecords = [];

      // 为每个采购记录获取物流信息
      for (const purchase of purchases) {
        try {
          // 注意：这里将方法名称从getLogisticsRecords更正为getLogisticsBySku
          // 并传递sku作为参数
          const logistics = await blockchainService.getLogisticsBySku(purchase.sku);

          // 将获取的物流记录添加到列表中
          if (logistics && logistics.length > 0) {
            this.logisticsRecords.push(...logistics);
          }
        } catch (error) {
          console.error(`获取产品 ${purchase.sku} 的物流记录失败:`, error);
        }
      }

      this.loading = false;
    } catch (error) {
      console.error("获取物流记录失败:", error);
      this.loading = false;
      ElMessage.error("获取物流记录失败: " + error.message);
    }
  }

  // 获取完整产品信息（产品+物流）
  async getCompleteProductInfo(sku) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const result = await this.contract.getCompleteProductInfo(sku);

      // 格式化产品数据
      const productData = result.product;
      const product = {
        sku: productData.sku,
        productName: productData.product_name,
        brand: productData.brand,
        specification: productData.specification,
        productionDate: new Date(Number(productData.production_date) * 1000).toISOString(),
        expirationDate: new Date(Number(productData.expiration_date) * 1000).toISOString(),
        rawMaterialSource: productData.raw_material_source,
        processingMethod: productData.processing_method,
        logisticsTemp: productData.logistics_temp,
        storageCondition: productData.storage_condition,
        qcReport: productData.qc_report,
        manufacturerID: productData.manufacturer_id,
        imageURL: productData.image_url,
        batchNumber: productData.batchNumber,
        quantity: Number(productData.quantity),
        currentHolder: productData.current_holder,
        currentLocation: productData.current_location,
        isValid: productData.isValid,
      };

      // 格式化物流记录
      const logisticsRecords = result.logisticsRecords.map(record => ({
        sku: record.sku,
        transportMethod: record.transportMethod,
        carrier: record.carrier,
        temperature: record.temperature,
        departure: record.departure,
        destination: record.destination,
        operator: record.operator,
        operatorName: record.operatorName,
        isValid: record.isValid
      }));

      return {
        product,
        logisticsRecords
      };
    } catch (error) {
      console.error(`获取产品 ${sku} 的完整信息失败:`, error);
      throw error;
    }
  }

  // 创建采购记录
  async createPurchaseRecord(purchaseData) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const tx = await this.contract.createPurchaseRecord(
          purchaseData.sku,
          purchaseData.productName,
          purchaseData.quantity,
          purchaseData.sellerId,
          purchaseData.buyerId
      );

      const receipt = await tx.wait();
      console.log("采购记录创建成功，交易哈希:", receipt.hash);
      return receipt;
    } catch (error) {
      console.error("创建采购记录失败:", error);
      throw error;
    }
  }

  // 确认采购
  async confirmPurchase(purchaseId) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const tx = await this.contract.confirmPurchase(purchaseId);

      const receipt = await tx.wait();
      console.log("采购确认成功，交易哈希:", receipt.hash);
      return receipt;
    } catch (error) {
      console.error("确认采购失败:", error);
      throw error;
    }
  }

  // 获取买家的所有采购记录ID
  async getBuyerPurchaseIds(buyerAddress) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const purchaseIds = await this.contract.getBuyerPurchaseIds(buyerAddress);
      return purchaseIds.map(id => Number(id));
    } catch (error) {
      console.error(`获取买家 ${buyerAddress} 的采购记录ID失败:`, error);
      throw error;
    }
  }

  // 获取卖家的所有采购记录ID
  async getSellerPurchaseIds(sellerAddress) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const purchaseIds = await this.contract.getSellerPurchaseIds(sellerAddress);
      return purchaseIds.map(id => Number(id));
    } catch (error) {
      console.error(`获取卖家 ${sellerAddress} 的采购记录ID失败:`, error);
      throw error;
    }
  }

  // 获取指定采购记录
  async getPurchaseRecord(purchaseId) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const purchaseData = await this.contract.getPurchaseRecord(purchaseId);

      return {
        purchaseId: Number(purchaseData.purchaseId),
        sku: purchaseData.sku,
        productName: purchaseData.product_name,
        quantity: Number(purchaseData.quantity),
        purchaseDate: new Date(Number(purchaseData.purchase_date) * 1000).toISOString(),
        status: purchaseData.status,
        sellerId: purchaseData.seller_id,
        buyerId: purchaseData.buyer_id,
        isValid: purchaseData.isValid
      };
    } catch (error) {
      console.error(`获取采购记录 ${purchaseId} 失败:`, error);
      throw error;
    }
  }

  // 获取买家的完整采购历史
  async getBuyerCompletePurchaseHistory(buyerAddress) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const result = await this.contract.getBuyerCompletePurchaseHistory(buyerAddress);

      // 格式化采购记录
      const purchases = result.purchases.map(purchase => ({
        purchaseId: Number(purchase.purchaseId),
        sku: purchase.sku,
        productName: purchase.product_name,
        quantity: Number(purchase.quantity),
        purchaseDate: new Date(Number(purchase.purchase_date) * 1000).toISOString(),
        status: purchase.status,
        sellerId: purchase.seller_id,
        buyerId: purchase.buyer_id,
        isValid: purchase.isValid
      }));

      // 格式化产品数据
      const products = result.products.map(product => ({
        sku: product.sku,
        productName: product.product_name,
        brand: product.brand,
        specification: product.specification,
        productionDate: new Date(Number(product.production_date) * 1000).toISOString(),
        expirationDate: new Date(Number(product.expiration_date) * 1000).toISOString(),
        rawMaterialSource: product.raw_material_source,
        processingMethod: product.processing_method,
        logisticsTemp: product.logistics_temp,
        storageCondition: product.storage_condition,
        qcReport: product.qc_report,
        manufacturerID: product.manufacturer_id,
        imageURL: product.image_url,
        batchNumber: product.batchNumber,
        quantity: Number(product.quantity),
        currentHolder: product.current_holder,
        currentLocation: product.current_location,
        isValid: product.isValid
      }));

      return { purchases, products };
    } catch (error) {
      console.error(`获取买家 ${buyerAddress} 的完整采购历史失败:`, error);
      throw error;
    }
  }

  // 获取卖家的完整销售历史
  async getSellerCompleteSalesHistory(sellerAddress) {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const result = await this.contract.getSellerCompleteSalesHistory(sellerAddress);

      // 格式化采购记录
      const purchases = result.purchases.map(purchase => ({
        purchaseId: Number(purchase.purchaseId),
        sku: purchase.sku,
        productName: purchase.product_name,
        quantity: Number(purchase.quantity),
        purchaseDate: new Date(Number(purchase.purchase_date) * 1000).toISOString(),
        status: purchase.status,
        sellerId: purchase.seller_id,
        buyerId: purchase.buyer_id,
        isValid: purchase.isValid
      }));

      // 格式化产品数据
      const products = result.products.map(product => ({
        sku: product.sku,
        productName: product.product_name,
        brand: product.brand,
        specification: product.specification,
        productionDate: new Date(Number(product.production_date) * 1000).toISOString(),
        expirationDate: new Date(Number(product.expiration_date) * 1000).toISOString(),
        rawMaterialSource: product.raw_material_source,
        processingMethod: product.processing_method,
        logisticsTemp: product.logistics_temp,
        storageCondition: product.storage_condition,
        qcReport: product.qc_report,
        manufacturerID: product.manufacturer_id,
        imageURL: product.image_url,
        batchNumber: product.batchNumber,
        quantity: Number(product.quantity),
        currentHolder: product.current_holder,
        currentLocation: product.current_location,
        isValid: product.isValid
      }));

      return { purchases, products };
    } catch (error) {
      console.error(`获取卖家 ${sellerAddress} 的完整销售历史失败:`, error);
      throw error;
    }
  }

  // 获取所有可用产品
  async getAllAvailableProducts() {
    try {
      if (!(await this.ensureInitialized())) {
        throw new Error("区块链服务未初始化");
      }

      const productsData = await this.contract.getAllAvailableProducts();

      // 格式化产品数据
      const products = productsData.map(product => ({
        sku: product.sku,
        productName: product.product_name,
        brand: product.brand,
        specification: product.specification,
        productionDate: new Date(Number(product.production_date) * 1000).toISOString(),
        expirationDate: new Date(Number(product.expiration_date) * 1000).toISOString(),
        rawMaterialSource: product.raw_material_source,
        processingMethod: product.processing_method,
        logisticsTemp: product.logistics_temp,
        storageCondition: product.storage_condition,
        qcReport: product.qc_report,
        manufacturerID: product.manufacturer_id,
        imageURL: product.image_url,
        batchNumber: product.batchNumber,
        quantity: Number(product.quantity),
        currentHolder: product.current_holder,
        currentLocation: product.current_location,
        isValid: product.isValid
      }));

      return products;
    } catch (error) {
      console.error("获取所有可用产品失败:", error);
      throw error;
    }
  }

  // 获取当前用户地址
  async getAccountAddress() {
    try {
      if (!this.initialized) {
        await this.init();
      }

      if (this.provider) {
        const signer = await this.provider.getSigner();
        this.accountAddress = await signer.getAddress();
        return this.accountAddress;
      }

      return null;
    } catch (error) {
      console.error("获取账户地址失败:", error);
      throw error;
    }
  }
}

// 导出单例实例
export default new BlockchainService();