<template>
    <div id="container"></div>
</template>




<script>
export default {
  data() {
    return {
      map: null,
      markers: [],
      points: [],
      // 假设有一个站点数组，这里用随机生成的点代替
      stations: this.generateRandomStations(100),
    };
  },
  mounted() {
    this.initMap();
  },
  methods: {
    initMap() {
      this.map = new AMap.Map('container', {
        resizeEnable: true,
        center: [116.397428, 39.90923],
        zoom: 13
      });

      this.map.on('click', (e) => {
        this.placeMarker(e.lnglat);
      });
    },
    placeMarker(lnglat) {
      // 清除之前的标记和点
      if (this.markers.length >= 2){
        console.log(1);
        this.clearMarkersAndPoints();
      }
      

      let marker = new AMap.Marker({
        position: lnglat,
        map: this.map
      });

      this.markers.push(marker);
      this.points.push(lnglat);


      if (this.markers.length === 2) {
        this.findNearestStations();
      }
    },
    clearMarkersAndPoints() {
        this.markers.forEach(function(marker) {
            marker.setMap(null);
            
        });
        this.markers = [];
        this.points = [];
    },
    findNearestStations() {
      // 这里使用简单的距离计算方法来模拟查找最近站点
      // 实际应用中，你可能需要调用高德地图的API或使用数据库查询
      let nearestStations = this.stations
        .map(station => ({
          ...station,
          distanceToStart: AMap.GeometryUtil.distance(this.points[0], station.position),
          distanceToEnd: AMap.GeometryUtil.distance(this.points[1], station.position)
        }))
        .sort((a, b) => a.distanceToStart - b.distanceToStart)
        .slice(0, 2);

      this.showNearestStations(nearestStations);
    },
    showNearestStations(stations) {
      stations.forEach(station => {
        let marker = new AMap.Marker({
          position: station.position,
          map: this.map,
          content: '<div style="background-color: #FFA500; padding: 2px; border-radius: 3px;">站点</div>',
          offset: new AMap.Pixel(-15, -15)
        });
        this.markers.push(marker);
      });
      // 返回或处理最近站点数据
      console.log(stations);
    },
    generateRandomStations(count) {
      let stations = [];
      for (let i = 0; i < count; i++) {
        stations.push({
          position: [116.397428 + Math.random() * 0.1 - 0.05, 39.90923 + Math.random() * 0.1 - 0.05]
        });
      }
      return stations;
    }
  }
};
</script>



<style scoped>
#container{
    width:500px;
    height: 500px;
}
</style>