<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Dashboard</span>
      </el-card>
    </el-header>
    <el-main>
      <Dashboard/>
    </el-main>
  </el-container>
</template>

<script>
import Dashboard from "../components/Dashboard";

export default {
  name: 'Home',
  components: {
    Dashboard,
  }
};
</script>

<style scoped>
  .el-header {
    overflow: visible;
    margin-bottom: 10px;
  }

  .el-header .el-card {
    min-width: 650px;
  }

  .el-main {
    overflow: visible;
  }

  .el-main .el-card {
    min-height: 500px;
    min-width: 650px;
  }
</style>
