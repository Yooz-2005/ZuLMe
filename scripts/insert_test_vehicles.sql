-- 插入豪华车辆测试数据
-- 包含不同商家、不同类型、不同品牌的豪华车辆，每辆车都有多张高清图片

-- 假设我们有以下商家ID (需要根据实际情况调整)
-- 商家1: ID=1 (北京豪车租赁), 商家2: ID=2 (上海尊享汽车), 商家3: ID=3 (深圳奢华出行)

-- 假设我们有以下车辆类型ID (需要根据实际情况调整)
-- 豪华轿车: ID=3, 豪华SUV: ID=4, 超级跑车: ID=5

-- 首先确保豪华品牌存在，如果不存在则插入
INSERT IGNORE INTO `vehicle_brands` (`id`, `name`, `english_name`, `country`, `description`, `status`, `sort`, `is_hot`, `created_at`, `updated_at`) VALUES
(26, '路虎', 'Land Rover', '英国', '英国豪华SUV品牌，以全地形能力著称', 1, 75, 1, NOW(), NOW()),
(27, '宾利', 'Bentley', '英国', '英国超豪华汽车品牌，手工打造的奢华典范', 1, 74, 1, NOW(), NOW()),
(28, '劳斯莱斯', 'Rolls-Royce', '英国', '世界顶级豪华汽车品牌，奢华之巅', 1, 73, 1, NOW(), NOW());

