<template>
  <el-card class="box-card">
    <el-table
            :data="table"
            stripe
            style="width: 100%">
      <el-table-column
              type="index">
      </el-table-column>
      <el-table-column
          prop="uuid"
          label="UUID"
          width="280">
      </el-table-column>
      <el-table-column
              prop="alias"
              label="Name"
              width="180">
      </el-table-column>
      <el-table-column
              label="UpTime"
              width="100">
        <template slot-scope="scope">
          <span>{{ prettyTime(scope.row.uptime) }}</span>
        </template>
      </el-table-column>
      <el-table-column
          label="Address"
          min-width="180">
        <template slot-scope="scope">
          <el-link :href="link(scope.row.address)">{{ scope.row.address }}</el-link>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: 'VSwitchList',
  data: function() {
    return {
      table: [],
    };
  },
  methods: {
    link: function (address) {
      return "https://"+address.split(":")[0]+":10000";
    },
    prettyTime(second) {
      let min = Math.floor(second / 60);
      if (min < 60) {
        return min + "m" + (second % 60) + "s";
      }
      let hour = Math.floor(min / 60);
      if (hour < 24) {
        return hour + "h" + (min % 60) + "s";
      }
      let day = Math.floor(hour / 24);
      return day + "d" + (hour % 24) + "h";
    }
  },
  mounted: function() {
    axios.get("/api/vswitch").then((resp) => {
      this.table = resp.data;
    }).catch((error) => {
      console.log(`GET /api/vswitch: ${error}`);
    });
  },
};
</script>

<style scoped>
</style>
