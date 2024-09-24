<template>

  <div class="line-menu">
    <button @click="getuse" class="mb">查看并选择司机</button>


  </div>
  <div class="mainline">

		<el-table :data="users" height="100%" style="width: 30%" >
      <el-table-column prop="name" label="Name" width="100px" ></el-table-column>
    <el-table-column prop="staff_id" label="Staff" width="100px"   />
      <el-table-column prop="post" label="Post" width="100px"   />

      <el-table-column :label="count" >
        <template #default="scope">
        <el-checkbox
          v-model="scope.row.check"
          @change="checkbx(scope.row)">
        </el-checkbox>
        </template>
      </el-table-column>



  </el-table>
<el-transfer
    v-model="value"
    filterable
    :filter-method="filterMethod"
    filter-placeholder="State Abbreviations"
    :data="data"
  />
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import axios, {all} from 'axios';
import router from "@/router/index.js";



interface Option {
  key: number
  label: string
  initial: string
}

const generateData = () => {
  const data: Option[] = []
  const states = [
    'California',
    'Illinois',
    'Maryland',
    'Texas',
    'Florida',
    'Colorado',
    'Connecticut ',
  ]
  const initials = ['PP', 'IL', 'MD', 'TX', 'FL', 'CO', 'CT']
  states.forEach((city, index) => {
    data.push({
      label: city,
      key: index,
      initial: initials[index],
    })
  })
  return data
}

const data = ref<Option[]>(generateData())
const value = ref([])

const filterMethod = (query, item) => {
  return item.initial.toLowerCase().includes(query.toLowerCase())
}




  const users=ref("")
 const count=ref(0)

    const getuse=()=>{
      axios.get("goapi/api/getusers",{headers:{'Authorization': localStorage.getItem("Authorization")}})
          .then((response)=>{
            if(response.data.code==1) {
                users.value=response.data.drivers;
                for(var i=0;i<users.value.length;i++){
                  users.value[i].check= false
                }


            }else{
              alert(response.data.error)
            }
          })
          .catch((error)=>{
            console.log(error)
          })
    }
    const checkbx=(row)=> {
      if(row.check==true){
        count.value=count.value+1
      }else{
        count.value=count.value-1
      }
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