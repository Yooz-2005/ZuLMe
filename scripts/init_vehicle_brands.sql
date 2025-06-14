-- 车辆品牌表初始化脚本
-- 创建车辆品牌表
CREATE TABLE IF NOT EXISTS `vehicle_brands` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '品牌名称',
  `english_name` varchar(50) DEFAULT NULL COMMENT '英文名称',
  `logo` varchar(255) DEFAULT NULL COMMENT '品牌Logo URL',
  `country` varchar(50) DEFAULT NULL COMMENT '品牌国家',
  `description` varchar(500) DEFAULT NULL COMMENT '品牌描述',
  `status` tinyint DEFAULT '1' COMMENT '状态 1:启用 0:禁用',
  `sort` int DEFAULT '0' COMMENT '排序',
  `is_hot` tinyint DEFAULT '0' COMMENT '是否热门 1:是 0:否',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_vehicle_brands_deleted_at` (`deleted_at`),
  KEY `idx_status` (`status`),
  KEY `idx_is_hot` (`is_hot`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='车辆品牌表';

-- 插入常见汽车品牌数据
INSERT INTO `vehicle_brands` (`name`, `english_name`, `country`, `description`, `status`, `sort`, `is_hot`) VALUES
('奔驰', 'Mercedes-Benz', '德国', '德国豪华汽车品牌，以高品质和创新技术著称', 1, 100, 1),
('宝马', 'BMW', '德国', '德国豪华汽车制造商，专注于驾驶乐趣和性能', 1, 99, 1),
('奥迪', 'Audi', '德国', '德国豪华汽车品牌，以先进技术和优雅设计闻名', 1, 98, 1),
('大众', 'Volkswagen', '德国', '德国汽车制造商，全球最大的汽车集团之一', 1, 97, 1),
('丰田', 'Toyota', '日本', '日本汽车制造商，以可靠性和燃油经济性著称', 1, 96, 1),
('本田', 'Honda', '日本', '日本汽车制造商，以技术创新和可靠性闻名', 1, 95, 1),
('日产', 'Nissan', '日本', '日本汽车制造商，在电动汽车领域有重要地位', 1, 94, 1),
('马自达', 'Mazda', '日本', '日本汽车制造商，以驾驶乐趣和创驰蓝天技术著称', 1, 93, 0),
('福特', 'Ford', '美国', '美国汽车制造商，历史悠久的汽车品牌', 1, 92, 0),
('雪佛兰', 'Chevrolet', '美国', '美国汽车品牌，通用汽车旗下品牌', 1, 91, 0),
('别克', 'Buick', '美国', '美国豪华汽车品牌，通用汽车旗下品牌', 1, 90, 0),
('凯迪拉克', 'Cadillac', '美国', '美国豪华汽车品牌，以豪华和性能著称', 1, 89, 0),
('特斯拉', 'Tesla', '美国', '美国电动汽车制造商，电动汽车行业的领导者', 1, 88, 1),
('比亚迪', 'BYD', '中国', '中国新能源汽车制造商，在电动汽车和电池技术方面领先', 1, 87, 1),
('吉利', 'Geely', '中国', '中国汽车制造商，拥有沃尔沃等国际品牌', 1, 86, 0),
('长城', 'Great Wall', '中国', '中国汽车制造商，以SUV和皮卡车型著称', 1, 85, 0),
('奇瑞', 'Chery', '中国', '中国汽车制造商，在自主研发方面有重要地位', 1, 84, 0),
('长安', 'Changan', '中国', '中国汽车制造商，历史悠久的自主品牌', 1, 83, 0),
('现代', 'Hyundai', '韩国', '韩国汽车制造商，以性价比和设计著称', 1, 82, 0),
('起亚', 'Kia', '韩国', '韩国汽车制造商，现代汽车集团旗下品牌', 1, 81, 0),
('沃尔沃', 'Volvo', '瑞典', '瑞典豪华汽车品牌，以安全性著称', 1, 80, 0),
('保时捷', 'Porsche', '德国', '德国豪华跑车制造商，以性能和工艺著称', 1, 79, 1),
('法拉利', 'Ferrari', '意大利', '意大利超级跑车制造商，赛车运动的象征', 1, 78, 1),
('兰博基尼', 'Lamborghini', '意大利', '意大利超级跑车制造商，以极致性能和设计著称', 1, 77, 1),
('玛莎拉蒂', 'Maserati', '意大利', '意大利豪华汽车品牌，以优雅和性能著称', 1, 76, 0);

-- 为车辆表添加品牌ID字段（如果不存在）
ALTER TABLE `vehicles` 
ADD COLUMN IF NOT EXISTS `brand_id` bigint NOT NULL DEFAULT 0 COMMENT '品牌ID' AFTER `type_id`,
ADD INDEX IF NOT EXISTS `idx_brand_id` (`brand_id`);

-- 更新现有车辆数据的品牌ID（根据品牌名称匹配）
UPDATE `vehicles` v 
INNER JOIN `vehicle_brands` vb ON v.brand = vb.name 
SET v.brand_id = vb.id 
WHERE v.brand_id = 0;

-- 对于没有匹配到的品牌，创建新的品牌记录
INSERT INTO `vehicle_brands` (`name`, `english_name`, `country`, `description`, `status`, `sort`, `is_hot`)
SELECT DISTINCT 
    v.brand,
    v.brand,
    '未知',
    CONCAT('自动创建的品牌：', v.brand),
    1,
    0,
    0
FROM `vehicles` v
LEFT JOIN `vehicle_brands` vb ON v.brand = vb.name
WHERE vb.id IS NULL AND v.brand != '' AND v.brand IS NOT NULL;

-- 再次更新车辆表的品牌ID
UPDATE `vehicles` v 
INNER JOIN `vehicle_brands` vb ON v.brand = vb.name 
SET v.brand_id = vb.id 
WHERE v.brand_id = 0;
