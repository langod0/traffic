<template>

  <div class="line-menu">
    <button @click="getline" class="mb">查看线路信息</button>
    <button @click="getsta"  class="mb">查看站点信息</button>
    <button @click="createLine" class="mb">创建线路</button>
  </div>
  <div class="mainline">

		<el-table :data="lines" height="100%" style="width: 100%" v-if="f1" >
    <el-table-column prop="line_id" label="Line_id" width="200"  />
    <el-table-column prop="name" label="Name" width="200"   />
<!--    <el-table-column prop="address" label="Address" />-->
  </el-table>
    <el-table :data="stations" height="100%" style="width: 100%" v-if="f2" >
    <el-table-column prop="id" label="id" width="200"  />
    <el-table-column prop="lon" label="longitude" width="200"   />
      <el-table-column prop="lat" label="latitude" width="200"   />
      <el-table-column prop="name" label="Name" width="200"   />
<!--    <el-table-column prop="address" label="Address" />-->
  </el-table>
  <div class="line-create" v-if="f3">
    <h2>创建新线路</h2>
    <form @submit.prevent="createLine" >
      <div class="form-content">
        <label for="line-name">线路名称:</label>
        <input type="text" id="line-name" v-model="newlinename" placeholder="线路名称" required>
      </div>
      <button type="submit">提交</button>
    </form>
  </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import router from "@/router/index.js";

    const f1=ref(false)
    const f2=ref(false)
    const f3=ref(false)
    const newlinename=ref('')
const lines=ref("")
const stations=ref("")
const  trains=ref("")
const user=ref("")

  const createLine=()=> {
    // Logic to create a new line and store it in Vuex or local state
    // console.log('创建线路:', this.newLineName);
    f1.value = false
    f2.value = false
    f3.value = true
  }

    const getline=()=>{
      f1.value=true
      f2.value=false
      f3.value=false
      axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              console.log(response.data.lines)
              lines.value = response.data.lines
              stations.value = response.data.stations
              trains.value = response.data.trains
              user.value = response.data.user
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    }
    const getsta=()=>{
      f1.value=false
      f2.value=true
      f3.value=false
      axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              console.log(response.data.lines)
              lines.value = response.data.lines
              stations.value = response.data.stations
              trains.value = response.data.trains
              user.value = response.data.user
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    }



</script>

<style scoped>
.linese{
  text-align: right;
  //margin-right: 40px;
}
.mainline{
width: 100%;
  height: 620px;
  //border: 3px solid #ccc;
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
  background-color: #ccc;
  color: #fff;
  font-size: 15px;

}
.line-menu{
  margin-top: 20px;
  //border: 3px solid #ccc;
  width: 100%;
  height: 60px;
}
.line-create {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  margin-top: 100px;
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