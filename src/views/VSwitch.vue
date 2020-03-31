<template>
  <el-container>
    <el-header>
      <el-card class="box-card">
        <span style="font-size: smaller;">Virtual Switch</span>
        <el-divider direction="vertical"></el-divider>
        <el-dropdown>
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
      <div class="home">
        <VSwitchList/>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import VSwitchList from "../components/VSwitchList";

export default {
  name: 'Home',
  components: {
    VSwitchList,
  }
};
</script>

<style scoped>
</style>
