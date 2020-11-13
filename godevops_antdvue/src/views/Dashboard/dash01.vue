<template>
  <a-card>
    <a-row :gutter="6">
      <a-col :span="6">
        <a-card>
          <a-statistic title="Feedback" :value="1128" style="margin-right: 50px">
            <template #suffix>
              <a-icon type="like" />
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic title="Unmerged" :value="93" class="demo-class">
            <template #suffix>
              <span>/ 100</span>
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic title="Unmerged" :value="93" class="demo-class">
            <template #suffix>
              <span>/ 100</span>
            </template>
          </a-statistic>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <a-statistic title="Unmerged" :value="93" class="demo-class">
            <template #suffix>
              <span>/ 100</span>
            </template>
          </a-statistic>
        </a-card>
      </a-col>
    </a-row>
    <a-card>
      <v-chart :forceFit="true" :height="height" :data="data" :scale="scale">
        <v-tooltip :showTitle="false" data-key="item*percent" />
        <v-axis />
        <v-legend data-key="item" />
        <v-pie position="percent" color="item" :v-style="pieStyle" :label="labelConfig" />
        <v-coord type="theta" />
      </v-chart>
    </a-card>
  </a-card>
</template>

<script>
const DataSet = require("@antv/data-set");

const sourceData = [
  { item: "事例一", count: 40 },
  { item: "事例二", count: 21 },
  { item: "事例三", count: 17 },
  { item: "事例四", count: 13 },
  { item: "事例五", count: 9 }
];

const scale = [
  {
    dataKey: "percent",
    min: 0,
    formatter: ".0%"
  }
];

const dv = new DataSet.View().source(sourceData);
dv.transform({
  type: "percent",
  field: "count",
  dimension: "item",
  as: "percent"
});
const data = dv.rows;

export default {
  data() {
    return {
      data,
      scale,
      height: 400,
      pieStyle: {
        stroke: "#fff",
        lineWidth: 1
      },
      labelConfig: [
        "percent",
        {
          formatter: (val, item) => {
            return item.point.item + ": " + val;
          }
        }
      ]
    };
  }
};
</script>