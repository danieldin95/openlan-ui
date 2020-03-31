<template>
  <div id="app">
    <el-container>
      <el-aside width="200px">
        <el-menu default-active="1">
          <el-menu-item index="1">
            <i class="el-icon-pie-chart"></i>
            <router-link to="/">Dashboard</router-link>
          </el-menu-item>
          <el-menu-item index="1">
            <i class="el-icon-s-home"></i>
            <router-link to="/vswitch">Virtual switch</router-link>
          </el-menu-item>
          <el-menu-item index="2">
            <i class="el-icon-s-platform"></i>
            <router-link to="/point">Access point</router-link>
          </el-menu-item>
          <el-menu-item index="3">
            <i class="el-icon-s-marketing"></i>
            <router-link to="/topo">Topology</router-link>
          </el-menu-item>
          </el-menu>
      </el-aside>
      <router-view/>
    </el-container>
  </div>
</template>

<script>
export default {
};
</script>

<style scoped>
  #app {
    font-family: Arial,sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  .el-aside {
    border-right: solid 1px #e6e6e6;
  }

  .el-menu {
    border-right: none;
    min-height: 600px;
  }

  .el-menu-item a {
    padding: 15px 0;
    text-decoration-line: none;
  }

  a.router-link-exact-active {
    color: #42b983;
  }
</style>
