<template>
  <a-card>
    <a-form layout="inline" :form="form" @submit="handleSubmit">
      <a-form-item>
        <!-- <a-input placeholder="input client abs path" allow-clear @change="onChange" /> -->
        <a-input
          v-decorator="[
          'client_path',
        ]"
          placeholder="client Path"
          allow-clear
          v-model="client_path"
        />
      </a-form-item>

      <a-form-item>
        <!-- <a-input placeholder="input server abs path" allow-clear @change="onChange" /> -->
        <a-input
          v-decorator="[
          'server_path',
        ]"
          placeholder="server Path"
          allow-clear
          v-model="server_path"
        />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" @click="postIds()">批量打包</a-button>
      </a-form-item>
    </a-form>

    <!-- showModal 模态框 -->
    <div>
      <a-button icon="plus" type="primary" @click="showModal">新建</a-button>
    </div>
  </a-card>
</template>

<script>
export default {};
</script>

<style>
</style>