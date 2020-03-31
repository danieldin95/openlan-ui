<template>
  <el-container>
    <el-header>
      <el-card class="box-card">
        <span style="font-size: smaller;">Topology</span>
      </el-card>
    </el-header>
    <el-main>
      <div class="graph">
        <Topology/>
      </div>
    </el-main>
  </el-container>
</template>
<script>
import Topology from "../components/Topology.vue";

export default {
  name: "Graph",
  components: {
    Topology
  },
};
</script>

<style scoped>
</style>