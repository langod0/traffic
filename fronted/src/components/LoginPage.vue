<template>
    <div class="backhome">
        <a href="/">首页</a>
    </div>
    <div class="shell">
        <div
            :class="['container', activeContainer === 'a' ? 'a-container' : 'b-container', activeContainer === 'a' ? 'is-z' : '']">
            <form @submit.prevent="submitForm()"  class="form">
                <h2 class="form_title title">{{ activeContainer === 'a' ? '司机登入账号' : '管理员登入账号' }}</h2>
                <input type="text" class="form_input" placeholder="邮箱/工号" v-model="email" />
                <input type="password" class="form_input" placeholder="密码" v-model="password" />
                <div class="rout">
                    <!-- <a href="#" class="form_link"><button @click="forgotpassword()">忘记密码</button></a> -->
                    <router-link to="/forget" class="rou">忘记密码</router-link>
                     <a href="/register" class="rou" style="text-decoration:none;" @click="toregister" >注册账号</a>
<!--                    <router-link to="/register" class="rou">前往注册</router-link>-->
                </div>

                <button class="form_button button submit"  @click="LoginFunc">登录</button>
            </form>
        </div>
        <div class="switch" @click="switchForm">
            <div class="switch_circle" :class="{ 'is-txr': activeContainer === 'b' }"></div>
            <div class="switch_circle switch_circle-t" :class="{ 'is-txr': activeContainer === 'b' }"></div>
            <div class="switch_container" :class="{ 'is-hidden': activeContainer === 'a' }">
                <h2 class="switch_title title" style="letter-spacing: 0;">I AM AN ADMINISTRATOR!</h2>
                <p class="switch_description description"></p>
                <button class="switch_button button">我是司机</button>
            </div>

            <div class="switch_container" :class="{ 'is-hidden': activeContainer === 'b' }">
                <h2 class="switch_title title" style="letter-spacing: 0;">I AM A DRIVER!</h2>
                <p class="switch_description description"></p>
                <button class="switch_button button">我是管理员</button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import axios, {all} from 'axios';

const activeContainer = ref('a');
const email = ref('');
const password = ref('');
const dt=ref("")

const LoginFunc=()=>{
  if (activeContainer.value==='a'){
    axios.post("goapi/api/login",{"staff_id":email.value,"password":password.value,"usertype":activeContainer.value})
        .then((response) =>{
            console.log(response.data)
            if(response.data.code==0){
              alert(response.data.message)
            }else{

              localStorage.setItem("Authorization", response.data.token);
              router.push('/main2').then(()=>{
              window.location.reload();
              });
            }
        }
        )
  }else{
    axios.post("goapi/api/login",{"staff_id":email.value,"password":password.value,"usertype":activeContainer.value})
        .then((response) =>
        {
          console.log(1)
            if(response.data.code==0){
              alert(response.data.message)
            }else{
              console.log(1)
              localStorage.setItem("Authorization", response.data.token);
              router.push('/main').then(()=>{
              window.location.reload();
              });
            }
        }
        )
    // axios.get("goapi/api/getinfo",{
    //   headers: {
    //     'Authorization': localStorage.getItem("Authorization")
    //   }
    // })
    //     .then((response) => {
    //         if(response.data.code==0){
    //           alert(response.data.message)
    //         }else{
    //           router.push('/main').then(()=>{
    //           window.location.reload();
    //           });
    //         }
    //     }
    //     )
  }
}
const switchForm = () => {
    activeContainer.value = activeContainer.value === 'a' ? 'b' : 'a';
};


import { useRouter } from 'vue-router';

const router = useRouter();


// function submitForm()  {
//
//     console.log('Email:', email.value);
//     console.log('Password:', password.value);
//     //Add your form submission logic here
//
//     //test code////////////////////////////////////////////////
//     if (activeContainer.value === 'b'){
//         router.push('/main').then(()=>{
//           window.location.reload();
//       });
//     }
//     else{
//         router.push('/main2').then(()=>{
//           window.location.reload();
//         });
//     }
//
// };


function forgotpassword() {
    router.push('/forgot-password').then(()=>{
    window.location.reload();
  });
}

function toregister() {
    // router.push('/register').then(()=>{
    //   window.location.reload();
    // });
}
</script>

<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    /* 字体无法选中 */
    user-select: none;
}
.backhome{
    position:absolute;
    top:0px;
    left:0px;
    z-index: 100;
}
.backhome a{
    font-size: 30px;
    font-family: 楷体;
    text-decoration: none;
    
}
body {
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 12px;
    background-color: #ecf0f3;
    color: #a0a5a8;
}

