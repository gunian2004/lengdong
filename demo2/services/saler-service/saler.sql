use reval_system;
CREATE TABLE purchase_records (
    purchase_id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '进货ID',
    buyer_id BIGINT NOT NULL COMMENT '买家ID',
    seller_id BIGINT NOT NULL COMMENT '卖家ID',
    sku VARCHAR(50) NOT NULL COMMENT '产品SKU',
    batch_id BIGINT NOT NULL COMMENT '批次ID',
    quantity INT UNSIGNED NOT NULL COMMENT '数量',
    purchase_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '进货时间',
    FOREIGN KEY (buyer_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (seller_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (sku) REFERENCES frozen_products(sku) ON DELETE CASCADE,
    FOREIGN KEY (batch_id) REFERENCES product_batches(batch_id) ON DELETE CASCADE,
    INDEX idx_buyer_seller (buyer_id, seller_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;