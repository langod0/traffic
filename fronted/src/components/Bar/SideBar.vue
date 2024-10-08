<template>
    <body>
        <div id="main_div">
            <div class="menu-box">
                <div class="emotion-box">
                    <img src="../../assets/笑脸.png" id="emotion1">
                    <img src="../../assets/闭眼.png" id="emotion2">
                </div>
                <ul class="table">
                    <li>
                        <router-link to="#" id="router">
                            <img src="../../assets/首页.png">
                            <p class="content">首页</p>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/lines" id="router">
                            <img src="../../assets/线路.png">
                            <p class="content">创建线路</p>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/info" id="router">
                            <img src="../../assets/填写.png">
                            <p class="content">填写信息</p>
                        </router-link>
                    </li>
                    <li>
                        <router-link to="/schedule" id="router">
                            <img src="../../assets/查看.png">
                            <p class="content">查看排班</p>
                        </router-link>
                    </li>
                </ul>
            </div>
        </div>
    </body>
</template>

<style>
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

#main_div {
    height: 100vh;
    background-color: rgba(0, 0, 0);
    display: flex;
    justify-content: center;
    align-items: center;

}

.menu-box {
    width: 60px;
    height: 100vh;
    background-color: #22243a;
    border-radius: 30px;
    box-shadow: 1px 1px 2px rgba(0, 0, 0, .3);
    /* 添加过渡 */
    transition: all .2s ease-out;
    overflow: hidden;
    margin-right: auto;
    z-index: 1000;
}

.menu-box:hover {
    width: 120px;
    border-radius: 60px;
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
    width: 30px;
    height: 30px;
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
    width: 20px;
    height: 20px;
}

#router {
    display: flex;
    height: 50px;
    align-items:center;
    justify-content: center;
    
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
</style>