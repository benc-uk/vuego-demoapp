//
// All API logic here, mixin with a Component that needs to call the API
//
export default {

  data: function () {
    return {
      apiEndpoint: "/api"
    }
  },
  
  methods: {
    apiGetMetrics: function() {  
      return fetch(`${this.apiEndpoint}/metrics`)
        .then(resp => {
          return resp.json();
        })
        .catch(err => {
          console.log(`### API Error! ${err}`);
        })
    },

    apiGetInfo: function() {  
      return fetch(`${this.apiEndpoint}/info`)
        .then(resp => {
          return resp.json();
        })
        .catch(err => {
          console.log(`### API Error! ${err}`);
        })
    }    
  }
}