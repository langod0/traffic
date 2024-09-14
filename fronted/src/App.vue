<template>
<div id="app">
    <headbar></headbar>
    <div class="maincontent">
        <sidebar></sidebar>
        
    </div>
    
    
    <section class="app-main">
      <router-view></router-view>
    </section>
  <el-footer >
    1
  </el-footer>
  </div>

</template>


<script setup>
import sidebar from './components/SideBar.vue'
import headbar from './components/HeadBar.vue'




</script>



<style scoped>
@import "./assets/css/app.css";
/* maincontent */
.maincontent{
  position: absolute;
  left:0;
  width:100%;
  height: 80%;
  background-color: #fff;
  z-index: 0;
}

/* 内容框 */
.app-main{
  position:absolute;
  top:200px;
  left:500px;
}
body{
  /* 溢出隐藏 */
  overflow: hidden;
}

#app{
  height: 100%;
  width:100%;
}


</style>