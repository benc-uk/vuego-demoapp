<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <i class="fas fa-cogs"></i>&nbsp; System Info
    </div>
    <div class="card-body">
      <spinner v-if="!infoComputed"></spinner>

      <table v-if="infoComputed" class="table table-hover">
        <tbody>
          <tr v-for="(val, key) in infoComputed">
            <td ><b>{{ key | titleify }}</b></td>
            <td>{{ val }}</td>
          </tr>
        </tbody>        
      </table>

      <h4>Environment Variables</h4>
      <pre>{{envVars}}</pre>

    </div>
  </div>
</template>

<script>
import Spinner from "./Spinner.vue";
const info = null;

export default {
  data: function() {
    return {
      info: info
    };
  },

  components: {
    Spinner
  },

  created() {
    this.getInfo();
    setInterval(this.getInfo, 5000);
  },

  methods: {
    getInfo: function() {
      fetch("http://localhost:4000/api/info")
        .then(resp => {
          return resp.json();
        })
        .then(json => {
          this.info = json;
        });
    }
  },

  filters: {
    titleify: function(value) {
      if (!value) return "";
      value = value.toString();
      value = value.replace(/([A-Z])/g, ' $1')
      value = value.replace(/^./, function(str){ return str.toUpperCase(); });
      return value;
    }
  },

  computed: {
    infoComputed: function () {
      var result = {};
      for (let k in this.info) {
        if(k != "envVars") result[k] = this.info[k]
      }
      return result;
    },

    envVars: function () {
      if(!this.info) return "";
      var result = "";
      for (let e of this.info.envVars) {
        result += e + "\n";
      }
      return result;
    }    
  }  
}
</script>

<style>
  pre {
    background-color: #222;
    color:rgb(59, 190, 33);
    padding: 10px;
    max-height: 500px;
    font-family: 'Lucida Console', monospace
  }
</style>