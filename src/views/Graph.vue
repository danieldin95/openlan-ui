<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Topology</span>
      </el-card>
    </el-header>
    <el-main>
      <Topology/>
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