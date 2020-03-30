<template>
  <div class="graph">
    <h1>This is an Graph page</h1>
    <ECharts :theme="theme" :options="lineOpts" :autoresize="true"></ECharts>
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
      theme: 'default',
    };
  },
  created() {
    console.log('Graph.created');
  },
  mounted() {
    this.lineOpts = {
      xAxis: {
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
      },
      yAxis: {
        type: 'value'
      },
      series: [{
        data: [820, 932, 901, 934, 1290, 1330, 1320],
        type: 'line',
        smooth: true
      }]
    };

    axios.get('/api/webkit')
      .then((response) => {
        let webkit = response.data;
        this.graphOpts = {
          theme: 'dark',
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
            focusNodeAdjacency: true,
            data: webkit.nodes,
            categories: webkit.categories,
            links: webkit.links,
            label: { position: 'right' },
            force: { edgeLength: 180 },
          }]
        };
      })
      .catch((error) => {
        console.log(`Error! Could not reach the API: ${error}`);
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