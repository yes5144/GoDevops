<template>
  <div>
    <a-breadcrumb style="margin: 16px 0">
      <a-breadcrumb-item>appdeploy</a-breadcrumb-item>
      <a-breadcrumb-item>apppack</a-breadcrumb-item>
    </a-breadcrumb>

    <a-card class="card">
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
          <a-button type="primary" @click="postVersionIds()">批量打包</a-button>
        </a-form-item>
      </a-form>

      <!-- showModal 模态框 -->
      <div>
        <a-button icon="plus" type="primary" @click="showModal">新建</a-button>
      </div>
      <!-- version data table -->
      <a-table :row-selection="rowSelection" :columns="columns" :data-source="data" row-key="id">
        <!-- <a slot="id" slot-scope="text">{{ text }}</a> -->
        <template slot="id" slot-scope="text, record">
          <a-button-group :size="size" v-if="data.length">
            <a-button type="primary" @click="() => packVersionOne(record.id)">打包</a-button>
            <a-button type="danger">nothing</a-button>
          </a-button-group>
        </template>
      </a-table>
    </a-card>

    <!-- 自选录入input -->
    <a-card title="打包记录">
      <a-table :columns="columns2" :data-source="data2">
        <a slot="name" slot-scope="text">{{ text }}</a>
        <span slot="customTitle">
          <a-icon type="smile-o" />Name
        </span>
        <span slot="tags" slot-scope="tags">
          <a-tag
            v-for="tag in tags"
            :key="tag"
            :color="tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'"
          >{{ tag.toUpperCase() }}</a-tag>
        </span>
        <span slot="action" slot-scope="text, record">
          <a>Invite 一 {{ record.name }}</a>
          <a-divider type="vertical" />
          <a>Delete</a>
          <a-divider type="vertical" />
          <a class="ant-dropdown-link">
            More actions
            <a-icon type="down" />
          </a>
        </span>
      </a-table>
    </a-card>

    <!-- 模态对话框 -->
    <div>
      <a-modal
        title="Title"
        :visible="visible"
        :confirm-loading="confirmLoading"
        @ok="handleOk"
        @cancel="handleCancel"
      >
        <p>{{ ModalText }}</p>
      </a-modal>
    </div>
  </div>
</template>

<script>
import { postPackIds, getAllPackVersions } from "../../api";
// import { postPackIds } from "../../api";
const columns = [
  {
    title: "项目",
    dataIndex: "project"
  },
  {
    title: "大区",
    dataIndex: "zone"
  },
  {
    title: "版本号",
    dataIndex: "version"
  },
  {
    title: "上次打包日期",
    dataIndex: "update_time"
  },
  {
    title: "操作",
    dataIndex: "id",
    scopedSlots: { customRender: "id" }
  }
];

// const data = [];

const columns2 = [
  {
    dataIndex: "name",
    key: "name",
    slots: { title: "customTitle" },
    scopedSlots: { customRender: "name" }
  },
  {
    title: "Age",
    dataIndex: "age",
    key: "age"
  },
  {
    title: "Address",
    dataIndex: "address",
    key: "address"
  },
  {
    title: "Tags",
    key: "tags",
    dataIndex: "tags",
    scopedSlots: { customRender: "tags" }
  },
  {
    title: "Action",
    key: "action",
    scopedSlots: { customRender: "action" }
  }
];

const data2 = [
  {
    key: "1",
    name: "John Brown",
    age: 32,
    address: "New York No. 1 Lake Park",
    tags: ["nice", "developer"]
  },
  {
    key: "2",
    name: "Jim Green",
    age: 42,
    address: "London No. 1 Lake Park",
    tags: ["loser"]
  },
  {
    key: "3",
    name: "Joe Black",
    age: 32,
    address: "Sidney No. 1 Lake Park",
    tags: ["cool", "teacher"]
  }
];

// for (let i = 0; i < 22; i++) {
//   data.push({
//     key: i,
//     project: `nz`,
//     zone: "hf",
//     version: `version ${i}`,
//     newVersion: `newVersion ${i}`,
//     update_time: "2020-05-02"
//   });
// }

