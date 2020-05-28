import Vue from 'vue'
import Vuex from 'vuex'


Vue.use(Vuex)

// init state
const state = {
    token: null,
    user: null,
    collapsed: false,
};

// define mutations
const mutations = {
    LOGIN: (state, data) => {
        state.token = data.token;
        state.user = null;
    },
    LOGOUT: (state) => {
        state.token = null;
        state.user = null;
    }
}

// create store
const store = new Vuex.Store({
    state,
    mutations
});

export default store