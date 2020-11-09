<!--
Weather page, fetches weather data from API
----------------------------------------------
Ben C, Sept 2018
-->

<template>
  <div class="card">
    <div class="card-header bg-success h3">
      <fa icon="umbrella" />&nbsp; Weather
    </div>
    <div class="card-body">
      <b-alert v-if="error" show variant="warning">
        <h4>There was a problem 😥</h4>
        <div class="errmsg">
          {{ error }}
        </div>
      </b-alert>

      <spinner v-if="!weather && !error" />

      <div v-if="weather" id="weather-div">
        <skycon :condition="weather.weather.currently.icon" :width="256" :height="256" color="#223322" />
      </div>

      <table v-if="weather" class="table table-hover">
        <tbody>
          <tr>
            <td><b>IP Address:</b></td>
            <td>{{ weather.ipAddress }}</td>
          </tr>
          <tr>
            <td><b>Location:</b></td>
            <td>{{ weather.location.city }} / {{ weather.location.country_name }}</td>
          </tr>
          <tr>
            <td><b>Summary:</b></td>
            <td>{{ weather.weather.currently.summary }}</td>
          </tr>
          <tr>
            <td><b>Temperature:</b></td>
            <td>{{ weather.weather.currently.temperature }} &deg;C</td>
          </tr>
          <tr>
            <td><b>Precipitation:</b></td>
            <td>{{ weather.weather.currently.precipProbability }}%</td>
          </tr>
          <tr>
            <td><b>Wind Speed:</b></td>
            <td>{{ weather.weather.currently.windSpeed }} km/h</td>
          </tr>
          <tr>
            <td><b>UV Index:</b></td>
            <td>{{ weather.weather.currently.uvIndex }}</td>
          </tr>
          <tr>
            <td><b>Humidity:</b></td>
            <td>{{ weather.weather.currently.humidity }}</td>
          </tr>
        </tbody>
      </table>
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
  mixins: [ apiMixin ],

  data: function() {
    return {
      weather: null,
      error: null
    }
  },

  computed: {
  },

  created() {
    this.update()
  },

  methods: {
    // Update the data with an API call
    update: async function() {
      this.error = null
      try {
        this.weather = await this.apiGetWeather()
      } catch (err) {
        this.error = err
        if (err.toString().includes('Not Implemented')) {
          this.error = 'Feature not enabled on the server, WEATHER_API_KEY & IPSTACK_API_KEY need to be set'
        }
      }
    }
  }
}
</script>

<style>
#weather-div {
  width: 100%;
  text-align:center;
}
</style>
