import React, { useState, useEffect } from "react";
import { ethers } from "ethers";
import ProductList from "./components/ProductList";
import ProductForm from "./components/ProductForm";
import FrozenProductABI from "./contracts/FrozenProduct.json";

function App() {
  const [account, setAccount] = useState("");
  const [contract, setContract] = useState(null);
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [isOwner, setIsOwner] = useState(false);

  useEffect(() => {
    const init = async () => {
      try {
        // 检查是否安装了MetaMask
        if (window.ethereum) {
          // 请求用户授权
          const accounts = await window.ethereum.request({
            method: "eth_requestAccounts",
          });
          setAccount(accounts[0]);

          // 创建Provider
          const provider = new ethers.BrowserProvider(window.ethereum);

          // 获取签名者
          const signer = await provider.getSigner();

          // 合约地址（部署后需要更新）
          const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3"; // 这是Hardhat本地网络的默认地址，需要替换为实际部署的地址

          // 创建合约实例
          const frozenProductContract = new ethers.Contract(
            contractAddress,
            FrozenProductABI.abi,
            signer
          );

          setContract(frozenProductContract);

          // 检查当前用户是否是合约所有者
          const owner = await frozenProductContract.owner();
          setIsOwner(owner.toLowerCase() === accounts[0].toLowerCase());

          // 加载产品列表
          await loadProducts(frozenProductContract);

          // 监听账户变化
          window.ethereum.on("accountsChanged", (newAccounts) => {
            setAccount(newAccounts[0]);
            checkOwner(frozenProductContract, newAccounts[0]);
          });
        } else {
          alert("请安装MetaMask钱包以使用此应用！");
        }
      } catch (error) {
        console.error("初始化错误:", error);
      } finally {
        setLoading(false);
      }
    };

    init();

    return () => {
      // 清理事件监听器
      if (window.ethereum) {
        window.ethereum.removeAllListeners("accountsChanged");
      }
    };
  }, []);

  // 检查是否是所有者
  const checkOwner = async (contract, address) => {
    try {
      const owner = await contract.owner();
      setIsOwner(owner.toLowerCase() === address.toLowerCase());
    } catch (error) {
      console.error("检查所有者错误:", error);
      setIsOwner(false);
    }
  };

  // 加载产品列表
  const loadProducts = async (contractInstance) => {
    try {
      setLoading(true);
      const productIds = await contractInstance.getAllProductIds();

      const productDetails = [];
      for (const id of productIds) {
        const product = await contractInstance.getProduct(id);
        productDetails.push({
          id: id,
          name: product.productName,
          category: product.category,
          manufacturer: product.manufacturer,
          productionDate: new Date(
            Number(product.productionDate) * 1000
          ).toLocaleDateString(),
          expiryDate: new Date(
            Number(product.expiryDate) * 1000
          ).toLocaleDateString(),
          storageTemp: Number(product.storageTemp),
          batchNumber: product.batchNumber,
          description: product.description,
        });
      }

      setProducts(productDetails);
    } catch (error) {
      console.error("加载产品错误:", error);
    } finally {
      setLoading(false);
    }
  };

  // 添加产品
  const addProduct = async (product) => {
    try {
      setLoading(true);
      const tx = await contract.addProduct(
        product.id,
        product.name,
        product.category,
        product.manufacturer,
        Math.floor(new Date(product.productionDate).getTime() / 1000),
        Math.floor(new Date(product.expiryDate).getTime() / 1000),
        product.storageTemp,
        product.batchNumber,
        product.description
      );

      await tx.wait();
      alert("产品添加成功！");
      await loadProducts(contract);
    } catch (error) {
      console.error("添加产品错误:", error);
      alert("添加产品失败: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  // 更新产品
  const updateProduct = async (product) => {
    try {
      setLoading(true);
      const tx = await contract.updateProduct(
        product.id,
        product.name,
        product.category,
        product.manufacturer,
        Math.floor(new Date(product.productionDate).getTime() / 1000),
        Math.floor(new Date(product.expiryDate).getTime() / 1000),
        product.storageTemp,
        product.batchNumber,
        product.description
      );

      await tx.wait();
      alert("产品更新成功！");
      await loadProducts(contract);
    } catch (error) {
      console.error("更新产品错误:", error);
      alert("更新产品失败: " + error.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mt-5">
      <h1 className="text-center mb-4">冷冻品信息上链系统</h1>

      <div className="card mb-4">
        <div className="card-body">
          <h5 className="card-title">账户信息</h5>
          <p className="card-text">
            <strong>当前账户:</strong> {account || "未连接"}
          </p>
          <p className="card-text">
            <strong>身份:</strong> {isOwner ? "管理员" : "普通用户"}
          </p>
        </div>
      </div>

      {loading ? (
        <div className="text-center">
          <div className="spinner-border" role="status">
            <span className="visually-hidden">加载中...</span>
          </div>
        </div>
      ) : (
        <>
          {isOwner && (
            <div className="card mb-4">
              <div className="card-header">
                <h5 className="mb-0">添加/更新产品</h5>
              </div>
              <div className="card-body">
                <ProductForm onSubmit={addProduct} onUpdate={updateProduct} />
              </div>
            </div>
          )}

          <div className="card">
            <div className="card-header">
              <h5 className="mb-0">产品列表</h5>
            </div>
            <div className="card-body">
              <ProductList products={products} />
            </div>
          </div>
        </>
      )}
    </div>
  );
}

export default App;
