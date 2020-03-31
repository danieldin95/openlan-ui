<template>
  <el-container>
    <el-header>
      <el-card class="box-card" shadow="hover">
        <span style="font-size: smaller;">Accessed Point</span>
        <el-dropdown style="float: right;">
          <i class="el-icon-setting"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>Edit</el-dropdown-item>
            <el-dropdown-item>Add</el-dropdown-item>
            <el-dropdown-item>Delete</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-card>
    </el-header>
    <el-main>
      <PointList/>
    </el-main>
  </el-container>
</template>

<script>
import PointList from '../components/PointList.vue';

export default {
  name: 'Point',
  components: {
    PointList
  }
};
</script>
<style scoped>
  .el-header {
    overflow: visible;
    margin-bottom: 10px;
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
