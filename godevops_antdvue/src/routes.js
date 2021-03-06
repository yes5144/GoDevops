import Vue from "vue"
import VueRouter from "vue-router"
import store from './vuex/store';

// 引入组件
import Layout from "./views/Layout.vue"
import Login from "./views/Login.vue"
import Register from "./views/Register.vue"
import HelloWorld from "./components/HelloWorld.vue"
import Dash01 from "./views/Dashboard/dash01.vue"
import Dash02 from "./views/Dashboard/dash02.vue"
import DeployDetail from "./views/AppDeploy/DeployDetail.vue"
import AppPack from "./views/AppDeploy/AppPack.vue"

// 要告诉 vue 使用 vueRouter
Vue.use(VueRouter)

const routes = [
    {
        path: "/",
        component: Layout,
        name: '',
        icon: 'home',
        children: [
            { path: "hello", component: HelloWorld }
        ]
    },
    {
        path: "/dashboard",
        component: Layout,
        name: "Dashboard",
        icon: 'dashboard',
        children: [
            { path: "dash01", component: Dash01, name: "dash01" },
            { path: "dash02", component: Dash02, name: "dash02" }
        ]
    },
    {
        path: "/appdeploy",
        component: Layout,
        name: "appdeploy",
        icon: 'gitlab',
        children: [
            { path: "deploydetail", component: DeployDetail, name: "部署管理" },
            { path: "apppack", component: AppPack, name: "应用打包" }
        ]
    },
    {
        path: "/system",
        component: Layout,
        name: "system",
        icon: 'setting',
        children: [
            { path: "Account", name: "账户管理", component: () => import('./views/System/Account') },
            { path: "Role", name: "角色管理", component: () => import('./views/System/Role') },
            { path: "Setting", name: "系统设置", component: () => import('./views/System/Setting') }
        ]
    },
    {
        path: "/login",
        component: Login,
        name: '',
        icon: '',
    },
    {
        path: "/register",
        component: Register,
        name: '',
        icon: '',
    }
]

// 页面刷新时，重新赋值token和username
if (window.localStorage.getItem('token')) {
    store.commit("LOGIN", { token: window.localStorage.getItem('token'), user: window.localStorage.getItem('user') });
}


var router = new VueRouter({
    mode: 'history',
    routes: routes
})

router.beforeEach((to, from, next) => {
    if (to.path == '/login') {
        store.commit("LOGOUT");
    }
    console.log('route.js 判断有没有token：', store.state);

    let token = store.state.token;
    if (!token && to.path != '/login') {
        next({
            path: '/login',
            //query: {redirect: to.fullPath}
            query: { redirect: '/welcome' }
        });
    } else {
        // // 判断是否需要折叠导航栏
        // switch (to.path) {
        //     case '/zabbix/map':
        //         store.dispatch('collapsedClose');
        //         break;
        //     default:
        //         store.dispatch('collapsedOpen');
        // }
        next();
    }
});

export default router
