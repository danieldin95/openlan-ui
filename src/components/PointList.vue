<template>
  <el-card class="point box-card">
    <el-table
            :data="tableData"
            stripe
            style="width: 100%"
            :default-sort = "{prop: 'uptime', order: 'ascending'}">
      <el-table-column
              prop="uptime"
              label="UpTime"
              sortable
              width="100">
      </el-table-column>
      <el-table-column
              prop="alias"
              label="Alias"
              width="200">
      </el-table-column>
      <el-table-column
              prop="switch"
              label="Switch"
              sortable
              width="200">
      </el-table-column>
      <el-table-column
              prop="tenant"
              label="Tenant"
              sortable
              width="100">
      </el-table-column>
      <el-table-column
              prop="server"
              label="Address"
              width="180">
      </el-table-column>
      <el-table-column
              prop="device"
              label="Device"
              width="120">
      </el-table-column>
      <el-table-column
              prop="state"
              label="State">
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: 'PointList',
  data() {
    return {
      tableData: [],
    };
  },
  created() {
    axios.get("/api/point").then((resp) => {
      this.tableData = resp.data;
    }).catch((error) => {
      console.log(`GET /api/point: ${error}`);
    });
  },
};
</script>

<style scoped>
.point {
  margin: 0 auto;
}
</style>
