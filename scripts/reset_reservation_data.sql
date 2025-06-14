-- 重置预订数据脚本
-- 用于清理测试过程中产生的脏数据

-- 1. 重置预订记录的order_id为0（如果对应的订单不存在）
UPDATE vehicle_inventories 
SET order_id = 0 
WHERE order_id > 0 
AND order_id NOT IN (SELECT id FROM orders);

-- 2. 删除没有对应预订的订单记录
DELETE FROM orders 
WHERE reservation_id > 0 
AND reservation_id NOT IN (SELECT id FROM vehicle_inventories);

-- 3. 查看当前预订状态
SELECT 
    vi.id as reservation_id,
    vi.vehicle_id,
    vi.order_id,
    vi.status,
    vi.start_date,
    vi.end_date,
    vi.created_by,
    o.id as actual_order_id,
    o.order_sn
FROM vehicle_inventories vi
LEFT JOIN orders o ON vi.order_id = o.id
WHERE vi.deleted_at IS NULL
ORDER BY vi.id DESC;

-- 4. 查看当前订单状态
SELECT 
    o.id,
    o.order_sn,
    o.reservation_id,
    o.status,
    o.total_amount,
    vi.id as actual_reservation_id
FROM orders o
LEFT JOIN vehicle_inventories vi ON o.reservation_id = vi.id
WHERE o.deleted_at IS NULL
ORDER BY o.id DESC;
