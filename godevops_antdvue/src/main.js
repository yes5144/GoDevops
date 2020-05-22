import Vue from 'vue'

import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
Vue.config.productionTip = false

import Viser from 'viser-vue'
Vue.use(Viser)

Vue.use(Antd)

// 引入路由
import router from './routes.js'

import App from './App.vue'
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
