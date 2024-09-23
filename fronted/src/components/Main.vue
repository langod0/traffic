<template>
    <div id="app">
        <div class="headbar">
            <img src="../assets/headbarimg2.png">
            <div class="mask"></div>
            <div :class="['navbar', { 'slide': isSliding }]">
                <svg 
 xmlns="http://www.w3.org/2000/svg"
 xmlns:xlink="http://www.w3.org/1999/xlink"
 width="4181px" height="113px">
<image  x="0px" y="0px" width="4181px" height="113px"  xlink:href="data:img/png;base64,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" />
                </svg>
                <div class="nav-content">
                    <div class="nav-content-left">
                        <img src="../assets/幽灵.jpg">
                        <router-link to="/contact" class="nav-content-rl">联系我们</router-link>
                    </div>
                    <div class="nav-content-center">
                        <img src="../assets/地铁系统logo.png" class="logo">
                    </div>
                    <div class="exit">
                        <a href="/"><span>退出</span></a>
                        <img src="../assets/幽灵.jpg">
                    </div>
                </div>
            </div>
        </div>
        <div class="maincontent">
            <div id="main_div">
                <div class="menu-box">
                    <div class="emotion-box">
                        <img src="../assets/笑脸.png" id="emotion1">
                        <img src="../assets/闭眼.png" id="emotion2">
                    </div>
                    <ul class="table">
                        <li>
                            <router-link to="#" id="back">
                                <img src="../assets/首页.png">
                                <p class="content">首页</p>
                            </router-link>
                        </li>
                        <li >
                            <button @click="addLines" id="router">
                                <img src="../assets/线路.png">
                                <p class="content">创建线路</p>
                            </button>
                        </li>
                        <li>
                            <button @click="addInfo" id="router">
                                <img src="../assets/填写.png">
                                <p class="content">填写信息</p>
                            </button>
                        </li>
                        <li>
                            <button @click="addTable" id="router">
                                <img src="../assets/查看.png">
                                <p class="content">查看排班</p>
                            </button>
                        </li>
                        <li>
                            <button @click="showdata" id="router">
                                <img src="../assets/ksh.png">
                                <p class="content">可视化分析</p>
                            </button>
                        </li>
                        <li>
                            <button @click="predata" id="router">
                                <img src="../assets/yc.png">
                                <p class="content">预测客流量</p>
                            </button>
                        </li>
                    </ul>
                </div>
            </div>
            <dev class="app-main">
                <div class="main-lines" v-if="islines">
                    <linesvue/>
                </div>
                <div class="main-info" v-if="isinfo">
                    <infovue/>
                    <mapvue/>
                </div>
                <div class="main-table" v-if="istable">
                    <tablevue/>
                </div>
                <div class="main-show" v-if="isshow">
                    <showvue/>
                </div>
                <div class="main-show" v-if="ispre">
                    <prevue/>
                </div>
            </dev>
        </div>
        
        
        
      </div>
    </template>
    
    
<script setup>
import { ref, onMounted } from 'vue';
import linesvue from './items/LineCreate.vue'
import infovue from './items/InfoForm.vue'
import tablevue from './items/ScheduleTable.vue'
import mapvue from './items/Map.vue'
import showvue from './items/Show.vue'
import prevue from './items/Predata.vue'
//地铁滑动
let isSliding = ref(false);
onMounted(() => {
setTimeout(() => {
    isSliding.value = true;
}, 1000); // 延迟 1 秒后开始动画
});

let islines = ref(false);
let istable = ref(false);
let isinfo = ref(false);
let isshow = ref(false);
let ispre= ref(false);
//lines组件

function addLines(){
    islines.value = true;
    isinfo.value = false;
    istable.value = false;
    isshow.value = false;
    ispre.value = false;
    console.log('1');
}
function predata() {
    islines.value = false;
    isinfo.value = false;
    istable.value = false;
    isshow.value = false;
    ispre.value = true;
    console.log('1');
}
function showdata(){
    islines.value = false;
    isinfo.value = false;
    istable.value = false;
    isshow.value = true;
    ispre.value = false;
    console.log('1');

}
//lnfo组件

function addInfo(){
    isinfo.value = true;
    islines.value = false;
    istable.value = false;
    isshow.value = false;
    ispre.value = false;
    console.log('1');
}

//table组件

function addTable(){
    istable.value = true;
    isinfo.value = false;
    islines.value = false;
    isshow.value = false;
    ispre.value = false;
    console.log('1');
}


</script>
    
    
    
<style scoped>
*{
    /* 初始化 取消内外边距*/
    margin: 0;
    padding: 0;
    list-style: none;
    text-decoration: none;
}

body{
    /* 溢出隐藏 */
    overflow: hidden;
}

#app{
    height: 100%;
    width:100%;
}



/* headbar */
.headbar {
position: absolute;
top: 0;
left: 0;
width: 100%;
height: 220px;
overflow: hidden;
}

.headbar>img {
    position: absolute;
    top: -120px;
    left: 0;
    width: 100%;
    height: 400px;
    object-fit: fill;
    object-position: center center;
    z-index: -2000;
}

.headbar .mask {
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    height: 200px;
    background-color: black;
    background: linear-gradient(to top, rgba(13, 13, 13, 0.6) 25%, transparent 100%);
    z-index: -1000;
}

