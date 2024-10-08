<template>
  <div id="container"></div>
</template>

<script setup>
import {ref, onMounted, nextTick} from 'vue';
import axios from "axios";
import * as echarts from "echarts";

const map = ref(null);
const stations=ref([]);

const refresh = ()=>{
  axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
      .then((response)=>{
        if(response.data.code==1) {
          stations.value = response.data.stations
          console.log(stations.value[0]["lon"])
initMap();
          }
      })
      .catch((error)=>{
        console.log(error)
      })

}

const initMap = () => {
  // 初始化地图
  map.value = new AMap.Map('container', {
    resizeEnable: true,
    center: [120.204748,30.257602], // 地图中心坐标
    zoom: 13
  });

  // 创建标记点
console.log([stations.value[0]["lon"], stations.value[0]["lat"]])
  for(var i=0;i<stations.value.length;i++){
     const marker = new AMap.Marker({
     position: [stations.value[i]["lon"], stations.value[i]["lat"]], // 标记点的坐标
        map: map.value,
       label: {
        content: stations.value[i]["name"],    // 站点名称
        direction: 'top',       // 标签显示在标记点的顶部
        offset: new AMap.Pixel(0, -35) // 标签的偏移量，调整其相对位置
      }
      });

  }

};

onMounted(() => {
  // 确保在 DOM 加载完成后再初始化地图
  refresh();


});
</script>

<style scoped>
#container {
  width: 100%;
  height: 500px;
}
</style>
