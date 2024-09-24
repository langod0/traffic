<template>
<div class="man">
  <div class="line-menu">
    <label style="margin-left: 20px">已添加的司机：</label>
    <button @click="getuse" class="mb">+</button>
    <button @click="getuse" class="mb">-</button>
<!--<button @click="getuse" class="mb">+</button>-->
  </div>

<el-dialog v-model="f1"  width="600px" height="600px">
<div  class="tran" v-if="f1">
<el-transfer
    v-model="values"
    filterable
    :filter-method="filterMethod"
    filter-placeholder="State Abbreviations"
    :data="data"
    :titles="['可选司机', '已选司机']"
     class="cusr"
  />
  <button class="mbt" @click="dradd">提交</button>
  </div>
  </el-dialog>
  <div class="tb" >
  <el-table :data="nowdr" height="100%" style="width: 100%"  >
      <el-table-column prop="staff_id" label="staff" width="150" ></el-table-column>
    <el-table-column prop="name" label="Name" width="150"   />
  </el-table>
    </div>
  </div>

  <div class="man">
<div class="line-menu">
  <label style="margin-left: 20px">已添加的列车：</label>
  <button @click="gettrain" class="mb">+</button>
    <button @click="gettrain" class="mb">-</button>
</div>

<el-dialog v-model="f2" class="dlog">
<div  class="tran" v-if="f2">
<el-transfer
    v-model="values1"
    filterable
    :filter-method="filterMethod1"
    :data="data1"
    :titles="['可选列车', '已选列车']"
     class="cusr"
  />
  <button class="mbt" @click="caradd">提交</button>
  </div>
</el-dialog>
    <div class="tb" >
  <el-table :data="nowcar" height="100%" style="width: 100%"  >
      <el-table-column prop="id" label="Id" width="150" ></el-table-column>
    <el-table-column prop="line_id" label="Line_id" width="150"   />
  </el-table>
    </div>
  </div>
  <div  class="mesg" >
      开始时间：<input type="text" v-model="styear" placeholder="年份"/>年
    <input type="text" v-model="stmo"  placeholder="月份"/>月
    <input type="text" v-model="stday" placeholder="日期"/>日<br>
  </div>
  <div class="btv">
  <button class="btf" @click="sch">
    排班
  </button>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import router from "@/router/index.js";
// import {constrast} from "./themes";
  const users=ref("")
const trainid=ref([])
const staff=ref([])
const name=ref([])
const values = ref([]);
const train=ref("")
const values1 = ref([]);
const f1=ref(false);
const f2=ref(false);
const nowdr=ref([])
  const nowcar=ref([])
interface Option {
  key: number;
  label: string;
  initial: string;
}
interface Option1 {
  key: number;
  label: string;
  initial: string;
}
const data = ref<Option[]>([]);
const data1 = ref<Option1[]>([]);




// 生成 Transfer 所需的数据
const generateData = () => {
  const result: Option[] = [];
  name.value.forEach((city, index) => {
    result.push({
      label: city,
      key: index,
      initial: staff.value[index],
    });
  });
  data.value = result;  // 更新 data
};


const generateData1 = () => {
  const result: Option1[] = [];
  trainid.value.forEach((city, index) => {
    result.push({
      label: city,
      key: index,
      initial: "",
    });
  });
  data1.value = result;  // 更新 data
};
const filterMethod = (query, item) => {
  return item.initial.toLowerCase().includes(query.toLowerCase());
};
const filterMethod1 = (query, item) => {
  return item.initial.toLowerCase().includes(query.toLowerCase());
};



const dradd=()=>{
  nowdr.value=[]
  f1.value=false;
  for(var i=0;i<values.value.length;i++){
    console.log(values.value[i])
    nowdr.value.push(users.value[values.value[i]])

  }

}
const caradd=()=>{
  nowcar.value=[]
  f2.value=false;
  for(var i=0;i<values1.value.length;i++){
    console.log(values1.value[i])
    nowcar.value.push(train.value[values1.value[i]])
  }

}
const sch=()=>{


  console.log(nowcar.value)
  // axios.post("goapi/api/",{},{headers:{'Authorization': localStorage.getItem("Authorization")}})
  //         .then((response)=>{
  //           if(response.data.code==1) {
  //
  //           }else{
  //             alert(response.data.error)
  //           }
  //         })
  //         .catch((error)=>{
  //           console.log(error)
  //         })
}

const gettrain=()=>{
  f2.value=true
      trainid.value=[];
      axios.get("goapi/api/getinfo",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                train.value=response.data.trains;
                for(var i=0;i<train.value.length;i++){
                  trainid.value.push(train.value[i]["id"])
                }
                console.log(trainid.value)
                generateData1()



            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    }
    const getuse=()=>{
      f1.value=true;
      staff.value=[];
      name.value=[];
      axios.get("goapi/api/getusers",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                users.value=response.data.drivers;
                for(var i=0;i<users.value.length;i++){

                  staff.value.push(users.value[i]["staff_id"]);
                  name.value.push(users.value[i]["name"]);
                }
                console.log(staff.value)
                generateData()



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
.dlog{
  width: 600px;
  height: 600px;
}
.mbt{
  width: 100px;
  height:50px;
  background-color: #9acfea;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 20px;

}
.mbt:hover{
  background-color: #63b4dd;
}
.tb{
    height:300px;
  width: 45%;
}
.mesg{
  width:100%;
  height: 100px;
  float: left;
  border: 3px solid #c1d7ea;
}
.btv{
  height: 50px;
  width: 100%;
  text-align: center;
  float: left;
}
.btf{
  height: 50px;
  width:130px;
  font-size: 20px;
  border-radius: 15px;
background-color: #c1d7ea;
  border: none;
  transition: 0.7s;
  cursor: pointer;
}
.btf:hover{
  background-color: #238ded;

}
.man{
  width:45%;
  height: 367px;
margin-left: 5%;
  float: left;
}
.tran{
  width:100%;
  height: 330px;
  text-align: center;
}
.cusr{

}
.mainline{
width: 96%;
  height: 650px;
}
.mb{
  width: 30px;
  height: 30px;
  margin-left: 30px;
  cursor: pointer;
  transition:0.4s;
  border: none;
  font-size: 25px;

}
.mb:hover{
  background-color: #ccc;
  color: #fff;
  font-size: 15px;

}
.line-menu{
  margin-top: 5px;
  width: 45%;
  height: 37px;

  background-color: white;

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