.headbar .navbar {
    position: absolute;
    bottom: 0px;
    right: 100%;
    transform: translateY(-50%);
    width: 100%;
    height: 90px;
    background-color: transparent;
    transition: all 1s ease-in-out;
    z-index: 1000;
}

.headbar .navbar.slide {
    right: 0;
    /* 滑动到左侧 */
}


.headbar .navbar svg {
    position: absolute;
    right: 5px;
}



/* nav-content */
.headbar .nav-content {
    position: absolute;
    bottom: 0px;
    width: 100%;
    height: 85px;
    align-content: center;
    display: flex;
    z-index: 2000;
}


.headbar .nav-content img {
    width: 50px;
    height: 50px;
    border-radius: 50%;
}

.headbar .nav-content-left {
    position: relative;
    left: 0;
    width: 200px;
    height: 70px;
    display: flex;
    align-content: center;
    z-index: 4000;
}


.headbar .nav-content-left img {
    position: relative;
    top: 17px;
    left: 10px;
}


.headbar .nav-content-left .nav-content-rl {
    color: #fff;
    position: relative;
    left: 50px;
    top: 30px;
    font-size:20px;
}



.headbar .nav-content-center {
    position: relative;
    left: 950px;
    width: 200px;
    align-content: center;
}


.headbar .nav-content-center img {
    width:250px;
    height: 100px;
}
.exit{
    position: relative;
    margin-left:650px ;
    margin-top:15px;
    /* border: 3px solid blue; */
}
.exit a{
    color: black;
}


    
    
    
/* nav-content
.headbar .nav-content{
    position: absolute;
    bottom: 0px;
    width:100%;
    height: 85px;
    align-content: center;
    display: flex;
    z-index: 2000;
}


.headbar .nav-content img{
    width: 50px;
    height: 50px;
    border-radius: 50%;
}

.headbar .nav-content-left{
    position: relative;
    left:0;
    width:200px;
    height: 70px;
    display: flex;
    align-content: center;
    z-index:4000;
    background-color: black;
}


.headbar .nav-content-left img{
    position: relative;
    top:17px;
    left:10px;
}


.headbar .nav-content-left .nav-content-rl{
    color:#fff;
    position: relative;
    left:50px;
    top:30px;

}



.headbar .nav-content-center{
    background-color: black;
    position: relative;
    left:550px;
    width:200px;
    align-content: center;
}


.headbar .nav-content-center span{
    color: #fff;
}



.headbar .nav-content-right{
    background-color: black;
    position: relative;
    left:850px;
    width:220px;
    display: flex;
    align-content: center;
}


.headbar .nav-content-right span{
    color:#fff;
    position: relative;
    left:100px;
    top:30px;
}


.headbar .nav-content-right img{
    position: relative;
    left:130px;
    top:17px;
} */











/* maincontent */
.maincontent{
    position: absolute;
    top:200px;
    left:0;
    width:100%;
    height: 100%;
    background-color: rgba(0,0,0);
    z-index: 0;
}


/* sidebar */

#main_div {
height: 100vh;
background-color: rgb(11, 12, 29);
display: flex;
justify-content: center;
align-items: center;
}

.menu-box {
    width: 100px;
    height: 100vh;
    background-color: rgba(255, 255, 255, 0.3);
    /* border-radius: 30px; */
    box-shadow: 1px 1px 2px rgba(0, 0, 0, .3);
    /* 添加过渡 */
    transition: all .2s ease-out;
    overflow: hidden;
    margin-right: auto;
    z-index: 1000;
}

.menu-box:hover {
    width: 200px;
    /* border-radius: 60px; */
}

.menu-box:hover>ul>li p {
    display: block;
}

.emotion-box {
    width: 100%;
    text-align: center;
    margin: 10px auto;
}

#emotion1,
#emotion2 {
    width: 50px;
    height: 50px;
    animation: jitter .3s linear forwards;

}

#emotion2 {
    display: none;
    margin-left: auto;
    margin-right: auto;

}

.menu-box:hover #emotion1 {
    display: none;
}

.menu-box:hover #emotion2 {
    display: block;
}

ul {
    width: 100%;
    height: 75%;
}

li {
    list-style: none;
    margin-top: 20px;
    margin-bottom: 20px;
}

li> {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50px;
    text-decoration: none;
    user-select: none;
}

li:hover {
    background-color: #343850;
    cursor: pointer;
}

li:hover .content,
li:hover img {
    color: #fff;
}

li:hover img {
    animation: jitter .3s linear forwards;
}

.content {
    display: none;
    padding-left: 5px;
    color: #fff;
    white-space: nowrap;
    transition: all .2s ease-out;
}




.table img {
    width: 30px;
    height: 30px;
}


.table p{
    font-size:20px;
}

#router,#back {
    background-color: transparent;
    border:none;
    display: flex;
    height: 50px;
    align-items:center;
    justify-content: center;
    margin:20px auto;
}


@keyframes jitter {
    0% {
        transform: scale(1.2) rotate(10deg);
    }

    50% {
        transform: scale(0.8) rotate(-10deg);
    }

    0% {
        transform: scale(1.0) rotate(0deg);
    }
}


/* app-main */
.app-main{
    height: 740px;
    width:1300px;
    position: absolute;
    /* border-radius: 30px; */
    top:0px;
    left:200px;
    background-color: rgba(255, 255, 255, 0.5);
}

</style>