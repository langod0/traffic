<template>
  <div class="info-form-container">
    <h2 class="form-title">线路及发车时刻表编号</h2>
    <form @submit.prevent="generateSchedule" class="info-form">
      <div class="form-group">
        <label for="line-id">线编号:</label>
        <input type="number" id="line-id" v-model.number="lineId" min="1" required>
      </div>

      <div class="form-group">
        <label for="schedule-id">时刻表编号:</label>
        <input type="number" id="schedule-id" v-model.number="scheduleId" min="1" required>
      </div>

      <div class="form-group">
        <label for="peak-time-start">早高峰时间段:</label>
        <input type="time" id="peak-time-start" v-model="peakStartTime" required>
        <label for="peak-time-end">至:</label>
        <input type="time" id="peak-time-end" v-model="peakEndTime" required>
      </div>

      <div class="form-group">
        <label for="off-peak-time-start">晚高峰时间段:</label>
        <input type="time" id="off-peak-time-start" v-model="offPeakStartTime" required>
        <label for="off-peak-time-end">至:</label>
        <input type="time" id="off-peak-time-end" v-model="offPeakEndTime" required>
      </div>

      <div class="form-group">
        <label for="peak-break-time">高峰休息时间:</label>
        <input type="number" id="peak-break-time" v-model.number="peakBreakTime" min="0" step="1" required>
        <label for="off-peak-break-time">平峰休息时间:</label>
        <input type="number" id="off-peak-break-time" v-model.number="offPeakBreakTime" min="0" step="1" required>
      </div>

      <div class="form-group">
        <label for="disinfect-interval">消毒间隔次数:</label>
        <input type="number" id="disinfect-interval" v-model.number="disinfectInterval" min="0" step="1" required>
        <label for="disinfect-extra-break-time">消毒额外休息时间:</label>
        <input type="number" id="disinfect-extra-break-time" v-model.number="disinfectExtraBreakTime" min="0" step="1"
          required>
      </div>

      <div class="form-group">
        <label for="max-bus-count">最大车辆数:</label>
        <input type="number" id="max-bus-count" v-model.number="maxBusCount" min="1" required>
      </div>
      <div>
        <label for="upstream-max-bus-count">上行最大车辆数:</label>
        <input type="number" id="upstream-max-bus-count" v-model.number="upstreamMaxBusCount" min="1" required>
        <label for="downstream-max-bus-count">下行最大车辆数:</label>
        <input type="number" id="downstream-max-bus-count" v-model.number="downstreamMaxBusCount" min="1" required>
      </div>
      <div class="form-group">
        <label for="max-work-hours">最大工作时长:</label>
        <select id="max-work-hours" v-model="maxWorkHours">
          <option value="8">8小时</option>
          <option value="10">10小时</option>
          <option value="11">11小时</option>
          <option value="12">12小时</option>
        </select>
      </div>

      <div class="form-group">
        <button type="submit">生成行车计划</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lineId: null,
      scheduleId: null,
      peakStartTime: null,
      peakEndTime: null,
      offPeakStartTime: null,
      offPeakEndTime: null,
      peakBreakTime: null,
      offPeakBreakTime: null,
      disinfectInterval: null,
      disinfectExtraBreakTime: null,
      maxBusCount: null,
      upstreamMaxBusCount: null,
      downstreamMaxBusCount: null,
      maxWorkHours: null
    };
  },
  methods: {
    generateSchedule() {
      // Logic to generate the schedule based on user inputs
      console.log('生成行车计划:', {
        lineId: this.lineId,
        scheduleId: this.scheduleId,
        peakStartTime: this.peakStartTime,
        peakEndTime: this.peakEndTime,
        offPeakStartTime: this.offPeakStartTime,
        offPeakEndTime: this.offPeakEndTime,
        peakBreakTime: this.peakBreakTime,
        offPeakBreakTime: this.offPeakBreakTime,
        disinfectInterval: this.disinfectInterval,
        disinfectExtraBreakTime: this.disinfectExtraBreakTime,
        maxBusCount: this.maxBusCount,
        upstreamMaxBusCount: this.upstreamMaxBusCount,
        downstreamMaxBusCount: this.downstreamMaxBusCount,
        maxWorkHours: this.maxWorkHours
      });
    }
  }
};
</script>

<style scoped>
.info-form-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f4f4f4;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgb(167, 167, 167);
  /* font-family: 'Arial', sans-serif; */
  color: #ffffff;
}

.form-title {
  text-align: center;
  margin-bottom: 20px;
  color: #000000;
}

.info-form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.form-group {
  display: flex;
  align-items: center;
  gap: 10px;
}

label {
  flex-shrink: 0;
  width: 120px;
  /* Adjust the width as needed */
  font-weight: bold;
  text-align: right;
  margin-bottom: 5px;
  color: #000000;
}

input,
select {
  flex-grow: 1;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #1a1a1a;
  border-radius: 4px;
  /* background-color: #cfcfcf; */
  /* color: #fff; */
  transition: all 0.3s ease-in-out;
}

input:focus,
select:focus {
  outline: none;
  border-color: #3b3b3b;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

button {
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}

button:active {
  transform: translateY(1px);
}
</style>