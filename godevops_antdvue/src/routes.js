import Vue from "vue"
import VueRouter from "vue-router"

// 引入组件
import Layout from "./views/Layout.vue"
import HelloWorld from "./components/HelloWorld.vue"
import Dash01 from "./views/Dashboard/dash01.vue"
import Dash02 from "./views/Dashboard/dash02.vue"
import DeployDetail from "./views/AppDeploy/DeployDetail.vue"

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
            { path: "", component: DeployDetail, name: "部署管理" }
        ]
    }
]

var router = new VueRouter({
    mode: 'history',
    routes: routes
})
export default router
