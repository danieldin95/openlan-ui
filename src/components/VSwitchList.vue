<template>
  <div class="vswitch">
    <el-table
            :data="tableData"
            stripe
            style="width: 100%">
      <el-table-column
              prop="name"
              label="Name"
              width="180">
      </el-table-column>
      <el-table-column
              prop="url"
              label="Url"
              width="280">
      </el-table-column>
      <el-table-column
              prop="password"
              label="Token">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: 'VSwitchList',
  data() {
    return {
      tableData: [],
    };
  },
  created() {
    axios.get("/api/vswitch").then((resp) => {
      this.tableData = resp.data;
    }).catch((error) => {
      console.log(`GET /api/vswitch: ${error}`);
    });
  },
};
</script>

<style scoped>
.vswitch {
  margin: 0 auto;
  max-width: 1024px;
}
</style>
