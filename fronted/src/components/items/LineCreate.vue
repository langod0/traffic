<template>
  <el-tabs type="border-card" v-model="activate" style="height: 100%" @tab-change="handleChange"  >

    <el-tab-pane label="查看线路信息" name="1" >
      <button @click="linevisible = true;" class="mb">创建线路</button>
      <button @click="mapshow" class="mb">线路查看</button>
      <div class="mainline" style="display: inline-flex">

        <el-table :data="lines" stripe height="100%" style="width: 450px;margin-right: 10px;color: #02040f"  >
          <el-table-column prop="line_id" label="线路编号" width="150" ></el-table-column>
          <el-table-column prop="name" label="线路名称" width="200"   />
          <el-table-column width="100">

            <template #default="scope">
              <el-button  @click="updata1(scope.row)">...</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div ref="chart" style="width: calc(100% - 450px - 20px);height: 600px;margin-top: 30px;margin-left: 10px" v-if="chartVis"></div>

      </div>
    </el-tab-pane>
    <el-tab-pane label="查看站点信息"  name="2" >
      <button @click="stationvisible=true" class="mb">创建站点</button>
<button @click="mapshow" class="mb">线路查看</button>
      <div class="mainline">
        <el-table :data="stations" stripe height="100%" style="width:950px"  >
          <el-table-column prop="id" label="站点编号" width="150"  />
          <el-table-column prop="lon" label="经度" width="150"   />
          <el-table-column prop="lat" label="纬度" width="150"   />
          <el-table-column prop="name" label="站点名称" width="200"   />
          <el-table-column  label="线路编号" width="200"   ><template #default="scope">
             {{fdline(scope.row)}}
            </template></el-table-column>
          <el-table-column width="100">
            <template #default="scope">
              <el-button  @click="updata2(scope.row)">...</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-tab-pane>

    <el-tab-pane label="查看列车信息"  name="3" >
      <button @click="trainvisible = true" class="mb">创建列车</button>

      <div class="mainline">
      <el-table :data="trains" stripe height="100%" style="width: 750px"  >
        <el-table-column prop="id" label="列车编号" width="150"  />
        <el-table-column prop="line_id" label="所属线路" width="150"   />
        <el-table-column prop="capacity" label="载客量" width="150"   />
        <el-table-column width="100">
          <template #default="scope">
            <el-button  @click="updata3(scope.row)">...</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    </el-tab-pane>
  </el-tabs>
    <el-dialog v-model="linevisible" title="创建新线路" width="500">
      <div class="line-create" >
        <h2>创建新线路</h2>
        <form @submit.prevent="createLine" >
          <div class="form-content">
            <label for="line-name">线路名称:</label>
            <input type="text" id="line-name" v-model="new_line_name" placeholder="线路名称" required>
          </div>
          <button type="submit" @click="upline">提交</button>
        </form>
      </div>
    </el-dialog>
    <el-dialog v-model="stationvisible" title="创建新站点" width="500">
      <div class="line-create" >
        <h2>创建新站点</h2>
        <form @submit.prevent="createLine" >
          <div class="form-content">
            <label for="line-name">站点名称:</label>
            <input type="text" id="line-name" v-model="new_station_name" placeholder="站点名称" required>
            <label for="line-name">站点经度:</label>
            <input type="text" id="line-name" v-model="new_station_lon" placeholder="站点经度" required>
            <label for="line-name">站点纬度:</label>
            <input type="text" id="line-name" v-model="new_station_lat" placeholder="站点纬度" required>


          </div>
          <button type="submit" @click="upstation">提交</button>
        </form>
      </div>

    </el-dialog>
    <el-dialog v-model="trainvisible" title="创建新列车" width="500">
      <div class="line-create" >
        <h2>创建新列车</h2>
        <form @submit.prevent="createLine" >
          <div class="form-content">
            <label for="line-name">列车ID</label>
            <input type="text" id="line-name" v-model="new_train_id" placeholder="列车ID" required>
            <label for="line-name">列车载客量:</label>
            <input type="text" id="line-name" v-model="new_train_cap" placeholder="列车载客量" required>
            <label for="line-name">列车所在线路编号:</label>
            <input type="text" id="line-name" v-model="new_train_lineid" placeholder="列车所在线路编号" required>
          </div>
          <button type="submit" @click="uptrain">提交</button>
        </form>
      </div>

    </el-dialog>
    <el-dialog v-model="lineexvisible" title="编辑线路" width="500">

