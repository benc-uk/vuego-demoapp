import * as msal from '@azure/msal-browser'

// MSAL object used for signing in users with MS identity platform
let msalApp

export default {
  //
  // Configure with clientId or empty string/null to set in "demo" mode
  //
  async configure(clientId) {
    // Can only call configure once
    if (msalApp) {
      return
    }
    // Skip if clientId isn't set
    if (!clientId) {
      return
    }

    const config = {
      auth: {
        clientId: clientId,
        redirectUri: window.location.origin,
        authority: 'https://login.microsoftonline.com/common'
      },
      cache: {
        cacheLocation: 'localStorage'
      }
      // Only uncomment when you *really* need to debug what is going on in MSAL
      /* system: {
        logger: new msal.Logger(
          (logLevel, msg) => { console.log(msg) },
          {
            level: msal.LogLevel.Verbose
          }
        )
      } */
    }
    console.log('### Azure AD sign-in: enabled\n', config)

    // Create our shared/static MSAL app object
    msalApp = new msal.PublicClientApplication(config)
  },

  //
  // Return the configured client id
  //
  clientId() {
    if (!msalApp) {
      return null
    }

    return msalApp.clientId
  },

  //
  // Login a user with a popup
  //
  async login() {
    if (!msalApp) {
      return
    }

    const LOGIN_SCOPES = ['user.read', 'openid', 'profile', 'email']
    await msalApp.loginPopup({
      scopes: LOGIN_SCOPES,
      prompt: 'select_account'
    })
  },

  //
  // Logout any stored user
  //
  logout() {
    if (!msalApp) {
      return
    }

    msalApp.logoutPopup()
  },

  //
  // Call to get user, probably cached and stored locally by MSAL
  //
  user() {
    if (!msalApp) {
      return null
    }

    const currentAccounts = msalApp.getAllAccounts()
    if (!currentAccounts || currentAccounts.length === 0) {
      // No user signed in
      return null
    } else if (currentAccounts.length > 1) {
      return currentAccounts[0]
    } else {
      return currentAccounts[0]
    }
  },

  //
  // Call through to acquireTokenSilent or acquireTokenPopup
  //
  async acquireToken(scopes = ['user.read']) {
    if (!msalApp) {
      return null
    }

    // Set scopes for token request
    const accessTokenRequest = {
      scopes,
      account: this.user()
    }

    let tokenResp
    try {
      // 1. Try to acquire token silently
      tokenResp = await msalApp.acquireTokenSilent(accessTokenRequest)
      console.log('### MSAL acquireTokenSilent was successful')
    } catch (err) {
      // 2. Silent process might have failed so try via popup
      tokenResp = await msalApp.acquireTokenPopup(accessTokenRequest)
      console.log('### MSAL acquireTokenPopup was successful')
    }

    // Just in case check, probably never triggers
    if (!tokenResp.accessToken) {
      throw new Error("### accessToken not found in response, that's bad")
    }

    return tokenResp.accessToken
  },

  //
  // Clear any stored/cached user
  //
  clearLocal() {
    if (msalApp) {
      for (let entry of Object.entries(localStorage)) {
        let key = entry[0]
        if (key.includes('login.windows')) {
          localStorage.removeItem(key)
        }
      }
    }
  },

  isConfigured() {
    return msalApp != null
  }
}
