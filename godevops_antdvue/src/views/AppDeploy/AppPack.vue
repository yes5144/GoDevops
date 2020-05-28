<template>
  <div>
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

    <a-card>
      <div>
        <a-table :row-selection="rowSelection" :columns="columns" :data-source="data" />
      </div>
    </a-card>
  </div>
</template>

<script>
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
    title: "新版本号",
    dataIndex: "newVersion"
  },
  {
    title: "上次打包日期",
    dataIndex: "lastDate"
  }
];

const data = [];
for (let i = 0; i < 22; i++) {
  data.push({
    key: i,
    project: `nz`,
    zone: "hf",
    version: `version ${i}`,
    newVersion: `newVersion ${i}`,
    lastDate: "2020-05-02"
  });
}

export default {
  data() {
    return {
      data,
      columns,
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
              this.selectedRowKeys = [...Array(46).keys()]; // 0...45
            }
          },
          {
            key: "odd",
            text: "Select Odd Row",
            onSelect: changableRowKeys => {
              let newSelectedRowKeys = [];
              newSelectedRowKeys = changableRowKeys.filter((key, index) => {
                if (index % 2 !== 0) {
                  return false;
                }
                return true;
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
  methods: {
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
    }
  }
};
</script>
