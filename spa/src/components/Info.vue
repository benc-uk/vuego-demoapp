<!--
Info page, fetches system info data from API
--------------------------------------------
Ben C, April 2018
-->

<template>
  <div>
    <b-alert v-if="error" show variant="warning">
      <h4>There was a problem 😥</h4>
      <div class="errmsg">
        {{ error }}
      </div>
    </b-alert>

    <spinner v-if="!info && !error" />

    <div class="card">
      <div class="card-header bg-success h3">
        <fa icon="cogs" />&nbsp; System Information
      </div>
      <div class="card-body">
        <table v-if="info" class="table table-hover table-sm table-striped">
          <tbody>
            <tr><td><fa :icon="['fab', 'docker']" fixed-width /> Containerized</td><td>{{ info.isContainer ? 'Running in a container! 😁' : 'Not running in a container 😢' }}</td></tr>
            <tr><td><fa icon="dharmachakra" fixed-width /> Kubernetes</td><td>{{ info.isKubernetes ? 'Running in Kubernetes! 😄' : 'Not running in Kubernetes 😪' }}</td></tr>
            <tr><td><fa icon="home" fixed-width /> Hostname</td><td>{{ info.hostname }}</td></tr>
            <tr><td><fa icon="wrench" fixed-width /> Platform</td><td>{{ info.platform }} {{ info.architecture }}</td></tr>
            <tr><td><fa icon="laptop-code" fixed-width /> Operating System</td><td>{{ info.os }}</td></tr>
            <tr><td><fa icon="microchip" fixed-width /> Processors</td><td>{{ info.cpuCount }} x {{ info.cpuModel }}</td></tr>
            <tr><td><fa icon="memory" fixed-width /> Memory</td><td>{{ (info.mem / (1024*1024*1024)).toFixed(2) }} GB</td></tr>
            <tr><td><fa icon="flask" fixed-width /> Go Version</td><td>{{ info.goVersion }}</td></tr>
            <tr><td><fa icon="project-diagram" fixed-width /> Network Address</td><td>{{ info.netHost }}</td></tr>
          </tbody>
        </table>
      </div>
    </div>

    <br>

    <div class="card">
      <div class="card-header bg-success h3">
        <fa icon="comment-dollar" />&nbsp; Environment Variables
      </div>
      <div class="card-body">
        <table v-if="info" class="vartable">
          <tbody>
            <tr v-for="(envVar, index) in envVars" :key="index">
              <td>{{ envVar.name }}</td><td class="value">
                {{ envVar.value }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import apiMixin from '../mixins/apiMixin.js'
import Spinner from './Spinner.vue'
const info = null

export default {
  components: {
    Spinner
  },

  // Adds functions to call the API
  mixins: [ apiMixin ],

  data: function() {
    return {
      info: info,
      error: null
    }
  },

  computed: {
    envVars: function () {
      let vars = []
      for (const envVar of this.info.envVars) {
        const parts = envVar.split('=')
        const name = parts[0]
        const value = parts[1]
        if (name.includes('PATH')) { continue }
        if (name.includes('NPM_')) { continue }
        if (name.includes('VSCODE_')) { continue }
        vars.push({ name, value })
      }
      vars = vars.sort((e1, e2) => {
        if (e1.name > e2.name) { return 1 }
        return -1
      })
      return vars
    }
  },

  created() {
    this.getInfo()
  },

  methods: {
    getInfo: async function() {
      try {
        this.info = await this.apiGetInfo()
      } catch (err) {
        this.error = err
      }
    }
  }
}
</script>

<style scoped>
  .vartable {
    font-size: 19px;
    background-color: #141414;
    border-top: none;
    border-bottom: none;
    color: #afafaf;
    overflow-x: scroll;
    width: 100%;
    table-layout: fixed;
    border-radius: 0.3rem;
  }

  .vartable tr {
    border-top: solid 1px #444444;
    font-family: 'Consolas', 'Courier New', Courier, monospace;
  }
  .vartable tr:hover {
    background-color: #222222;
  }
  .vartable tr:first-child {
    border-top: none;
  }
  .vartable td {
    padding: 0.5rem;
  }

  .value {
    color: #2dc22d;
  }
</style>