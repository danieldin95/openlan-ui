<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Virtual Switch</span>
        <div style="float: right;">
          <i class="el-icon-refresh" @click="refresh"></i>
          <el-divider direction="vertical"></el-divider>
          <el-dropdown>
            <i class="el-icon-setting"></i>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>Edit</el-dropdown-item>
              <el-dropdown-item>Add</el-dropdown-item>
              <el-dropdown-item>Delete</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </el-card>
    </el-header>
    <el-main>
      <VSwitchList :key="key"/>
    </el-main>
  </el-container>
</template>

<script>
import VSwitchList from "../components/VSwitchList";

export default {
  name: 'Home',
  components: {
    VSwitchList,
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
    padding-right: 10px;
  }

  .el-main {
    overflow: visible;
  }

  .el-main .el-card {
    min-height: 500px;
    min-width: 650px;
    padding-right: 10px;
  }
</style>
