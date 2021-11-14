<!--
Dial/gauge component, wrapper around gauge.js library
-----------------------------------------------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="card border-primary">
    <canvas ref="can">{{ value }}</canvas>
    <h5>{{ title }}: {{ valComputed }}<span v-if="percentage">%</span></h5>
  </div>
</template>

<script>
import { Gauge } from 'gaugeJS'

const staticZones = [
  { strokeStyle: '#30B32D', min: 0, max: 70 },
  { strokeStyle: '#FFDD00', min: 70, max: 90 },
  { strokeStyle: '#F03E3E', min: 90, max: 100 }
]

const opts = {
  angle: -0.15,
  lineWidth: 0.3,
  radiusScale: 1,
  pointer: {
    length: 0.5,
    strokeWidth: 0.045,
    color: '#4ce1ff'
  },
  colorStart: '#2063b9',
  limitMax: false,
  limitMin: false,
  strokeColor: '#ddd',
  generateGradient: true,
  highDpiSupport: true,
  fontSize: 40
}

export default {
  name: 'Dial',

  props: {
    value: {
      default: 0,
      type: Number
    },
    title: {
      default: 'No Label',
      type: String
    },
    percentage: {
      default: true,
      type: Boolean
    }
  },

  data: function () {
    return {
      gauge: null,
      opts: opts
    }
  },

  computed: {
    valComputed: function () {
      const rounded = Math.round(this.value * Math.pow(10, 2)) / Math.pow(10, 2)
      if (this.value && this.gauge) {
        this.gauge.set(rounded)
      }
      return rounded
    }
  },

  mounted: function () {
    const target = this.$refs.can

    if (this.percentage) {
      this.opts.staticZones = staticZones
    } else {
      this.opts.staticZones = null
    }
    this.opts.limitMax = this.percentage
    this.gauge = new Gauge(target).setOptions(this.opts)
    this.gauge.maxValue = this.percentage ? 100 : 1
    this.gauge.setMinValue(0)
    this.gauge.set(this.value)
  }
}
</script>

<style scoped>
canvas {
  width: 100%;
  height: 100%;
}

h5 {
  text-align: center;
  color: var(--bs-cyan);
}

div {
  border: 3px solid #ccc;
  border-radius: 6px;
  margin: 5px;
  background-color: rgba(0, 0, 0, 0.4);
}
</style>
