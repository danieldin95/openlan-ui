<template>
  <el-container>
    <el-header>
      <el-card class="box-card">
        <span style="font-size: smaller;">Dashboard</span>
      </el-card>
    </el-header>
    <el-main>
      <div class="home">
        <VSwitchList/>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import VSwitchList from "../components/VSwitchList";

export default {
  name: 'Home',
  components: {
    VSwitchList,
  }
};
</script>

<style scoped>
</style>
