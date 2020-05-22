<template>
  <div>
    <a-row type="flex" :gutter="6">
      <a-card>
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
      </a-card>
    </a-row>

    <!-- next -->

    <template>
      <div>
        <a-row type="flex">
          <a-col :span="6" :order="4">1 col-order-4</a-col>
          <a-col :span="6" :order="3">2 col-order-3</a-col>
          <a-col :span="6" :order="2">3 col-order-2</a-col>
          <a-col :span="6" :order="1">4 col-order-1</a-col>
        </a-row>
      </div>
    </template>
  </div>
</template>



<script>
export default {
  methods: {
    onChange (date, dateString) {
      console.log(date, dateString)
    },

    handleChange (value) {
      console.log(`selected ${value}`)
    },
    // dd
    handleChange2 (value) {
      console.log(`selected ${value}`)
    },
    handleBlur () {
      console.log('blur')
    },
    handleFocus () {
      console.log('focus')
    },
    filterOption (input, option) {
      return (
        option.componentOptions.children[0].text.toLowerCase().indexOf(input.toLowerCase()) >= 0
      )
    },
  },



};
</script>
