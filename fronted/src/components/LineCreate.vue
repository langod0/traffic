<template>
  <div class="line-create">
    <h2>创建新线路</h2>
    <form @submit.prevent="createLine">
      <div class="form-content">
        <label for="line-name">线路名称:</label>
        <input type="text" id="line-name" v-model="newLineName" placeholder="线路名称" required>
      </div>
      <button type="submit">提交</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newLineName: ''
    };
  },
  methods: {
    createLine() {
      // Logic to create a new line and store it in Vuex or local state
      console.log('创建线路:', this.newLineName);
    }
  }
};
</script>

<style scoped>

</style>