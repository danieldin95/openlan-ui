<template>
  <div class="graph">
    <h1>OpenLAN Topology</h1>
    <ECharts :theme="theme" :options="graphOpts" :autoresize="true"></ECharts>
  </div>
</template>
<script>
import ECharts from 'vue-echarts';
import 'echarts/lib/chart/line';
import 'echarts/lib/chart/graph';
import axios from "axios";

export default {
  name: "Graph",
  components: {
    ECharts,
  },
  data () {
    return {
      lineOpts: {},
      graphOpts: {},
      theme: 'dark',
    };
  },
  created() {
    axios.get('/api/graph/default').then((response) => {
      let webkit = response.data;
      this.graphOpts = {
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
          links: webkit.links,
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
.echarts {
  margin-right: auto;
  margin-left: auto;
}
</style>