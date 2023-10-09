const SUPADUPAPAPA = {
    HOST: 'http://localhost:8080',
    URLS: {
        LOGIN_PAGE: 'http://localhost:8080' + '/login', 
        PROFILE_PAGE: 'http://localhost:8080' + '/profile', 
        LOGIN: 'http://localhost:8080' + '/api/v1/token',  // POST
        LOGOUT: 'http://localhost:8080' + '/api/v1/token'  // DELETE
    }
}

/**
 * token is a arbitary string that act as key to access restricted resources
 */
function getToken() {
    return window.localStorage.getItem("token")
}

/**
 * token is a arbitary string that act as key to access restricted resources
 */
function setToken(token) {
    if(null == token) {
        window.localStorage.clear("token")

        return
    }
    window.localStorage.setItem("token", token)
}

/**
 * functionality when user logging out from the system
 */
async function logout() {
    if(getToken()) {
        setToken(null);
    }

    window.location.replace(SUPADUPAPAPA.HOST)
}

/**
 * functionality when user logging in to the system
 */
async function login(username, password) {
    try {
        const response = await fetch(SUPADUPAPAPA.URLS.LOGIN, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username: username, password: password })
        })

        const result = await response.json()

        setToken(result.data.token)
        window.location.replace(SUPADUPAPAPA.HOST)
    } catch(err) {
        console.log("[AUTH] failed to logging in", err)
        console.alert("[AUTH] failed to logging in")
    }
}
