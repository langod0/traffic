
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import axios from "axios"
//需要挂载到Vue原型上
// Vue.prototype.$echarts = echarts
const app = createApp(App);
// axios.defaults.baseURL = "http://localhost:8001"
// axios.defaults.baseURL= 'http://127.0.0.1:7234/'
// app.config.globalProperties.$axios = axios
app.config.globalProperties.$http = axios

app.use(router);

app.mount('#app');