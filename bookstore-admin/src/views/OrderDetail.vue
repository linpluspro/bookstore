<template>
  <el-card class="order-container">
    <div class="data">
      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>订单状态</span>
          </div>
        </template>
        <div>
          {{ state.data.orderStatusString }}
        </div>
      </el-card>
      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>创建时间</span>
          </div>
        </template>
        <div>
          {{ state.data.createTime }}
        </div>
      </el-card>
      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>订单号</span>
          </div>
        </template>
        <div>
          {{ state.data.orderNo }}
        </div>
      </el-card>

      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>收货人</span>
          </div>
        </template>
        <div>
          {{ state.username }}
        </div>
      </el-card>

      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>收货电话</span>
          </div>
        </template>
        <div>
          {{ state.tel }}
        </div>
      </el-card>

      <el-card class="data-item" shadow="hover">
        <template #header>
          <div class="card-header">
            <span>收货地址</span>
          </div>
        </template>
        <div>
          {{ state.address }}
        </div>
      </el-card>

    </div>
    <el-table
        :data="state.tableData"
        tooltip-effect="dark"
        style="width: 100%"
    >
      <el-table-column
          label="图书图片"
      >
        <template #default="scope">
          <img style="width: 100px" :key="scope.row.booksId" :src="$filters.prefix(scope.row.booksCoverImg)" alt="图书主图">
        </template>
      </el-table-column>
      <el-table-column
          prop="booksId"
          label="图书编号"
      >
      </el-table-column>
      <el-table-column
          prop="booksName"
          label="图书名"
      ></el-table-column>
      <el-table-column
          prop="booksCount"
          label="图书数量"
      >
      </el-table-column>
      <el-table-column
          prop="sellingPrice"
          label="价格"
      >
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import {onMounted, reactive} from 'vue'
import {useRoute} from 'vue-router'
import axios from '@/utils/axios'

const route = useRoute()
const {id} = route.query
const state = reactive({
  data: {},
  tableData: [],
  username: '',
  tel: '',
  address: ''
})
onMounted(() => {
  axios.get(`/orders/${id}`).then(res => {
    state.data = res
    state.tableData = res.bookStoreOrderItemVOS
    state.username = state.data.bookStoreOrderAddressVOS.userName
    state.tel = state.data.bookStoreOrderAddressVOS.userPhone
    state.address = state.data.bookStoreOrderAddressVOS.provinceName+'/'+state.data.bookStoreOrderAddressVOS.regionName +'/' +state.data.bookStoreOrderAddressVOS.detailAddress
  })
})
</script>

<style scoped>
.data {
  display: flex;
  margin-bottom: 50px;
}

.data .data-item {
  flex: 1;
  margin: 0 10px;
}

.el-table {
  border: 1px solid #EBEEF5;
  border-bottom: none;
}

.has-gutter th {
  border-right: 1px solid #EBEEF5;
}

.has-gutter th:last-child {
  border-right: none;
}

.el-table__row td {
  border-right: 1px solid #EBEEF5;
}

.el-table__row td:last-child {
  border-right: none;
}
</style>