# Quick Deploy

Deploy as a Azure Container App

```bash
RES_GRP=demoapps
REGION=northeurope
az group create --name $RES_GRP --location $REGION -o table
az deployment group create --template-file container-app.bicep --resource-group $RES_GRP
```

Optional parameters
 - **weatherKey** - Set to an OpenWeather API key, see main docs
 - **azureClientId** - Set to an Azure AD Client ID, see main docs

