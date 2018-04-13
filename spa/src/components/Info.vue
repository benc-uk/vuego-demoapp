<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <i class="fas fa-cogs"></i>&nbsp; System Info
    </div>
    <div class="card-body">
      <spinner v-if="!info"></spinner>

      <table v-if="info" class="table table-hover">
        <tbody>
          <tr v-for="(val, key) in info">
            <td><b>{{ key }}</b></td>
            <td>{{ val }}</td>
          </tr>
        </tbody>        
      </table>
    </div>
  </div>
</template>

<script>
import Spinner from './Spinner.vue'
const info = null;

export default {
  data: function () {
    return {
      info: info
    }
  },

  components: {
    Spinner,
  },

  created () {
    this.getInfo();
    setInterval(this.getInfo, 5000);
  },

  methods: {
    getInfo: function() {
      fetch('http://localhost:4000/api/info')
        .then(resp => {
          return resp.json();
        })
        .then(json => {
          this.info = json
        });
    }
  }
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
</script>