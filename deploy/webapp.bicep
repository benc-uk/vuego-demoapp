param location string = resourceGroup().location

param planName string = 'app-plan-linux'
param planTier string = 'P1v2'

param webappName string = 'vuego-demoapp'
param webappImage string = 'ghcr.io/benc-uk/vuego-demoapp:latest'
param weatherKey string = ''
param ipStackKey string = ''
param authClientId string = ''
param releaseInfo string = 'Released on ${utcNow('f')}'

resource appServicePlan 'Microsoft.Web/serverfarms@2020-10-01' = {
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

resource webApp 'Microsoft.Web/sites@2020-10-01' = {
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
        {
          name: 'AUTH_CLIENT_ID'
          value: authClientId
        }        
      ]
      linuxFxVersion: 'DOCKER|${webappImage}'
    }
  }
}
