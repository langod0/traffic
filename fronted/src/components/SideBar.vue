<template>
    <body>
        <div id="main_div">
            <div class="menu-box">
                <div class="emotion-box">
                    <img src="../assets/幽灵.jpg" id="emotion1">
                    <img src="../assets/幽灵.jpg" id="emotion2">
                </div>
                <el-menu
                    :router="true"
                    class="table"
                >
                    <el-menu-item index="1" route="#">
                      <div class="flex-container">
                          <img src="../assets/幽灵.jpg">
                          <span class="content">首页</span>
                      </div>
                    </el-menu-item>
                    <el-sub-menu
                    >
                      <template #title>
                        <div  class="flex-container">
                          <img src="../assets/幽灵.jpg" alt="">
                          <span class="content">公交线路</span>
                        </div>
                      </template>
                      <el-menu-item index="2" route="/lines">
                        <div  class="flex-container">
                          <img src="../assets/幽灵.jpg" alt="">
                          <span class="content">创建线路</span>
                        </div>
                      </el-menu-item>
                      <el-menu-item index="3" route="/info">
                        <div  class="flex-container">
                          <img src="../assets/幽灵.jpg">
                          <span class="content">填写信息</span>
                        </div>
                      </el-menu-item>
                      <el-menu-item index="/schedule" route="/schedule">
                        <div  class="flex-container">
                          <img src="../assets/幽灵.jpg">
                          <span class="content">查看排班</span>
                        </div>
                      </el-menu-item >
                    </el-sub-menu>


                </el-menu>
            </div>
        </div>
    </body>
</template>


<script setup>
</script>



<style>
@import "../assets/css/siderbar.css";
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    list-style: none;
}

body {
    /* 溢出隐藏 */
    overflow: visible;
}
.el-submenu__title .el-icon {
  font-size: 14px; /* 修改成你需要的大小 */
}
#main_div {
    height: 100vh;
    background-color: rgba(0, 0, 0, 0);
    display: flex;
    justify-content: center;
    align-items: center;

}
.flex-container {
  display: flex;
  align-items: center;
}

</style>