<!--    <div style="width: 80%;height: 80%;margin-left: 100px;margin-top: 60px;background-color: white;float: left" >-->
      <div class="line-create">
        线路编号：<label   class="inp">{{line_id}}</label><br>
        线路名称：<input type="text" v-model="line_name" class="inp"><br>
        <el-button @click="updatan1" >保存</el-button>
        <el-button @click="deline = true">删除信息</el-button>
      </div>
    </el-dialog>
    <el-dialog v-model="stationexvisible" title="编辑站点" width="500">

      <div class="line-create">
        站点编号：<label   class="inp">{{station_id}}</label><br>
        &nbsp经&nbsp&nbsp&nbsp&nbsp度&nbsp：<input type="text" v-model="station_lon" class="inp"><br>
        &nbsp纬&nbsp&nbsp&nbsp&nbsp度&nbsp：<input type="text" v-model="station_lat" class="inp"><br>
        站点名称：<input type="text" v-model="station_name" class="inp"><br>

        <el-button @click="updatan2" >保存</el-button>
        <el-button  @click="destation=true">删除信息</el-button>
      </div>

    </el-dialog>
    <el-dialog v-model="trainexvisible" title="编辑列车" width="500">

      <div class="line-create">
        列车编号：<label   class="inp">{{train_id}}</label><br>
        线路编号：<input type="text" v-model="train_line_id" class="inp"><br>
        capacity：<input type="text" v-model="train_capacity" class="inp"><br>

        <el-button @click="updatan3" >保存</el-button>
        <el-button  @click="detrain=true">删除信息</el-button>


    </div>
    </el-dialog>
  <el-dialog v-model="deline"  width="400" style="margin-top: 400px">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header" style="align-items: center;display: flex;">
         警告<el-icon :size="30" color="red"><WarnTriangleFilled /></el-icon>
      </div>
    </template>
    <div>
      <span>
        确定要删除吗？
      </span>
    </div>
    <br>
    <el-button size="small" @click="deline=false">取消</el-button>
    <el-button
        type="danger"
        size="small"
        @click="deline=false,delete1(now)"
    >
      确定
    </el-button>
  </el-dialog>
  <el-dialog v-model="destation"  width="400" style="margin-top: 400px">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header" style="align-items: center;display: flex;">
        警告<el-icon :size="30" color="red"><WarnTriangleFilled /></el-icon>
      </div>
    </template>
    <div>
      <span>
        确定要删除吗？
      </span>
    </div>
    <br>
    <el-button size="small" @click="destation=false">取消</el-button>
    <el-button
        type="danger"
        size="small"
        @click="destation=false,delete2(now)"
    >
      确定
    </el-button>
  </el-dialog>
  <el-dialog v-model="detrain"  width="400" style="margin-top: 400px">
    <template #header="{ close, titleId, titleClass }">
      <div class="my-header" style="align-items: center;display: flex;">
        警告<el-icon :size="30" color="red"><WarnTriangleFilled /></el-icon>
      </div>
    </template>
    <div>
      <span>
        确定要删除吗？
      </span>
    </div>
    <br>
    <el-button size="small" @click="detrain=false">取消</el-button>
    <el-button
        type="danger"
        size="small"
        @click="detrain=false,delete3(now)"
    >
      确定
    </el-button>
  </el-dialog>
  <el-dialog v-model="mapflag" width="800px"  style="margin-top: 200px">
    <mapvue/>
  </el-dialog>
</template>

<script setup>
import {nextTick, onMounted, ref} from 'vue';
import * as echarts from "echarts";

