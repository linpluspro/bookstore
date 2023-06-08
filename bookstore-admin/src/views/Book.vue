<template>
  <el-card class="book-container">
    <template #header>
      <div class="header">
        <el-button type="primary" :icon="Plus" @click="handleAdd">新增图书</el-button>
      </div>
    </template>
    <el-table
        :load="state.loading"
        :data="state.tableData"
        tooltip-effect="dark"
        style="width: 100%"
    >
      <el-table-column
          prop="booksId"
          label="图书编号"
      >
      </el-table-column>
      <el-table-column
          prop="booksName"
          label="图书名"
          width="150px"
          :show-overflow-tooltip="true"
      >
      </el-table-column>
      <el-table-column
          prop="booksAuthor"
          label="图书作者"
          width="150px"
          :show-overflow-tooltip="true"
      >
      </el-table-column>
      <el-table-column
          prop="booksAuthorCountry"
          label="图书作者国籍"
      >
      </el-table-column>
      <el-table-column
          prop="booksPublish"
          label="出版社出版时间"
      >
      </el-table-column>
      <el-table-column
          prop="booksIntro"
          label="图书简介"
          width="300px"
          :show-overflow-tooltip="true"
      >
      </el-table-column>
      <el-table-column
          label="图书图片"
          width="150px"
      >
        <template #default="scope">
          <img style="width: 100px; height: 100px;" :key="scope.row.booksId"
               :src="$filters.prefix(scope.row.booksCoverImg)" alt="图书主图">
        </template>
      </el-table-column>
      <el-table-column
          prop="stockNum"
          label="图书库存"
      >
      </el-table-column>
      <el-table-column
          prop="sellingPrice"
          label="图书售价"
      >
      </el-table-column>
      <el-table-column
          label="上架状态"
      >
        <template #default="scope">
          <span style="color: green;" v-if="scope.row.booksSellStatus === 0">销售中</span>
          <span style="color: red;" v-else>已下架</span>
        </template>
      </el-table-column>

      <el-table-column
          label="操作"
          width="100"
      >
        <template #default="scope">
          <a style="cursor: pointer; margin-right: 10px" @click="handleEdit(scope.row.booksId)">修改</a>
          <a style="cursor: pointer; margin-right: 10px" v-if="scope.row.booksSellStatus === 0"
             @click="handleStatus(scope.row.booksId, 1)">下架</a>
          <a style="cursor: pointer; margin-right: 10px" v-else @click="handleStatus(scope.row.booksId, 0)">上架</a>
        </template>
      </el-table-column>
    </el-table>
    <!--总数超过一页，再展示分页器-->
    <el-pagination
        background
        layout="prev, pager, next"
        :total="state.total"
        :page-size="state.pageSize"
        :current-page="state.currentPage"
        @current-change="changePage"
    />
  </el-card>
</template>

<script setup>
import {onMounted, reactive, getCurrentInstance} from 'vue'
import axios from '@/utils/axios'
import {ElMessage} from 'element-plus'
import {Plus, Delete} from '@element-plus/icons-vue'
import {useRouter} from 'vue-router'

const app = getCurrentInstance()
const {goTop} = app.appContext.config.globalProperties
const router = useRouter()
const state = reactive({
  loading: false,
  tableData: [], // 数据列表
  total: 0, // 总条数
  currentPage: 1, // 当前页
  pageSize: 10 // 分页大小
})
onMounted(() => {
  getBookList()
})
// 获取轮播图列表
const getBookList = () => {
  state.loading = true
  axios.get('/books/list', {
    params: {
      pageNumber: state.currentPage,
      pageSize: state.pageSize
    }
  }).then(res => {
    state.tableData = res.list
    state.total = res.totalCount
    state.currentPage = res.currPage
    state.loading = false
    goTop && goTop()
  })
}
const handleAdd = () => {
  router.push({path: '/add'})
}
const handleEdit = (id) => {
  router.push({path: '/add', query: {id}})
}
const changePage = (val) => {
  state.currentPage = val
  getBookList()
}
const handleStatus = (id, status) => {
  axios.put(`/books/status/${status}`, {
    ids: id ? [id] : []
  }).then(() => {
    ElMessage.success('修改成功')
    getBookList()
  })
}
</script>

<style scoped>
.book-container {
  min-height: 100%;
}

.el-card.is-always-shadow {
  min-height: 100% !important;
}
</style>
