<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Topology</span>
        <div style="float: right;">
          <i class="el-icon-refresh" @click="refresh"></i>
        </div>
      </el-card>
    </el-header>
    <el-main>
      <Topology :key="key"/>
    </el-main>
  </el-container>
</template>
<script>
import Topology from "../components/Topology";

export default {
  name: "Graph",
  components: {
    Topology
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
    min-height: 500px;
    min-width: 650px;
  }
</style>