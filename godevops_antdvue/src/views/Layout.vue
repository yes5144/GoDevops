<template>
  <a-layout id="components-layout-demo-side" style="min-height: 100vh">
    <a-layout-sider v-model="collapsed" collapsible>
      <div class="logo" />
      <a-menu theme="dark" :default-selected-keys="['1']" mode="inline">
        <!-- home -->
        <a-menu-item key="1">
          <a-icon type="home" />
          <span>Hello SRE</span>
          <router-link to="/hello">手机</router-link>
        </a-menu-item>

        <!-- dashboard -->
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

        <!-- appdeploy -->
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

        <!-- system -->
        <a-sub-menu key="sub3">
          <span slot="title">
            <a-icon type="setting" />
            <span>System</span>
          </span>
          <a-menu-item key="10">
            <router-link to="/system/account">账户管理</router-link>
          </a-menu-item>
          <a-menu-item key="11">
            <router-link to="/system/role">角色管理</router-link>
          </a-menu-item>
          <a-menu-item key="12">
            <router-link to="/system/setting">系统设置</router-link>
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <div>your name</div>
      </a-layout-header>
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