import axios, {all} from 'axios';
import router from "@/router/index.js";
import { InfoFilled } from '@element-plus/icons-vue'
import mapvue from './Map.vue'
const new_line_name=ref('')
const lines=ref("")
const stations=ref("")
const  trains=ref("")
const user=ref("")
const now=ref("")
const line_id=ref("")
const line_name=ref("")
const station_id=ref("")
const station_name=ref("")
const station_lon=ref("")
const station_lat=ref("")
const train_id=ref("")
const train_line_id=ref("")
const train_capacity=ref("")
const new_train_name=ref("")
const new_train_cap=ref("")
const new_train_id=ref("")
const new_train_lineid=ref("")
const new_station_name=ref("")
const new_station_lon=ref("")
const new_station_lat=ref("")
const linevisible = ref(false)
const stationvisible = ref(false)
const trainvisible = ref(false)
const lineexvisible = ref(false)
const stationexvisible = ref(false)
const trainexvisible = ref(false)
const deline = ref(false)
const destation = ref(false)
const detrain = ref(false)
const relations = ref([])
const xData = ref([])
const trData = ref([])
const StData = ref([])
const chartVis = ref(true)
const avaliableTr = ref(0)
const mapflag=ref(false)
const stline=ref({})
let v = 1;
const activate = ref("1")
const handleChange = (newTab)=>{
    if (newTab === '1') {
      ChartInit();
    }
}
const updata1=(e)=>{
  lineexvisible.value=true;
  now.value=e;
    line_id.value=now.value["line_id"]
    line_id.value=Number(line_id.value)
    line_name.value=now.value["name"]
}
const labelOption = {
  show: true,
  rotate: 0,
  formatter: (params) => {
    // 如果值为 0 或 null，则返回空字符串，不显示标签
    return params.value === 0 || params.value === null ? '0' :`${params.value} `;
  },
  fontSize: 16,
  color: '#FFFFFF',
  position: 'top',
  distance: 0,
  offset: [0, 0],
  rich: {
    name: {}
  }
};
let chartInstance = null
const chart = ref(null)
const ChartInit=()=>{
  chartInstance.clear();
  const options = {
    title:{
      text:"线网分析" ,
      textStyle:{
        color:'#FFFFFF',
        fontSize: 36,
      }
    },
    ss:v^1,
    backgroundColor:'#2d3455',
    legend: {
      data: ['站点数量', '列车数量'], // 图例中展示的标签，和系列名称一致
      top: '5%', // 设置图例的位置，可以调整位置
      textStyle: {
        color: '#FFFFFF', // 图例文字颜色
      },
    },
    itemStyle: {
      barBorderRadius: 5,
      borderWidth: 1,
      borderType: 'solid',
      shadowBlur: 3
    },
    xAxis: {
      data: xData.value,
      axisLabel: {
        textStyle: {
          color: '#FFFFFF', // 设置 x 轴标签字体颜色为黑色
        },
      },
    },
    yAxis:{
      axisLabel:{
        textStyle:{
          color:'#FFFFFF',
        }
      }
    },
    series:[
        {
          name: '站点数量',
          type: 'bar',
          label: labelOption,
          emphasis: {
            focus: 'series'
          },
          barWidth:24,
          data: StData.value,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              { offset: 0, color: '#fb834b' },
              { offset: 1, color: '#fed201' }
            ])
          }
        },
        {
          name: '列车数量',
          type: 'bar',
          barWidth:24,
          label: labelOption,
          emphasis: {
            focus: 'series'
          },
          data: trData.value,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              { offset: 0, color: '#16beaf' },
              { offset: 1, color: '#6dab63' }
            ])
          }
        }
    ],
    graphic: {
      elements: [
        {
          type: 'text', // 添加文本类型的标签
          left: '75%', // 水平方向的位置
          top: '5%', // 垂直方向的位置
          style: {
            text: avaliableTr.value>0?`有 ${avaliableTr.value} 辆列车空闲 `:'', // 文本内容
            fontSize: 24,
            fontWeight: 'bold',
            fill: 'rgba(231,225,225,0.97)', // 字体颜色
          }
        }
      ]
    }
  }
  chartInstance.setOption(options)
}

