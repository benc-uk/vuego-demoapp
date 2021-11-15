```bash
RES_GRP=vuego-demoapp
REGION=northeurope
az group create -n $RES_GRP -l $REGION -o table
az deployment group create --template-file container-app.bicep -g $RES_GRP
```