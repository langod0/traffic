<template>

  <div class="line-menu">
    <button @click="getline" class="mb">查看并选择司机</button>


  </div>
  <div class="mainline">

		<el-table :data="lines" height="100%" style="width: 100%" >
      <el-table-column prop="line_id" label="Line_id" width="200" ></el-table-column>
    <el-table-column prop="name" label="Name" width="200"   />
      <el-table-column >
        <template #default="scope">
        <el-button  @click="updata1(scope.row)">更改信息</el-button>
        </template>
      </el-table-column>
      <el-table-column >
        <template #default="scope">
        <el-button  @click="delete1(scope.row)">删除信息</el-button>
        </template>
      </el-table-column>


  </el-table>

  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import router from "@/router/index.js";

  const staff=ref("")

    const getline=()=>{
      axios.get("goapi/api/getusers",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {

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

.mainline{
width: 96%;
  height: 650px;
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

  width: 100%;
  height: 60px;
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