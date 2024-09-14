<template>
  <div class="schedule-table">
    <h2>排班表</h2>
    <table>
      <thead>
        <tr>
          <th>时间</th>
          <th>车辆编号</th>
          <th>司机</th>
          <th>路线</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(schedule, index) in schedules" :key="index">
          <td>{{ schedule.time }}</td>
          <td>{{ schedule.busNo }}</td>
          <td>{{ schedule.driver }}</td>
          <td>{{ schedule.route }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      schedules: [
        { time: '08:00', busNo: 'B123', driver: '张三', route: 'A -> B' },
        // Add more schedules as needed
      ]
    };
  }
};
</script>

<style scoped>
.schedule-table {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f4f4f4;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.schedule-table h2 {
  text-align: center;
  margin-bottom: 20px;
}

.schedule-table table {
  width: 100%;
  border-collapse: collapse;
}

.schedule-table th, .schedule-table td {
  padding: 10px;
  text-align: left;
  border: 1px solid #ccc;
}

.schedule-table th {
  background-color: #007bff;
  color: white;
}

.schedule-table tr:nth-child(even) {
  background-color: #f2f2f2;
}
</style>