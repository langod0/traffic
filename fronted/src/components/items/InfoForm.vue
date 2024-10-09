<template>
<div class="man">
  <div class="line-menu">
    <label style="margin-left: 20px">已添加的司机：</label>
    <button @click="getuse" class="mb">+</button>
    <button @click="getuse" class="mb">-</button>
    <label style="margin-left: 20px">司机数：{{nowdr.length}}</label>
<!--<button @click="getuse" class="mb">+</button>-->
  </div>

<el-dialog v-model="f1"  width="600px"   style="margin-top: 200px">
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
  <label style="margin-left: 20px">列车数：{{nowcar.length}}</label>
</div>

<el-dialog v-model="f2" style="margin-top: 200px">
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
    <div class="nr">
      <el-date-picker
          v-model="duration"
          type="datetimerange"
          @change="debug"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
          format="YYYY-MM-DD"
          date-format="YYYY/MM/DD"
      />
      <br/>
    排班名称：<input type="text" v-model="schname"  placeholder="名称" class="schin"/><br>
    选择班型：<select v-model="val" @change="hdv($event)" class="schin">
      <option value="1">四班二倒</option>
    <option value="3">四班三倒</option>
    <option value="2">三班二倒</option>
  </select><br>
<button class="btf" @click="sch" >
    排班
  </button>
</div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import  {formatDateTime} from '@/assets/js/funcs';
import router from "@/router/index.js";
// import {constrast} from "./themes";
import Child from "@/global";
const schtable=ref("")

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
const st=ref(''),ed = ref('')
const bx=ref("")
const val=ref("")
const schname=ref("")
const nowdrid=ref([])
const nowcarid=ref([])
const duration = ref('')
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

const debug = ()=>{
}


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
  nowdrid.value=[]
  f1.value=false;
  for(var i=0;i<values.value.length;i++){
    console.log(values.value[i])
    nowdr.value.push(users.value[values.value[i]])
    nowdrid.value.push(values.value[i]);
  }
}

const hdv=(e)=>{
  val.value=e.target.value;
console.log(val.value)
}
const caradd=()=>{
  nowcar.value=[]
  nowcarid.value=[]
  f2.value=false;
  for(var i=0;i<values1.value.length;i++){
    console.log(values1.value[i])
    nowcar.value.push(train.value[values1.value[i]])
    nowcarid.value.push(values1.value[i])

  }

}
const sch=()=>{

  console.log(nowcar.value)
  st.value = formatDateTime(duration.value[0])
  ed.value = formatDateTime(duration.value[1])
  const postdr=ref([])
  const postcar=ref([])
  for(var i=0;i<nowdr.value.length;i++){
    postdr.value.push(String(nowdr.value[i]["staff_id"]));
  }
  for(var i=0;i<nowcar.value.length;i++){
    postcar.value.push(String(nowcar.value[i]["id"]));
  }
  console.log(postdr.value)
  console.log(postcar.value)
  console.log(val.value)
  console.log(st.value,ed.value)
  axios.post("goapi/api/calcschedule",{"name":schname.value,"startTime":st.value,"endTime":ed.value,"type":String(val.value),"drivers":postdr.value,"trains":postcar.value},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              Child.schtb.value=response.data.schedule;
              alert("创建成功！请前往查看！")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  console.log(Child.schtb.value)
}

const gettrain=()=>{
  f2.value=true
  values1.value=nowcarid.value
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
      values.value=nowdrid.value
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
.schin{
  height: 30px;
  margin-top: 10px;
  border-radius: 8px;
  border-color: #9acfea;
}

.nr{
  width: 90%;
  height: 100%;
  margin-top: 20px;
}
.dlog{
  width: 600px;
  height: 600px;
}
.mbt{
  width: 70px;
  height:25px;
  background-color: #cbe3ef;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-size: 17px;
  transition: 0.5S;

}
.mbt:hover{
  background-color: #63b4dd;
}
.tb{
    height:300px;
  width: 90%;
  background-color: white;

}
.mesg{
  background-color: #ffffff;
  width:100%;
  height: 370px;
  float: left;
  text-align: center;

}
.btv{
  height: 50px;
  width: 100%;
  //text-align: center;
  float: left;
}
.btf{
  height: 40px;
  width:130px;
  font-size: 20px;
  border-radius: 15px;
background-color: #c1d7ea;
  border: none;
  transition: 0.7s;
  cursor: pointer;
  margin-top: 10px;
}
.btf:hover{
  background-color: #238ded;

}
.man{
  width:46.7%;
  height: 367px;
border: 2px solid #c0c0cf;
  background-color: white;
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
  width: 100%;
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