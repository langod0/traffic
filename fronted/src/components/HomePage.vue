<template>
    
    <div class="total">
        <div class="top">
            <div class="header">
                <img src="../assets/地铁系统logo.png" class="logo">
                <div class="cta">
                    <div class="but"><button @click="navigateToLogin">登录</button></div>
                    <div class="but"><button @click="navigateToRegister">注册</button></div>
                    <!-- <router-link class="btbt"><button>登录</button></router-link>
                    <router-link class="btbt"><button>注册</button></router-link> -->
                </div>
            </div>
        </div>
        <div class="container">
            <section id="img-one" class="img">
                <h1>地铁系统</h1>
            </section>
            <section class="content-one">
                <div class="text-box">
                    <div class="cnt">
                        <h2>Welcome</h2>
                        <p>
                            欢迎踏入未来出行的便捷之门——地铁智能系统首页。在这里，每一次点击都引领您迈向更加高效、舒适的出行体验。

                        </p>

                    </div>
                </div>
            </section>
            <section class="content-two">
                <div class="text-box">
                    <div class="cnt">
                        <h2>智能化</h2>
                        <p>
                            我们运用了很多的预测算法和排班方案来优化我们的地铁智能系统！
                        </p>

                    </div>
                </div>
            </section>
            <section class="content-three">
                <div class="text-box">
                    <div class="cnt">
                        <h2></h2>
                        <p>

                        </p>

                    </div>
                </div>
            </section>
        </div>
    </div>
</template>


<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

function navigateToLogin() {
    router.push('/login');
}

function navigateToRegister() {
    router.push('/register');
}



</script>


<style>
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

.container {
    width: 100vw;
    height: auto;
}

section,
#footer {
    height: 100vh;
    position: relative;
    color: rgba(255, 255, 255, .8);
    font-size: 10vh;
    /* 设置字体 */
    font-family: '微软雅黑',
        serif;
    /* 文字阴影 */
    text-shadow: 4px 4px rgb(46, 43, 43);
    /* 弹性布局 */
    display: flex;
    /* 水平居中 */
    justify-content: center;
    /* 垂直居中 */
    align-items: center;
}

#footer {
    background-color: #c0b5aa;
    height: 50vh;
    color: #fff;
    font-size: 20vh;
}

.img {
    overflow: hidden;
    /* 背景图片不重复 */
    background-repeat: no-repeat;
    /* 背景图片位置：居中 */
    background-position: center;
    /* 背景图片固定 */
    background-attachment: fixed;
    background-size: cover;
    /* 禁止用户误选中 */
    user-select: none;

}


#img-one {
    background-image: url(../assets/img2.jpg);
    background-size: cover;
}

#img-two {
    background-image: url(../assets/img4.jpg);
}

.content-one {
    background: url(../assets/img1.jpg) no-repeat fixed center;
    background-size: cover;
}

.content-two {
    background: url(../assets/img2.jpg) no-repeat fixed center;
    background-size: cover;
}

.content-three {
    background: url(../assets/img3.jpg) no-repeat fixed center;
    background-size: cover;
}

.text-box {
    width: 100%;
    height: 50%;
    /* 绝对定位 */
    position: absolute;
    bottom: 25%;
    /* 文字居中 */
    text-align: center;
    background-color: rgba(0, 0, 0, .3);
    /* 旋转中心 */
    transform-origin: center bottom;
    /* 扭曲斜切 */
    transform: skew(0, -8deg);

}

.text-box .cnt {
    display: inline-block;
    text-align: right;
    margin: 40px 0 10px;
    width: 90%;
    max-width: 990px;
    transform-origin: right bottom;
    transform: skew(0, 8deg);
}

.cnt h2 {
    margin-top: 100px;
    font-size: 36px;
    font-family: '微软雅黑',
        serif;
    text-shadow: 1px 0px 0px #323232;
    letter-spacing: 0.3em;
    /* 大写字母 */
    text-transform: uppercase;
    color: #fff;
    margin-bottom: 5px;
}

.cnt p {
    font-size: 20px;
    line-height: 1.7;
    margin-bottom: 10px;
    color: #fff;
}

.cnt a.btn {
    font-size: 90%;
    font-family: '微软雅黑',
        serif;
    letter-spacing: 0.3em;
    text-shadow: 1px 0px 0px black;
    line-height: 2.8;
    display: inline-block;
    background-color: #eaeaea;
    color: #000;
    border: 3px solid #000;
    white-space: nowrap;
    text-decoration: none;
    padding: 0 20px;

}

/* ---------------------------------------------- */
.top {
    z-index: 1000;
    display: block;
    width: 100%;
}

.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 30px 50px 30px 10px;
    height: 120px;
    background-color: rgba(53, 53, 53, 0.45);
}

.logo {
    width: 10vw;
    min-width: 100px;
}

.header>.cta {
    margin-left: auto;


}

button {
    padding: 9px 10px;
    background-color: rgba(0, 136, 169, 1);
    border: none;
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s ease 0s;
    width: 4vw;
    min-width: 50px;
    margin: 10px;

}

button:hover {
    background-color: rgba(0, 136, 169, 0.7);
}


.total {
    position: relative;
}

.top {
    position: fixed;
    ;
    top: 0;
    left: 0;
}

.container {
    position: absolute;
    top: 0;
    left: 0;
}

.but {
    display: inline;
}


.but button{
    width:100px;
    height: 50px;
    font-size:20px;
}

</style>

