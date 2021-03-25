param location string = resourceGroup().location

param planName string = 'app-plan-linux'
param planTier string = 'P1v2'

param webappName string = 'vuego-demoapp'
param webappImage string = 'ghcr.io/benc-uk/vuego-demoapp:latest'
param weatherKey string = ''
param ipStackKey string = ''

resource appServicePlan 'Microsoft.Web/serverFarms@2020-06-01' = {
  name: planName
  location: location
  kind: 'linux'
  sku: {
    name: planTier
  }
  properties: {
    reserved: true
  }
}

resource webApp 'Microsoft.Web/sites@2018-11-01' = {
  name: webappName
  location: location
  properties: {
    serverFarmId: appServicePlan.id
    siteConfig: {
      appSettings:[
        {
          name: 'WEATHER_API_KEY'
          value: weatherKey
        }
        {
          name: 'IPSTACK_API_KEY'
          value: ipStackKey
        }        
      ]
      linuxFxVersion: 'DOCKER|${webappImage}'
    }
  }
}