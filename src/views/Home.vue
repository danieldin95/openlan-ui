<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Dashboard</span>
        <div style="float: right;">
          <i class="el-icon-refresh" @click="refresh"></i>
        </div>
      </el-card>
    </el-header>
    <el-main>
      <Dashboard :key="dashboardKey"/>
      <Message :key="messageKey"/>
    </el-main>
  </el-container>
</template>

<script>
import Dashboard from "../components/Dashboard";
import Message from "../components/Message";

export default {
  name: 'Home',
  computed: {
    dashboardKey() {
      return "dashboard-" + this.key;
    },
    messageKey() {
      return "message-" + this.key;
    }
  },
  components: {
    Message,
    Dashboard,
  },
  methods: {
    refresh() {
      this.key += 1;
    }
  },
  data: function() {
    return {
      key: 0,
    };
  },
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
