<template>
  <el-card class="box-card">
    <el-table
            :data="table"
            stripe
            style="width: 100%"
            :default-sort = "{prop: 'switch', order: 'ascending'}">
      <el-table-column
              type="index">
      </el-table-column>
      <el-table-column
              label="Uptime"
              width="100">
        <template slot-scope="scope">
          <span>{{ prettyTime(scope.row.uptime) }}</span>
        </template>
      </el-table-column>
      <el-table-column
              prop="alias"
              label="Alias"
              width="200">
      </el-table-column>
      <el-table-column
              sortable
              prop="switch"
              label="Switch"
              width="200">
      </el-table-column>
      <el-table-column
              sortable
              prop="network"
              label="Network"
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
              label="State"
              min-width="120">
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: 'PointList',
  methods: {
    prettyTime(second) {
      let min = Math.floor(second / 60);
      if (min < 60) {
        return min + "m" + (second % 60) + "s";
      }
      let hour = Math.floor(min / 60);
      if (hour < 24) {
        return hour + "h" + (min % 60) + "m";
      }
      let day = Math.floor(hour / 24);
      return day + "d" + (hour % 24) + "h";
    }
  },
  data: function() {
    return {
      table: [],
    };
  },
  mounted: function() {
    axios.get("/api/point").then((resp) => {
      this.table = resp.data;
    }).catch((error) => {
      console.log(`GET /api/point: ${error}`);
    });
  },
};
</script>

<style scoped>
</style>
