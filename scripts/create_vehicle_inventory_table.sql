-- 创建车辆库存管理表
-- 用于管理车辆的预订、租用、维护等状态

CREATE TABLE IF NOT EXISTS `vehicle_inventories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `vehicle_id` bigint unsigned NOT NULL COMMENT '车辆ID',
  `start_date` date NOT NULL COMMENT '开始日期',
  `end_date` date NOT NULL COMMENT '结束日期',
  `status` tinyint DEFAULT '1' COMMENT '库存状态 1:可租用 2:已预订 3:租用中 4:维护中 5:不可用',
  `order_id` bigint unsigned DEFAULT '0' COMMENT '关联订单ID，0表示无订单',
  `quantity` int DEFAULT '1' COMMENT '数量(通常为1，支持同型号多辆车)',
  `notes` varchar(500) DEFAULT NULL COMMENT '备注',
  `created_by` bigint unsigned DEFAULT NULL COMMENT '创建人ID',
  `updated_by` bigint unsigned DEFAULT NULL COMMENT '更新人ID',
  PRIMARY KEY (`id`),
  KEY `idx_vehicle_inventories_deleted_at` (`deleted_at`),
  KEY `idx_vehicle_id` (`vehicle_id`),
  KEY `idx_start_date` (`start_date`),
  KEY `idx_end_date` (`end_date`),
  KEY `idx_status` (`status`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_date_range` (`vehicle_id`, `start_date`, `end_date`),
  KEY `idx_status_date` (`status`, `start_date`, `end_date`),
  CONSTRAINT `fk_vehicle_inventories_vehicle` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='车辆库存管理表';

-- 插入一些示例库存数据
-- 注意：这里的vehicle_id需要根据实际的vehicles表中的ID进行调整

-- 为一些车辆添加维护计划
INSERT INTO `vehicle_inventories` (`vehicle_id`, `start_date`, `end_date`, `status`, `notes`, `created_by`, `updated_by`, `created_at`, `updated_at`) VALUES
-- 车辆1的维护计划
(1, '2024-01-15', '2024-01-17', 4, '定期保养维护', 1, 1, NOW(), NOW()),
(1, '2024-03-01', '2024-03-03', 4, '季度检查维护', 1, 1, NOW(), NOW()),

-- 车辆2的维护计划
(2, '2024-01-20', '2024-01-22', 4, '轮胎更换维护', 1, 1, NOW(), NOW()),
(2, '2024-04-10', '2024-04-12', 4, '空调系统检查', 1, 1, NOW(), NOW()),

-- 车辆3的维护计划
(3, '2024-02-05', '2024-02-07', 4, '发动机保养', 1, 1, NOW(), NOW()),

-- 一些已有的预订记录（示例）
(4, '2024-01-25', '2024-01-28', 2, '春节期间预订', 2, 2, NOW(), NOW()),
(5, '2024-02-14', '2024-02-16', 2, '情人节预订', 3, 3, NOW(), NOW()),
(6, '2024-03-15', '2024-03-20', 2, '商务出行预订', 4, 4, NOW(), NOW());

-- 创建库存状态统计视图
CREATE OR REPLACE VIEW `vehicle_inventory_stats` AS
SELECT 
    v.merchant_id,
    COUNT(DISTINCT v.id) as total_vehicles,
    COUNT(DISTINCT CASE 
        WHEN vi.status = 2 AND vi.start_date <= CURDATE() AND vi.end_date >= CURDATE() 
        THEN v.id 
    END) as reserved_vehicles,
    COUNT(DISTINCT CASE 
        WHEN vi.status = 3 AND vi.start_date <= CURDATE() AND vi.end_date >= CURDATE() 
        THEN v.id 
    END) as rented_vehicles,
    COUNT(DISTINCT CASE 
        WHEN vi.status = 4 AND vi.start_date <= CURDATE() AND vi.end_date >= CURDATE() 
        THEN v.id 
    END) as maintenance_vehicles,
    COUNT(DISTINCT v.id) - COUNT(DISTINCT CASE 
        WHEN vi.status IN (2,3,4) AND vi.start_date <= CURDATE() AND vi.end_date >= CURDATE() 
        THEN v.id 
    END) as available_vehicles
FROM vehicles v
LEFT JOIN vehicle_inventories vi ON v.id = vi.vehicle_id AND vi.deleted_at IS NULL
WHERE v.deleted_at IS NULL AND v.status = 1
GROUP BY v.merchant_id;

