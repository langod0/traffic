<template>
    <div class="mainshow">
      <div class="bs">
        <button @click="Show_data" class="showb">客流量预测</button>
        <!-- <span>aaa{{ Stnumin }}</span> -->
      </div>
  
      <div class="op" v-if="isss">
        日期：2019.1.29
        
        站点编号：
        <select class="slt" name="slt2" id="slt2" @change="selectValue2($event.target.value)">
          <option> 请选择编号</option>
          <option v-for="item in Stationnum" :value="item" :key="item">
            {{ item }}
          </option>
          <option value="all">全站点</option>
        </select>
        上下车：
        <select class="slt" name="slt3" id="slt3" @change="selectValue3($event.target.value)">
          <option> 请选择范围</option>
          <option value="in">上车数量</option>
          <option value="out">下车数量</option>
          <option value="all">都包含</option>
        </select>
        <button class="btns" @click="showdata">可视化展示</button>
      </div>
      <!-- <br> -->
      <div ref="chart" class="shows" v-if="isst"></div>
    </div>
  </template>
  
  <script setup>
  import { onMounted, ref , nextTick} from "vue";
  import * as echarts from "echarts";
  import axios from "axios";
  let Numin = ref("");
  let Numout = ref("");
  let Stnumin = ref("");
  let Stnumout = ref("");
  let Tm = ref("");
  let Stationnum = ref(0);
  const isss = ref(false);
  const isst = ref(false);
  let dd=ref("");
  let ss=ref(0);
  let tit=ref("");
  let xdata=ref([]);
  let ydata=ref([]);
  let stat=ref("");
  function Show_data() {
    isss.value = true;
    console.log(isss.value);
    axios
      .post("pyapi/pre_api/pre/", {})
      .then((response) => {
        Numin.value = response.data.Numin;
        // console.log(1)
        Numout.value = response.data.Numout;
        Stationnum.value = response.data.Stationnum;
        // alert(Stationnum)
        Stnumin.value = response.data.station_num_in;
        Stnumout.value = response.data.station_num_out;
        Tm.value = response.data.Tm;
        // console.log(Stnumin.value)
        console.log(1);
        
      })
      .catch((error) => {
        console.error(error);
      });
  }
  dd.value="29";
  function selectValue2(sl){
    ss.value=sl
  }
  function selectValue3(sl){
    stat.value=sl
  }
  function showdata(){
    
    ydata.value=[]
    isst.value =true;
    console.log(dd.value);
    console.log(typeof(ss.value));
    console.log(stat.value);
      let ii=ref(0)
      ii.value=Number(dd.value)
      console.log(ii.value)
      console.log(typeof(ii.value))
      console.log(ss.value)
      console.log(stat.value)
      xdata.value =Tm.value;
      if(ss.value=="all"){
        if(stat.value=="all"){
          tit.value="2019.1."+dd.value+"号每个时间段总客流量";
          for(var i=0;i<Numin.value.length;i++){
            ydata.value.push(Numin.value[i]+Numout.value[i])
          }
          // ydata.value=Numin[i]
        }else if(stat.value=="in"){
          tit.value="2019.1."+dd.value+"号每个时间段上车客流量";
          for(var i=0;i<Numin.value.length;i++){
            ydata.value.push(Numin.value[i])
          }
        }else{
          tit.value="2019.1."+dd.value+"每个时间段下车客流量";
          for(var i=0;i<Numin.value.length;i++){
            ydata.value.push(Numout.value[i])
          }
        }
      }else{
        let jj=ref(0)
        jj.value=Number(ss.value)-1;
        console.log(11)
        if(stat.value=="all"){
          console.log(11)
          tit.value="2019.1."+dd.value+","+String(jj.value)+"号站点每个时间段总客流量";
          for(var i=0;i<Stnumin.value[jj.value].length;i++){
            ydata.value.push(Stnumin.value[jj.value][i]+Stnumout.value[jj.value][i])
          }
        }else if(stat.value=="in"){
          tit.value="2019.1."+dd.value+","+String(jj.value)+"号站点每个时间段上车客流量";
          for(var i=0;i<Stnumin.value[jj.value].length;i++){
            ydata.value.push(Stnumin.value[jj.value][i])
          }
        }else{
          tit.value="2019.1."+dd.value+","+String(jj.value)+"号站点每个时间段下车客流量";
          for(var i=0;i<Stnumin.value[jj.value].length;i++){
            ydata.value.push(Stnumout.value[jj.value][i])
          }
        }
      }
    
    console.log(xdata.value)
    console.log(ydata.value)
  
    console.log(dd,ss);
    nextTick(() => {
          initChart(); // 在显示图表之后调用初始化
        });
  }
  // 引用 ECharts 组件的 DOM
  const chart = ref(null);
  
  // 初始化 ECharts 图表
  const initChart = () => {
    const chartInstance = echarts.init(chart.value);
  
    // 配置图表选项
    const options = {
      title: {
        text: tit.value,
        textStyle: {
          color: '#000', // 设置标题字体颜色为黑色
        },
      },
      tooltip: {
        textStyle: {
          color: '#000', // 设置标题字体颜色为黑色
        },
      },
      xAxis: {
        data: xdata.value,
        axisLabel: {
          textStyle: {
            color: '#000', // 设置 x 轴标签字体颜色为黑色
          },
        },
      },
      yAxis: {
        axisLabel: {
          textStyle: {
            color: '#000', // 设置 x 轴标签字体颜色为黑色
          },
        },
      },
      series: [
        {
          type: "line",
          data: ydata.value,
          lineStyle: {
            width: 8, // 设置线条宽度，值越大线条越粗
            color: '#FFFFFF',
          },
          itemStyle: {
            
            color: '#FF0000', // 设置数据点的颜色（可选）
          },
          symbolSize: 10,
        },
      ],
    };
  
    // 设置图表选项
    chartInstance.setOption(options);
  };
  
  // 当组件挂载时初始化图表
  onMounted(() => {
    nextTick(initChart);
  });
  </script>
  
  <style scoped>
  button {
    cursor: pointer;
    transition: 0.7s;
  }
  /* 可选：为图表容器添加样式 */
  .showb {
    background-color: rgb(185, 235, 235);
    width: 120px;
    height: 40px;
    border-radius: 10px;
    border: none;
    font-family: 楷体;
    font-size: 20px;
    font-weight: bold;
  }
  .showb:hover {
    background-color: rgb(132, 159, 205);
  }
  .bs {
    float: left;
    width: 120px;
    /* border: 3px solid blue; */
  }
  .op {
    margin-left: 200px;
    width: 850px;
    height: 40px;
    float: left;
    /* border: 3px solid blue; */
    font-size: 22px;
  }
  .shows {
    float: left;
    margin-left: 200px;
    margin-top: 50px;
    width: 900px;
    height: 600px;
    /* border: 3px solid blue; */
  }
  .mainshow {
    width: 1300px;
    height: 740px;
    /* border: 3px solid blue; */
  }
  .slt {
    width: 120px;
    height: 35px;
    background-color: rgb(213, 225, 237);
    font-family: 楷体;
    font-size: 18px;
  }
  .btns {
    margin-left: 50px;
    background-color: rgb(206, 216, 234);
    width: 120px;
    height: 38px;
    border-radius: 10px;
    border: none;
    font-family: 楷体;
    font-size: 17px;
    font-weight: bold;
  }
  .btns:hover {
    background-color: rgb(132, 159, 205);
    font-size: 19px;
  }
  </style>
  