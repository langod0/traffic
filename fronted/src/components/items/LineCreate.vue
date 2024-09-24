<template>

  <div class="line-menu">
    <button @click="getline" class="mb">查看线路信息</button>
    <button @click="getsta"  class="mb">查看站点信息</button>
    <button @click="gettrains" class="mb">查看列车信息</button>
    <button @click="createLine" class="mb">创建线路</button>
    <button @click="createstation" class="mb">创建站点</button>
    <button @click="createtrain" class="mb">创建列车</button>

  </div>
  <div class="mainline">

		<el-table :data="lines" height="100%" style="width: 100%" v-if="f1" >
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

<!--    <el-table-column prop="address" label="Address" />-->
  </el-table>
    <el-table :data="stations" height="100%" style="width: 100%" v-if="f2" >
    <el-table-column prop="id" label="id" width="200"  />
    <el-table-column prop="lon" label="longitude" width="200"   />
      <el-table-column prop="lat" label="latitude" width="200"   />
      <el-table-column prop="name" label="Name" width="200"   />
      <el-table-column >
        <template #default="scope">
        <el-button  @click="updata2(scope.row)">更改信息</el-button>
        </template>
      </el-table-column>c
      <el-table-column >
        <template #default="scope">
        <el-button  @click="delete2(scope.row)">删除信息</el-button>
        </template>
      </el-table-column>

<!--    <el-table-column prop="address" label="Address" />-->
  </el-table>

    <el-table :data="trains" height="100%" style="width: 100%" v-if="f4" >
    <el-table-column prop="id" label="id" width="200"  />
    <el-table-column prop="line_id" label="line_id" width="200"   />
      <el-table-column prop="capacity" label="capacity" width="200"   />
      <el-table-column >
        <template #default="scope">
        <el-button  @click="updata3(scope.row)">更改信息</el-button>
        </template>
      </el-table-column>
      <el-table-column >
        <template #default="scope">
        <el-button  @click="delete3(scope.row)">删除信息</el-button>
        </template>
      </el-table-column>
<!--      <el-table-column prop="name" label="Name" width="200"   />-->
<!--    <el-table-column prop="address" label="Address" />-->
  </el-table>
  <div class="line-create" v-if="f3">
    <h2>创建新线路</h2>
    <form @submit.prevent="createLine" >
      <div class="form-content">
        <label for="line-name">线路名称:</label>
        <input type="text" id="line-name" v-model="new_line_name" placeholder="线路名称" required>
      </div>
      <button type="submit" @click="upline">提交</button>
    </form>
  </div>
  <div style="width: 80%;height: 80%;margin-left: 100px;margin-top: 60px;background-color: white;float: left" v-if="f5">
      <div style="width: 400px;height: 300px;margin:0 auto;margin-top: 60px">

        <p style="text-align: center">线路信息更改</p>
        线路编号：<label   class="inp">{{line_id}}</label><br>
        线路名称：<input type="text" v-model="line_name" class="inp"><br>
        <button @click="updatan1" class="bt" >提交</button>

        </div>

    </div>
    <div style="width: 80%;height: 80%;margin-left: 100px;margin-top: 60px;background-color: white;float: left" v-if="f6">
      <div style="width: 400px;height: 300px;margin:0 auto;margin-top: 60px">

        <p style="text-align: center">站点信息更改</p>
        站点编号：<label   class="inp">{{station_id}}</label><br>
        &nbsp经&nbsp&nbsp&nbsp&nbsp度&nbsp：<input type="text" v-model="station_lon" class="inp"><br>
        &nbsp纬&nbsp&nbsp&nbsp&nbsp度&nbsp：<input type="text" v-model="station_lat" class="inp"><br>
        站点名称：<input type="text" v-model="station_name" class="inp"><br>
        <button @click="updatan2" class="bt" >提交</button>

        </div>

    </div>
    <div style="width: 80%;height: 80%;margin-left: 100px;margin-top: 60px;background-color: white;float: left" v-if="f7">
      <div style="width: 400px;height: 300px;margin:0 auto;margin-top: 60px">

        <p style="text-align: center">列车信息更改</p>
        列车编号：<label   class="inp">{{train_id}}</label><br>
        线路编号：<input type="text" v-model="train_line_id" class="inp"><br>
        capacity：<input type="text" v-model="train_capacity" class="inp"><br>

        <button @click="updatan3" class="bt" >提交</button>

        </div>

    </div>
    <div class="line-create" v-if="f8">
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
    <div class="line-create" v-if="f9">
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
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import router from "@/router/index.js";

    const f1=ref(false)
    const f2=ref(false)
    const f3=ref(false)
