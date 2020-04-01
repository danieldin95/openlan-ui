<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Dashboard</span>
      </el-card>
    </el-header>
    <el-main>
      <Dashboard/>
      <Message/>
    </el-main>
  </el-container>
</template>

<script>
import Dashboard from "../components/Dashboard";
import Message from "../components/Message";

export default {
  name: 'Home',
  components: {
    Message,
    Dashboard,
  }
};
</script>

<style scoped>
  .el-header {
    overflow: visible;
  }

  .el-header .el-card {
    min-width: 650px;
  }

  .el-main {
    overflow: visible;
  }

  .el-main .el-card {
    min-width: 650px;
    margin-bottom: 15px;
  }
</style>
