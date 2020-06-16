<template>
  <div>
    <a-breadcrumb style="margin: 16px 0">
      <a-breadcrumb-item>appdeploy</a-breadcrumb-item>
      <a-breadcrumb-item>apppack</a-breadcrumb-item>
    </a-breadcrumb>

    <a-card>
      <a-col :span="8" :order="1">
        <!-- <a-input placeholder="input client abs path" allow-clear @change="onChange" /> -->
        <a-input
          v-decorator="[
          'client_path',
        ]"
          placeholder="client Path"
          allow-clear
          v-model="client_path"
        />
      </a-col>
      <a-col :span="8" :order="2">
        <!-- <a-input placeholder="input server abs path" allow-clear @change="onChange" /> -->
        <a-input
          v-decorator="[
          'server_path',
        ]"
          placeholder="server Path"
          allow-clear
          v-model="server_path"
        />
      </a-col>
      <a-button color="blue" type="primary; margin-right: 5px;">Add</a-button>
      <a-button type="primary" @click="postIds()">批量打包</a-button>
    </a-card>

    <a-card>
      <div>
        <a-table :row-selection="rowSelection" :columns="columns" :data-source="data" row-key="id">
          <a slot="id" slot-scope="text">{{ text }}</a>
        </a-table>
      </div>
    </a-card>
    <a-card>
      <a-row type="flex" :gutter="6">
        <a-col :span="6" :order="1">
          <a-select
            mode="multiple"
            :default-value="['a1', 'b2']"
            style="width: 100%"
            placeholder="Please select"
            allow-clear
            @change="handleChange"
          >
            <a-select-option
              v-for="i in 25"
              :key="(i + 9).toString(36) + i"
            >{{ (i + 9).toString(36) + i }}</a-select-option>
          </a-select>
        </a-col>
        <a-col :span="6" :order="3">
          <a-select
            show-search
            placeholder="Select a person"
            option-filter-prop="children"
            style="width: 200px"
            :filter-option="filterOption"
            @focus="handleFocus"
            @blur="handleBlur"
            @change="handleChange2"
          >
            <a-select-option value="jack">Jack</a-select-option>
            <a-select-option value="lucy">Lucy</a-select-option>
            <a-select-option value="tom">Tom</a-select-option>
          </a-select>
        </a-col>
        <a-col :span="6" :order="2">
          <!-- <a-input placeholder="input with clear icon" allow-clear @change="onChange" /> -->
          <a-input placeholder="input placeholder" />
        </a-col>
        <a-col :span="6" :order="4">
          <a-range-picker @change="onChange">
            <a-icon slot="suffixIcon" type="smile" />
          </a-range-picker>
        </a-col>
      </a-row>

      <!-- <a-row>
            <a-col :span="8">col-8</a-col>
            <a-col :span="8" :offset="8">col-8</a-col>
      </a-row>-->
    </a-card>

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
    postIds() {
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
    }
  }
};
</script>