-- 创建存储过程：检查车辆可用性
DELIMITER //
CREATE PROCEDURE CheckVehicleAvailability(
    IN p_vehicle_id BIGINT,
    IN p_start_date DATE,
    IN p_end_date DATE,
    OUT p_is_available BOOLEAN
)
BEGIN
    DECLARE conflict_count INT DEFAULT 0;
    
    -- 检查是否有冲突的预订或租用记录
    SELECT COUNT(*) INTO conflict_count
    FROM vehicle_inventories
    WHERE vehicle_id = p_vehicle_id
      AND status IN (2, 3, 4) -- 已预订、租用中、维护中
      AND deleted_at IS NULL
      AND (
          (start_date <= p_start_date AND end_date >= p_start_date) OR
          (start_date <= p_end_date AND end_date >= p_end_date) OR
          (start_date >= p_start_date AND end_date <= p_end_date)
      );
    
    -- 如果没有冲突，则可用
    SET p_is_available = (conflict_count = 0);
END //
DELIMITER ;

-- 创建存储过程：获取可用车辆列表
DELIMITER //
CREATE PROCEDURE GetAvailableVehicles(
    IN p_start_date DATE,
    IN p_end_date DATE,
    IN p_merchant_id BIGINT,
    IN p_type_id BIGINT,
    IN p_brand_id BIGINT
)
BEGIN
    SELECT DISTINCT v.*
    FROM vehicles v
    WHERE v.status = 1
      AND v.deleted_at IS NULL
      AND (p_merchant_id IS NULL OR v.merchant_id = p_merchant_id)
      AND (p_type_id IS NULL OR v.type_id = p_type_id)
      AND (p_brand_id IS NULL OR v.brand_id = p_brand_id)
      AND v.id NOT IN (
          SELECT DISTINCT vi.vehicle_id
          FROM vehicle_inventories vi
          WHERE vi.status IN (2, 3, 4) -- 已预订、租用中、维护中
            AND vi.deleted_at IS NULL
            AND (
                (vi.start_date <= p_start_date AND vi.end_date >= p_start_date) OR
                (vi.start_date <= p_end_date AND vi.end_date >= p_end_date) OR
                (vi.start_date >= p_start_date AND vi.end_date <= p_end_date)
            )
      )
    ORDER BY v.created_at DESC;
END //
DELIMITER ;

-- 创建触发器：当订单状态变化时自动更新库存状态
DELIMITER //
CREATE TRIGGER update_inventory_on_order_status
AFTER UPDATE ON orders
FOR EACH ROW
BEGIN
    -- 当订单状态变为已确认时，将预订状态改为租用中
    IF NEW.status = 'confirmed' AND OLD.status != 'confirmed' THEN
        UPDATE vehicle_inventories 
        SET status = 3, updated_at = NOW()
        WHERE order_id = NEW.id AND status = 2;
    END IF;
    
    -- 当订单状态变为已完成时，删除库存记录
    IF NEW.status = 'completed' AND OLD.status != 'completed' THEN
        UPDATE vehicle_inventories 
        SET deleted_at = NOW()
        WHERE order_id = NEW.id AND status = 3;
    END IF;
    
    -- 当订单状态变为已取消时，删除预订记录
    IF NEW.status = 'cancelled' AND OLD.status != 'cancelled' THEN
        UPDATE vehicle_inventories 
        SET deleted_at = NOW()
        WHERE order_id = NEW.id AND status = 2;
    END IF;
END //
DELIMITER ;

-- 创建索引优化查询性能
CREATE INDEX idx_vehicle_date_status ON vehicle_inventories(vehicle_id, start_date, end_date, status);
CREATE INDEX idx_merchant_vehicle ON vehicles(merchant_id, id);
CREATE INDEX idx_vehicle_type_brand ON vehicles(type_id, brand_id, status);

-- 插入库存状态说明数据（用于前端显示）
CREATE TABLE IF NOT EXISTS `inventory_status_types` (
  `id` int NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '状态名称',
  `description` varchar(200) DEFAULT NULL COMMENT '状态描述',
  `color` varchar(20) DEFAULT NULL COMMENT '显示颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='库存状态类型表';

INSERT INTO `inventory_status_types` (`id`, `name`, `description`, `color`) VALUES
(1, '可租用', '车辆空闲，可以预订', 'green'),
(2, '已预订', '车辆被预订但还未取车', 'orange'),
(3, '租用中', '车辆正在被使用', 'blue'),
(4, '维护中', '车辆在保养或维修', 'purple'),
(5, '不可用', '车辆暂时下架', 'red');

-- 执行说明
-- 1. 执行前请确保vehicles表和orders表已存在
-- 2. 根据实际的vehicle_id调整示例数据
-- 3. 触发器需要orders表存在，如果没有可以注释掉
-- 4. 可以根据业务需求调整库存状态和规则