.shell {
    position: relative;
    width: 1000px;
    min-width: 1000px;
    min-height: 600px;
    height: 600px;
    padding: 25px;
    background-color: #ecf0f3;
    box-shadow: 10px 10px 10px #d1d9e6, -10px -10px 10px #f9f9f9;
    border-radius: 12px;
    overflow: hidden;
}

/* 设置响应式 */
@media (max-width: 1200px) {
    .shell {
        transform: scale(0.7);
    }
}

@media (max-width: 1000px) {
    .shell {
        transform: scale(0.6);
    }
}

@media (max-width: 800px) {
    .shell {
        transform: scale(0.5);
    }
}

@media (max-width: 600px) {
    .shell {
        transform: scale(0.4);
    }
}

.container {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    width: 600px;
    height: 100%;
    padding: 25px;
    background-color: #ecf0f3;
    transition: 1.25s;
}

.form {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    width: 100%;
    height: 100%;
}




.form_input {
    width: 350px;
    height: 40px;
    margin: 4px 0;
    padding-left: 25px;
    font-size: 13px;
    letter-spacing: 0.15px;
    border: none;
    outline: none;
    background-color: #ecf0f3;
    transition: 0.25s ease;
    border-radius: 8px;
    box-shadow: inset 2px 2px 4px #d1d9e6, inset -2px -2px 4px #f9f9f9;
}

.form_input:focus {
    box-shadow: inset 4px 4px 4px #d1d9e6, inset -4px -4px 4px #f9f9f9;
}

.form_span {
    margin-top: 30px;
    margin-bottom: 12px;
}

/* .form_link {
    color: #181818;
    font-size: 15px;
    margin-top: 25px;
    border-bottom: 1px solid #a0a5a8;
    line-height: 2;
    cursor: pointer;
} */

.title {
    font-size: 34px;
    font-weight: 700;
    line-height: 3;
    color: #181818;
    letter-spacing: 10px;
}

.description {
    font-size: 14px;
    letter-spacing: 0.25px;
    text-align: center;
    line-height: 1.6;
}

.button {
    width: 180px;
    height: 50px;
    border-radius: 25px;
    margin-top: 50px;
    font-weight: 700;
    font-size: 14px;
    letter-spacing: 1.15px;
    background-color: #4B70E2;
    color: #f9f9f9;
    box-shadow: 8px 8px 16px #d1d9e6, -8px -8px 16px #f9f9f9;
    border: none;
    outline: none;
}

.a-container {
    z-index: 100;
    left: calc(100% - 600px);
}

.b-container {
    left: calc(100% - 600px);
    z-index: 0;
}

.switch {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 400px;
    padding: 50px;
    z-index: 200;
    transition: 1.25s;
    background-color: #ecf0f3;
    overflow: hidden;
    box-shadow: 4px 4px 10px #d1d9e6, -4px -4px 10px #d1d9e6;
}

.switch_circle {
    position: absolute;
    width: 500px;
    height: 500px;
    border-radius: 50%;
    background-color: #ecf0f3;
    box-shadow: inset 8px 8px 12px #b8bec7, inset -8px -8px 12px #fff;
    bottom: -60%;
    left: -60%;
    transition: 1.25s;
}

.switch_circle-t {
    top: -30%;
    left: 60%;
    width: 300px;
    height: 300px;
}

.switch_container {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    position: absolute;
    width: 400px;
    padding: 50px 55px;
    transition: 1.25s;
}

.switch_button {
    cursor: pointer;
}

.switch_button:hover,
.submit:hover {
    box-shadow: 6px 6px 10px #d1d9e6, -6px -6px 10px #f9f9f9;
    transform: scale(0.985);
    transition: 0.25s;
}

.switch_button:active,
.switch_button:focus {
    box-shadow: 2px 2px 6px #d1d9e6, -2px -2px 6px #f9f9f9;
    transform: scale(0.97);
    transition: 0.25s;
}

.is-txr {
    left: calc(100% - 400px);
    transition: 1.25s;
    transform-origin: left;
}

.is-txl {
    left: 0;
    transition: 1.25s;
    transform-origin: right;
}

.is-z {
    z-index: 200;
    transition: 1.25s;
}

.is-hidden {
    visibility: hidden;
    opacity: 0;
    position: absolute;
    transition: 1.25s;
}

.is-gx {
    animation: is-gx 1.25s;
}

@keyframes is-gx {

    0%,
    10%,
    100% {
        width: 400px;
    }

    30%,
    50% {
        width: 500px;
    }
}

.rout {
    display: flex;
    padding: 5px;
}

.rou {
    margin: 5px;
    text-decoration: none;
    font-family: '微软雅黑';
}
</style>