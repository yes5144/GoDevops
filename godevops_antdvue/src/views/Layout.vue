<template>
  <a-layout id="components-layout-demo-side" style="min-height: 100vh">
    <a-layout-sider v-model="collapsed" collapsible>
      <div class="logo" />
      <a-menu theme="dark" :default-selected-keys="['1']" mode="inline">
        <a-menu-item key="1">
          <a-icon type="home" />
          <span>Hello SRE</span>
          <router-link to="/hello">手机</router-link>
        </a-menu-item>
        <a-sub-menu key="sub1">
          <span slot="title">
            <a-icon type="dashboard" />
            <span>Dashboard</span>
          </span>
          <a-menu-item key="3">
            <router-link to="/dashboard/dash01">dash01</router-link>
          </a-menu-item>
          <a-menu-item key="4">
            <router-link to="/dashboard/dash02">dash02</router-link>
          </a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="sub2">
          <span slot="title">
            <a-icon type="gitlab" />
            <span>AppDeploy</span>
          </span>
          <a-menu-item key="8">
            <router-link to="/appdeploy/apppack">应用打包</router-link>
          </a-menu-item>
          <a-menu-item key="9">
            <router-link to="/appdeploy/deploydetail">应用部署</router-link>
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0" />
      <a-layout-content style="margin: 0 16px">
        <!-- <div :style="{ padding: '24px', background: '#fff', minHeight: '360px' }">Bill is a cat.</div> -->
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer style="text-align: center">Ant Design ©2018 Created by Ant UED</a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
export default {
  data() {
    return {
      collapsed: false
    };
  }
};
</script>

<style>
#components-layout-demo-side .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
</style>
