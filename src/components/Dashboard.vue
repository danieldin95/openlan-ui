<template>
  <el-card class="box-card">
    <el-row :gutter="20">
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <div class="card">
          <div class="value">
            <el-row>
              <el-col :span="8" style="padding: 5px;">
                <div class="title">Total</div>
                <div class="info">{{point.total}}</div>
                <div class="title">Success</div>
                <div class="success">{{point.success}}</div>
              </el-col>
              <el-col :span="16" style="">
                <el-progress type="circle" :width="100"
                             :percentage="point.percentage" :color="colors">
                </el-progress>
              </el-col>
            </el-row>
          </div>
          <div class="label">Access point</div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="8" :lg="6">
        <div class="card">
          <div class="value">
            <el-row>
              <el-col :span="8" style="padding: 5px;">
                <div class="title">Total</div>
                <div class="info">{{vswitch.total}}</div>
                <div class="title">Success</div>
                <div class="success">{{vswitch.success}}</div>
              </el-col>
              <el-col :span="16" style="">
                <el-progress type="circle" :width="100"
                             :percentage="vswitch.percentage" :color="colors">
                </el-progress>
              </el-col>
            </el-row>
          </div>
          <div class="label">Virtual switch</div>
        </div>
      </el-col>
    </el-row>
  </el-card>
</template>
<script>
import axios from "axios";
import 'echarts/lib/chart/pie';

export default {
  name: "Dashboard",
  components: {
  },
  data: function() {
    return {
      point: {
        total: 0,
        success: 0,
        percentage: 100,
      },
      vswitch: {
        total: 0,
        success: 0,
        percentage: 100,
      },
      colors: [
        {color: '#f50d0a', percentage: 20},
        {color: '#e62f12', percentage: 40},
        {color: '#e66c09', percentage: 60},
        {color: '#b88514', percentage: 80},
        {color: '#b8a412', percentage: 99},
        {color: '#67c23a', percentage: 100}
      ]
    };
  },
  mounted: function() {
    axios.get("/api/point").then((resp) => {
      this.point.total = resp.data.length;
      this.point.success = resp.data.filter((e) => e.state === 'authenticated').length;
      if (this.point.total > 0) {
        this.point.percentage = this.point.success * 100 / this.point.total;
      }
    }).catch((error) => {
      console.log(`GET /api/point: ${error}`);
    });

    axios.get("/api/vswitch").then((resp) => {
      this.vswitch.total = resp.data.length;
      this.vswitch.success = resp.data.filter((e) => e.state === '').length;
      if (this.vswitch.total > 0) {
        this.vswitch.percentage = this.vswitch.success * 100 / this.vswitch.total;
      }
    }).catch((error) => {
      console.log(`GET /api/vswitch: ${error}`);
    });
  }
};
</script>

<style scoped>
  .card {
    background: #fbfcfd;
    height: 180px;
    text-align: center;
    padding: 15px;
  }

  .card .label {
    color: #909399;
    font-size: 18px;
  }

  .card .value {
    padding-top: 12px;
    padding-bottom: 24px;
    color: #303133;
    font-size: 24px;
  }

  .card .value .title {
    color: #303133;
    font-size: 12px;
    padding-bottom: 5px;
  }

  .card .value .success {
    color: #67C23A;
    font-weight: bolder;
  }

  .card .value .warn {
    color: #E6A23C;
    font-weight: bolder;
  }

  .card .value .info {
    color: #409EFF;
    font-weight: bolder;
  }
</style>