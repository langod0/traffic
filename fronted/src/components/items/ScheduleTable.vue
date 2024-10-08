<template>
  <section></section>
  <div class="sel">
  请选择日期：
    <select class="selo" @change="selectValue($event.target.value)">
      <option v-for="item in tim" :value="item" :key="item">
            {{ item }}
          </option>
    </select>
    <button @click="look" class="btl">查看</button>
    <button @click="" class="btl">导出</button>
  </div>
  <div class="schedule-table">
    <h2>{{titl}}排班表</h2>
    <table>
      <thead>
        <tr>
          <th>班次</th>
          <th>班组号</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(schedule, index) in tab[Number(nowid)]" :key="index">
          <td>{{ schedule.shift }}</td>
          <td>{{ schedule.class }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'


import Child from "@/global";
const id=ref([{}])
const tim=ref([])
const tab=ref([])
const schtab=ref([])
const nowid=ref("")
const dat=ref("")
const titl=ref("")
onMounted(()=>{
    schtab.value=Child.schtb.value
  for(var i = 0;i<schtab.value.length;i++){
      tim.value.push(schtab.value[i]["date"])
      tab.value.push(schtab.value[i]["allot"])
      id.value[schtab.value[i]["date"]]=i;
  }
})

const look=()=>{
  nowid.value=dat.value[0]
  console.log(nowid.value)
  titl.value=tim.value[nowid.value]
}
const selectValue=(e)=>{
  dat.value=[id.value[e]]
  console.log(dat.value)
}


</script>

<style scoped>
.btl{
  margin-top: 5px;
  width: 100px;
  height: 30px;
  transition: 0.6s;
  cursor: pointer;
  border-radius: 5px;
  border: none;
  margin-left: 5px;
}
.btl:hover{
  background-color: #9acfea;
}
.schedule-table {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f4f4f4;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

}

.schedule-table h2 {
  text-align: center;
  margin-bottom: 20px;
}

.schedule-table table {
  width: 100%;
  border-collapse: collapse;
}

.schedule-table th, .schedule-table td {
  padding: 10px;
  text-align: left;
  border: 1px solid #ccc;
}

.schedule-table th {
  background-color: #007bff;
  color: white;
}

.schedule-table tr:nth-child(even) {
  background-color: #f2f2f2;
}
.sel{
  float: left;
  width: 550px;
  height: 100px;
  text-align: center;
}
.selo{
  margin-top: 10px;
  width: 120px;
  height: 30px;
}
</style>