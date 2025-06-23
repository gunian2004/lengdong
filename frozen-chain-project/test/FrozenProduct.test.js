const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("FrozenProduct合约", function () {
  let FrozenProduct;
  let frozenProduct;
  let owner;
  let addr1;
  let addr2;

  // 在每个测试前部署合约
  beforeEach(async function () {
    // 获取合约工厂
    FrozenProduct = await ethers.getContractFactory("FrozenProduct");

    // 获取账户
    [owner, addr1, addr2] = await ethers.getSigners();

    // 部署合约
    frozenProduct = await FrozenProduct.deploy();
  });

  // 测试合约部署
  describe("部署", function () {
    it("应该设置正确的所有者", async function () {
      expect(await frozenProduct.owner()).to.equal(owner.address);
    });

    it("初始产品数量应为0", async function () {
      expect(await frozenProduct.getProductCount()).to.equal(0);
    });
  });

  // 测试添加产品功能
  describe("添加产品", function () {
    it("所有者应该能添加产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await frozenProduct.addProduct(
        "P001",
        "冰鲜三文鱼",
        "海鲜",
        "海洋食品公司",
        now,
        expiryDate,
        -18,
        "B20230501",
        "新鲜冰鲜三文鱼"
      );

      expect(await frozenProduct.getProductCount()).to.equal(1);

      const productIds = await frozenProduct.getAllProductIds();
      expect(productIds[0]).to.equal("P001");

      const product = await frozenProduct.getProduct("P001");
      expect(product.productName).to.equal("冰鲜三文鱼");
      expect(product.category).to.equal("海鲜");
      expect(product.manufacturer).to.equal("海洋食品公司");
      expect(product.productionDate).to.equal(now);
      expect(product.expiryDate).to.equal(expiryDate);
      expect(product.storageTemp).to.equal(-18);
      expect(product.batchNumber).to.equal("B20230501");
      expect(product.description).to.equal("新鲜冰鲜三文鱼");
    });

    it("非所有者不应该能添加产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await expect(
        frozenProduct
          .connect(addr1)
          .addProduct(
            "P002",
            "冰鲜牛肉",
            "肉类",
            "肉食品公司",
            now,
            expiryDate,
            -20,
            "B20230502",
            "新鲜冰鲜牛肉"
          )
      ).to.be.reverted;
    });

    it("不应该能添加重复ID的产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await frozenProduct.addProduct(
        "P003",
        "冰鲜鸡肉",
        "肉类",
        "肉食品公司",
        now,
        expiryDate,
        -18,
        "B20230503",
        "新鲜冰鲜鸡肉"
      );

      await expect(
        frozenProduct.addProduct(
          "P003",
          "冰鲜猪肉",
          "肉类",
          "肉食品公司",
          now,
          expiryDate,
          -18,
          "B20230504",
          "新鲜冰鲜猪肉"
        )
      ).to.be.revertedWith("Product ID already exists");
    });
  });

  // 测试更新产品功能
  describe("更新产品", function () {
    it("所有者应该能更新产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await frozenProduct.addProduct(
        "P004",
        "冰鲜虾",
        "海鲜",
        "海洋食品公司",
        now,
        expiryDate,
        -18,
        "B20230505",
        "新鲜冰鲜虾"
      );

      const newExpiryDate = now + 60 * 24 * 60 * 60; // 60天后过期

      await frozenProduct.updateProduct(
        "P004",
        "优质冰鲜虾",
        "高级海鲜",
        "高级海洋食品公司",
        now,
        newExpiryDate,
        -20,
        "B20230505-A",
        "优质新鲜冰鲜虾"
      );

      const product = await frozenProduct.getProduct("P004");
      expect(product.productName).to.equal("优质冰鲜虾");
      expect(product.category).to.equal("高级海鲜");
      expect(product.manufacturer).to.equal("高级海洋食品公司");
      expect(product.expiryDate).to.equal(newExpiryDate);
      expect(product.storageTemp).to.equal(-20);
      expect(product.batchNumber).to.equal("B20230505-A");
      expect(product.description).to.equal("优质新鲜冰鲜虾");
    });

    it("非所有者不应该能更新产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await frozenProduct.addProduct(
        "P005",
        "冰鲜蟹",
        "海鲜",
        "海洋食品公司",
        now,
        expiryDate,
        -18,
        "B20230506",
        "新鲜冰鲜蟹"
      );

      await expect(
        frozenProduct
          .connect(addr1)
          .updateProduct(
            "P005",
            "优质冰鲜蟹",
            "高级海鲜",
            "高级海洋食品公司",
            now,
            expiryDate,
            -20,
            "B20230506-A",
            "优质新鲜冰鲜蟹"
          )
      ).to.be.reverted;
    });

    it("不应该能更新不存在的产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await expect(
        frozenProduct.updateProduct(
          "P999",
          "不存在的产品",
          "未知",
          "未知",
          now,
          expiryDate,
          -18,
          "B999",
          "不存在的产品"
        )
      ).to.be.revertedWith("Product does not exist");
    });
  });

  // 测试获取产品功能
  describe("获取产品", function () {
    it("应该能获取存在的产品", async function () {
      const now = Math.floor(Date.now() / 1000);
      const expiryDate = now + 30 * 24 * 60 * 60; // 30天后过期

      await frozenProduct.addProduct(
        "P006",
        "冰鲜鱿鱼",
        "海鲜",
        "海洋食品公司",
        now,
        expiryDate,
        -18,
        "B20230507",
        "新鲜冰鲜鱿鱼"
      );

      const product = await frozenProduct.getProduct("P006");
      expect(product.productName).to.equal("冰鲜鱿鱼");
    });

    it("不应该能获取不存在的产品", async function () {
      await expect(frozenProduct.getProduct("P999")).to.be.revertedWith(
        "Product does not exist"
      );
    });
  });
});
