package logic

import (
	"Common/global"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"models/model_mysql"
	"strconv"
	"strings"
	"time"
)

// VehicleESDoc ES中的车辆文档结构
type VehicleESDoc struct {
	ID          int64     `json:"id"`
	MerchantID  int64     `json:"merchant_id"`
	TypeID      int64     `json:"type_id"`
	BrandID     int64     `json:"brand_id"`
	Brand       string    `json:"brand"`
	Style       string    `json:"style"`
	Year        int32     `json:"year"`
	Color       string    `json:"color"`
	Mileage     int32     `json:"mileage"`
	Price       float64   `json:"price"`
	Status      int32     `json:"status"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Contact     string    `json:"contact"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const VehicleIndexName = "vehicles"

// CreateVehicleIndex 创建车辆索引
func CreateVehicleIndex() error {
	// 定义索引映射
	mapping := `{
		"mappings": {
			"properties": {
				"id": {"type": "long"},
				"merchant_id": {"type": "long"},
				"type_id": {"type": "long"},
				"brand_id": {"type": "long"},
				"brand": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"style": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"year": {"type": "integer"},
				"color": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"mileage": {"type": "integer"},
				"price": {"type": "double"},
				"status": {"type": "integer"},
				"description": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"location": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"contact": {"type": "keyword"},
				"created_at": {"type": "date"},
				"updated_at": {"type": "date"}
			}
		}
	}`

	// 检查索引是否存在
	res, err := global.Es.Indices.Exists([]string{VehicleIndexName})
	if err != nil {
		return fmt.Errorf("检查索引失败: %v", err)
	}

	if res.StatusCode == 200 {
		log.Printf("索引 %s 已存在", VehicleIndexName)
		return nil
	}

	// 创建索引
	res, err = global.Es.Indices.Create(
		VehicleIndexName,
		global.Es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		return fmt.Errorf("创建索引失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("创建索引失败: %s", res.String())
	}

	log.Printf("索引 %s 创建成功", VehicleIndexName)
	return nil
}

// IndexVehicleToES 将车辆信息索引到ES
func IndexVehicleToES(vehicle *model_mysql.Vehicle) error {
	// 转换为ES文档格式
	doc := VehicleESDoc{
		ID:          int64(vehicle.ID),
		MerchantID:  vehicle.MerchantID,
		TypeID:      vehicle.TypeID,
		BrandID:     vehicle.BrandID,
		Brand:       vehicle.Brand,
		Style:       vehicle.Style,
		Year:        int32(vehicle.Year),
		Color:       vehicle.Color,
		Mileage:     int32(vehicle.Mileage),
		Price:       vehicle.Price,
		Status:      int32(vehicle.Status),
		Description: vehicle.Description,
		Location:    vehicle.Location,
		Contact:     vehicle.Contact,
		CreatedAt:   vehicle.CreatedAt,
		UpdatedAt:   vehicle.UpdatedAt,
	}

	// 序列化为JSON
	docJSON, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("序列化车辆文档失败: %v", err)
	}

	// 索引文档到ES
	res, err := global.Es.Index(
		VehicleIndexName,
		bytes.NewReader(docJSON),
		global.Es.Index.WithDocumentID(strconv.FormatInt(int64(vehicle.ID), 10)),
		global.Es.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("索引车辆到ES失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("索引车辆到ES失败: %s", res.String())
	}

	log.Printf("车辆 ID:%d 成功索引到ES", vehicle.ID)
	return nil
}

// UpdateVehicleInES 更新ES中的车辆信息
func UpdateVehicleInES(vehicle *model_mysql.Vehicle) error {
	return IndexVehicleToES(vehicle) // 更新和创建使用相同的逻辑
}

// DeleteVehicleFromES 从ES中删除车辆
func DeleteVehicleFromES(vehicleID uint) error {
	res, err := global.Es.Delete(
		VehicleIndexName,
		strconv.FormatInt(int64(vehicleID), 10),
		global.Es.Delete.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("从ES删除车辆失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() && res.StatusCode != 404 { // 404表示文档不存在，这是正常的
		return fmt.Errorf("从ES删除车辆失败: %s", res.String())
	}

	log.Printf("车辆 ID:%d 成功从ES删除", vehicleID)
	return nil
}

// SearchVehiclesInES 在ES中搜索车辆
func SearchVehiclesInES(keyword string, page, pageSize int64, filters map[string]interface{}) ([]VehicleESDoc, int64, error) {
	// 构建搜索查询
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{},
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
		"sort": []interface{}{
			map[string]interface{}{
				"created_at": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	mustQueries := []interface{}{}

	// 关键词搜索
	if keyword != "" {
		mustQueries = append(mustQueries, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"brand^3", "style^2", "description", "location", "color"},
				"type":   "best_fields",
			},
		})
	}

	// 添加过滤条件
	if merchantID, ok := filters["merchant_id"]; ok && merchantID != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"term": map[string]interface{}{
				"merchant_id": merchantID,
			},
		})
	}

	if typeID, ok := filters["type_id"]; ok && typeID != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"term": map[string]interface{}{
				"type_id": typeID,
			},
		})
	}

	if brandID, ok := filters["brand_id"]; ok && brandID != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"term": map[string]interface{}{
				"brand_id": brandID,
			},
		})
	}

	if status, ok := filters["status"]; ok && status != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"term": map[string]interface{}{
				"status": status,
			},
		})
	}

	// 价格范围
	if priceMin, ok := filters["price_min"]; ok && priceMin != nil {
		if priceMax, ok := filters["price_max"]; ok && priceMax != nil {
			mustQueries = append(mustQueries, map[string]interface{}{
				"range": map[string]interface{}{
					"price": map[string]interface{}{
						"gte": priceMin,
						"lte": priceMax,
					},
				},
			})
		} else {
			mustQueries = append(mustQueries, map[string]interface{}{
				"range": map[string]interface{}{
					"price": map[string]interface{}{
						"gte": priceMin,
					},
				},
			})
		}
	} else if priceMax, ok := filters["price_max"]; ok && priceMax != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"range": map[string]interface{}{
				"price": map[string]interface{}{
					"lte": priceMax,
				},
			},
		})
	}

	// 年份范围
	if yearMin, ok := filters["year_min"]; ok && yearMin != nil {
		if yearMax, ok := filters["year_max"]; ok && yearMax != nil {
			mustQueries = append(mustQueries, map[string]interface{}{
				"range": map[string]interface{}{
					"year": map[string]interface{}{
						"gte": yearMin,
						"lte": yearMax,
					},
				},
			})
		} else {
			mustQueries = append(mustQueries, map[string]interface{}{
				"range": map[string]interface{}{
					"year": map[string]interface{}{
						"gte": yearMin,
					},
				},
			})
		}
	} else if yearMax, ok := filters["year_max"]; ok && yearMax != nil {
		mustQueries = append(mustQueries, map[string]interface{}{
			"range": map[string]interface{}{
				"year": map[string]interface{}{
					"lte": yearMax,
				},
			},
		})
	}

	// 如果没有任何查询条件，使用match_all
	if len(mustQueries) == 0 {
		query["query"] = map[string]interface{}{
			"match_all": map[string]interface{}{},
		}
	} else {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = mustQueries
	}

	// 序列化查询
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, 0, fmt.Errorf("序列化查询失败: %v", err)
	}

	// 执行搜索
	res, err := global.Es.Search(
		global.Es.Search.WithIndex(VehicleIndexName),
		global.Es.Search.WithBody(bytes.NewReader(queryJSON)),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("ES搜索失败: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("ES搜索失败: %s", res.String())
	}

	// 解析搜索结果
	var searchResult struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source VehicleESDoc `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return nil, 0, fmt.Errorf("解析搜索结果失败: %v", err)
	}

	// 提取车辆数据
	vehicles := make([]VehicleESDoc, 0, len(searchResult.Hits.Hits))
	for _, hit := range searchResult.Hits.Hits {
		vehicles = append(vehicles, hit.Source)
	}

	return vehicles, searchResult.Hits.Total.Value, nil
}
