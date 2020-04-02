<template>
  <div id="app">
    <el-container>
      <el-aside width="auto">
        <div @click="onButton" :align="'right'">
          <span style="padding: 22px;">
            <i :class="iconButton"></i>
          </span>
        </div>
        <el-menu default-active="1" :collapse="collapse" class="el-menu-app" router>
          <el-menu-item index="/">
            <i class="el-icon-pie-chart"></i>
            <span slot="title">Dashboard</span>
          </el-menu-item>
          <el-menu-item index="/vswitch">
            <i class="el-icon-cpu"></i>
            <span slot="title">Virtual switch</span>
          </el-menu-item>
          <el-menu-item index="/point">
            <i class="el-icon-s-platform"></i>
            <span slot="title">Access point</span>
          </el-menu-item>
          <el-menu-item index="/topo">
            <i class="el-icon-s-marketing"></i>
            <span slot="title">Topology</span>
          </el-menu-item>
          </el-menu>
      </el-aside>
      <router-view/>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      collapse: true,
      iconButton: 'el-icon-d-arrow-right',
    };
  },
  methods: {
    onButton() {
      this.collapse = !this.collapse;
      if (this.collapse) {
        this.iconButton = 'el-icon-d-arrow-right';
      } else {
        this.iconButton = 'el-icon-d-arrow-left';
      }
    },
  }
};
</script>

<style>
  #app {
    font-family: Arial,sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  .el-main .el-card {
    max-height: 500px;
    overflow: auto;
  }

  .el-table .level-error {
    color: #F56C6C;
  }

  .el-table .level-warn {
    color: #E6A23C;
  }

  .el-table .el-info {
    color: #303133;
  }
</style>

<style scoped>
  .el-aside {
    border-right: solid 1px #e6e6e6;
  }

  .el-menu-app:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
    border-right: none;
  }

  .el-menu {
    min-height: 600px;
    border-right: none;
  }

  .el-menu-item span {
    padding: 15px 0;
  }
</style>
