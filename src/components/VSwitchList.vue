<template>
  <el-card class="box-card">
    <el-table
            :data="table"
            stripe
            style="width: 100%">
      <el-table-column
              prop="name"
              label="Name"
              width="180">
      </el-table-column>
      <el-table-column
              label="Url"
              width="280">
        <template slot-scope="scope">
          <el-link :href="scope.row.url">{{ scope.row.url }}</el-link>
        </template>
      </el-table-column>
      <el-table-column
              prop="password"
              label="Token">
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: 'VSwitchList',
  data() {
    return {
      table: [],
    };
  },
  mounted() {
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
