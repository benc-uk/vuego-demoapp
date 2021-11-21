// ============================================================================
// Deploy a container app with app container environment and log analytics
// ============================================================================

@description('Name of container app')
param appName string = 'vuego-demoapp'

@description('Region to deploy into')
param location string = resourceGroup().location

@description('Container image to deploy')
param image string = 'ghcr.io/benc-uk/vuego-demoapp:latest'

@description('Optional feature: OpenWeather API Key')
param weatherApiKey string = ''

@description('Optional feature: Azure AD Client ID')
param authClientId string = ''

// ===== Variables ============================================================

var logWorkspaceName = '${resourceGroup().name}-logs'
var environmentName = '${resourceGroup().name}-environment'

// ===== Modules & Resources ==================================================

resource logWorkspace 'Microsoft.OperationalInsights/workspaces@2020-08-01' = {
  location: location
  name: logWorkspaceName
  properties:{
    sku:{
      name: 'Free'
    }
  }
}

resource kubeEnv 'Microsoft.Web/kubeEnvironments@2021-02-01' = {
  location: location
  name: environmentName
  kind: 'containerenvironment'
  
  properties: {
    type: 'Managed'
    appLogsConfiguration: {
      destination: 'log-analytics'
      logAnalyticsConfiguration: {
        customerId: logWorkspace.properties.customerId 
        sharedKey: logWorkspace.listKeys().primarySharedKey
      }
    }
  }
}

resource containerApp 'Microsoft.Web/containerApps@2021-03-01' = {
  location: location
  name: appName

  properties: {
    kubeEnvironmentId: kubeEnv.id
    template: {
      containers: [
        {
          image: image
          name: appName
          resources: {
            cpu: json('0.25')
            memory: '0.5Gi'
          }
          env: [
            {
              name: 'WEATHER_API_KEY'
              value: weatherApiKey
            }
            {
              name: 'AUTH_CLIENT_ID'
              value: authClientId
            }
          ]
        }
      ]
    }

    configuration: {
      ingress: {
        external: true
        targetPort: 4000
      }
    }
  }
}

// ===== Outputs ==============================================================

output appURL string = 'https://${containerApp.properties.configuration.ingress.fqdn}'
