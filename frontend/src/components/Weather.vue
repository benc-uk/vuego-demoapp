<!--
Weather page, fetches weather data from API
----------------------------------------------
Ben C, April 2018, Updated for Vue3 2021
-->

<template>
  <div class="card border-light bg-dark mb-3 rounded-3">
    <h4 class="card-header bg-info">
      <i class="fas fa-umbrella fa-fw me-3"></i>
      Weather
    </h4>
    <div class="card-body">
      <div v-if="error" class="alert text-dark alert-warning">
        <h4>There was a problem 😥</h4>
        <div class="errmsg">
          {{ error }}
        </div>
      </div>

      <spinner v-if="!weather && !error" />

      <div v-if="weather" class="row">
        <div class="col">
          <table class="table table-hover table-dark table-striped">
            <tbody>
              <tr>
                <td><b>Location:</b></td>
                <td>{{ weather.name }} ({{ weather.sys.country }})</td>
              </tr>
              <tr>
                <td><b>Summary:</b></td>
                <td class="text-capitalize">{{ weather.weather[0].description }}</td>
              </tr>
              <tr>
                <td><b>Temperature:</b></td>
                <td>{{ weather.main.temp }} °C (Feels like {{ weather.main.feels_like }} °C)</td>
              </tr>
              <tr>
                <td><b>Cloud Cover:</b></td>
                <td>{{ weather.clouds.all }}%</td>
              </tr>
              <tr>
                <td><b>Rain:</b></td>
                <td>{{ weather.rain ? weather.rain['1h'] : 0 }}mm</td>
              </tr>
              <tr>
                <td><b>Wind:</b></td>
                <td>{{ weather.wind.speed }} m/s</td>
              </tr>
              <tr>
                <td><b>Humidity:</b></td>
                <td>{{ weather.main.humidity }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="col d-flex justify-content-center">
          <i class="owi weather-icon" :class="'owi-' + weather.weather[0].icon"></i>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import apiMixin from '../mixins/apiMixin.js'
import Spinner from './Spinner.vue'

export default {
  components: {
    Spinner
  },

  // Adds functions to call the API
  mixins: [apiMixin],

  data: function () {
    return {
      weather: null,
      error: null
    }
  },

  computed: {},

  created() {
    this.update()
  },

  methods: {
    // Update the data with an API call
    update: async function () {
      this.error = null
      try {
        if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition(
            // Callback when position is found
            async (pos) => {
              this.fetchWeather(pos.coords.latitude, pos.coords.longitude)
            },
            // Callback when position error
            async (err) => {
              let errMessage = err.message
              // API only allowed on localhost and HTTPS domains
              if (err.message.startsWith('Only secure origins are allowed')) {
                errMessage = 'getCurrentPosition API only works on secure (HTTPS) domains'
              }
              this.error = errMessage + '. Will fall back to showing weather for London'
              this.fetchWeather()
            }
          )
        } else {
          this.error = "Geolocation is not supported by this browser. Maybe it's time to upgrade!"
          this.fetchWeather()
        }
      } catch (err) {
        this.error = err
        if (err.toString().includes('Not Implemented')) {
          this.error = 'Feature not enabled on the server, WEATHER_API_KEY needs to be set'
        }
      }
    },

    // Fetch the weather data from the API, with London as fallback
    fetchWeather: async function (lat = 51.5072, long = 0.1276) {
      try {
        this.weather = await this.apiGetWeather(lat, long)
      } catch (err) {
        this.error = err.message
      }
    }
  }
}
</script>

<style scoped>
.table {
  font-size: 1.3rem;
}
.weather-icon {
  font-size: clamp(15rem, 25vw, 25rem);
}
</style>
