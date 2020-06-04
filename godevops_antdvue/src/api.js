import axios from 'axios'
import store from './vuex/store'
import router from './routes'

// localStorage.baseUrl = idcUrlConfig.baseIdc
// localStorage.baseUrl = 'http://127.0.0.1:8080'
localStorage.baseUrl = process.env.VUE_APP_BASE_URL

axios.defaults.timeout = 3600000
axios.defaults.baseURL = localStorage.baseUrl


// http request 拦截器
axios.interceptors.request.use(
    config => {
        console.log("store.state.token", store.state.token);
        if (store.state.token) {
            // let url = config.url.split('/')[2]
            let needToken = localStorage.token
            console.log("localStorage: token", 'Bearer ' + needToken);
            config.headers.Authorization = `Bearer ${needToken}`
        }
        // console.log(store.state.token)
        console.log("header--config:", config);

        return config
    },
    err => {
        return Promise.reject(err);
    }
)


// http response 拦截器
axios.interceptors.response.use(
    response => {
        return response
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 401:
                    store.commit("LOGOUT")
                    router.replace({
                        path: 'login',
                        query: { redirect: router.currentRoute.path }
                    })
            }
        }
        return Promise.reject(error.response.data)
    }
)

// register 
export const requestRegister = params => {
    return axios.post(`/api/auth/register`, params).then(res => res.data)
}

// login
export const requestLogin = params => {
    return axios.post(`/api/auth/login`, params).then(res => res.data)
}

// get userinfo
export const getUserInfo = () => {
    return axios.get(`/api/auth/info`).then(res => res.data)
}

//navigation
export const getNavigation = () => {
    return axios.get(`/api/v1/navigation`).then(res => res.data);
};