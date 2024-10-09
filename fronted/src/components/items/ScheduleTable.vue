<template>
  <section style="display: block">

  <div class="sel">
  请选择工作表：
    <select class="selo" @change="selectValue($event.target.value)">
      <option v-for="item in tim" :value="item" :key="item">
            {{ item }}
          </option>
    </select>
    <button @click="look" class="btl">查看</button>
    <button @click="" class="btl">导出</button>
  </div>
  <br>
  <div class="schedule-table">
    <h2>{{titl}}排班表</h2>
    <el-table
      :data="tableData"
    >
      <el-table-column  prop="date" label="日期/班次"></el-table-column>
      <el-table-column  label="A班">
        <template #default="scope">
          <span>{{ scope.row.classA }}</span>
          <el-icon v-if="scope.row.classA === '白班'">
            <sunny />
          </el-icon>
          <el-icon v-else-if="scope.row.classA === '夜班'">
            <moon />
          </el-icon>
          <el-icon v-else-if="scope.row.classA === '休息'">
            <timer />
          </el-icon>
        </template>
      </el-table-column>
      <el-table-column  label="B班">
        <template #default="scope">
          <span>{{ scope.row.classB }}</span>
          <el-icon v-if="scope.row.classB === '白班'">
            <sunny />
          </el-icon>
          <el-icon v-else-if="scope.row.classB === '夜班'">
            <moon />
          </el-icon>
          <el-icon v-else-if="scope.row.classB === '休息'">
            <timer />
          </el-icon>
        </template>
      </el-table-column>
      <el-table-column  label="C班">
        <template #default="scope">
          <span>{{ scope.row.classC }}</span>
          <el-icon v-if="scope.row.classC === '白班'">
            <sunny />
          </el-icon>
          <el-icon v-else-if="scope.row.classC === '夜班'">
            <moon />
          </el-icon>
          <el-icon v-else-if="scope.row.classC === '休息'">
            <timer />
          </el-icon>
        </template>
      </el-table-column>
      <el-table-column label="D班">
        <template #default="scope">
          <span>{{ scope.row.classD }}</span>
          <el-icon v-if="scope.row.classD === '白班'">
            <sunny />
          </el-icon>
          <el-icon v-else-if="scope.row.classD === '夜班'">
            <moon />
          </el-icon>
          <el-icon v-else-if="scope.row.classD === '休息'">
            <timer />
          </el-icon>
        </template>
      </el-table-column>
    </el-table>
  </div>
  </section>
</template>

<script setup>
import {onMounted, ref} from 'vue'


import Child from "@/global";
const id=ref([{}])
const tim=ref([])
const tab=ref([])
const schtab=ref('')
const nowid=ref("")
const dat=ref("")
const titl=ref("")
const tmp = [{"date":"2024-09-24","allot":[{"shift":"白班","class":"A"},{"shift":"夜班","class":"B"},{"shift":"休息","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-09-25","allot":[{"shift":"夜班","class":"A"},{"shift":"休息","class":"B"},{"shift":"休息","class":"C"},{"shift":"白班","class":"D"}]},{"date":"2024-09-26","allot":[{"shift":"休息","class":"A"},{"shift":"休息","class":"B"},{"shift":"白班","class":"C"},{"shift":"夜班","class":"D"}]},{"date":"2024-09-27","allot":[{"shift":"休息","class":"A"},{"shift":"白班","class":"B"},{"shift":"夜班","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-09-28","allot":[{"shift":"白班","class":"A"},{"shift":"夜班","class":"B"},{"shift":"休息","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-09-29","allot":[{"shift":"夜班","class":"A"},{"shift":"休息","class":"B"},{"shift":"休息","class":"C"},{"shift":"白班","class":"D"}]},{"date":"2024-09-30","allot":[{"shift":"休息","class":"A"},{"shift":"休息","class":"B"},{"shift":"白班","class":"C"},{"shift":"夜班","class":"D"}]},{"date":"2024-10-01","allot":[{"shift":"休息","class":"A"},{"shift":"白班","class":"B"},{"shift":"夜班","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-10-02","allot":[{"shift":"白班","class":"A"},{"shift":"夜班","class":"B"},{"shift":"休息","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-10-03","allot":[{"shift":"夜班","class":"A"},{"shift":"休息","class":"B"},{"shift":"休息","class":"C"},{"shift":"白班","class":"D"}]},{"date":"2024-10-04","allot":[{"shift":"休息","class":"A"},{"shift":"休息","class":"B"},{"shift":"白班","class":"C"},{"shift":"夜班","class":"D"}]},{"date":"2024-10-05","allot":[{"shift":"休息","class":"A"},{"shift":"白班","class":"B"},{"shift":"夜班","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-10-06","allot":[{"shift":"白班","class":"A"},{"shift":"夜班","class":"B"},{"shift":"休息","class":"C"},{"shift":"休息","class":"D"}]},{"date":"2024-10-07","allot":[{"shift":"夜班","class":"A"},{"shift":"休息","class":"B"},{"shift":"休息","class":"C"},{"shift":"白班","class":"D"}]},{"date":"2024-10-08","allot":[{"shift":"休息","class":"A"},{"shift":"休息","class":"B"},{"shift":"白班","class":"C"},{"shift":"夜班","class":"D"}]},{"date":"2024-10-09","allot":[{"shift":"休息","class":"A"},{"shift":"白班","class":"B"},{"shift":"夜班","class":"C"},{"shift":"休息","class":"D"}]}]
const data = ref({})
const tableData = ref(
    ''
);
onMounted(()=>{
    schtab.value=tmp
  tim.value.push("无")
    tim.value.push("123")
    tab.value.push(schtab.value)
    id.value[0]=0;
    // tableData.value = data.value.map(item => {
    //   return {
    //     date: item.date,
    //     classA: item.allot.find(allot => allot.class === 'A')?.shift || '',
    //     classB: item.allot.find(allot => allot.class === 'B')?.shift || '',
    //     classC: item.allot.find(allot => allot.class === 'C')?.shift || '',
    //     classD: item.allot.find(allot => allot.class === 'D')?.shift || '',
    //   };
    // })
})


const look=()=>{
  // nowid.value=dat.value[0]
  // console.log(nowid.value)
  // titl.value=tim.value[nowid.value]
  data.value = tmp
  tableData.value = data.value.map(item => {
    return {
      date: item.date,
      classA: item.allot.find(allot => allot.class === 'A')?.shift || '',
      classB: item.allot.find(allot => allot.class === 'B')?.shift || '',
      classC: item.allot.find(allot => allot.class === 'C')?.shift || '',
      classD: item.allot.find(allot => allot.class === 'D')?.shift || '',
    };
  })
}
const selectValue=(e)=>{
  // dat.value=[id.value[e]]

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