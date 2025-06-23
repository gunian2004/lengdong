import React, { useState } from "react";

function ProductForm({ onSubmit, onUpdate }) {
  const [formData, setFormData] = useState({
    id: "",
    name: "",
    category: "",
    manufacturer: "",
    productionDate: "",
    expiryDate: "",
    storageTemp: -18,
    batchNumber: "",
    description: "",
  });

  const [formMode, setFormMode] = useState("add"); // 'add' 或 'update'

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    // 表单验证
    if (
      !formData.id ||
      !formData.name ||
      !formData.category ||
      !formData.manufacturer ||
      !formData.productionDate ||
      !formData.expiryDate ||
      !formData.batchNumber
    ) {
      alert("请填写所有必填字段！");
      return;
    }

    // 根据模式调用不同的处理函数
    if (formMode === "add") {
      onSubmit(formData);
    } else {
      onUpdate(formData);
    }

    // 重置表单
    resetForm();
  };

  const resetForm = () => {
    setFormData({
      id: "",
      name: "",
      category: "",
      manufacturer: "",
      productionDate: "",
      expiryDate: "",
      storageTemp: -18,
      batchNumber: "",
      description: "",
    });
    setFormMode("add");
  };

  const switchToUpdateMode = () => {
    setFormMode("update");
  };

  const switchToAddMode = () => {
    setFormMode("add");
    resetForm();
  };

  return (
    <div>
      <div className="mb-3">
        <div className="btn-group" role="group">
          <button
            type="button"
            className={`btn ${
              formMode === "add" ? "btn-primary" : "btn-outline-primary"
            }`}
            onClick={switchToAddMode}
          >
            添加产品
          </button>
          <button
            type="button"
            className={`btn ${
              formMode === "update" ? "btn-primary" : "btn-outline-primary"
            }`}
            onClick={switchToUpdateMode}
          >
            更新产品
          </button>
        </div>
      </div>

      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="id" className="form-label">
            产品ID *
          </label>
          <input
            type="text"
            className="form-control"
            id="id"
            name="id"
            value={formData.id}
            onChange={handleChange}
            readOnly={formMode === "update"}
            required
          />
          {formMode === "update" && (
            <small className="text-muted">更新模式下产品ID不可修改</small>
          )}
        </div>

        <div className="mb-3">
          <label htmlFor="name" className="form-label">
            产品名称 *
          </label>
          <input
            type="text"
            className="form-control"
            id="name"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label htmlFor="category" className="form-label">
            产品类别 *
          </label>
          <input
            type="text"
            className="form-control"
            id="category"
            name="category"
            value={formData.category}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label htmlFor="manufacturer" className="form-label">
            生产厂家 *
          </label>
          <input
            type="text"
            className="form-control"
            id="manufacturer"
            name="manufacturer"
            value={formData.manufacturer}
            onChange={handleChange}
            required
          />
        </div>

        <div className="row">
          <div className="col-md-6 mb-3">
            <label htmlFor="productionDate" className="form-label">
              生产日期 *
            </label>
            <input
              type="date"
              className="form-control"
              id="productionDate"
              name="productionDate"
              value={formData.productionDate}
              onChange={handleChange}
              required
            />
          </div>

          <div className="col-md-6 mb-3">
            <label htmlFor="expiryDate" className="form-label">
              过期日期 *
            </label>
            <input
              type="date"
              className="form-control"
              id="expiryDate"
              name="expiryDate"
              value={formData.expiryDate}
              onChange={handleChange}
              required
            />
          </div>
        </div>

        <div className="mb-3">
          <label htmlFor="storageTemp" className="form-label">
            存储温度(°C) *
          </label>
          <input
            type="number"
            className="form-control"
            id="storageTemp"
            name="storageTemp"
            value={formData.storageTemp}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label htmlFor="batchNumber" className="form-label">
            批次号 *
          </label>
          <input
            type="text"
            className="form-control"
            id="batchNumber"
            name="batchNumber"
            value={formData.batchNumber}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label htmlFor="description" className="form-label">
            产品描述
          </label>
          <textarea
            className="form-control"
            id="description"
            name="description"
            value={formData.description}
            onChange={handleChange}
            rows="3"
          ></textarea>
        </div>

        <div className="d-flex justify-content-between">
          <button type="submit" className="btn btn-primary">
            {formMode === "add" ? "添加产品" : "更新产品"}
          </button>
          <button
            type="button"
            className="btn btn-secondary"
            onClick={resetForm}
          >
            重置表单
          </button>
        </div>
      </form>
    </div>
  );
}

export default ProductForm;
