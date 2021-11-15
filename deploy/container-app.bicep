// ============================================================================
// Deploy a container app with app container environment and log analytics
// ============================================================================

@description('Name used for resource group, and default base name for all resources')
param appName string = 'vuego-demoapp'

@description('Azure region for all resources')
param location string = resourceGroup().location

@description('Container image')
param image string = 'ghcr.io/benc-uk/vuego-demoapp:latest'

@description('Optional featiure: OpenWeather API Key')
param weatherKey string = ''

@description('Optional featiure: Azure AD Client ID')
param azureClientId string = ''

// ===== Variables ============================================================

// ===== Modules & Resources ==================================================

resource logWorkspace 'Microsoft.OperationalInsights/workspaces@2020-08-01' = {
  location: location
  name: appName
  properties:{
    sku:{
      name: 'Free'
    }
  }
}

resource kubeEnv 'Microsoft.Web/kubeEnvironments@2021-02-01' = {
  location: location
  name: appName
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
              value: weatherKey
            }
            {
              name: 'AZURE_AD_CLIENT_ID'
              value: azureClientId
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
