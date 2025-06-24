/**
 * 高德地图工具类
 */
class MapUtils {
  constructor() {
    this.map = null;
    this.markers = [];
    this.infoWindows = [];
  }

  /**
   * 检查高德地图API是否已加载
   * @returns {boolean}
   */
  static isAmapLoaded() {
    return typeof window.AMap !== 'undefined';
  }

  /**
   * 等待高德地图API加载完成
   * @returns {Promise}
   */
  static waitForAmapLoad() {
    return new Promise((resolve, reject) => {
      if (MapUtils.isAmapLoaded()) {
        resolve();
        return;
      }

      const checkInterval = setInterval(() => {
        if (MapUtils.isAmapLoaded()) {
          clearInterval(checkInterval);
          resolve();
        }
      }, 100);

      // 10秒超时
      setTimeout(() => {
        clearInterval(checkInterval);
        reject(new Error('高德地图API加载超时'));
      }, 10000);
    });
  }

  /**
   * 创建地图实例
   * @param {string} containerId - 地图容器ID
   * @param {Object} options - 地图配置选项
   * @returns {Promise<AMap.Map>}
   */
  async createMap(containerId, options = {}) {
    await MapUtils.waitForAmapLoad();

    const defaultOptions = {
      center: [116.397428, 39.90923], // 北京天安门
      zoom: 13,
      mapStyle: 'amap://styles/normal',
      features: ['bg', 'road', 'building', 'point'],
      viewMode: '2D'
    };

    const mapOptions = { ...defaultOptions, ...options };
    this.map = new window.AMap.Map(containerId, mapOptions);

    // 添加默认控件
    this.addDefaultControls();

    return this.map;
  }

  /**
   * 添加默认控件
   */
  addDefaultControls() {
    if (!this.map) return;

    // 比例尺
    this.map.addControl(new window.AMap.Scale({
      position: 'LB'
    }));

    // 工具条
    this.map.addControl(new window.AMap.ToolBar({
      position: 'RT'
    }));

    // 控制条
    this.map.addControl(new window.AMap.ControlBar({
      position: 'RT'
    }));
  }

  /**
   * 添加标记点
   * @param {Array} position - 经纬度 [lng, lat]
   * @param {Object} options - 标记选项
   * @returns {AMap.Marker}
   */
  addMarker(position, options = {}) {
    if (!this.map) return null;

    const defaultOptions = {
      position: position,
      anchor: 'bottom-center',
      draggable: false
    };

    const markerOptions = { ...defaultOptions, ...options };
    const marker = new window.AMap.Marker(markerOptions);

    this.map.add(marker);
    this.markers.push(marker);

    return marker;
  }

  /**
   * 添加信息窗体
   * @param {Array} position - 经纬度 [lng, lat]
   * @param {string} content - 窗体内容
   * @param {Object} options - 窗体选项
   * @returns {AMap.InfoWindow}
   */
  addInfoWindow(position, content, options = {}) {
    if (!this.map) return null;

    const defaultOptions = {
      content: content,
      anchor: 'bottom-center',
      offset: new window.AMap.Pixel(0, -30)
    };

    const infoWindowOptions = { ...defaultOptions, ...options };
    const infoWindow = new window.AMap.InfoWindow(infoWindowOptions);

    infoWindow.open(this.map, position);
    this.infoWindows.push(infoWindow);

    return infoWindow;
  }

  /**
   * 清除所有标记
   */
  clearMarkers() {
    if (!this.map) return;

    this.markers.forEach(marker => {
      this.map.remove(marker);
    });
    this.markers = [];
  }

  /**
   * 清除所有信息窗体
   */
  clearInfoWindows() {
    this.infoWindows.forEach(infoWindow => {
      infoWindow.close();
    });
    this.infoWindows = [];
  }

  /**
   * 地理编码 - 地址转坐标
   * @param {string} address - 地址
   * @returns {Promise<Object>}
   */
  geocode(address) {
    return new Promise((resolve, reject) => {
      if (!this.map) {
        reject(new Error('地图未初始化'));
        return;
      }

      const geocoder = new window.AMap.Geocoder();
      geocoder.getLocation(address, (status, result) => {
        if (status === 'complete' && result.geocodes.length > 0) {
          const geocode = result.geocodes[0];
          resolve({
            longitude: geocode.location.lng,
            latitude: geocode.location.lat,
            address: geocode.formattedAddress,
            level: geocode.level
          });
        } else {
          reject(new Error('地址解析失败'));
        }
      });
    });
  }