const f4=ref(false)
const f5=ref(false)
const f6=ref(false)
const f7=ref(false)
const f8=ref(false)
const f9=ref(false)
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
const updata1=(e)=>{
      now.value=e;
      f1.value = false
    f2.value = false
    f3.value = false
    f4.value=false
    f5.value=true
  f6.value=false
  f7.value=false
  f8.value=false
  f9.value=false
  line_id.value=now.value["line_id"]
  line_id.value=Number(line_id.value)
  line_name.value=now.value["name"]

}
const updatan1=()=>{
     axios.post("goapi/api/updateline",{"line_id":line_id.value,"name":line_name.value,"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("更新成功")
            }else{
              alert(response.data.error)
            }
          })

}
const updatan2=()=>{
   axios.post("goapi/api/updatestation",{"id":station_id.value,"name":station_name.value,"lon":station_lon.value,"lat":station_lat.value,"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                alert("更新成功")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
}
const updatan3=()=>{
    axios.post("goapi/api/updatetrain",{"id":String(train_id.value),"line_id":Number(train_line_id.value),"cap":Number(train_capacity.value),"use":0},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("更新成功")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
}
const updata2=(e)=>{
      now.value=e;
      f1.value = false
    f2.value = false
    f3.value = false
    f4.value=false
    f5.value=false
  f6.value=true
  f7.value=false
  f8.value=false
  f9.value=false
  station_id.value=now.value["id"]
  station_name.value=now.value["name"]
  station_lon.value=now.value["lon"]
  station_lat.value=now.value["lat"]
}
const updata3=(e)=>{
      now.value=e;
     f1.value = false
    f2.value = false
    f3.value = false
    f4.value=false
    f5.value=false
  f6.value=false
  f7.value=true
  f8.value=false
  f9.value=false
train_id.value=now.value["id"]
train_line_id.value=now.value["line_id"]
train_capacity.value=now.value["capacity"]
}
const delete1=(e)=>{
      now.value=e;
    axios.post("goapi/api/updateline",{"line_id":now.value["line_id"],"name":now.value["name"],"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("删除成功")
              getline()
            }else{
              alert(response.data.error)
            }
          })
}
const delete2=(e)=>{
      now.value=e;
       axios.post("goapi/api/updatestation",{"id":now.value["id"],"name":now.value["name"],"lon":now.value["lon"],"lat":now.value["lat"],"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                alert("删除成功")
              getsta()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
}
const  delete3=(e)=>{
      now.value=e
      axios.post("goapi/api/updatetrain",{"id":String(now.value["id"]),"line_id":Number(now.value["line_id"]),"cap":Number(now.value["capacity"]),"use":-1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("删除成功")
              gettrains()
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
}
  const createLine=()=> {
    // Logic to create a new line and store it in Vuex or local state
    //  console.log('创建线路:', new_line_name.value);
    f1.value = false
    f2.value = false
    f3.value = true
    f4.value=false
    f5.value=false
    f6.value=false
  f7.value=false
    f8.value=false
  f9.value=false


  }
  const createstation=()=> {
      f1.value = false
    f2.value = false
    f3.value = false
    f4.value=false
    f5.value=false
    f6.value=false
  f7.value=false
    f8.value=true
  f9.value=false
  }
  const createtrain=()=> {
          f1.value = false
    f2.value = false
    f3.value = false
    f4.value=false
    f5.value=false
    f6.value=false
  f7.value=false
    f8.value=false
  f9.value=true
  }
  const upline=()=>{
      axios.post("goapi/api/updateline",{"line_id":10000,"name":new_line_name.value,"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  }
  const upstation=()=>{
      axios.post("goapi/api/updatestation",{"id":1000000,"name":new_station_name.value,"lon":new_station_lon.value,"lat":new_station_lat.value,"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  }
  const uptrain=()=>{
      axios.post("goapi/api/updatetrain",{"id":new_train_id.value,"cap":Number(new_train_cap.value),"line_id":Number(new_train_lineid.value),"use":1},{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
              alert("添加成功")
            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
  }
 const gettrains=()=>{
      f1.value=false
      f2.value=false
      f3.value=false
      f4.value=true
   f5.value=false
   f6.value=false
  f7.value=false
   f8.value=false
  f9.value=false
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
    const getline=()=>{
      f1.value=true
      f2.value=false
      f3.value=false
      f4.value=false
      f5.value=false
      f6.value=false
  f7.value=false
      f8.value=false
  f9.value=false
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
      f4.value=false
      f5.value=false
      f6.value=false
  f7.value=false
      f8.value=false
  f9.value=false
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
.main{
  width: 100%;
  height: 650px;
}
.linese{
  text-align: right;
}
.mainlines{
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
  border-radius: 10px
}
.linese{
  text-align: right;
}
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