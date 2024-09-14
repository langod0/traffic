import { createApp } from 'vue';
import App from './App.vue';
import ElementPlus from'element-plus'
import router from './router'; // 导入路由

const app = createApp(App);

app.use(router); // 使用路由
app.use(ElementPlus);
app.mount('#app');