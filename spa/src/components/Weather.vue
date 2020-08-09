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
      <spinner v-if="!weather" />

      <div v-if="weather" id="weather-div">
        <skycon :condition="weather.weather.currently.icon" :width="256" :height="256" />
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
  mixins: [apiMixin],

  data: function() {
    return {
      weather: null,
      skycons: null
    }
  },

  computed: {
  },

  created() {
    this.update()
  },

  methods: {
    update: function() {
      this.apiGetWeather()
        .then((json) => {
          this.weather = json
        })
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