export default {
  data() {
    return {
      ModalText: "Content of the modal",
      visible: false,
      confirmLoading: false,
      client_path: "",
      server_path: "",
      data: [],
      columns,
      data2,
      columns2,
      selectedRowKeys: [] // Check here to configure the default column
    };
  },
  computed: {
    rowSelection() {
      const { selectedRowKeys } = this;
      return {
        selectedRowKeys,
        onChange: this.onSelectChange,
        hideDefaultSelections: true,
        selections: [
          {
            key: "all-data",
            text: "Select All Data",
            onSelect: () => {
              // this.selectedRowKeys = [...Array(this.data.length).keys()]; // 0...45
              this.selectedRowKeys = [...Map(this.data.id)]; // 0...45
            }
          },
          {
            key: "odd",
            text: "Select Odd Row",
            onSelect: changableRowKeys => {
              let newSelectedRowKeys = [];
              // newSelectedRowKeys = changableRowKeys.filter((key, index) => {
              newSelectedRowKeys = changableRowKeys.map(key => {
                // if (index % 2 !== 0) {
                //   return false;
                // }
                // return true;
                return this.data[key].id;
              });
              this.selectedRowKeys = newSelectedRowKeys;
            }
          },
          {
            key: "even",
            text: "Select Even Row",
            onSelect: changableRowKeys => {
              let newSelectedRowKeys = [];
              newSelectedRowKeys = changableRowKeys.filter((key, index) => {
                if (index % 2 !== 0) {
                  return true;
                }
                return false;
              });
              this.selectedRowKeys = newSelectedRowKeys;
            }
          }
        ],
        onSelection: this.onSelection
      };
    }
  },
  created() {
    this.getAllVersion();
  },
  methods: {
    // getAllVersion
    getAllVersion() {
      getAllPackVersions().then(res => {
        // http code
        let { code, msg, data } = res;
        console.log(code, msg, data);
        this.data = data.data;
      });
    },

    onSelectChange(selectedRowKeys) {
      console.log("selectedRowKeys changed: ", selectedRowKeys);
      this.selectedRowKeys = selectedRowKeys;
    },

    // df
    onChange(date, dateString) {
      console.log(date, dateString);
    },

    handleChange(value) {
      console.log(`selected ${value}`);
    },
    // dd
    handleChange2(value) {
      console.log(`selected ${value}`);
    },
    handleBlur() {
      console.log("blur");
    },
    handleFocus() {
      console.log("focus");
    },
    filterOption(input, option) {
      return (
        option.componentOptions.children[0].text
          .toLowerCase()
          .indexOf(input.toLowerCase()) >= 0
      );
    },
    postVersionIds() {
      const apppackIds = this.selectedRowKeys;
      if (apppackIds.length > 0) {
        console.log(apppackIds);
        const packParm = {
          packIds: apppackIds.toString(),
          serverpath: this.server_path,
          clientpath: this.client_path
        };
        console.log("packparam---------", packParm);
        postPackIds(packParm).then(res => {
          // http code
          let { code, msg, data } = res;
          console.log(code, msg, data);
        });
      } else {
        alert("select a one, then try again");
      }
    },

    // appPack
    packVersionOne(key) {
      console.log("onDelete--------key", key);
      const dataSource = [...this.data];
      this.dataSource = dataSource.filter(item => item.key !== key);

      const packParm = {
        packIds: key.toString(),
        serverpath: this.server_path,
        clientpath: this.client_path
      };
      console.log("packparam---------", packParm);
      postPackIds(packParm).then(res => {
        // http code
        let { code, msg, data } = res;
        console.log(code, msg, data);
      });
    },

    // modal
    showModal() {
      this.visible = true;
    },
    handleOk(e) {
      this.ModalText = "The modal will be closed after two seconds";
      this.confirmLoading = true;
      console.log("handleOk", e);

      setTimeout(() => {
        this.visible = false;
        this.confirmLoading = false;
      }, 2000);
    },
    handleCancel(e) {
      console.log("Clicked cancel button", e);
      this.visible = false;
    }
    // over
  }
};
</script>

<style scoped>
.card {
  margin-bottom: 24px;
}
</style>
