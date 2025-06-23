import React from "react";

function ProductList({ products }) {
  return (
    <div className="table-responsive">
      {products.length === 0 ? (
        <div className="alert alert-info">暂无产品信息</div>
      ) : (
        <table className="table table-striped table-hover">
          <thead>
            <tr>
              <th>产品ID</th>
              <th>产品名称</th>
              <th>类别</th>
              <th>生产厂家</th>
              <th>生产日期</th>
              <th>过期日期</th>
              <th>存储温度(°C)</th>
              <th>批次号</th>
              <th>描述</th>
            </tr>
          </thead>
          <tbody>
            {products.map((product) => (
              <tr key={product.id}>
                <td>{product.id}</td>
                <td>{product.name}</td>
                <td>{product.category}</td>
                <td>{product.manufacturer}</td>
                <td>{product.productionDate}</td>
                <td>{product.expiryDate}</td>
                <td>{product.storageTemp}</td>
                <td>{product.batchNumber}</td>
                <td>{product.description}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}

export default ProductList;
