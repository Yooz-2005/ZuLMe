// MongoDB查询脚本 - 检查评论删除状态
// 在MongoDB Compass或mongo shell中运行

// 查看所有评论及其状态
db.comments.find({}, {
  order_id: 1,
  user_id: 1,
  vehicle_id: 1,
  content: 1,
  status: 1,
  created_at: 1,
  updated_at: 1
}).pretty()

// 查看已删除的评论 (status = 3)
db.comments.find({status: 3}, {
  order_id: 1,
  user_id: 1,
  vehicle_id: 1,
  content: 1,
  status: 1,
  created_at: 1,
  updated_at: 1
}).pretty()

// 查看正常状态的评论 (status = 1)
db.comments.find({status: 1}, {
  order_id: 1,
  user_id: 1,
  vehicle_id: 1,
  content: 1,
  status: 1,
  created_at: 1,
  updated_at: 1
}).pretty()

// 统计各状态的评论数量
db.comments.aggregate([
  {
    $group: {
      _id: "$status",
      count: { $sum: 1 }
    }
  },
  {
    $project: {
      status: {
        $switch: {
          branches: [
            { case: { $eq: ["$_id", 1] }, then: "正常" },
            { case: { $eq: ["$_id", 2] }, then: "隐藏" },
            { case: { $eq: ["$_id", 3] }, then: "已删除" }
          ],
          default: "未知状态"
        }
      },
      count: 1
    }
  }
])