  /**
   * 逆地理编码 - 坐标转地址
   * @param {Array} position - 经纬度 [lng, lat]
   * @returns {Promise<Object>}
   */
  regeocode(position) {
    return new Promise((resolve, reject) => {
      if (!this.map) {
        reject(new Error('地图未初始化'));
        return;
      }

      const geocoder = new window.AMap.Geocoder();
      geocoder.getAddress(position, (status, result) => {
        if (status === 'complete' && result.regeocode) {
          const regeocode = result.regeocode;
          resolve({
            longitude: position[0],
            latitude: position[1],
            address: regeocode.formattedAddress,
            addressComponent: regeocode.addressComponent
          });
        } else {
          reject(new Error('坐标解析失败'));
        }
      });
    });
  }

  /**
   * 地点搜索
   * @param {string} keyword - 搜索关键词
   * @param {Object} options - 搜索选项
   * @returns {Promise<Array>}
   */
  searchPlace(keyword, options = {}) {
    return new Promise((resolve, reject) => {
      if (!this.map) {
        reject(new Error('地图未初始化'));
        return;
      }

      const defaultOptions = {
        pageSize: 10,
        pageIndex: 1,
        city: '全国'
      };

      const searchOptions = { ...defaultOptions, ...options };
      const placeSearch = new window.AMap.PlaceSearch(searchOptions);

      placeSearch.search(keyword, (status, result) => {
        if (status === 'complete' && result.poiList && result.poiList.pois.length > 0) {
          const pois = result.poiList.pois.map(poi => ({
            id: poi.id,
            name: poi.name,
            address: poi.address,
            longitude: poi.location.lng,
            latitude: poi.location.lat,
            type: poi.type,
            tel: poi.tel
          }));
          resolve(pois);
        } else {
          reject(new Error('搜索无结果'));
        }
      });
    });
  }

  /**
   * 获取当前位置
   * @returns {Promise<Object>}
   */
  getCurrentPosition() {
    return new Promise((resolve, reject) => {
      if (!this.map) {
        reject(new Error('地图未初始化'));
        return;
      }

      const geolocation = new window.AMap.Geolocation({
        enableHighAccuracy: true,
        timeout: 10000,
        maximumAge: 0,
        convert: true
      });

      geolocation.getCurrentPosition((status, result) => {
        if (status === 'complete') {
          resolve({
            longitude: result.position.lng,
            latitude: result.position.lat,
            address: result.formattedAddress,
            accuracy: result.accuracy
          });
        } else {
          reject(new Error(result.message || '定位失败'));
        }
      });
    });
  }

  /**
   * 设置地图中心点
   * @param {Array} center - 中心点坐标 [lng, lat]
   * @param {number} zoom - 缩放级别
   */
  setCenter(center, zoom) {
    if (!this.map) return;

    this.map.setCenter(center);
    if (zoom !== undefined) {
      this.map.setZoom(zoom);
    }
  }

  /**
   * 适应显示所有标记
   */
  fitView() {
    if (!this.map || this.markers.length === 0) return;

    this.map.setFitView(this.markers);
  }

  /**
   * 计算两点间距离
   * @param {Array} point1 - 点1坐标 [lng, lat]
   * @param {Array} point2 - 点2坐标 [lng, lat]
   * @returns {number} 距离(米)
   */
  static calculateDistance(point1, point2) {
    if (!window.AMap) return 0;

    const lngLat1 = new window.AMap.LngLat(point1[0], point1[1]);
    const lngLat2 = new window.AMap.LngLat(point2[0], point2[1]);
    
    return lngLat1.distance(lngLat2);
  }

  /**
   * 销毁地图
   */
  destroy() {
    if (this.map) {
      this.clearMarkers();
      this.clearInfoWindows();
      this.map.destroy();
      this.map = null;
    }
  }
}

export default MapUtils;
