<!--
Monitoring page, fetches metrics data from API
----------------------------------------------
Ben C, April 2018
-->

<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <fa icon="tachometer-alt" />&nbsp; Monitoring
    </div>
    <div class="card-body">
      <spinner v-if="!metrics" />

      <b-container v-if="metrics" fluid>
        <b-row align-h="around">
          <b-col md>
            <dial :value="cpu" :title="'CPU Load'" :percentage="true" />
          </b-col>
          <b-col md>
            <dial :value="mem" :title="'Memory Used'" :percentage="true" />
          </b-col>
        </b-row>
        <b-row align-h="around">
          <b-col md>
            <dial :value="disk" :title="'Disk Used'" :percentage="true" />
          </b-col>
          <b-col md>
            <dial :value="net" :title="'Net I/O'" :percentage="false" />
          </b-col>
        </b-row>
      </b-container>
    </div>
  </div>
</template>

<script>
import apiMixin from '../mixins/apiMixin.js'
import Spinner from './Spinner.vue'
import Dial from './Dial.vue'

let prevNetBytes

export default {
  components: {
    Spinner,
    Dial
  },

  mixins: [apiMixin],

  data: function() {
    return {
      metrics: null,
      prevNetBytes: null
    }
  },

  computed: {
    cpu: function() {
      if (!this.metrics) { return 0 }
      return this.metrics.cpuPerc
    },
    mem: function() {
      if (!this.metrics) { return 0 }
      return (this.metrics.memUsed / this.metrics.memTotal) * 100
    },
    disk: function() {
      if (!this.metrics) { return 0 }
      return 100 - ((this.metrics.diskFree / this.metrics.diskTotal) * 100)
    },
    net: function() {
      let newTot = this.metrics.netBytesSent + this.metrics.netBytesRecv
      let delta = newTot - prevNetBytes
      if (prevNetBytes) {
        prevNetBytes = newTot
        // Scaling factor here is mostly arbitrary to get nice looking numbers
        return delta / 1000
      } else {
        prevNetBytes = newTot
        return 0
      }
    }
  },

  created() {
    this.update()
    setInterval(this.update, 2500)
  },

  methods: {
    update: function() {
      this.apiGetMetrics()
        .then((json) => { this.metrics = json })
    }
  }
}
</script>

<style scoped>
  .card-body {
    background-color: #ddd;
  }
</style>