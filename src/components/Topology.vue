<template>
  <el-card class="topology box-card">
    <ECharts :theme="theme" :options="options" :autoresize="true"></ECharts>
  </el-card>
</template>
<script>
import axios from "axios";
import ECharts from 'vue-echarts';
import 'echarts/lib/chart/graph';

export default {
  name: "Topology",
  components: {
    ECharts,
  },
  data () {
    return {
      options: {},
      theme: 'dark',
    };
  },
  created() {
    axios.get('/api/graph/default').then((response) => {
      let webkit = response.data;
      this.options = {
        legend: {
          data: webkit.categories.map(function (node) {
            return node.name;
          })
        },
        series: [{
          type: 'graph',
          layout: 'force',
          roam: true,
          draggable: true,
          animation: true,
          focusNodeAdjacency: true,
          data: webkit.nodes,
          categories: webkit.categories,
          edges: webkit.links,
          label: {
            position: 'right',
            show: true,
          },
          force: {
            edgeLength: 160,
            repulsion: 400,
          },
        }]
      };
    }).catch((error) => {
      console.log(`GET /api/graph/default: ${error}`);
    });
  }
};
</script>

<style scoped>
.topology {
  margin: 0 auto;
}

.echarts {
  width: 100%;
  min-width: 1024px;
  min-height: 500px;
}
</style>