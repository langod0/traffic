<template>
  <el-tabs type="border-card" v-model="activate" style="height: 100%" @tab-change="handleChange"  >


    <el-tab-pane label="查看站点信息"  name="1" >
      <button @click="stationvisible=true" class="mb">提交新站点</button>
<button @click="mapshow" class="mb">站点查看</button>
      <div class="mainline">
        <el-table :data="stations" stripe height="100%" style="width:950px"  >
          <el-table-column prop="id" label="站点编号" width="150"  />
          <el-table-column prop="lon" label="经度" width="150"   />
          <el-table-column prop="lat" label="纬度" width="150"   />
          <el-table-column prop="name" label="站点名称" width="200"   />
          <el-table-column  label="线路编号" width="200"   ><template #default="scope">
             {{fdline(scope.row)}}
            </template></el-table-column>


        </el-table>
      </div>
    </el-tab-pane>
<el-tab-pane label="查看已提交的信息"  name="2" >

      <div class="mainline">
        <el-table :data="stationsub" stripe height="100%" style="width:1150px"  >
          <el-table-column prop="id" label="站点编号" width="150"  />
          <el-table-column prop="data.lon" label="经度" width="150"   />
          <el-table-column prop="data.lat" label="纬度" width="150"   />
          <el-table-column prop="data.name" label="站点名称" width="200"   />
          <el-table-column  label="线路编号" width="200"   ><template #default="scope">
             {{fdline(scope.row)}}
            </template></el-table-column>
          <el-table-column prop="userid" label="申请人工号" width="200"   />
          <el-table-column prop="done" label="状态" width="200"   />


        </el-table>
      </div>
    </el-tab-pane>

  </el-tabs>

    <el-dialog v-model="stationvisible" title="提交新站点" width="500">
      <div class="line-create" >
        <h2>提交新站点</h2>
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

    <el-dialog v-model="stationexvisible" title="编辑申请" width="500">

      <div class="line-create">
        站点编号：<label   class="inp">{{station_id}}</label><br>
        &nbsp经&nbsp&nbsp&nbsp&nbsp度&nbsp：<label   class="inp">{{station_lon}}</label><br>
        &nbsp纬&nbsp&nbsp&nbsp&nbsp度&nbsp：<label   class="inp">{{station_lat}}</label><br>
        站点名称：<label type="text"  class="inp">{{station_name}}</label><br>

        <el-button  @click="destation=true">删除申请</el-button>
      </div>

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
const stationsub=ref([])
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
const lineFilter = ref([])



const refresh = ()=>{
stline.value={}
  axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
      .then((response)=>{
        if(response.data.code==1) {
          // console.log(response.data)
          lines.value = response.data.lines
          stations.value = response.data.stations
          console.log(stations.value[0])
          stationsub.value=response.data.subs
          console.log(stationsub.value[0])
          for(var i=0;i<stationsub.value.length;i++){
              stationsub.value[i]["data"]=JSON.parse(stationsub.value[i]["data"])
              if(stationsub.value[i]["done"]==false) stationsub.value[i]["done"]="未审核"
              else  stationsub.value[i]["done"]="已审核"

          }
          for(var i=0;i<stations.value.length;i++){
            stline.value[stations.value[i]["id"]]="";

          }
          trains.value = response.data.trains
          user.value = response.data.user
          axios.get("goapi/api/getrelations",{headers:{'Authorization': localStorage.getItem("Authorization")}}).then((response)=>{
              if (response.data.code==1){
                relations.value = response.data.result
                // console.log(relations.value[0])
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
                  // console.log(stline.value[relations.value[i]["SubwayStationId"]])
                  stline.value[relations.value[i]["SubwayStationId"]]=stline.value[relations.value[i]["SubwayStationId"]]+String(relations.value[i]["SubwayLineId"])+" ";
                  if (mm.has(relations.value[i].SubwayLineId)){
                    StData.value[mm.get(relations.value[i].SubwayLineId)]++
                  }
                }
                console.log(trData.value,StData.value,mm)

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

    refresh()

})

const updata2=(e)=>{
  stationexvisible.value=true
      now.value=e;
  station_id.value=now.value["id"]
  station_name.value=now.value["name"]
  station_lon.value=now.value["lon"]
  station_lat.value=now.value["lat"]
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
  // refresh()
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
    // refresh()
  }

const mapshow=()=>{
    mapflag.value=true
    console.log(mapflag.value)
}
const fdline=(e)=>{
// console.log(stline.value[e["id"]])
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