-- 奔驰S级豪华轿车 (brand_id=1)
INSERT INTO `vehicles` (`merchant_id`, `type_id`, `brand_id`, `brand`, `style`, `year`, `color`, `mileage`, `price`, `status`, `description`, `images`, `location`, `contact`, `created_at`, `updated_at`) VALUES
(1, 7, 1, '奔驰', 'S500L', 2024, '曜石黑', 2000, 1299.00, 1, '全新奔驰S500L旗舰豪华轿车，配备魔毯悬挂、柏林之声音响、行政级后排座椅，商务出行的不二之选', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800', '北京市朝阳区国贸CBD', '13800138001', NOW(), NOW()),

(2, 7, 1, '奔驰', 'S680 迈巴赫', 2023, '珍珠白', 3500, 2999.00, 1, '奔驰迈巴赫S680，V12双涡轮增压发动机，手工打造内饰，极致奢华体验', 'https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800', '上海市浦东新区陆家嘴', '13800138002', NOW(), NOW()),

-- 宝马7系豪华轿车 (brand_id=2)
(3, 7, 2, '宝马', '760Li', 2024, '矿石灰', 1800, 1599.00, 1, '宝马760Li xDrive，V12发动机，水晶档把，天使之翼迎宾灯，科技与豪华并存', 'https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800', '深圳市南山区科技园', '13800138003', NOW(), NOW()),

(1, 7, 2, '宝马', 'X7 M50i', 2023, '宝石青', 4200, 1199.00, 1, '宝马X7 M50i大型豪华SUV，7座布局，全景天窗，Bowers & Wilkins钻石环绕音响', 'https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800', '广州市天河区珠江新城', '13800138004', NOW(), NOW()),

-- 奥迪A8L豪华轿车 (brand_id=3)
(2, 7, 3, '奥迪', 'A8L 60 TFSI', 2024, '冰川白', 2800, 1099.00, 1, '奥迪A8L 60 TFSI quattro，48V轻混系统，Matrix LED大灯，Bang & Olufsen 3D音响', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800', '杭州市西湖区', '13800138005', NOW(), NOW()),

(3, 7, 3, '奥迪', 'Q8 55 TFSI', 2023, '曼哈顿灰', 3600, 899.00, 1, '奥迪Q8 55 TFSI quattro轿跑SUV，运动外观，豪华内饰，quattro全时四驱', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '成都市锦江区', '13800138006', NOW(), NOW()),

-- 路虎揽胜豪华SUV (需要先添加路虎品牌，假设brand_id=26)
(1, 2, 26, '路虎', '揽胜 5.0 V8', 2024, '圣托里尼黑', 1500, 1399.00, 1, '路虎揽胜5.0 V8 SC，全地形反馈适应系统，Meridian音响，真皮内饰，越野之王', 'https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '武汉市江汉区', '13800138007', NOW(), NOW()),

(2, 2, 26, '路虎', '揽胜运动版 SVR', 2023, '富士白', 2800, 1699.00, 1, '路虎揽胜运动版SVR，5.0L V8机械增压，575马力，性能与豪华的完美结合', 'https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '南京市鼓楼区', '13800138008', NOW(), NOW()),

-- 宾利豪华轿车 (需要先添加宾利品牌，假设brand_id=27)
(3, 7, 27, '宾利', '飞驰 V8', 2024, '月光银', 1200, 2199.00, 1, '宾利飞驰V8，手工缝制真皮内饰，Naim音响，钻石纹理装饰，英伦奢华典范', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '西安市雁塔区', '13800138009', NOW(), NOW()),

(1, 7, 27, '宾利', '添越 V8', 2023, '翡翠绿', 3200, 1899.00, 1, '宾利添越V8豪华SUV，全地形能力，奢华内饰，Bentley Dynamic Ride主动防侧倾系统', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '重庆市渝中区', '13800138010', NOW(), NOW()),

-- 保时捷豪华跑车 (brand_id=22)
(2, 1, 22, '保时捷', '911 Turbo S', 2024, '火山灰', 800, 2599.00, 1, '保时捷911 Turbo S，3.8L双涡轮增压水平对置6缸发动机，650马力，传奇跑车', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '天津市和平区', '13800138011', NOW(), NOW()),

(3, 1, 22, '保时捷', 'Cayenne Turbo', 2023, '白金银', 2400, 1799.00, 1, '保时捷Cayenne Turbo，4.0L V8双涡轮增压，541马力，运动型豪华SUV', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '青岛市市南区', '13800138012', NOW(), NOW()),

-- 法拉利超级跑车 (brand_id=23)
(1, 1, 23, '法拉利', 'F8 Tributo', 2023, '法拉利红', 1500, 3999.00, 1, '法拉利F8 Tributo，3.9L V8双涡轮增压，720马力，意大利超跑艺术品', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '苏州市工业园区', '13800138013', NOW(), NOW()),

(2, 1, 23, '法拉利', 'Roma', 2024, '银石灰', 600, 3599.00, 1, '法拉利Roma GT跑车，3.9L V8涡轮增压，612马力，优雅与性能的完美融合', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '厦门市思明区', '13800138014', NOW(), NOW()),

-- 兰博基尼超级跑车 (brand_id=24)
(3, 1, 24, '兰博基尼', 'Huracán EVO', 2023, '橙色', 1200, 3799.00, 1, '兰博基尼Huracán EVO，5.2L V10自然吸气，640马力，意式狂野美学', 'https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '长沙市岳麓区', '13800138015', NOW(), NOW()),

(1, 1, 24, '兰博基尼', 'Urus', 2024, '珍珠胶囊白', 800, 2999.00, 1, '兰博基尼Urus超级SUV，4.0L V8双涡轮增压，650马力，最快的SUV', 'https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '郑州市金水区', '13800138016', NOW(), NOW()),

-- 玛莎拉蒂豪华轿车 (brand_id=25)
(2, 7, 25, '玛莎拉蒂', 'Quattroporte', 2023, '蓝色激情', 2600, 1599.00, 1, '玛莎拉蒂Quattroporte，3.0L V6双涡轮增压，意式优雅与性能的象征', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '三亚市天涯区', '13800138017', NOW(), NOW()),

(3, 7, 25, '玛莎拉蒂', 'Levante', 2024, '黑曜石', 1800, 1299.00, 1, '玛莎拉蒂Levante豪华SUV，3.0L V6双涡轮增压，意式风情与实用性并存', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '大连市中山区', '13800138018', NOW(), NOW()),

-- 劳斯莱斯顶级豪华轿车 (需要先添加劳斯莱斯品牌，假设brand_id=28)
(1, 7, 28, '劳斯莱斯', '幻影', 2024, '英伦白', 500, 4999.00, 1, '劳斯莱斯幻影，6.75L V12双涡轮增压，手工打造，奢华之巅', 'https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '海口市美兰区', '13800138019', NOW(), NOW()),

(2, 7, 28, '劳斯莱斯', '库里南', 2023, '钻石黑', 1200, 3999.00, 1, '劳斯莱斯库里南豪华SUV，全地形能力，至尊奢华体验，SUV中的王者', 'https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8?w=800,https://images.unsplash.com/photo-1555215695-3004980ad54e?w=800,https://images.unsplash.com/photo-1617886322207-d4c1cc4e4ea1?w=800,https://images.unsplash.com/photo-1617814076367-b759c7d7e738?w=800,https://images.unsplash.com/photo-1606664515524-ed2f786a0bd6?w=800', '昆明市五华区', '13800138020', NOW(), NOW());

-- ==================== 执行说明 ====================
-- 1. 执行前请确保以下表已存在：
--    - merchants (商家表)
--    - vehicle_types (车辆类型表)
--    - vehicle_brands (车辆品牌表)
--    - vehicles (车辆表)

-- 2. 请根据实际情况调整以下ID：
--    - merchant_id: 商家ID (1, 2, 3)
--    - type_id: 车辆类型ID (3=豪华轿车, 4=豪华SUV, 5=超级跑车)
--    - brand_id: 品牌ID (根据vehicle_brands表中的实际ID)

-- 3. 图片URL使用的是Unsplash的高质量汽车图片，可以正常显示

-- 4. 价格单位为人民币/天，已设置为豪华车合理价格区间

-- 5. 车辆状态默认为1(上架)，可根据需要调整

-- 6. 执行命令示例：
--    mysql -u username -p database_name < insert_test_vehicles.sql

-- ==================== 车辆统计 ====================
-- 总计插入20辆豪华车辆：
-- - 奔驰: 2辆 (S500L, S680迈巴赫)
-- - 宝马: 2辆 (760Li, X7 M50i)
-- - 奥迪: 2辆 (A8L, Q8)
-- - 路虎: 2辆 (揽胜, 揽胜运动版SVR)
-- - 宾利: 2辆 (飞驰V8, 添越V8)
-- - 保时捷: 2辆 (911 Turbo S, Cayenne Turbo)
-- - 法拉利: 2辆 (F8 Tributo, Roma)
-- - 兰博基尼: 2辆 (Huracán EVO, Urus)
-- - 玛莎拉蒂: 2辆 (Quattroporte, Levante)
-- - 劳斯莱斯: 2辆 (幻影, 库里南)

-- 价格区间：1099-4999元/天
-- 覆盖城市：北京、上海、深圳、广州、杭州、成都、武汉、南京、西安、重庆、天津、青岛、苏州、厦门、长沙、郑州、三亚、大连、海口、昆明
