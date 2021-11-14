<!--
Info page, fetches system info data from API
--------------------------------------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div>
    <div v-if="error" class="alert text-dark alert-warning">
      <h4>There was a problem 😥</h4>
      <div class="errmsg">
        {{ error }}
      </div>
    </div>

    <spinner v-if="!info && !error" />

    <div v-if="info" class="card border-light bg-dark mb-3 rounded-3">
      <h4 class="card-header bg-info"><i class="fas fa-cogs fa-fw me-3"></i> System Information</h4>
      <div class="card-body">
        <table class="table table-hover table-striped fs-5">
          <tbody>
            <tr class="table-dark">
              <td><i class="fab fa-docker fa-fw me-1"></i> Containerized</td>
              <td>{{ info.isContainer ? 'Running in a container! 😁' : 'Not running in a container 😢' }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-dharmachakra fa-fw me-1"></i> Kubernetes</td>
              <td>{{ info.isKubernetes ? 'Running in Kubernetes! 😄' : 'Not running in Kubernetes 😪' }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-home fa-fw me-1"></i> Hostname</td>
              <td>{{ info.hostname }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-wrench fa-fw me-1"></i> Platform</td>
              <td>{{ info.platform }} {{ info.architecture }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-laptop-code fa-fw me-1"></i> Operating System</td>
              <td>{{ info.os }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-microchip fa-fw me-1"></i> Processors</td>
              <td>{{ info.cpuCount }} x {{ info.cpuModel }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-memory fa-fw me-1"></i> Memory</td>
              <td>{{ (info.mem / (1024 * 1024 * 1024)).toFixed(2) }} GB</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-flask fa-fw me-1"></i> Go Version</td>
              <td>{{ info.goVersion }}</td>
            </tr>
            <tr class="table-dark">
              <td><i class="fas fa-project-diagram fa-fw me-1"></i> Network Address</td>
              <td>{{ info.netHost }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <br />

    <div v-if="envVars" class="card border-light bg-dark mb-3 rounded-3">
      <h4 class="card-header bg-info"><i class="fas fa-comment-dollar fa-fw me-3"> </i> Environment Variables</h4>
      <div class="card-body">
        <table class="vartable">
          <tbody>
            <tr v-for="(envVar, index) in envVars" :key="index">
              <td>{{ envVar.name }}</td>
              <td class="value">
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
  mixins: [apiMixin],

  data: function () {
    return {
      info: info,
      error: null
    }
  },

  computed: {
    envVars: function () {
      let vars = []
      if (!this.info) return null
      for (const envVar of this.info.envVars) {
        const parts = envVar.split('=')
        const name = parts[0]
        const value = parts[1]
        if (name.includes('PATH')) {
          continue
        }
        if (name.includes('NPM_')) {
          continue
        }
        if (name.includes('VSCODE_')) {
          continue
        }
        vars.push({ name, value })
      }
      vars = vars.sort((e1, e2) => {
        if (e1.name > e2.name) {
          return 1
        }
        return -1
      })
      return vars
    }
  },

  created() {
    this.getInfo()
  },

  methods: {
    getInfo: async function () {
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
