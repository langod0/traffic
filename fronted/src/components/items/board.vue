
<template>
  <section class="mainline" >

    <h1>欢迎回来</h1>
    <br>
    <section style="width: 1000px">
      <el-descriptions
          title=""
          direction="vertical"
          border
          style="margin-top: 20px"
      >
        <el-descriptions-item
            :rowspan="2"
            :width="140"
            label="Photo"
            align="center"
        >
          <el-image
              style="width: 100px; height: 100px"
              src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          />
        </el-descriptions-item>
        <el-descriptions-item label="Username">{{ user.name }}</el-descriptions-item>
        <el-descriptions-item label="Telephone">{{ user["phone"] }}</el-descriptions-item>
        <el-descriptions-item label="Staff ID">{{user.staff_id}}</el-descriptions-item>
        <el-descriptions-item label="Email">
          {{user.email}}
        </el-descriptions-item>
        <el-descriptions-item label="">
        </el-descriptions-item>
      </el-descriptions>
    </section>
    <br>
    <section class="statics">
      <el-row :gutter="16" style="height: calc(100% - 80)">
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic :value="outputLines" :value-style="{ color: '#FFFFFF' }">
              <template #title>
                <div style="display: inline-flex; align-items: center;font-size: 20px;color: #FFFFFF;">
                  已有
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>投入使用</span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic :value="outputStations" :value-style="{ color: '#FFFFFF' }">
              <template #title>
                <div style="display: inline-flex; align-items: center;font-size: 20px;color: #FFFFFF;">
                  总计
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>开放使用</span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic :value="outputTrains" :value-style="{ color: '#FFFFFF' }" >
              <template #title>
                <div style="display: inline-flex; align-items: center;font-size: 20px;color: #FFFFFF;">
                  已有
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span>投入使用</span>
              </div>
            </div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="statistic-card">
            <el-statistic :value="outputAva" :value-style="{ color: '#FFFFFF' }">
              <template #title>
                <div style="display: inline-flex; align-items: center ;font-size: 20px;color: #FFFFFF;">
                  现有
                </div>
              </template>
            </el-statistic>
            <div class="statistic-footer">
              <div class="footer-item">
                <span class="green">空闲</span>
                /
                <span class="red">
                  维护
            </span>
              </div>
              <div class="footer-item">
                <el-icon :size="14">
                  <ArrowRight />
                </el-icon>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
    </section>
  </section>
</template>
<script setup>
import {nextTick, onMounted, ref} from "vue";
import axios from "axios";
import {useTransitionWithUnit} from "@/assets/js/funcs.js";
const lines=ref([])
const stations=ref("")
const  trains=ref("")
const user=ref("")
const avaliable=ref(0)
const numline= ref(0)
const numSt = ref(0)
const numTr = ref(0)
const outputLines =  useTransitionWithUnit(numline,' 条线路',{duration:1500,transition: [0.75, 0.25, 0.25, 1],})
const outputStations = useTransitionWithUnit(numSt,' 个站点',{duration:1500,transition: [0.75, 0.25, 0.25, 1],})
const outputTrains = useTransitionWithUnit(numTr,' 辆列车',{duration:1500,transition: [0.75, 0.25, 0.25, 1],})
const outputAva = useTransitionWithUnit(avaliable,' 辆列车',{duration:1500,transition: [0.75, 0.25, 0.25, 1],})
const refresh = ()=>{
  axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
      .then((response)=>{
        if(response.data.code==1) {
          console.log(response.data.lines)
          lines.value = response.data.lines
          stations.value = response.data.stations
          trains.value = response.data.trains
          user.value = response.data.user
          numSt.value = stations.value.length
          numTr.value = trains.value.length
          numline.value = lines.value.length
          console.log(user,response.data.user)
          // outputLines = useTransitionWithUnit(lines.value.length,'  条线路', { duration: 35000,})
          for(let i = 0;i<trains.value.length;i++){
            if (trains.value[i].line_id==0){
              avaliable.value++
            }
          }
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
  nextTick(()=>{

  })
})
</script>

<style scoped>
@import "../../assets/style/div.css";

.statics{
  background-color: #474e6e;
  height: 240px;
  width:1000px;
}

:global(h2#card-usage ~ .example .example-showcase) {
  background-color: var(--el-fill-color) !important;
}

.el-statistic {
  --el-statistic-content-font-size: 28px;
}

.statistic-card {
  height: calc(100% - 20px);
  padding: 20px;
  margin-top: 20px;
  margin-left: 10px;
  margin-right: 10px;
  border-radius: 4px;
  background-color: #232942;
}

.statistic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  font-size: 20px;
  color: #FFFFFF;
  margin-top: 16px;
}

.statistic-footer .footer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.statistic-footer .footer-item span:last-child {
  display: inline-flex;
  align-items: center;
  margin-left: 4px;
}

.green {
  color: var(--el-color-success);
}
.red {
  color: var(--el-color-error);
}
</style>