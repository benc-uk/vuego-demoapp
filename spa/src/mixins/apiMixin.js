//
// Mixin added to components, all API logic here
// ----------------------------------------------
// Ben C, April 2018
//

export default {
  methods: {
    apiGetWeather: function () {
      return apiCall('/weather')
    },

    apiGetMetrics: function () {
      return apiCall('/metrics')
    },

    apiGetInfo: function () {
      return apiCall('/info')
    },

  }
}

//
// ===== Base fetch wrapper, not exported =====
//
async function apiCall(apiPath, method = 'get', data = null) {
  let headers = {}
  const url = `${(process.env.VUE_APP_API_ENDPOINT || '/api')}${apiPath}`
  //console.log(`### API CALL ${method} ${url}`)

  // Build request
  const request = {
    method,
    headers,
  }

  // Add payload if required
  if (data) {
    request.body = JSON.stringify(data)
  }

  // Make the HTTP request
  const resp = await fetch(url, request)

  // Decode error message when non-HTTP OK (200~299) & JSON is received
  if (!resp.ok) {
    let error = `API call to ${url} failed with ${resp.status} ${resp.statusText}`
    if (resp.headers.get('Content-Type') === 'application/json') {
      error = `Status: ${resp.statusText}\n`
      let errorObj = await resp.json()
      for (const [key, value] of Object.entries(errorObj)) {
        error += `${key}: '${value}\n', `
      }
    }
    throw new Error(error)
  }

  // Attempt to return response body as data object if JSON
  if (resp.headers.get('Content-Type') === 'application/json') {
    return resp.json()
  } else {
    return resp.text()
  }
}