const refresh = ()=>{
  axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
      .then((response)=>{
        if(response.data.code==1) {
          console.log(response.data.lines)
          lines.value = response.data.lines
          stations.value = response.data.stations
          for(var i=0;i<stations.value.length;i++){
            stline.value[stations.value[i]["id"]]="";
          }
          trains.value = response.data.trains
          user.value = response.data.user
          axios.get("goapi/api/getrelations",{headers:{'Authorization': localStorage.getItem("Authorization")}}).then((response)=>{
              if (response.data.code==1){
                relations.value = response.data.result
                console.log(relations.value[0])
                let mm = new Map()
                let idx = 0
                xData.value = []
                for (let i = 0; i < lines.value.length; i++) {
                  xData.value.push(lines.value[i].name)
                  mm.set(lines.value[i].line_id,idx)
                  idx++
                }
                avaliableTr.value=0
                trData.value = new Array(idx).fill(0)
                StData.value = new Array(idx).fill(0)
                for (let i = 0; i < trains.value.length; i++) {
                    if (mm.has(trains.value[i].line_id)){
                      trData.value[mm.get(trains.value[i].line_id)]++
                    }else{
                      avaliableTr.value++
                    }
                }
                for (let i = 0;i< relations.value.length;i++) {
                  stline.value[relations.value[i]["SubwayStationId"]]=stline.value[relations.value[i]["SubwayStationId"]]+String(relations.value[i]["SubwayLineId"])+" ";
                  if (mm.has(relations.value[i].SubwayLineId)){
                    StData.value[mm.get(relations.value[i].SubwayLineId)]++
                  }
                }
                console.log(trData.value,StData.value,mm)
                  ChartInit()
              }
          }).catch((error)=>{
            console.log(error);
          })
        }else{
          alert(response.data.error)
        }
      })
      .catch((error)=>{
        console.log(error)
      })

}
onMounted(()=>{
  nextTick(()=>{
    chartInstance = echarts.init(chart.value);
    refresh()
  })
})
const updatan1=()=>{
     axios.post("goapi/api/updateline",{"line_id":line_id.value,"name":line_name.value,"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              refresh()
              alert("更新成功")
            }else{
              alert(response.data.error)
            }
          })
    refresh()
}
const updatan2=()=>{
   axios.post("goapi/api/updatestation",{"id":station_id.value,"name":station_name.value,"lon":station_lon.value,"lat":station_lat.value,"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("更新成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  refresh()
}
const updatan3=()=>{
    axios.post("goapi/api/updatetrain",{"id":String(train_id.value),"line_id":Number(train_line_id.value),"cap":Number(train_capacity.value),"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("更新成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  refresh()
}
const updata2=(e)=>{
  stationexvisible.value=true
      now.value=e;
  station_id.value=now.value["id"]
  station_name.value=now.value["name"]
  station_lon.value=now.value["lon"]
  station_lat.value=now.value["lat"]
  refresh()
}
const updata3=(e)=>{
  trainexvisible.value=true
  now.value=e;
  train_id.value=now.value["id"]
  train_line_id.value=now.value["line_id"]
  train_capacity.value=now.value["capacity"]
  refresh()
}
const delete1=(e)=>{
      now.value=e;
      axios.post("goapi/api/updateline",{"line_id":now.value["line_id"],"name":now.value["name"],"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("删除成功")
              refresh()
              // getline()
            }else{
              alert(response.data.error)
            }
          })
  lineexvisible.value= false
  refresh()
}
const delete2=(e)=>{
      now.value=e;
       axios.post("goapi/api/updatestation",{"id":now.value["id"],"name":now.value["name"],"lon":now.value["lon"],"lat":now.value["lat"],"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                alert("删除成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  stationexvisible.value = false
  refresh()
}
const  delete3=(e)=>{
      now.value=e
      axios.post("goapi/api/updatetrain",{"id":String(now.value["id"]),"line_id":Number(now.value["line_id"]),"cap":Number(now.value["capacity"]),"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("删除成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  trainexvisible.value = false
  refresh()
}
  const upline=()=>{
      axios.post("goapi/api/updateline",{"line_id":10000,"name":new_line_name.value,"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    refresh()
  }
  const upstation=()=>{
      axios.post("goapi/api/updatestation",{"id":1000000,"name":new_station_name.value,"lon":new_station_lon.value,"lat":new_station_lat.value,"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    refresh()
  }
  const uptrain=()=>{
      axios.post("goapi/api/updatetrain",{"id":new_train_id.value,"cap":Number(new_train_cap.value),"line_id":Number(new_train_lineid.value),"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
              refresh()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    refresh()
  }
const mapshow=()=>{
    mapflag.value=true
    console.log(mapflag.value)
}
const fdline=(e)=>{
  return stline.value[e["id"]]

}

</script>

<style scoped>
.bt{
  margin-left:150px;
  margin-top: 30px;
  width: 100px;
  height: 30px;
  transition: 0.6s;
  border-radius: 10px;
  cursor: pointer;
}
.bt:hover{
  background-color: #9acfea;
}
.inp{
  width: 250px;
  height: 30px;
  border-radius: 5px;
  margin-top: 40px;


}
.mb{
  width: 120px;
  height: 40px;
  margin-left: 40px;
  cursor: pointer;
  transition:0.4s;
  border: none;
  border-radius: 10px
}
.mainline{
  width: 96%;
  height: 800px;
  margin-left:50px ;
}
.mb{
  width: 120px;
  height: 40px;
  margin-left: 40px;
  cursor: pointer;
  transition:0.4s;
  border: none;
  border-radius: 10px;
}
.mb:hover{
  background-color: rgba(101, 250, 250, 0.53);
  color: #0c0c0c;
  font-size: 15px;

}
.line-create {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  margin-top: 10px;
  background-color: #f4f4f4;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  overflow: auto; /* Ensure content doesn't overflow the container */
}

.line-create h2 {
  text-align: center;
  margin-bottom: 20px;
}

.form-content {
  display: flex;
  flex-direction: column;
  margin-bottom: 15px; /* Adjust the margin as needed */
}

.line-create label {
  font-weight: bold;
  margin-bottom: 8px; /* Adjust the margin as needed for better spacing */
}

.line-create input {
  padding: 10px;
  margin-bottom: 15px; /* Keep the same or adjust as needed */
  border: 1px solid #ccc;
  border-radius: 4px;
}

.line-create button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.line-create button:hover {
  background-color: #0056b3;
}
</style>