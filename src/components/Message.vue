<template>
  <el-card class="box-card">
    <el-table
            :data="table"
            style="width: 100%;"
            :row-class-name="rowClassName">
      <el-table-column
              type="index">
      </el-table-column>
      <el-table-column
              prop="date"
              label="Date"
              width="180">
      </el-table-column>
      <el-table-column
              prop="level"
              label="Level"
              sortable
              width="100">
      </el-table-column>
      <el-table-column
              prop="message"
              label="Message">
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import axios from "axios";

export default {
  name: 'Message',
  methods: {
    rowClassName({row, index}) {
      return 'level-' + row.level.toLowerCase();
    },
  },
  data() {
    return {
      table: [],
    };
  },
  mounted() {
    axios.get("/api/message?size=64").then((resp) => {
      this.table = resp.data;
    }).catch((error) => {
      console.log(`GET /api/message: ${error}`);
    });
  },
};
</script>

<style scoped>
</style>
